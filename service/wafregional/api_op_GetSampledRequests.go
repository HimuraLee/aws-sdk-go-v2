// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetSampledRequestsInput struct {
	_ struct{} `type:"structure"`

	// The number of requests that you want AWS WAF to return from among the first
	// 5,000 requests that your AWS resource received during the time range. If
	// your resource received fewer requests than the value of MaxItems, GetSampledRequests
	// returns information about all of them.
	//
	// MaxItems is a required field
	MaxItems *int64 `min:"1" type:"long" required:"true"`

	// RuleId is one of three values:
	//
	//    * The RuleId of the Rule or the RuleGroupId of the RuleGroup for which
	//    you want GetSampledRequests to return a sample of requests.
	//
	//    * Default_Action, which causes GetSampledRequests to return a sample of
	//    the requests that didn't match any of the rules in the specified WebACL.
	//
	// RuleId is a required field
	RuleId *string `min:"1" type:"string" required:"true"`

	// The start date and time and the end date and time of the range for which
	// you want GetSampledRequests to return a sample of requests. Specify the date
	// and time in the following format: "2016-09-27T14:50Z". You can specify any
	// time range in the previous three hours.
	//
	// TimeWindow is a required field
	TimeWindow *TimeWindow `type:"structure" required:"true"`

	// The WebACLId of the WebACL for which you want GetSampledRequests to return
	// a sample of requests.
	//
	// WebAclId is a required field
	WebAclId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSampledRequestsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSampledRequestsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSampledRequestsInput"}

	if s.MaxItems == nil {
		invalidParams.Add(aws.NewErrParamRequired("MaxItems"))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if s.RuleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleId"))
	}
	if s.RuleId != nil && len(*s.RuleId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleId", 1))
	}

	if s.TimeWindow == nil {
		invalidParams.Add(aws.NewErrParamRequired("TimeWindow"))
	}

	if s.WebAclId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebAclId"))
	}
	if s.WebAclId != nil && len(*s.WebAclId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WebAclId", 1))
	}
	if s.TimeWindow != nil {
		if err := s.TimeWindow.Validate(); err != nil {
			invalidParams.AddNested("TimeWindow", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetSampledRequestsOutput struct {
	_ struct{} `type:"structure"`

	// The total number of requests from which GetSampledRequests got a sample of
	// MaxItems requests. If PopulationSize is less than MaxItems, the sample includes
	// every request that your AWS resource received during the specified time range.
	PopulationSize *int64 `type:"long"`

	// A complex type that contains detailed information about each of the requests
	// in the sample.
	SampledRequests []SampledHTTPRequest `type:"list"`

	// Usually, TimeWindow is the time range that you specified in the GetSampledRequests
	// request. However, if your AWS resource received more than 5,000 requests
	// during the time range that you specified in the request, GetSampledRequests
	// returns the time range for the first 5,000 requests.
	TimeWindow *TimeWindow `type:"structure"`
}

// String returns the string representation
func (s GetSampledRequestsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSampledRequests = "GetSampledRequests"

// GetSampledRequestsRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Gets detailed information about a specified number of requests--a sample--that
// AWS WAF randomly selects from among the first 5,000 requests that your AWS
// resource received during a time range that you choose. You can specify a
// sample size of up to 500 requests, and you can specify any time range in
// the previous three hours.
//
// GetSampledRequests returns a time range, which is usually the time range
// that you specified. However, if your resource (such as a CloudFront distribution)
// received 5,000 requests before the specified time range elapsed, GetSampledRequests
// returns an updated time range. This new time range indicates the actual period
// during which AWS WAF selected the requests in the sample.
//
//    // Example sending a request using GetSampledRequestsRequest.
//    req := client.GetSampledRequestsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetSampledRequests
func (c *Client) GetSampledRequestsRequest(input *GetSampledRequestsInput) GetSampledRequestsRequest {
	op := &aws.Operation{
		Name:       opGetSampledRequests,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSampledRequestsInput{}
	}

	req := c.newRequest(op, input, &GetSampledRequestsOutput{})
	return GetSampledRequestsRequest{Request: req, Input: input, Copy: c.GetSampledRequestsRequest}
}

// GetSampledRequestsRequest is the request type for the
// GetSampledRequests API operation.
type GetSampledRequestsRequest struct {
	*aws.Request
	Input *GetSampledRequestsInput
	Copy  func(*GetSampledRequestsInput) GetSampledRequestsRequest
}

// Send marshals and sends the GetSampledRequests API request.
func (r GetSampledRequestsRequest) Send(ctx context.Context) (*GetSampledRequestsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSampledRequestsResponse{
		GetSampledRequestsOutput: r.Request.Data.(*GetSampledRequestsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSampledRequestsResponse is the response type for the
// GetSampledRequests API operation.
type GetSampledRequestsResponse struct {
	*GetSampledRequestsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSampledRequests request.
func (r *GetSampledRequestsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
