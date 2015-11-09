package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// GetRunningInstances returns the list of running instances.
func GetRunningInstances() ([]Instance, error) {
	var svc = ec2.New(session.New(), nil)

	request := ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("running")},
			},
		},
	}
	resp, err := svc.DescribeInstances(&request)
	if err != nil {
		return nil, err
	}

	var instances = make([]Instance, 0)
	for _, res := range resp.Reservations {
		for _, ins := range res.Instances {
			var tag = ""
			if len(ins.Tags) > 0 {
				tag = *ins.Tags[0].Value
			}
			var i = Instance{
				ID:        *ins.InstanceId,
				PublicIP:  *ins.PublicIpAddress,
				PrivateIP: *ins.PrivateIpAddress,
				Tag:       tag,
			}
			instances = append(instances, i)
		}
	}
	return instances, nil
}

// TerminateInstances terminates instances whose ID are contained in `ids`.
func TerminateInstances(ids []string) (*ec2.TerminateInstancesOutput, error) {
	var svc = ec2.New(session.New(), nil)
	var instanceIds = make([]*string, 0)
	for _, id := range ids {
		instanceIds = append(instanceIds, aws.String(id))
	}
	var req = ec2.TerminateInstancesInput{InstanceIds: instanceIds}
	var resp, err = svc.TerminateInstances(&req)
	if err != nil {
		return nil, err
	}
	return resp, err
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

// SendCommand sends command to host via SSH.
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
