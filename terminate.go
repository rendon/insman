package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

func terminate(c *cli.Context) {
	if len(c.Args()) < 1 {
		cli.ShowCommandHelp(c, "terminate")
		return
	}

	fmt.Printf("The following instances will be terminated:\n")
	for _, ins := range c.Args() {
		fmt.Printf("%s\n", ins)
	}
	fmt.Printf("Are you sure?[y/n]")
	var op rune
	if _, err := fmt.Scanf("%c", &op); err != nil || (op != 'y' && op != 'Y') {
		return
	}

	fmt.Printf("Terminating...\n")
	if _, err := TerminateInstances(c.Args()); err != nil {
		log.Fatalf("Failed to terminate instances: %s", err)
	}
}
