// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type TerminateInstancesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The IDs of the instances.
	//
	// Constraints: Up to 1000 instance IDs. We recommend breaking up this request
	// into smaller batches.
	//
	// InstanceIds is a required field
	InstanceIds []string `locationName:"InstanceId" locationNameList:"InstanceId" type:"list" required:"true"`
}

// String returns the string representation
func (s TerminateInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TerminateInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TerminateInstancesInput"}

	if s.InstanceIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TerminateInstancesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the terminated instances.
	TerminatingInstances []InstanceStateChange `locationName:"instancesSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s TerminateInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opTerminateInstances = "TerminateInstances"

// TerminateInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Shuts down the specified instances. This operation is idempotent; if you
// terminate an instance more than once, each call succeeds.
//
// If you specify multiple instances and the request fails (for example, because
// of a single incorrect instance ID), none of the instances are terminated.
//
// Terminated instances remain visible after termination (for approximately
// one hour).
//
// By default, Amazon EC2 deletes all EBS volumes that were attached when the
// instance launched. Volumes attached after instance launch continue running.
//
// You can stop, start, and terminate EBS-backed instances. You can only terminate
// instance store-backed instances. What happens to an instance differs if you
// stop it or terminate it. For example, when you stop an instance, the root
// device and any other devices attached to the instance persist. When you terminate
// an instance, any attached EBS volumes with the DeleteOnTermination block
// device mapping parameter set to true are automatically deleted. For more
// information about the differences between stopping and terminating instances,
// see Instance Lifecycle (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// For more information about troubleshooting, see Troubleshooting Terminating
// Your Instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstancesShuttingDown.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using TerminateInstancesRequest.
//    req := client.TerminateInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/TerminateInstances
func (c *Client) TerminateInstancesRequest(input *TerminateInstancesInput) TerminateInstancesRequest {
	op := &aws.Operation{
		Name:       opTerminateInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TerminateInstancesInput{}
	}

	req := c.newRequest(op, input, &TerminateInstancesOutput{})

	return TerminateInstancesRequest{Request: req, Input: input, Copy: c.TerminateInstancesRequest}
}

// TerminateInstancesRequest is the request type for the
// TerminateInstances API operation.
type TerminateInstancesRequest struct {
	*aws.Request
	Input *TerminateInstancesInput
	Copy  func(*TerminateInstancesInput) TerminateInstancesRequest
}

// Send marshals and sends the TerminateInstances API request.
func (r TerminateInstancesRequest) Send(ctx context.Context) (*TerminateInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TerminateInstancesResponse{
		TerminateInstancesOutput: r.Request.Data.(*TerminateInstancesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TerminateInstancesResponse is the response type for the
// TerminateInstances API operation.
type TerminateInstancesResponse struct {
	*TerminateInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TerminateInstances request.
func (r *TerminateInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
