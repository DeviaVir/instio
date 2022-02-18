// Package main is located in the cmd/instio directory and contains the code to build
// and operate the command line interface (CLI) to manage Instio systems.
package main

import (
	"time"
  "fmt"
  "os"
  "log"

  "github.com/urfave/cli/v2"
)

var year = fmt.Sprintf("%d", time.Now().Year())

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
  
  app.Usage = "see github.com/DeviaVir/instio for supported commands"
  app.Action = func(c *cli.Context) error {
    fmt.Println(`Instio is an open-source HTTP server framework and CMS,
released under the Apache 2.0 license.
(c) 2022 - ` + year + ` Chase Sillevis`)
    return nil
  }

  app.Flags = []cli.Flag {
    &cli.StringFlag{
      Name: "bind",
      Value: "localhost",
      Usage: "the hostname to listen on",
    },
    &cli.IntFlag{
      Name: "httpsport",
      Value: 443,
      Usage: "HTTPS port for instio",
    },
    &cli.IntFlag{
      Name: "port",
      Value: 8080,
      Usage: "HTTP port for instio",
    },
    &cli.BoolFlag{
      Name: "https",
      Value: false,
      Usage: "enable automatic TLS/SSL certificate management",
    },
    &cli.BoolFlag{
      Name: "devhttps",
      Value: false,
      Usage: "[dev environment] enable automatic TLS/SSL certificate management",
    },
    &cli.IntFlag{
      Name: "docsport",
      Value: 1234,
      Usage: "[dev environment] override the documentation server port",
    },
    &cli.BoolFlag{
      Name: "docs",
      Value: false,
      Usage: "[dev environment] run HTTP server to view local HTML documentation",
    },
    &cli.StringFlag{
      Name: "gocmd",
      Value: "go",
      Usage: "custom go command if using beta or new release of Go",
    },
  }

	app.Commands = []*cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "run [flags] <service(,service)>",
			Action: func(c *cli.Context) error {
        fmt.Println(c.String("bind"))
				fmt.Println("service: ", c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
