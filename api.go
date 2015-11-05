package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"

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

func getPublicKeys(file string) (ssh.AuthMethod, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(key), nil
}

func SendCommand(user, host, cmd string, auth ssh.AuthMethod) (string, error) {
	var config = &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{auth},
	}
	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		return "", fmt.Errorf("Failed to connect to server: %s", err)
	}

	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("Failed to create session: %s", err)
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(cmd); err != nil {
		return "", fmt.Errorf("Failed to run: %s", err)
	}
	return b.String(), nil
}
