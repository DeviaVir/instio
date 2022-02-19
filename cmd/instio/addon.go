package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

// Use `go get` to download addon and add to $GOPATH/src, useful
// for IDE auto-import and code completion, then copy entire directory
// tree to project's ./addons folder.
func getAddon(c *cli.Context) error {
	for _, p := range c.Args().Slice() {
		for _, path := range strings.Split(p, ",") {
			var cmdOptions []string

			// Go get
			cmdOptions = append(cmdOptions, "get", path)
			err := execAndWait(c.String("gocmd"), cmdOptions...)
			if err != nil {
				return err
			}

			// copy to ./addons folder
			// resolve GOPATH
			gopath, err := getGOPATH()
			if err != nil {
				return err
			}

			pwd, err := os.Getwd()
			if err != nil {
				return err
			}

			src := filepath.Join(gopath, "src", path)

			// Need to strip the addon name for copyAll?
			last := filepath.Base(path)
			dest := filepath.Join(pwd, "addons", strings.Replace(path, last, "", 1))

			err = replicateAll(src, dest)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// this is distinct from copyAll() in that files are copied, not moved,
// since we also need them to remain in $GOPATH/src
// thanks to @markc of stack overflow for the copyFile and copyFileContents functions
func replicateAll(src, dst string) error {
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
			// if its a file, copy file to dir of dst
			err = copyFile(path, target)
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

// copyFile copies a file from src to dst. if src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
// thanks to Stack Overflow
func copyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
// Thanks for Stack Overflow
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
