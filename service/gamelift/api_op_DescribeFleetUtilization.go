// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type DescribeFleetUtilizationInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for a fleet(s) to retrieve utilization data for. You
	// can use either the fleet ID or ARN value.
	FleetIds []string `min:"1" type:"list"`

	// The maximum number of results to return. Use this parameter with NextToken
	// to get results as a set of sequential pages. This parameter is ignored when
	// the request specifies one or a list of fleet IDs.
	Limit *int64 `min:"1" type:"integer"`

	// Token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this action. To start
	// at the beginning of the result set, do not specify a value. This parameter
	// is ignored when the request specifies one or a list of fleet IDs.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeFleetUtilizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFleetUtilizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFleetUtilizationInput"}
	if s.FleetIds != nil && len(s.FleetIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetIds", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type DescribeFleetUtilizationOutput struct {
	_ struct{} `type:"structure"`

	// A collection of objects containing utilization information for each requested
	// fleet ID.
	FleetUtilization []FleetUtilization `type:"list"`

	// Token that indicates where to resume retrieving results on the next call
	// to this action. If no token is returned, these results represent the end
	// of the list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeFleetUtilizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeFleetUtilization = "DescribeFleetUtilization"

// DescribeFleetUtilizationRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves utilization statistics for one or more fleets. You can request
// utilization data for all fleets, or specify a list of one or more fleet IDs.
// When requesting multiple fleets, use the pagination parameters to retrieve
// results as a set of sequential pages. If successful, a FleetUtilization object
// is returned for each requested fleet ID. When specifying a list of fleet
// IDs, utilization objects are returned only for fleets that currently exist.
//
// Some API actions may limit the number of fleet IDs allowed in one request.
// If a request exceeds this limit, the request fails and the error message
// includes the maximum allowed.
//
// Learn more
//
//  Working with Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html).
//
// Related operations
//
//    * CreateFleet
//
//    * ListFleets
//
//    * DeleteFleet
//
//    * Describe fleets: DescribeFleetAttributes DescribeFleetCapacity DescribeFleetPortSettings
//    DescribeFleetUtilization DescribeRuntimeConfiguration DescribeEC2InstanceLimits
//    DescribeFleetEvents
//
//    * UpdateFleetAttributes
//
//    * Manage fleet actions: StartFleetActions StopFleetActions
//
//    // Example sending a request using DescribeFleetUtilizationRequest.
//    req := client.DescribeFleetUtilizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeFleetUtilization
func (c *Client) DescribeFleetUtilizationRequest(input *DescribeFleetUtilizationInput) DescribeFleetUtilizationRequest {
	op := &aws.Operation{
		Name:       opDescribeFleetUtilization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeFleetUtilizationInput{}
	}

	req := c.newRequest(op, input, &DescribeFleetUtilizationOutput{})
	return DescribeFleetUtilizationRequest{Request: req, Input: input, Copy: c.DescribeFleetUtilizationRequest}
}

// DescribeFleetUtilizationRequest is the request type for the
// DescribeFleetUtilization API operation.
type DescribeFleetUtilizationRequest struct {
	*aws.Request
	Input *DescribeFleetUtilizationInput
	Copy  func(*DescribeFleetUtilizationInput) DescribeFleetUtilizationRequest
}

// Send marshals and sends the DescribeFleetUtilization API request.
func (r DescribeFleetUtilizationRequest) Send(ctx context.Context) (*DescribeFleetUtilizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFleetUtilizationResponse{
		DescribeFleetUtilizationOutput: r.Request.Data.(*DescribeFleetUtilizationOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeFleetUtilizationResponse is the response type for the
// DescribeFleetUtilization API operation.
type DescribeFleetUtilizationResponse struct {
	*DescribeFleetUtilizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFleetUtilization request.
func (r *DescribeFleetUtilizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
