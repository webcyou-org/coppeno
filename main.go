package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

func main() {
	app := cli.NewApp()
	app.Name = "coppeno"
	app.Usage = "Quick project kickstarter CLI tool."
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "Hello",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "path, p"},
			},
			Action:  func(c *cli.Context) error {
				log.Println("Hello, world!")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}