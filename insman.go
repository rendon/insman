// AWS Instance manager
package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	log.SetFlags(0)
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
		cli.Command{
			Name:      "exec",
			Usage:     "Execute command on instances",
			ArgsUsage: "<command> [<host1> <host2> ...]",
			Action:    exec,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all",
					Usage: "Execute command on all instances",
				},
			},
		},
		cli.Command{
			Name:      "terminate",
			Usage:     "Terminates one or more instances",
			ArgsUsage: "<InstanceID>[ <InstanceID>]*",
			Action:    terminate,
		},
	}
	app.Run(os.Args)
}
