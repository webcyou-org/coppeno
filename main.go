package main

import (
	"errors"
	"log"
	"os"
	"webcyou-org/coppeno/lib/copy"
	"webcyou-org/coppeno/lib/fetch"
	"webcyou-org/coppeno/lib/list"
	"webcyou-org/coppeno/lib/load"
	"webcyou-org/coppeno/lib/save"
	"webcyou-org/coppeno/lib/update"

	"github.com/manifoldco/promptui"
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
				validate := func(input string) error {
					if len(input) > 0 {
						return nil
					}
					return errors.New("Invalid Required input")
				}

				prompt := promptui.Prompt{
					Label:    "filename",
					Validate: validate,
				}
				filename, _ := prompt.Run()

				prompt = promptui.Prompt{
					Label:    "URL",
					Validate: validate,
				}
				url, _ := prompt.Run()

				err := save.Start(filename, url)
				// err := save.Start(c.Args().First(), c.Args().Get(1))
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
			Name:    "fetch",
			Aliases: []string{"f"},
			Usage:   "fetch",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "fetch, f"},
			},
			Action: func(c *cli.Context) error {
				err := fetch.Start(c.Args().First(), c.Args().Get(1))
				if err != nil {
					return err
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
