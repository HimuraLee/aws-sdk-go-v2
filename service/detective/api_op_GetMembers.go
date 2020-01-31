// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package detective

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetMembersInput struct {
	_ struct{} `type:"structure"`

	// The list of AWS account identifiers for the member account for which to return
	// member details.
	//
	// You cannot use GetMembers to retrieve information about member accounts that
	// were removed from the behavior graph.
	//
	// AccountIds is a required field
	AccountIds []string `min:"1" type:"list" required:"true"`

	// The ARN of the behavior graph for which to request the member details.
	//
	// GraphArn is a required field
	GraphArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMembersInput"}

	if s.AccountIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountIds"))
	}
	if s.AccountIds != nil && len(s.AccountIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountIds", 1))
	}

	if s.GraphArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("GraphArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMembersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountIds != nil {
		v := s.AccountIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "AccountIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.GraphArn != nil {
		v := *s.GraphArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GraphArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetMembersOutput struct {
	_ struct{} `type:"structure"`

	// The member account details that Detective is returning in response to the
	// request.
	MemberDetails []MemberDetail `type:"list"`

	// The requested member accounts for which Detective was unable to return member
	// details.
	//
	// For each account, provides the reason why the request could not be processed.
	UnprocessedAccounts []UnprocessedAccount `type:"list"`
}

// String returns the string representation
func (s GetMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.MemberDetails != nil {
		v := s.MemberDetails

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "MemberDetails", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.UnprocessedAccounts != nil {
		v := s.UnprocessedAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UnprocessedAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetMembers = "GetMembers"

// GetMembersRequest returns a request value for making API operation for
// Amazon Detective.
//
// Amazon Detective is currently in preview.
//
// Returns the membership details for specified member accounts for a behavior
// graph.
//
//    // Example sending a request using GetMembersRequest.
//    req := client.GetMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/detective-2018-10-26/GetMembers
func (c *Client) GetMembersRequest(input *GetMembersInput) GetMembersRequest {
	op := &aws.Operation{
		Name:       opGetMembers,
		HTTPMethod: "POST",
		HTTPPath:   "/graph/members/get",
	}

	if input == nil {
		input = &GetMembersInput{}
	}

	req := c.newRequest(op, input, &GetMembersOutput{})
	return GetMembersRequest{Request: req, Input: input, Copy: c.GetMembersRequest}
}

// GetMembersRequest is the request type for the
// GetMembers API operation.
type GetMembersRequest struct {
	*aws.Request
	Input *GetMembersInput
	Copy  func(*GetMembersInput) GetMembersRequest
}

// Send marshals and sends the GetMembers API request.
func (r GetMembersRequest) Send(ctx context.Context) (*GetMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMembersResponse{
		GetMembersOutput: r.Request.Data.(*GetMembersOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMembersResponse is the response type for the
// GetMembers API operation.
type GetMembersResponse struct {
	*GetMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMembers request.
func (r *GetMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
