/**
 * Retrieve list of running instances.
 *
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
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

func toTable(result *ec2.DescribeInstancesOutput) {
	var data = make([][]string, 0)
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			var id = *instance.InstanceId
			var publicIp = *instance.PublicIpAddress
			var privateIp = *instance.PrivateIpAddress
			var tag = instance.Tags[0].Value
			var record = []string{id, *tag, publicIp, privateIp}
			data = append(data, record)
		}
	}

	var table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Tag", "Public IP", "Private IP"})
	for _, record := range data {
		table.Append(record)
	}
	table.Render()
}

func toYaml(result *ec2.DescribeInstancesOutput) {
	fmt.Printf("---\n")
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			var id = *instance.InstanceId
			var publicIp = *instance.PublicIpAddress
			var privateIp = *instance.PrivateIpAddress
			var tag = instance.Tags[0].Value
			fmt.Printf("%s:\n", id)
			fmt.Printf("    %s\n", *tag)
			fmt.Printf("    %s\n", publicIp)
			fmt.Printf("    %s\n", privateIp)
		}
	}
}

func list(c *cli.Context) {
	defaults.DefaultConfig.Region = aws.String("us-west-2")
	svc := ec2.New(nil)

	var filters = []*ec2.Filter{
		&ec2.Filter{
			Name:   aws.String("instance-state-name"),
			Values: []*string{aws.String("running")},
		},
	}

	request := ec2.DescribeInstancesInput{Filters: filters}
	result, err := svc.DescribeInstances(&request)
	if err != nil {
		log.Fatalf("Error getting description: %s", err)
	}

	var format = "table"
	if c.String("format") != "" {
		format = strings.ToLower(c.String("format"))
	}

	switch format {
	case "table":
		toTable(result)
	case "yaml":
		toYaml(result)
	default:
		log.Fatalf("Format %q is not supported.", format)
	}
}
