/**
 * Retrieve list of running instances.

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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
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

	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			var id = *instance.InstanceId
			var publicIp = *instance.PublicIpAddress
			var privateIp = *instance.PrivateIpAddress
			var tag = instance.Tags[0].Value
			fmt.Printf("%s:", id)
			fmt.Printf("\n\ttag: %s", *tag)
			fmt.Printf("\n\tpublic_id: %s", publicIp)
			fmt.Printf("\n\tprivate_ip: %s", privateIp)
			fmt.Println()
		}
	}
}
