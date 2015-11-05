/*
 * Don't hard-code your credentials!
 * Export the following environment variables instead:
 *
 * export AWS_ACCESS_KEY_ID='AKID'
 * export AWS_SECRET_ACCESS_KEY='SECRET'
 *
 * This example loads credentials from ~/.aws/credentials:
 * [default]
 * aws_access_key_id = ...
 * aws_secret_access_key = ...
 */
package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/codegangsta/cli"
)

func init() {
	defaults.DefaultConfig.Region = aws.String("us-west-2")
	log.SetFlags(0)
}

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
		cli.Command{
			Name:   "exec",
			Usage:  "Execute command on instances",
			Action: exec,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all",
					Usage: "Execute command on all instances",
				},
			},
		},
	}

	app.Run(os.Args)
}
