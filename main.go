package main

import (
	"errors"
	"log"
	"os"

	"github.com/webcyou-org/coppeno/lib/fetch"
	initialization "github.com/webcyou-org/coppeno/lib/init"
	"github.com/webcyou-org/coppeno/lib/list"
	"github.com/webcyou-org/coppeno/lib/save"

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
	app.Usage = "Quick project kickstarter Simple File Manager CLI tool."
	app.Version = "0.5.5"

	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "Initialize coppeno.json.",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "init"},
			},
			Action: func(c *cli.Context) error {
				err := initialization.Start()
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "save",
			Aliases: []string{"s"},
			Usage:   "Save the file name and file path interactively.",
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
					Label: "group",
				}
				group, _ := prompt.Run()

				prompt = promptui.Prompt{
					Label:    "filename",
					Validate: validate,
				}
				filename, _ := prompt.Run()

				prompt = promptui.Prompt{
					Label:    "URL",
					Validate: validate,
				}
				url, _ := prompt.Run()

				err := save.Start(group, filename, url)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "load",
			Aliases: []string{"l"},
			Usage:   "Read the coppeno.json file that was created.",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "load, l"},
			},
			Action: func(c *cli.Context) error {
				if c.Args().First() != "" {
					err := save.File(c.Args().First())
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
			Usage:   "Download the file saved in coppeno.json",
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
			Name:  "list",
			Usage: "Check the list of files saved in coppeno.json.",
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
