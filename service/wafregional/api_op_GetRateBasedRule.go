// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetRateBasedRuleInput struct {
	_ struct{} `type:"structure"`

	// The RuleId of the RateBasedRule that you want to get. RuleId is returned
	// by CreateRateBasedRule and by ListRateBasedRules.
	//
	// RuleId is a required field
	RuleId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRateBasedRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRateBasedRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRateBasedRuleInput"}

	if s.RuleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleId"))
	}
	if s.RuleId != nil && len(*s.RuleId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetRateBasedRuleOutput struct {
	_ struct{} `type:"structure"`

	// Information about the RateBasedRule that you specified in the GetRateBasedRule
	// request.
	Rule *RateBasedRule `type:"structure"`
}

// String returns the string representation
func (s GetRateBasedRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRateBasedRule = "GetRateBasedRule"

// GetRateBasedRuleRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns the RateBasedRule that is specified by the RuleId that you included
// in the GetRateBasedRule request.
//
//    // Example sending a request using GetRateBasedRuleRequest.
//    req := client.GetRateBasedRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetRateBasedRule
func (c *Client) GetRateBasedRuleRequest(input *GetRateBasedRuleInput) GetRateBasedRuleRequest {
	op := &aws.Operation{
		Name:       opGetRateBasedRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRateBasedRuleInput{}
	}

	req := c.newRequest(op, input, &GetRateBasedRuleOutput{})
	return GetRateBasedRuleRequest{Request: req, Input: input, Copy: c.GetRateBasedRuleRequest}
}

// GetRateBasedRuleRequest is the request type for the
// GetRateBasedRule API operation.
type GetRateBasedRuleRequest struct {
	*aws.Request
	Input *GetRateBasedRuleInput
	Copy  func(*GetRateBasedRuleInput) GetRateBasedRuleRequest
}

// Send marshals and sends the GetRateBasedRule API request.
func (r GetRateBasedRuleRequest) Send(ctx context.Context) (*GetRateBasedRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRateBasedRuleResponse{
		GetRateBasedRuleOutput: r.Request.Data.(*GetRateBasedRuleOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRateBasedRuleResponse is the response type for the
// GetRateBasedRule API operation.
type GetRateBasedRuleResponse struct {
	*GetRateBasedRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRateBasedRule request.
func (r *GetRateBasedRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
