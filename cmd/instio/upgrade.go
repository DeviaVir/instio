package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli/v2"
)

func upgradeInstioProjectDir(c *cli.Context, path string) error {
	core := []string{
		".gitattributes",
		"LICENSE",
		"instio-banner.png",
		"README.md",
		"cmd",
		"pkg",
		"docs",
		"CONTRIBUTING.md",
		"examples",
	}

	stamp := fmt.Sprintf("instio-%d.bak", time.Now().Unix())
	temp := filepath.Join(os.TempDir(), stamp)
	err := os.Mkdir(temp, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	// Track non-Instio core items (added by user).
	var user []os.FileInfo
	list, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, item := range list {
		// Check if in core.
		var isCore bool
		for _, name := range core {
			if item.Name() == name {
				isCore = true
				break
			}
		}

		if !isCore {
			user = append(user, item)
		}
	}

	// Move non-Instio files to temp location.
	fmt.Println("Preserving files to be restored after upgrade...")
	for _, item := range user {
		src := filepath.Join(path, item.Name())
		if item.IsDir() {
			err := os.Mkdir(filepath.Join(temp, item.Name()), os.ModeDir|os.ModePerm)
			if err != nil {
				return err
			}
		}

		err := copyAll(src, temp)
		if err != nil {
			return err
		}

		fmt.Println(" [-]", item.Name())

	}

	// Remove all files in path.
	for _, item := range list {
		err := os.RemoveAll(filepath.Join(path, item.Name()))
		if err != nil {
			return fmt.Errorf("Failed to remove old Instio files.\n%s", err)
		}
	}

	// Development upgrade?
	if c.Bool("dev") {
		if c.String("fork") != "" {
			fmt.Println("Upgrading from " + c.String("fork"))
		} else {
			fmt.Println("Upgrading from 'dev' branch")
		}
	}
	err = createProjectInDir(c, path)
	if err != nil {
		fmt.Println("")
		fmt.Println("Upgrade failed...")
		fmt.Println("Your code is backed up at the following location:")
		fmt.Println(temp)
		fmt.Println("")
		fmt.Println("Manually create a new Instio project here and copy those files within it to fully restore.")
		fmt.Println("")
		return err
	}

	// Move non-instio files from temp location backed.
	restore, err := ioutil.ReadDir(temp)
	if err != nil {
		return err
	}

	fmt.Println("Restoring files preserved before upgrade...")
	for _, r := range restore {
		p := filepath.Join(temp, r.Name())
		err = copyAll(p, path)
		if err != nil {
			fmt.Println("Couldn't merge your previous project files with upgraded one.")
			fmt.Println("Manually copy your files from the following directory:")
			fmt.Println(temp)
			return err
		}

		fmt.Println(" [+]", r.Name())
	}

	// Clean-up.
	backups := []string{filepath.Join(path, stamp), temp}
	for _, bak := range backups {
		err := os.RemoveAll(bak)
		if err != nil {
			return err
		}
	}

	return nil
}
