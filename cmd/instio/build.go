package main

import (
	//"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

func buildInstioServer(c *cli.Context) error {
	/*
		  // Copy all ./content files to internal vendor directory.
			src := "content"
			dst := filepath.Join("cmd", "ponzu", "vendor", "github.com", "ponzu-cms", "ponzu", "content")
			err := emptyDir(dst)
			if err != nil {
				return err
			}
			err = copyFilesWarnConflicts(src, dst, []string{"doc.go"})
			if err != nil {
				return err
			}

			// copy all ./addons files & dirs to internal vendor directory
			src = "addons"
			dst = filepath.Join("cmd", "ponzu", "vendor")
			err = copyFilesWarnConflicts(src, dst, nil)
			if err != nil {
				return err
			}
	*/

	// execute go build -o instio-cms cmd/instio/*.go
	cmdPackageName := strings.Join([]string{".", "cmd", "instio"}, "/")
	buildOptions := []string{"build", "-o", buildOutputName(), cmdPackageName}
	return execAndWait(c.String("gocmd"), buildOptions...)
}
