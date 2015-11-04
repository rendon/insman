package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

func toTable(instances []Instance) {
	var data = make([][]string, 0)
	for _, item := range instances {
		var record = []string{item.ID, item.Tag, item.PublicIP, item.PrivateIP}
		data = append(data, record)
	}

	var table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Tag", "Public IP", "Private IP"})
	for _, record := range data {
		table.Append(record)
	}
	table.Render()
}

func toYaml(instances []Instance) {
	fmt.Printf("---\n")
	for _, item := range instances {
		fmt.Printf("%s:\n", item.ID)
		fmt.Printf("    %s\n", item.Tag)
		fmt.Printf("    %s\n", item.PublicIP)
		fmt.Printf("    %s\n", item.PrivateIP)
	}
}

func list(c *cli.Context) {
	var format = "table"
	if c.String("format") != "" {
		format = strings.ToLower(c.String("format"))
	}

	var instances []Instance
	var err error
	if instances, err = GetRunningInstances(); err != nil {
		log.Fatalf("Error retrieving instances: %s", err)
	}

	switch format {
	case "table":
		toTable(instances)
	case "yaml":
		toYaml(instances)
	default:
		log.Fatalf("Format %q is not supported.", format)
	}
}
