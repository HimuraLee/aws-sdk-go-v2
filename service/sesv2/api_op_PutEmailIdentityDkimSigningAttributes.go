// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to change the DKIM attributes for an email identity.
type PutEmailIdentityDkimSigningAttributesInput struct {
	_ struct{} `type:"structure"`

	// The email identity that you want to configure DKIM for.
	//
	// EmailIdentity is a required field
	EmailIdentity *string `location:"uri" locationName:"EmailIdentity" type:"string" required:"true"`

	// An object that contains information about the private key and selector that
	// you want to use to configure DKIM for the identity. This object is only required
	// if you want to configure Bring Your Own DKIM (BYODKIM) for the identity.
	SigningAttributes *DkimSigningAttributes `type:"structure"`

	// The method that you want to use to configure DKIM for the identity. There
	// are two possible values:
	//
	//    * AWS_SES – Configure DKIM for the identity by using Easy DKIM (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html).
	//
	//    * EXTERNAL – Configure DKIM for the identity by using Bring Your Own
	//    DKIM (BYODKIM).
	//
	// SigningAttributesOrigin is a required field
	SigningAttributesOrigin DkimSigningAttributesOrigin `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s PutEmailIdentityDkimSigningAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutEmailIdentityDkimSigningAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutEmailIdentityDkimSigningAttributesInput"}

	if s.EmailIdentity == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmailIdentity"))
	}
	if len(s.SigningAttributesOrigin) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("SigningAttributesOrigin"))
	}
	if s.SigningAttributes != nil {
		if err := s.SigningAttributes.Validate(); err != nil {
			invalidParams.AddNested("SigningAttributes", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutEmailIdentityDkimSigningAttributesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.SigningAttributes != nil {
		v := s.SigningAttributes

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SigningAttributes", v, metadata)
	}
	if len(s.SigningAttributesOrigin) > 0 {
		v := s.SigningAttributesOrigin

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SigningAttributesOrigin", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.EmailIdentity != nil {
		v := *s.EmailIdentity

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "EmailIdentity", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// If the action is successful, the service sends back an HTTP 200 response.
//
// The following data is returned in JSON format by the service.
type PutEmailIdentityDkimSigningAttributesOutput struct {
	_ struct{} `type:"structure"`

	// The DKIM authentication status of the identity. Amazon SES determines the
	// authentication status by searching for specific records in the DNS configuration
	// for your domain. If you used Easy DKIM (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html)
	// to set up DKIM authentication, Amazon SES tries to find three unique CNAME
	// records in the DNS configuration for your domain.
	//
	// If you provided a public key to perform DKIM authentication, Amazon SES tries
	// to find a TXT record that uses the selector that you specified. The value
	// of the TXT record must be a public key that's paired with the private key
	// that you specified in the process of creating the identity.
	//
	// The status can be one of the following:
	//
	//    * PENDING – The verification process was initiated, but Amazon SES hasn't
	//    yet detected the DKIM records in the DNS configuration for the domain.
	//
	//    * SUCCESS – The verification process completed successfully.
	//
	//    * FAILED – The verification process failed. This typically occurs when
	//    Amazon SES fails to find the DKIM records in the DNS configuration of
	//    the domain.
	//
	//    * TEMPORARY_FAILURE – A temporary issue is preventing Amazon SES from
	//    determining the DKIM authentication status of the domain.
	//
	//    * NOT_STARTED – The DKIM verification process hasn't been initiated
	//    for the domain.
	DkimStatus DkimStatus `type:"string" enum:"true"`

	// If you used Easy DKIM (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html)
	// to configure DKIM authentication for the domain, then this object contains
	// a set of unique strings that you use to create a set of CNAME records that
	// you add to the DNS configuration for your domain. When Amazon SES detects
	// these records in the DNS configuration for your domain, the DKIM authentication
	// process is complete.
	//
	// If you configured DKIM authentication for the domain by providing your own
	// public-private key pair, then this object contains the selector that's associated
	// with your public key.
	//
	// Regardless of the DKIM authentication method you use, Amazon SES searches
	// for the appropriate records in the DNS configuration of the domain for up
	// to 72 hours.
	DkimTokens []string `type:"list"`
}

// String returns the string representation
func (s PutEmailIdentityDkimSigningAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutEmailIdentityDkimSigningAttributesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.DkimStatus) > 0 {
		v := s.DkimStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DkimStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DkimTokens != nil {
		v := s.DkimTokens

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "DkimTokens", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opPutEmailIdentityDkimSigningAttributes = "PutEmailIdentityDkimSigningAttributes"

// PutEmailIdentityDkimSigningAttributesRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Used to configure or change the DKIM authentication settings for an email
// domain identity. You can use this operation to do any of the following:
//
//    * Update the signing attributes for an identity that uses Bring Your Own
//    DKIM (BYODKIM).
//
//    * Change from using no DKIM authentication to using Easy DKIM.
//
//    * Change from using no DKIM authentication to using BYODKIM.
//
//    * Change from using Easy DKIM to using BYODKIM.
//
//    * Change from using BYODKIM to using Easy DKIM.
//
//    // Example sending a request using PutEmailIdentityDkimSigningAttributesRequest.
//    req := client.PutEmailIdentityDkimSigningAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/PutEmailIdentityDkimSigningAttributes
func (c *Client) PutEmailIdentityDkimSigningAttributesRequest(input *PutEmailIdentityDkimSigningAttributesInput) PutEmailIdentityDkimSigningAttributesRequest {
	op := &aws.Operation{
		Name:       opPutEmailIdentityDkimSigningAttributes,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/email/identities/{EmailIdentity}/dkim/signing",
	}

	if input == nil {
		input = &PutEmailIdentityDkimSigningAttributesInput{}
	}

	req := c.newRequest(op, input, &PutEmailIdentityDkimSigningAttributesOutput{})
	return PutEmailIdentityDkimSigningAttributesRequest{Request: req, Input: input, Copy: c.PutEmailIdentityDkimSigningAttributesRequest}
}

// PutEmailIdentityDkimSigningAttributesRequest is the request type for the
// PutEmailIdentityDkimSigningAttributes API operation.
type PutEmailIdentityDkimSigningAttributesRequest struct {
	*aws.Request
	Input *PutEmailIdentityDkimSigningAttributesInput
	Copy  func(*PutEmailIdentityDkimSigningAttributesInput) PutEmailIdentityDkimSigningAttributesRequest
}

// Send marshals and sends the PutEmailIdentityDkimSigningAttributes API request.
func (r PutEmailIdentityDkimSigningAttributesRequest) Send(ctx context.Context) (*PutEmailIdentityDkimSigningAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutEmailIdentityDkimSigningAttributesResponse{
		PutEmailIdentityDkimSigningAttributesOutput: r.Request.Data.(*PutEmailIdentityDkimSigningAttributesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutEmailIdentityDkimSigningAttributesResponse is the response type for the
// PutEmailIdentityDkimSigningAttributes API operation.
type PutEmailIdentityDkimSigningAttributesResponse struct {
	*PutEmailIdentityDkimSigningAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutEmailIdentityDkimSigningAttributes request.
func (r *PutEmailIdentityDkimSigningAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
