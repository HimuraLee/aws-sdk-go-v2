// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTaskSetsInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the service that the task sets exist in.
	//
	// Cluster is a required field
	Cluster *string `locationName:"cluster" type:"string" required:"true"`

	// Specifies whether to see the resource tags for the task set. If TAGS is specified,
	// the tags are included in the response. If this field is omitted, tags are
	// not included in the response.
	Include []TaskSetField `locationName:"include" type:"list"`

	// The short name or full Amazon Resource Name (ARN) of the service that the
	// task sets exist in.
	//
	// Service is a required field
	Service *string `locationName:"service" type:"string" required:"true"`

	// The ID or full Amazon Resource Name (ARN) of task sets to describe.
	TaskSets []string `locationName:"taskSets" type:"list"`
}

// String returns the string representation
func (s DescribeTaskSetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTaskSetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTaskSetsInput"}

	if s.Cluster == nil {
		invalidParams.Add(aws.NewErrParamRequired("Cluster"))
	}

	if s.Service == nil {
		invalidParams.Add(aws.NewErrParamRequired("Service"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeTaskSetsOutput struct {
	_ struct{} `type:"structure"`

	// Any failures associated with the call.
	Failures []Failure `locationName:"failures" type:"list"`

	// The list of task sets described.
	TaskSets []TaskSet `locationName:"taskSets" type:"list"`
}

// String returns the string representation
func (s DescribeTaskSetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTaskSets = "DescribeTaskSets"

// DescribeTaskSetsRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Describes the task sets in the specified cluster and service. This is used
// when a service uses the EXTERNAL deployment controller type. For more information,
// see Amazon ECS Deployment Types (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
// in the Amazon Elastic Container Service Developer Guide.
//
//    // Example sending a request using DescribeTaskSetsRequest.
//    req := client.DescribeTaskSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DescribeTaskSets
func (c *Client) DescribeTaskSetsRequest(input *DescribeTaskSetsInput) DescribeTaskSetsRequest {
	op := &aws.Operation{
		Name:       opDescribeTaskSets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTaskSetsInput{}
	}

	req := c.newRequest(op, input, &DescribeTaskSetsOutput{})
	return DescribeTaskSetsRequest{Request: req, Input: input, Copy: c.DescribeTaskSetsRequest}
}

// DescribeTaskSetsRequest is the request type for the
// DescribeTaskSets API operation.
type DescribeTaskSetsRequest struct {
	*aws.Request
	Input *DescribeTaskSetsInput
	Copy  func(*DescribeTaskSetsInput) DescribeTaskSetsRequest
}

// Send marshals and sends the DescribeTaskSets API request.
func (r DescribeTaskSetsRequest) Send(ctx context.Context) (*DescribeTaskSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTaskSetsResponse{
		DescribeTaskSetsOutput: r.Request.Data.(*DescribeTaskSetsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTaskSetsResponse is the response type for the
// DescribeTaskSets API operation.
type DescribeTaskSetsResponse struct {
	*DescribeTaskSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTaskSets request.
func (r *DescribeTaskSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
