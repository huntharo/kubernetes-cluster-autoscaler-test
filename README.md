# Overview

This app is testing a possible bug in `cluster-autoscaler` that seems to result in a null requestParameters being sent on an `TerminateInstanceInAutoScalingGroup` API call.

# Building

`go build`

# Running Locally

* Login to AWS
  * `aws-okta exec [role-name] -- bash -l`
* Set Region
  * `export AWS_REGION=us-east-1`
* Run the app
  * `go run main.go`
