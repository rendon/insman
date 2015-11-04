package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetRunningInstances() ([]Instance, error) {
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
		return nil, fmt.Errorf("Error getting description: %s", err)
	}

	var instances = make([]Instance, 0)
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			var tag = ""
			if len(instance.Tags) > 0 {
				tag = *instance.Tags[0].Value
			}
			var ins = Instance{
				ID:        *instance.InstanceId,
				PublicIP:  *instance.PublicIpAddress,
				PrivateIP: *instance.PrivateIpAddress,
				Tag:       tag,
			}
			instances = append(instances, ins)
		}
	}
	return instances, nil
}
