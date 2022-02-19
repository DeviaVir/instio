package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

func execAndWait(command string, arg ...string) error {
	cmd := exec.Command(command, arg...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Start()
	if err != nil {
		return err

	}
	return cmd.Wait()
}

// buildOutputName returns the correct instio-server file name
// based on the host Operating System
func buildOutputName() string {
	if runtime.GOOS == "windows" {
		return "instio-server.exe"
	}

	return "instio-server"
}

// resolve GOPATH. In 1.8 can be default, or custom. A custom GOPATH can
// also contain multiple paths, in which case 'go get' uses the first
func getGOPATH() (string, error) {
	var gopath string
	gopath = os.Getenv("GOPATH")
	if gopath == "" {
		// not set, find the default
		usr, err := user.Current()
		if err != nil {
			return gopath, err
		}
		gopath = filepath.Join(usr.HomeDir, "go")
	} else {
		// parse out in case of multiple, retain first
		if runtime.GOOS == "windows" {
			gopaths := strings.Split(gopath, ";")
			gopath = gopaths[0]
		} else {
			gopaths := strings.Split(gopath, ":")
			gopath = gopaths[0]
		}
	}
	return gopath, nil
}

func copyFileNoRoot(src, dst string) error {
	noRoot := strings.Split(src, string(filepath.Separator))[1:]
	path := filepath.Join(noRoot...)
	dstFile, err := os.Create(filepath.Join(dst, path))
	defer dstFile.Close()
	if err != nil {
		return err
	}

	srcFile, err := os.Open(src)
	defer srcFile.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func copyFilesWarnConflicts(srcDir, dstDir string, conflicts []string) error {
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, conflict := range conflicts {
			if info.Name() == conflict {
				fmt.Println("Instio couldn't fully build your project:")
				fmt.Println("You must rename the following file, as it conflicts with Instio core:")
				fmt.Println(path)
				fmt.Println("")
				fmt.Println("Once the files above have been renamed, run '$ instio build' to retry.")
				return ErrInternalConflict
			}
		}

		if info.IsDir() {
			// don't copy root directory
			if path == srcDir {
				return nil
			}

			if len(path) > len(srcDir) {
				path = path[len(srcDir)+1:]
			}
			dir := filepath.Join(dstDir, path)
			err := os.MkdirAll(dir, os.ModeDir|os.ModePerm)
			if err != nil {
				return err
			}

			return nil
		}

		err = copyFileNoRoot(path, dstDir)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func emptyDir(path string) error {
	d, err := os.Open(path)
	if err != nil {
		return err
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(path, name))
		if err != nil {
			return err
		}
	}

	return nil
}

func copyAll(src, dst string) error {
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		sep := string(filepath.Separator)

		// base == the instio project dir + string(filepath.Separator)
		parts := strings.Split(src, sep)
		base := strings.Join(parts[:len(parts)-1], sep)
		base += sep

		target := filepath.Join(dst, path[len(base):])

		// if its a directory, make dir in dst
		if info.IsDir() {
			err := os.MkdirAll(target, os.ModeDir|os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			// if its a file, move file to dir of dst
			err = os.Rename(path, target)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func corePackages(path string) error {
	/*vendorPath := filepath.Join(path, "cmd", "ponzu", "vendor", "github.com", "ponzu-cms", "ponzu")
	err := os.MkdirAll(vendorPath, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	dirs := []string{"content", "management", "system"}
	for _, dir := range dirs {
		err = os.Rename(filepath.Join(path, dir), filepath.Join(vendorPath, dir))
		if err != nil {
			return err
		}
	}*/

	// create a user content directory at project root
	contentPath := filepath.Join(path, "content")
	err := os.Mkdir(contentPath, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func getAnswer() (string, error) {
	var answer string
	_, err := fmt.Scanf("%s\n", &answer)
	if err != nil {
		if err.Error() == "unexpected newline" {
			answer = ""
		} else {
			return "", err
		}
	}

	answer = strings.ToLower(answer)

	return answer, nil
}
