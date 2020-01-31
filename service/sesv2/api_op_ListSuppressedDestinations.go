// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to obtain a list of email destinations that are on the suppression
// list for your account.
type ListSuppressedDestinationsInput struct {
	_ struct{} `type:"structure"`

	// Used to filter the list of suppressed email destinations so that it only
	// includes addresses that were added to the list before a specific date. The
	// date that you specify should be in Unix time format.
	EndDate *time.Time `location:"querystring" locationName:"EndDate" type:"timestamp"`

	// A token returned from a previous call to ListSuppressedDestinations to indicate
	// the position in the list of suppressed email addresses.
	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`

	// The number of results to show in a single call to ListSuppressedDestinations.
	// If the number of results is larger than the number you specified in this
	// parameter, then the response includes a NextToken element, which you can
	// use to obtain additional results.
	PageSize *int64 `location:"querystring" locationName:"PageSize" type:"integer"`

	// The factors that caused the email address to be added to .
	Reasons []SuppressionListReason `location:"querystring" locationName:"Reason" type:"list"`

	// Used to filter the list of suppressed email destinations so that it only
	// includes addresses that were added to the list after a specific date. The
	// date that you specify should be in Unix time format.
	StartDate *time.Time `location:"querystring" locationName:"StartDate" type:"timestamp"`
}

// String returns the string representation
func (s ListSuppressedDestinationsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSuppressedDestinationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.EndDate != nil {
		v := *s.EndDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "EndDate",
			protocol.TimeValue{V: v, Format: protocol.ISO8601TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "PageSize", protocol.Int64Value(v), metadata)
	}
	if s.Reasons != nil {
		v := s.Reasons

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.QueryTarget, "Reason", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.StartDate != nil {
		v := *s.StartDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "StartDate",
			protocol.TimeValue{V: v, Format: protocol.ISO8601TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	return nil
}

// A list of suppressed email addresses.
type ListSuppressedDestinationsOutput struct {
	_ struct{} `type:"structure"`

	// A token that indicates that there are additional email addresses on the suppression
	// list for your account. To view additional suppressed addresses, issue another
	// request to ListSuppressedDestinations, and pass this token in the NextToken
	// parameter.
	NextToken *string `type:"string"`

	// A list of summaries, each containing a summary for a suppressed email destination.
	SuppressedDestinationSummaries []SuppressedDestinationSummary `type:"list"`
}

// String returns the string representation
func (s ListSuppressedDestinationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSuppressedDestinationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SuppressedDestinationSummaries != nil {
		v := s.SuppressedDestinationSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SuppressedDestinationSummaries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListSuppressedDestinations = "ListSuppressedDestinations"

// ListSuppressedDestinationsRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Retrieves a list of email addresses that are on the suppression list for
// your account.
//
//    // Example sending a request using ListSuppressedDestinationsRequest.
//    req := client.ListSuppressedDestinationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/ListSuppressedDestinations
func (c *Client) ListSuppressedDestinationsRequest(input *ListSuppressedDestinationsInput) ListSuppressedDestinationsRequest {
	op := &aws.Operation{
		Name:       opListSuppressedDestinations,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/email/suppression/addresses",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListSuppressedDestinationsInput{}
	}

	req := c.newRequest(op, input, &ListSuppressedDestinationsOutput{})
	return ListSuppressedDestinationsRequest{Request: req, Input: input, Copy: c.ListSuppressedDestinationsRequest}
}

// ListSuppressedDestinationsRequest is the request type for the
// ListSuppressedDestinations API operation.
type ListSuppressedDestinationsRequest struct {
	*aws.Request
	Input *ListSuppressedDestinationsInput
	Copy  func(*ListSuppressedDestinationsInput) ListSuppressedDestinationsRequest
}

// Send marshals and sends the ListSuppressedDestinations API request.
func (r ListSuppressedDestinationsRequest) Send(ctx context.Context) (*ListSuppressedDestinationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSuppressedDestinationsResponse{
		ListSuppressedDestinationsOutput: r.Request.Data.(*ListSuppressedDestinationsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSuppressedDestinationsRequestPaginator returns a paginator for ListSuppressedDestinations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSuppressedDestinationsRequest(input)
//   p := sesv2.NewListSuppressedDestinationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSuppressedDestinationsPaginator(req ListSuppressedDestinationsRequest) ListSuppressedDestinationsPaginator {
	return ListSuppressedDestinationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListSuppressedDestinationsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListSuppressedDestinationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSuppressedDestinationsPaginator struct {
	aws.Pager
}

func (p *ListSuppressedDestinationsPaginator) CurrentPage() *ListSuppressedDestinationsOutput {
	return p.Pager.CurrentPage().(*ListSuppressedDestinationsOutput)
}

// ListSuppressedDestinationsResponse is the response type for the
// ListSuppressedDestinations API operation.
type ListSuppressedDestinationsResponse struct {
	*ListSuppressedDestinationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSuppressedDestinations request.
func (r *ListSuppressedDestinationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
