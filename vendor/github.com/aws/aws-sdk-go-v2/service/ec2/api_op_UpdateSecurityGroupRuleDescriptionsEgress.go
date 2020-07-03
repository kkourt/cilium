// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateSecurityGroupRuleDescriptionsEgressInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the security group. You must specify either the security group
	// ID or the security group name in the request. For security groups in a nondefault
	// VPC, you must specify the security group ID.
	GroupId *string `type:"string"`

	// [Default VPC] The name of the security group. You must specify either the
	// security group ID or the security group name in the request.
	GroupName *string `type:"string"`

	// The IP permissions for the security group rule.
	//
	// IpPermissions is a required field
	IpPermissions []IpPermission `locationNameList:"item" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateSecurityGroupRuleDescriptionsEgressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSecurityGroupRuleDescriptionsEgressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSecurityGroupRuleDescriptionsEgressInput"}

	if s.IpPermissions == nil {
		invalidParams.Add(aws.NewErrParamRequired("IpPermissions"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateSecurityGroupRuleDescriptionsEgressOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, returns an error.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s UpdateSecurityGroupRuleDescriptionsEgressOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateSecurityGroupRuleDescriptionsEgress = "UpdateSecurityGroupRuleDescriptionsEgress"

// UpdateSecurityGroupRuleDescriptionsEgressRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// [VPC only] Updates the description of an egress (outbound) security group
// rule. You can replace an existing description, or add a description to a
// rule that did not have one previously.
//
// You specify the description as part of the IP permissions structure. You
// can remove a description for a security group rule by omitting the description
// parameter in the request.
//
//    // Example sending a request using UpdateSecurityGroupRuleDescriptionsEgressRequest.
//    req := client.UpdateSecurityGroupRuleDescriptionsEgressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/UpdateSecurityGroupRuleDescriptionsEgress
func (c *Client) UpdateSecurityGroupRuleDescriptionsEgressRequest(input *UpdateSecurityGroupRuleDescriptionsEgressInput) UpdateSecurityGroupRuleDescriptionsEgressRequest {
	op := &aws.Operation{
		Name:       opUpdateSecurityGroupRuleDescriptionsEgress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateSecurityGroupRuleDescriptionsEgressInput{}
	}

	req := c.newRequest(op, input, &UpdateSecurityGroupRuleDescriptionsEgressOutput{})

	return UpdateSecurityGroupRuleDescriptionsEgressRequest{Request: req, Input: input, Copy: c.UpdateSecurityGroupRuleDescriptionsEgressRequest}
}

// UpdateSecurityGroupRuleDescriptionsEgressRequest is the request type for the
// UpdateSecurityGroupRuleDescriptionsEgress API operation.
type UpdateSecurityGroupRuleDescriptionsEgressRequest struct {
	*aws.Request
	Input *UpdateSecurityGroupRuleDescriptionsEgressInput
	Copy  func(*UpdateSecurityGroupRuleDescriptionsEgressInput) UpdateSecurityGroupRuleDescriptionsEgressRequest
}

// Send marshals and sends the UpdateSecurityGroupRuleDescriptionsEgress API request.
func (r UpdateSecurityGroupRuleDescriptionsEgressRequest) Send(ctx context.Context) (*UpdateSecurityGroupRuleDescriptionsEgressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSecurityGroupRuleDescriptionsEgressResponse{
		UpdateSecurityGroupRuleDescriptionsEgressOutput: r.Request.Data.(*UpdateSecurityGroupRuleDescriptionsEgressOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSecurityGroupRuleDescriptionsEgressResponse is the response type for the
// UpdateSecurityGroupRuleDescriptionsEgress API operation.
type UpdateSecurityGroupRuleDescriptionsEgressResponse struct {
	*UpdateSecurityGroupRuleDescriptionsEgressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSecurityGroupRuleDescriptionsEgress request.
func (r *UpdateSecurityGroupRuleDescriptionsEgressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
