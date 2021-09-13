package main

import (
	"log"
	"os"
	"webcyou-org/coppeno/lib/copy"
	"webcyou-org/coppeno/lib/list"
	"webcyou-org/coppeno/lib/load"
	"webcyou-org/coppeno/lib/save"
	"webcyou-org/coppeno/lib/update"

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
			Name:    "copy",
			Aliases: []string{"c"},
			Usage:   "copy",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "copy, c"},
				cli.StringFlag{Name: "path, p"},
			},
			Action: func(c *cli.Context) error {
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
			Action: func(c *cli.Context) error {
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
			Action: func(c *cli.Context) error {
				if c.Args().First() != "" {
					err := load.File(c.Args().First())
					if err != nil {
						return err
					}
					return nil
				}
				return nil
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "update",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "update, u"},
			},
			Action: func(c *cli.Context) error {
				err := update.Start(c.Args().First(), c.Args().Get(1))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "list",
			Usage: "list",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "list"},
			},
			Action: func(c *cli.Context) error {
				err := list.Start()
				if err != nil {
					return err
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
