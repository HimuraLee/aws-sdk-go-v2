// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetReservationUtilizationInput struct {
	_ struct{} `type:"structure"`

	// Filters utilization data by dimensions. You can filter by the following dimensions:
	//
	//    * AZ
	//
	//    * CACHE_ENGINE
	//
	//    * DEPLOYMENT_OPTION
	//
	//    * INSTANCE_TYPE
	//
	//    * LINKED_ACCOUNT
	//
	//    * OPERATING_SYSTEM
	//
	//    * PLATFORM
	//
	//    * REGION
	//
	//    * SERVICE
	//
	//    * SCOPE
	//
	//    * TENANCY
	//
	// GetReservationUtilization uses the same Expression (http://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// object as the other operations, but only AND is supported among each dimension,
	// and nesting is supported up to only one level deep. If there are multiple
	// values for a dimension, they are OR'd together.
	Filter *Expression `type:"structure"`

	// If GroupBy is set, Granularity can't be set. If Granularity isn't set, the
	// response object doesn't include Granularity, either MONTHLY or DAILY. If
	// both GroupBy and Granularity aren't set, GetReservationUtilization defaults
	// to DAILY.
	//
	// The GetReservationUtilization operation supports only DAILY and MONTHLY granularities.
	Granularity Granularity `type:"string" enum:"true"`

	// Groups only by SUBSCRIPTION_ID. Metadata is included.
	GroupBy []GroupDefinition `type:"list"`

	// The token to retrieve the next set of results. AWS provides the token when
	// the response from a previous call has more results than the maximum page
	// size.
	NextPageToken *string `type:"string"`

	// Sets the start and end dates for retrieving RI utilization. The start date
	// is inclusive, but the end date is exclusive. For example, if start is 2017-01-01
	// and end is 2017-05-01, then the cost and usage data is retrieved from 2017-01-01
	// up to and including 2017-04-30 but not including 2017-05-01.
	//
	// TimePeriod is a required field
	TimePeriod *DateInterval `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetReservationUtilizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetReservationUtilizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetReservationUtilizationInput"}

	if s.TimePeriod == nil {
		invalidParams.Add(aws.NewErrParamRequired("TimePeriod"))
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(aws.ErrInvalidParams))
		}
	}
	if s.TimePeriod != nil {
		if err := s.TimePeriod.Validate(); err != nil {
			invalidParams.AddNested("TimePeriod", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetReservationUtilizationOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next set of retrievable results. AWS provides the token
	// when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string `type:"string"`

	// The total amount of time that you used your RIs.
	Total *ReservationAggregates `type:"structure"`

	// The amount of time that you used your RIs.
	//
	// UtilizationsByTime is a required field
	UtilizationsByTime []UtilizationByTime `type:"list" required:"true"`
}

// String returns the string representation
func (s GetReservationUtilizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetReservationUtilization = "GetReservationUtilization"

// GetReservationUtilizationRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Retrieves the reservation utilization for your account. Master accounts in
// an organization have access to member accounts. You can filter data by dimensions
// in a time period. You can use GetDimensionValues to determine the possible
// dimension values. Currently, you can group only by SUBSCRIPTION_ID.
//
//    // Example sending a request using GetReservationUtilizationRequest.
//    req := client.GetReservationUtilizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetReservationUtilization
func (c *Client) GetReservationUtilizationRequest(input *GetReservationUtilizationInput) GetReservationUtilizationRequest {
	op := &aws.Operation{
		Name:       opGetReservationUtilization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetReservationUtilizationInput{}
	}

	req := c.newRequest(op, input, &GetReservationUtilizationOutput{})
	return GetReservationUtilizationRequest{Request: req, Input: input, Copy: c.GetReservationUtilizationRequest}
}

// GetReservationUtilizationRequest is the request type for the
// GetReservationUtilization API operation.
type GetReservationUtilizationRequest struct {
	*aws.Request
	Input *GetReservationUtilizationInput
	Copy  func(*GetReservationUtilizationInput) GetReservationUtilizationRequest
}

// Send marshals and sends the GetReservationUtilization API request.
func (r GetReservationUtilizationRequest) Send(ctx context.Context) (*GetReservationUtilizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetReservationUtilizationResponse{
		GetReservationUtilizationOutput: r.Request.Data.(*GetReservationUtilizationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetReservationUtilizationResponse is the response type for the
// GetReservationUtilization API operation.
type GetReservationUtilizationResponse struct {
	*GetReservationUtilizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetReservationUtilization request.
func (r *GetReservationUtilizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
