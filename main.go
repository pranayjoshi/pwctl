package main

import (
	"fmt"
	"os"

	"pwctl/pkgs/status"

	"pwctl/pkgs/auth"
	"pwctl/pkgs/list"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "pwctl",
		Usage: "A CLI for pgwatch3",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "c",
				Usage:    "Server address",
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "status",
				Usage: "Check the server status",
				Action: func(c *cli.Context) error {
					serverAddress := c.String("c")
					if status.IsServerLive(serverAddress) {
						fmt.Println("pwctl is running")
					} else {
						fmt.Println("pwctl is not running")
					}
					return nil
				},
			},
			{
				Name:  "login",
				Usage: "Login to pwatch3.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "user",
						Aliases:  []string{"u"},
						Usage:    "Username",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "password",
						Aliases:  []string{"p"},
						Usage:    "Password",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					return auth.Login(c)
				},
			},
			{
				Name:  "list",
				Usage: "List commands for pgwatch3",

				Subcommands: []*cli.Command{
					{
						Name:  "sources",
						Usage: "List monitored sources",
						Action: func(c *cli.Context) error {
							return list.ListSources(c)
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
