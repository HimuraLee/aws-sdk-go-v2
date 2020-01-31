// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateVPCEConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the VPC endpoint configuration you want
	// to update.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`

	// The DNS (domain) name used to connect to your private service in your VPC.
	// The DNS name must not already be in use on the internet.
	ServiceDnsName *string `locationName:"serviceDnsName" type:"string"`

	// An optional description that provides details about your VPC endpoint configuration.
	VpceConfigurationDescription *string `locationName:"vpceConfigurationDescription" type:"string"`

	// The friendly name you give to your VPC endpoint configuration to manage your
	// configurations more easily.
	VpceConfigurationName *string `locationName:"vpceConfigurationName" type:"string"`

	// The name of the VPC endpoint service running in your AWS account that you
	// want Device Farm to test.
	VpceServiceName *string `locationName:"vpceServiceName" type:"string"`
}

// String returns the string representation
func (s UpdateVPCEConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateVPCEConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateVPCEConfigurationInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateVPCEConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains information about your VPC endpoint configuration.
	VpceConfiguration *VPCEConfiguration `locationName:"vpceConfiguration" type:"structure"`
}

// String returns the string representation
func (s UpdateVPCEConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateVPCEConfiguration = "UpdateVPCEConfiguration"

// UpdateVPCEConfigurationRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Updates information about an Amazon Virtual Private Cloud (VPC) endpoint
// configuration.
//
//    // Example sending a request using UpdateVPCEConfigurationRequest.
//    req := client.UpdateVPCEConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/UpdateVPCEConfiguration
func (c *Client) UpdateVPCEConfigurationRequest(input *UpdateVPCEConfigurationInput) UpdateVPCEConfigurationRequest {
	op := &aws.Operation{
		Name:       opUpdateVPCEConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateVPCEConfigurationInput{}
	}

	req := c.newRequest(op, input, &UpdateVPCEConfigurationOutput{})
	return UpdateVPCEConfigurationRequest{Request: req, Input: input, Copy: c.UpdateVPCEConfigurationRequest}
}

// UpdateVPCEConfigurationRequest is the request type for the
// UpdateVPCEConfiguration API operation.
type UpdateVPCEConfigurationRequest struct {
	*aws.Request
	Input *UpdateVPCEConfigurationInput
	Copy  func(*UpdateVPCEConfigurationInput) UpdateVPCEConfigurationRequest
}

// Send marshals and sends the UpdateVPCEConfiguration API request.
func (r UpdateVPCEConfigurationRequest) Send(ctx context.Context) (*UpdateVPCEConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateVPCEConfigurationResponse{
		UpdateVPCEConfigurationOutput: r.Request.Data.(*UpdateVPCEConfigurationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateVPCEConfigurationResponse is the response type for the
// UpdateVPCEConfiguration API operation.
type UpdateVPCEConfigurationResponse struct {
	*UpdateVPCEConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateVPCEConfiguration request.
func (r *UpdateVPCEConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
