package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func createInstance() {
	// Set your region for future requests.
	defaults.DefaultConfig.Region = aws.String("us-west-2")

	svc := ec2.New(nil)

	var minCount int64 = 1
	var maxCount int64 = 1
	params := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-989b7bab"),
		InstanceType: aws.String("t2.micro"),
		MinCount:     &minCount,
		MaxCount:     &maxCount,
	}

	runResult, err := svc.RunInstances(params)
	if err != nil {
		log.Println("Could not create instance", err)
		return
	}

	log.Println("Created instance", *runResult.Instances[0].InstanceId)

	// Add tags to the instance
	_, err = svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{runResult.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			&ec2.Tag{
				Key:   aws.String("Name"),
				Value: aws.String("MyInstanceName"),
			},
		},
	})
	if err != nil {
		log.Println("Could not create tags for instance", *runResult.Instances[0].InstanceId, err)
	}

	log.Println("Successfully tagged instance")
}

func listRunningInstance() {
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
			var keyName = *instance.KeyName
			fmt.Printf("%s\t%s\t%s\t%s\n", id, publicIp, privateIp, keyName)
		}
	}
}

func main() {
	listRunningInstance()
}
