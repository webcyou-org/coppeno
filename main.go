package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
	"webcyou-org/coppeno/lib/load"

	"webcyou-org/coppeno/lib/copy"
	"webcyou-org/coppeno/lib/save"
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
			Name:    "copy",
			Aliases: []string{"c"},
			Usage:   "copy",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "copy, c"},
				cli.StringFlag{Name: "path, p"},
			},
			Action:  func(c *cli.Context) error {
				err := copy.Start(c.Args().First())
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "save",
			Aliases: []string{"s"},
			Usage:   "save",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "save, s"},
			},
			Action:  func(c *cli.Context) error {
				err := save.Start(c.Args().First(), c.Args().Get(1))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "load",
			Aliases: []string{"l"},
			Usage:   "load",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "load, l"},
			},
			Action:  func(c *cli.Context) error {
				err := load.Start(c.Args().First(), c.Args().Get(1))
				if err != nil {
					return err
				}
				return nil
			},
		},
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