// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeFileSystemPolicyInput struct {
	_ struct{} `type:"structure"`

	// Specifies which EFS file system to retrieve the FileSystemPolicy for.
	//
	// FileSystemId is a required field
	FileSystemId *string `location:"uri" locationName:"FileSystemId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeFileSystemPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFileSystemPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFileSystemPolicyInput"}

	if s.FileSystemId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileSystemId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeFileSystemPolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FileSystemId != nil {
		v := *s.FileSystemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FileSystemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeFileSystemPolicyOutput struct {
	_ struct{} `type:"structure"`

	// Specifies the EFS file system to which the FileSystemPolicy applies.
	FileSystemId *string `type:"string"`

	// The JSON formatted FileSystemPolicy for the EFS file system.
	Policy *string `type:"string"`
}

// String returns the string representation
func (s DescribeFileSystemPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeFileSystemPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FileSystemId != nil {
		v := *s.FileSystemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FileSystemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Policy != nil {
		v := *s.Policy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Policy", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeFileSystemPolicy = "DescribeFileSystemPolicy"

// DescribeFileSystemPolicyRequest returns a request value for making API operation for
// Amazon Elastic File System.
//
// Returns the FileSystemPolicy for the specified EFS file system.
//
// This operation requires permissions for the elasticfilesystem:DescribeFileSystemPolicy
// action.
//
//    // Example sending a request using DescribeFileSystemPolicyRequest.
//    req := client.DescribeFileSystemPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/DescribeFileSystemPolicy
func (c *Client) DescribeFileSystemPolicyRequest(input *DescribeFileSystemPolicyInput) DescribeFileSystemPolicyRequest {
	op := &aws.Operation{
		Name:       opDescribeFileSystemPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-02-01/file-systems/{FileSystemId}/policy",
	}

	if input == nil {
		input = &DescribeFileSystemPolicyInput{}
	}

	req := c.newRequest(op, input, &DescribeFileSystemPolicyOutput{})
	return DescribeFileSystemPolicyRequest{Request: req, Input: input, Copy: c.DescribeFileSystemPolicyRequest}
}

// DescribeFileSystemPolicyRequest is the request type for the
// DescribeFileSystemPolicy API operation.
type DescribeFileSystemPolicyRequest struct {
	*aws.Request
	Input *DescribeFileSystemPolicyInput
	Copy  func(*DescribeFileSystemPolicyInput) DescribeFileSystemPolicyRequest
}

// Send marshals and sends the DescribeFileSystemPolicy API request.
func (r DescribeFileSystemPolicyRequest) Send(ctx context.Context) (*DescribeFileSystemPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFileSystemPolicyResponse{
		DescribeFileSystemPolicyOutput: r.Request.Data.(*DescribeFileSystemPolicyOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeFileSystemPolicyResponse is the response type for the
// DescribeFileSystemPolicy API operation.
type DescribeFileSystemPolicyResponse struct {
	*DescribeFileSystemPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFileSystemPolicy request.
func (r *DescribeFileSystemPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
