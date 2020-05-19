package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
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

	// Call with non-nill struct that has zero-length InstanceId
	fmt.Printf("Calling with zero-length InstanceId\n")
	makeTheCall(svc, params)

	// Call with non-nill struct that has single space InstanceId
	params = &autoscaling.TerminateInstanceInAutoScalingGroupInput{
		InstanceId:                     aws.String(" "),
		ShouldDecrementDesiredCapacity: aws.Bool(true),
	}
	fmt.Printf("\nCalling with zero-length InstanceId\n")
	makeTheCall(svc, params)

	// Call with nil struct
	params = nil
	fmt.Printf("\nCalling with nil input struct\n")
	makeTheCall(svc, params)
}

func makeTheCall(svc *autoscaling.AutoScaling, input *autoscaling.TerminateInstanceInAutoScalingGroupInput) {
	if input == nil {
		fmt.Println("input is nil")
	} else {
		fmt.Printf("input is NOT nil: %s\n", input.GoString())
	}

	// golang API docs: https://docs.aws.amazon.com/sdk-for-go/api/service/autoscaling/#AutoScaling.TerminateInstanceInAutoScalingGroup
	result, err := svc.TerminateInstanceInAutoScalingGroup(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case autoscaling.ErrCodeScalingActivityInProgressFault:
				fmt.Println(autoscaling.ErrCodeScalingActivityInProgressFault, aerr.Error())
			case autoscaling.ErrCodeResourceContentionFault:
				fmt.Println(autoscaling.ErrCodeResourceContentionFault, aerr.Error())
			default:
				fmt.Printf("Other AWS Error: %s\n", aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Printf("Non-AWS Error: %s\n", err.Error())
		}
		return
	}

	fmt.Printf("Result: %s\n", result)
}
