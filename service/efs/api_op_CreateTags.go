// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type CreateTagsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the file system whose tags you want to modify (String). This operation
	// modifies the tags only, not the file system.
	//
	// FileSystemId is a required field
	FileSystemId *string `location:"uri" locationName:"FileSystemId" type:"string" required:"true"`

	// An array of Tag objects to add. Each Tag object is a key-value pair.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s CreateTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTagsInput"}

	if s.FileSystemId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileSystemId"))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTagsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.FileSystemId != nil {
		v := *s.FileSystemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FileSystemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateTagsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTagsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCreateTags = "CreateTags"

// CreateTagsRequest returns a request value for making API operation for
// Amazon Elastic File System.
//
// Creates or overwrites tags associated with a file system. Each tag is a key-value
// pair. If a tag key specified in the request already exists on the file system,
// this operation overwrites its value with the value provided in the request.
// If you add the Name tag to your file system, Amazon EFS returns it in the
// response to the DescribeFileSystems operation.
//
// This operation requires permission for the elasticfilesystem:CreateTags action.
//
//    // Example sending a request using CreateTagsRequest.
//    req := client.CreateTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/CreateTags
func (c *Client) CreateTagsRequest(input *CreateTagsInput) CreateTagsRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, CreateTags, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opCreateTags,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-02-01/create-tags/{FileSystemId}",
	}

	if input == nil {
		input = &CreateTagsInput{}
	}

	req := c.newRequest(op, input, &CreateTagsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CreateTagsRequest{Request: req, Input: input, Copy: c.CreateTagsRequest}
}

// CreateTagsRequest is the request type for the
// CreateTags API operation.
type CreateTagsRequest struct {
	*aws.Request
	Input *CreateTagsInput
	Copy  func(*CreateTagsInput) CreateTagsRequest
}

// Send marshals and sends the CreateTags API request.
func (r CreateTagsRequest) Send(ctx context.Context) (*CreateTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTagsResponse{
		CreateTagsOutput: r.Request.Data.(*CreateTagsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTagsResponse is the response type for the
// CreateTags API operation.
type CreateTagsResponse struct {
	*CreateTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTags request.
func (r *CreateTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
