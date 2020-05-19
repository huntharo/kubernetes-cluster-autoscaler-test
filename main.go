package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

type autoScaling interface {
	TerminateInstanceInAutoScalingGroup(input *autoscaling.TerminateInstanceInAutoScalingGroupInput) (*autoscaling.TerminateInstanceInAutoScalingGroupOutput, error)
}

func main() {
	// Create a new autoscaling session
	svc := autoscaling.New(session.New())

	params := &autoscaling.TerminateInstanceInAutoScalingGroupInput{
		InstanceId:                     aws.String(""),
		ShouldDecrementDesiredCapacity: aws.Bool(true),
	}

	if params == nil {
		fmt.Println("params is nil")
	} else {
		fmt.Printf("params is NOT nil\n")
	}

	// golang API docs: https://docs.aws.amazon.com/sdk-for-go/api/service/autoscaling/#AutoScaling.TerminateInstanceInAutoScalingGroup
	resp, err := svc.TerminateInstanceInAutoScalingGroup(params)
	if err != nil {
		fmt.Printf("Error returned: %s\n", err)
	}

	fmt.Printf("Response as string: %s\n", resp.String())
}
