// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeAssetModelInput struct {
	_ struct{} `type:"structure"`

	// The ID of the asset model.
	//
	// AssetModelId is a required field
	AssetModelId *string `location:"uri" locationName:"assetModelId" min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeAssetModelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAssetModelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAssetModelInput"}

	if s.AssetModelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssetModelId"))
	}
	if s.AssetModelId != nil && len(*s.AssetModelId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("AssetModelId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeAssetModelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AssetModelId != nil {
		v := *s.AssetModelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "assetModelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeAssetModelOutput struct {
	_ struct{} `type:"structure"`

	// The ARN (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the asset model, which has the following format.
	//
	// arn:${Partition}:iotsitewise:${Region}:${Account}:asset-model/${AssetModelId}
	//
	// AssetModelArn is a required field
	AssetModelArn *string `locationName:"assetModelArn" min:"1" type:"string" required:"true"`

	// The date the asset model was created, in Unix epoch time.
	//
	// AssetModelCreationDate is a required field
	AssetModelCreationDate *time.Time `locationName:"assetModelCreationDate" type:"timestamp" required:"true"`

	// The asset model's description.
	//
	// AssetModelDescription is a required field
	AssetModelDescription *string `locationName:"assetModelDescription" min:"1" type:"string" required:"true"`

	// A list of asset model hierarchies that each contain a childAssetModelId and
	// a hierarchyId (named id). A hierarchy specifies allowed parent/child asset
	// relationships for an asset model.
	//
	// AssetModelHierarchies is a required field
	AssetModelHierarchies []AssetModelHierarchy `locationName:"assetModelHierarchies" type:"list" required:"true"`

	// The ID of the asset model.
	//
	// AssetModelId is a required field
	AssetModelId *string `locationName:"assetModelId" min:"36" type:"string" required:"true"`

	// The date the asset model was last updated, in Unix epoch time.
	//
	// AssetModelLastUpdateDate is a required field
	AssetModelLastUpdateDate *time.Time `locationName:"assetModelLastUpdateDate" type:"timestamp" required:"true"`

	// The name of the asset model.
	//
	// AssetModelName is a required field
	AssetModelName *string `locationName:"assetModelName" min:"1" type:"string" required:"true"`

	// The list of asset properties for the asset model.
	//
	// AssetModelProperties is a required field
	AssetModelProperties []AssetModelProperty `locationName:"assetModelProperties" type:"list" required:"true"`

	// The current status of the asset model, which contains a state and any error
	// message.
	//
	// AssetModelStatus is a required field
	AssetModelStatus *AssetModelStatus `locationName:"assetModelStatus" type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeAssetModelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeAssetModelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AssetModelArn != nil {
		v := *s.AssetModelArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "assetModelArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AssetModelCreationDate != nil {
		v := *s.AssetModelCreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "assetModelCreationDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.AssetModelDescription != nil {
		v := *s.AssetModelDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "assetModelDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AssetModelHierarchies != nil {
		v := s.AssetModelHierarchies

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "assetModelHierarchies", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.AssetModelId != nil {
		v := *s.AssetModelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "assetModelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AssetModelLastUpdateDate != nil {
		v := *s.AssetModelLastUpdateDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "assetModelLastUpdateDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.AssetModelName != nil {
		v := *s.AssetModelName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "assetModelName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AssetModelProperties != nil {
		v := s.AssetModelProperties

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "assetModelProperties", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.AssetModelStatus != nil {
		v := s.AssetModelStatus

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "assetModelStatus", v, metadata)
	}
	return nil
}

const opDescribeAssetModel = "DescribeAssetModel"

// DescribeAssetModelRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Retrieves information about an asset model.
//
//    // Example sending a request using DescribeAssetModelRequest.
//    req := client.DescribeAssetModelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/DescribeAssetModel
func (c *Client) DescribeAssetModelRequest(input *DescribeAssetModelInput) DescribeAssetModelRequest {
	op := &aws.Operation{
		Name:       opDescribeAssetModel,
		HTTPMethod: "GET",
		HTTPPath:   "/asset-models/{assetModelId}",
	}

	if input == nil {
		input = &DescribeAssetModelInput{}
	}

	req := c.newRequest(op, input, &DescribeAssetModelOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("model.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return DescribeAssetModelRequest{Request: req, Input: input, Copy: c.DescribeAssetModelRequest}
}

// DescribeAssetModelRequest is the request type for the
// DescribeAssetModel API operation.
type DescribeAssetModelRequest struct {
	*aws.Request
	Input *DescribeAssetModelInput
	Copy  func(*DescribeAssetModelInput) DescribeAssetModelRequest
}

// Send marshals and sends the DescribeAssetModel API request.
func (r DescribeAssetModelRequest) Send(ctx context.Context) (*DescribeAssetModelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAssetModelResponse{
		DescribeAssetModelOutput: r.Request.Data.(*DescribeAssetModelOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAssetModelResponse is the response type for the
// DescribeAssetModel API operation.
type DescribeAssetModelResponse struct {
	*DescribeAssetModelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAssetModel request.
func (r *DescribeAssetModelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
