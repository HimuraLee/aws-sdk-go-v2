// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Input for DeleteEndpoint action.
type DeleteEndpointInput struct {
	_ struct{} `type:"structure"`

	// EndpointArn of endpoint to delete.
	//
	// EndpointArn is a required field
	EndpointArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteEndpointInput"}

	if s.EndpointArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteEndpointOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteEndpoint = "DeleteEndpoint"

// DeleteEndpointRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Deletes the endpoint for a device and mobile app from Amazon SNS. This action
// is idempotent. For more information, see Using Amazon SNS Mobile Push Notifications
// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
//
// When you delete an endpoint that is also subscribed to a topic, then you
// must also unsubscribe the endpoint from the topic.
//
//    // Example sending a request using DeleteEndpointRequest.
//    req := client.DeleteEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/DeleteEndpoint
func (c *Client) DeleteEndpointRequest(input *DeleteEndpointInput) DeleteEndpointRequest {
	op := &aws.Operation{
		Name:       opDeleteEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteEndpointInput{}
	}

	req := c.newRequest(op, input, &DeleteEndpointOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteEndpointRequest{Request: req, Input: input, Copy: c.DeleteEndpointRequest}
}

// DeleteEndpointRequest is the request type for the
// DeleteEndpoint API operation.
type DeleteEndpointRequest struct {
	*aws.Request
	Input *DeleteEndpointInput
	Copy  func(*DeleteEndpointInput) DeleteEndpointRequest
}

// Send marshals and sends the DeleteEndpoint API request.
func (r DeleteEndpointRequest) Send(ctx context.Context) (*DeleteEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteEndpointResponse{
		DeleteEndpointOutput: r.Request.Data.(*DeleteEndpointOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteEndpointResponse is the response type for the
// DeleteEndpoint API operation.
type DeleteEndpointResponse struct {
	*DeleteEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteEndpoint request.
func (r *DeleteEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
