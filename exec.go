package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/goutil/arrays"
)

func exec(c *cli.Context) {
	if len(c.Args()) < 1 {
		cli.ShowCommandHelp(c, "exec")
		return
	}

	var cmd = c.Args()[0]
	var instances []Instance
	var err error
	if instances, err = GetRunningInstances(); err != nil {
		log.Fatalf("Error retrieving instances: %s", err)
	}

	auth, err := getPublicKeys(os.Getenv("AWS_KEY_FILE"))
	if err != nil {
		log.Fatalf("Error getting keys: %s", err)
	}

	if c.Bool("all") {
		for _, item := range instances {
			fmt.Printf("Executing %q on %s (%s)\n", cmd, item.ID, item.Tag)
			output, err := SendCommand("ubuntu", item.PublicIP, cmd, auth)
			if err != nil {
				log.Printf("Failed to execute command: %s", err)
			}
			fmt.Printf("Output: %s\n", output)
		}
	} else {
		var a = c.Args()
		for _, item := range instances {
			var id = item.ID
			var tag = item.Tag
			if !arrays.ContainsString(a, id) && !arrays.ContainsString(a, tag) {
				continue
			}
			fmt.Printf("Executing %q on %s (%s)\n", cmd, item.ID, item.Tag)
			output, err := SendCommand("ubuntu", item.PublicIP, cmd, auth)
			if err != nil {
				log.Printf("Failed to execute command: %s", err)
			}
			fmt.Printf("Output: %s\n", output)
		}
	}
}
