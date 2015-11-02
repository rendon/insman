package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	var app = cli.NewApp()
	app.Name = "insman"
	app.Version = "0.1.0"
	app.Usage = "EC2 Instance manager"
	app.Commands = []cli.Command{
		cli.Command{
			Name:   "list",
			Usage:  "List running instances",
			Action: list,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "format",
					Usage: "List format: table or yaml",
				},
			},
		},
	}

	app.Run(os.Args)
}
