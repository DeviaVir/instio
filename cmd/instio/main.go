// Package main is located in the cmd/instio directory and contains the code to build
// and operate the command line interface (CLI) to manage Instio systems.
package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

var (
	// ErrWrongOrMissingService informs a user that the services to run must be
	// explicitly specified when serve is called
	ErrWrongOrMissingService = errors.New("Instio currently support only 'admin' and 'api' services.")

	year = fmt.Sprintf("%d", time.Now().Year())
)

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

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "hostname",
			Value: "localhost",
			Usage: "the hostname to listen on (localhost by default)",
		},
		&cli.IntFlag{
			Name:  "https_port",
			Value: 443,
			Usage: "HTTPS port for instio (443 by default)",
		},
		&cli.IntFlag{
			Name:  "http_port",
			Value: 8080,
			Usage: "HTTP port for instio (8080 by default)",
		},
		&cli.BoolFlag{
			Name:  "https",
			Value: false,
			Usage: "enable automatic TLS/SSL certificate management (false by default)",
		},
		&cli.BoolFlag{
			Name:  "https_dev",
			Value: false,
			Usage: "[dev environment] enable automatic TLS/SSL certificate management (false by default)",
		},
		&cli.IntFlag{
			Name:  "docs_port",
			Value: 1234,
			Usage: "[dev environment] override the documentation server port (1234 by default)",
		},
		&cli.BoolFlag{
			Name:  "docs",
			Value: false,
			Usage: "[dev environment] run HTTP server to view local HTML documentation (false by default)",
		},
		&cli.StringFlag{
			Name:  "gocmd",
			Value: "go",
			Usage: "custom go command if using beta or new release of Go (go by default)",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "run [flags] <service(,service)>",
			Action: func(c *cli.Context) error {
				var services []string
				if c.Args().Len() > 0 {
					for _, s := range c.Args().Slice() {
						for _, ss := range strings.Split(s, ",") {
							services = append(services, ss)
						}
					}
				} else {
					services = []string{"admin", "api"}
				}

				// db.Init()
				// defer db.Close()

				// analytics.Init()
				// defer analytics.Close()

				for _, service := range services {
					switch service {
					case "api":
						fmt.Println("[DEBUG] would've started api")
						// api.Run()
						break
					case "admin":
						fmt.Println("[DEBUG] would've started admin")
						// admin.Run()
						break
					default:
						return ErrWrongOrMissingService
					}
				}

				if c.Bool("docs") {
					fmt.Println("[DEBUG] would've started docs")
					// admin.Docs(c.Int("docs_port"))
				}

				// go db.InitSearchIndex()

				// err := db.PutConfig("https_port", fmt.Sprintf("%d", c.Int("https_port")))
				// if err != nil {...}

				// cannot run production HTTPS and development HTTPS together
				if c.Bool("https_dev") {
					fmt.Println("[DEBUG] Enabling self-signed HTTPS...")

					// go tls.EnableDev(c.Int("https_port"))
					fmt.Println(fmt.Sprintf("Server listening on https://localhost:%d for requests... [DEV]", c.Int("https_port")))
					fmt.Println("----")
					fmt.Println("If your browser rejects HTTPS requests, try allowing insecure connections on localhost.")
					fmt.Println("on Chrome, visit chrome://flags/#allow-insecure-localhost")
				} else if c.Bool("https") {
					fmt.Println("[DEBUG] Enabling HTTPS...")

					// go tls.Enable(c.Int("https_port"))
					fmt.Printf("Server listening on :%d for HTTPS requests...\n", c.Int("https_port"))
				}

				// err = db.PutConfig("http_port", fmt.Sprintf("%d", c.Int("http_port")))
				// if err != nil {...}

				// err = db.PutConfig("bind_addr", c.String("hostname"))
				// if err != nil {...}

				fmt.Println(fmt.Sprintf("Server listening on %s:%d for HTTP requests...", c.String("hostname"), c.Int("http_port")))
				fmt.Println("Visit '/admin' to get started.")
				log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s:%d", c.String("hostname"), c.Int("http_port")), nil))
				return nil
			},
		},
		{
			Name:    "generate",
			Aliases: []string{"gen", "g"},
			Usage:   "generate <generator type (,...fields)>",
			Subcommands: []*cli.Command{
				{
					Name:    "content",
					Aliases: []string{"con", "c"},
					Usage:   "content <namespace> <field> <field>...",
					Action: func(c *cli.Context) error {
						generateContentType(c.Args().Slice())
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
