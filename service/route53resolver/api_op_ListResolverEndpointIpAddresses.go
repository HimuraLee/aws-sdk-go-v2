// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListResolverEndpointIpAddressesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of IP addresses that you want to return in the response
	// to a ListResolverEndpointIpAddresses request. If you don't specify a value
	// for MaxResults, Resolver returns up to 100 IP addresses.
	MaxResults *int64 `min:"1" type:"integer"`

	// For the first ListResolverEndpointIpAddresses request, omit this value.
	//
	// If the specified resolver endpoint has more than MaxResults IP addresses,
	// you can submit another ListResolverEndpointIpAddresses request to get the
	// next group of IP addresses. In the next request, specify the value of NextToken
	// from the previous response.
	NextToken *string `type:"string"`

	// The ID of the resolver endpoint that you want to get IP addresses for.
	//
	// ResolverEndpointId is a required field
	ResolverEndpointId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListResolverEndpointIpAddressesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListResolverEndpointIpAddressesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListResolverEndpointIpAddressesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ResolverEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResolverEndpointId"))
	}
	if s.ResolverEndpointId != nil && len(*s.ResolverEndpointId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResolverEndpointId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListResolverEndpointIpAddressesOutput struct {
	_ struct{} `type:"structure"`

	// The IP addresses that DNS queries pass through on their way to your network
	// (outbound endpoint) or on the way to Resolver (inbound endpoint).
	IpAddresses []IpAddressResponse `type:"list"`

	// The value that you specified for MaxResults in the request.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the specified endpoint has more than MaxResults IP addresses, you can
	// submit another ListResolverEndpointIpAddresses request to get the next group
	// of IP addresses. In the next request, specify the value of NextToken from
	// the previous response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListResolverEndpointIpAddressesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListResolverEndpointIpAddresses = "ListResolverEndpointIpAddresses"

// ListResolverEndpointIpAddressesRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// Gets the IP addresses for a specified resolver endpoint.
//
//    // Example sending a request using ListResolverEndpointIpAddressesRequest.
//    req := client.ListResolverEndpointIpAddressesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/ListResolverEndpointIpAddresses
func (c *Client) ListResolverEndpointIpAddressesRequest(input *ListResolverEndpointIpAddressesInput) ListResolverEndpointIpAddressesRequest {
	op := &aws.Operation{
		Name:       opListResolverEndpointIpAddresses,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListResolverEndpointIpAddressesInput{}
	}

	req := c.newRequest(op, input, &ListResolverEndpointIpAddressesOutput{})

	return ListResolverEndpointIpAddressesRequest{Request: req, Input: input, Copy: c.ListResolverEndpointIpAddressesRequest}
}

// ListResolverEndpointIpAddressesRequest is the request type for the
// ListResolverEndpointIpAddresses API operation.
type ListResolverEndpointIpAddressesRequest struct {
	*aws.Request
	Input *ListResolverEndpointIpAddressesInput
	Copy  func(*ListResolverEndpointIpAddressesInput) ListResolverEndpointIpAddressesRequest
}

// Send marshals and sends the ListResolverEndpointIpAddresses API request.
func (r ListResolverEndpointIpAddressesRequest) Send(ctx context.Context) (*ListResolverEndpointIpAddressesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResolverEndpointIpAddressesResponse{
		ListResolverEndpointIpAddressesOutput: r.Request.Data.(*ListResolverEndpointIpAddressesOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListResolverEndpointIpAddressesRequestPaginator returns a paginator for ListResolverEndpointIpAddresses.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListResolverEndpointIpAddressesRequest(input)
//   p := route53resolver.NewListResolverEndpointIpAddressesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListResolverEndpointIpAddressesPaginator(req ListResolverEndpointIpAddressesRequest) ListResolverEndpointIpAddressesPaginator {
	return ListResolverEndpointIpAddressesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListResolverEndpointIpAddressesInput
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

// ListResolverEndpointIpAddressesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListResolverEndpointIpAddressesPaginator struct {
	aws.Pager
}

func (p *ListResolverEndpointIpAddressesPaginator) CurrentPage() *ListResolverEndpointIpAddressesOutput {
	return p.Pager.CurrentPage().(*ListResolverEndpointIpAddressesOutput)
}

// ListResolverEndpointIpAddressesResponse is the response type for the
// ListResolverEndpointIpAddresses API operation.
type ListResolverEndpointIpAddressesResponse struct {
	*ListResolverEndpointIpAddressesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResolverEndpointIpAddresses request.
func (r *ListResolverEndpointIpAddressesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
