// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateTemplateInput struct {
	_ struct{} `type:"structure"`

	// The ID for the AWS account that the group is in. Currently, you use the ID
	// for the AWS account that contains your Amazon QuickSight account.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// A display name for the template.
	Name *string `min:"1" type:"string"`

	// A list of resource permissions to be set on the template. The shorthand syntax
	// should look similar to this: Shorthand Syntax: Principal=string,Actions=string,string
	// ...
	Permissions []ResourcePermission `min:"1" type:"list"`

	// The ARN of the source entity from which this template is being created. Templates
	// can be currently created from an analysis or another template. If the ARN
	// is for an analysis, you must include its dataset references.
	//
	// SourceEntity is a required field
	SourceEntity *TemplateSourceEntity `type:"structure" required:"true"`

	// Contains a map of the key-value pairs for the resource tag or tags assigned
	// to the resource.
	Tags []Tag `min:"1" type:"list"`

	// An ID for the template you want to create. This is unique per AWS region
	// per AWS account.
	//
	// TemplateId is a required field
	TemplateId *string `location:"uri" locationName:"TemplateId" min:"1" type:"string" required:"true"`

	// A description of the current template version being created. This API created
	// the first version of the template. Every time UpdateTemplate is called a
	// new version is created. Each version of the template maintains a description
	// of the version in the VersionDescription field.
	VersionDescription *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTemplateInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Permissions != nil && len(s.Permissions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Permissions", 1))
	}

	if s.SourceEntity == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceEntity"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if s.TemplateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateId"))
	}
	if s.TemplateId != nil && len(*s.TemplateId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateId", 1))
	}
	if s.VersionDescription != nil && len(*s.VersionDescription) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionDescription", 1))
	}
	if s.Permissions != nil {
		for i, v := range s.Permissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Permissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SourceEntity != nil {
		if err := s.SourceEntity.Validate(); err != nil {
			invalidParams.AddNested("SourceEntity", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Permissions != nil {
		v := s.Permissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Permissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.SourceEntity != nil {
		v := s.SourceEntity

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SourceEntity", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.VersionDescription != nil {
		v := *s.VersionDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VersionDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "TemplateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateTemplateOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for the template.
	Arn *string `type:"string"`

	// The template creation status.
	CreationStatus ResourceStatus `type:"string" enum:"true"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// The ID of the template.
	TemplateId *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) for the template, including the version information
	// of the first version.
	VersionArn *string `type:"string"`
}

// String returns the string representation
func (s CreateTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.CreationStatus) > 0 {
		v := s.CreationStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TemplateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionArn != nil {
		v := *s.VersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opCreateTemplate = "CreateTemplate"

// CreateTemplateRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Creates a template from an existing QuickSight analysis or template. The
// resulting template can be used to create a dashboard.
//
// A template is an entity in QuickSight which encapsulates the metadata required
// to create an analysis that can be used to create dashboard. It adds a layer
// of abstraction by use placeholders to replace the dataset associated with
// the analysis. You can use templates to create dashboards by replacing dataset
// placeholders with datasets which follow the same schema that was used to
// create the source analysis and template.
//
// To create a template from an existing analysis, use the analysis's ARN, aws-account-id,
// template-id, source-entity, and data-set-references.
//
// CLI syntax to create a template:
//
// aws quicksight create-template —cli-input-json file://create-template.json
//
// CLI syntax to create a template from another template in the same AWS account:
//
// aws quicksight create-template --aws-account-id 111122223333 --template-id
// reports_test_template --data-set-references DataSetPlaceholder=reports,DataSetArn=arn:aws:quicksight:us-west-2:111122223333:dataset/0dfc789c-81f6-4f4f-b9ac-7db2453eefc8
// DataSetPlaceholder=Elblogs,DataSetArn=arn:aws:quicksight:us-west-2:111122223333:dataset/f60da323-af68-45db-9016-08e0d1d7ded5
// --source-entity SourceAnalysis='{Arn=arn:aws:quicksight:us-west-2:111122223333:analysis/7fb74527-c36d-4be8-8139-ac1be4c97365}'
//
// To create template from another account’s template, you need to grant cross
// account resource permission for DescribeTemplate the account that contains
// the template.
//
// You can use a file to pass JSON to the function if you prefer.
//
//    // Example sending a request using CreateTemplateRequest.
//    req := client.CreateTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/CreateTemplate
func (c *Client) CreateTemplateRequest(input *CreateTemplateInput) CreateTemplateRequest {
	op := &aws.Operation{
		Name:       opCreateTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{AwsAccountId}/templates/{TemplateId}",
	}

	if input == nil {
		input = &CreateTemplateInput{}
	}

	req := c.newRequest(op, input, &CreateTemplateOutput{})
	return CreateTemplateRequest{Request: req, Input: input, Copy: c.CreateTemplateRequest}
}

// CreateTemplateRequest is the request type for the
// CreateTemplate API operation.
type CreateTemplateRequest struct {
	*aws.Request
	Input *CreateTemplateInput
	Copy  func(*CreateTemplateInput) CreateTemplateRequest
}

// Send marshals and sends the CreateTemplate API request.
func (r CreateTemplateRequest) Send(ctx context.Context) (*CreateTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTemplateResponse{
		CreateTemplateOutput: r.Request.Data.(*CreateTemplateOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTemplateResponse is the response type for the
// CreateTemplate API operation.
type CreateTemplateResponse struct {
	*CreateTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTemplate request.
func (r *CreateTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}