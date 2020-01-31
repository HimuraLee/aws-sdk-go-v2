// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateTrialComponentInput struct {
	_ struct{} `type:"structure"`

	// The name of the component as displayed. The name doesn't need to be unique.
	// If DisplayName isn't specified, TrialComponentName is displayed.
	DisplayName *string `min:"1" type:"string"`

	// When the component ended.
	EndTime *time.Time `type:"timestamp"`

	// The input artifacts for the component. Examples of input artifacts are datasets,
	// algorithms, hyperparameters, source code, and instance types.
	InputArtifacts map[string]TrialComponentArtifact `type:"map"`

	// The output artifacts for the component. Examples of output artifacts are
	// metrics, snapshots, logs, and images.
	OutputArtifacts map[string]TrialComponentArtifact `type:"map"`

	// The hyperparameters for the component.
	Parameters map[string]TrialComponentParameterValue `type:"map"`

	// When the component started.
	StartTime *time.Time `type:"timestamp"`

	// The status of the component. States include:
	//
	//    * InProgress
	//
	//    * Completed
	//
	//    * Failed
	Status *TrialComponentStatus `type:"structure"`

	// A list of tags to associate with the component. You can use Search API to
	// search on the tags.
	Tags []Tag `type:"list"`

	// The name of the component. The name must be unique in your AWS account and
	// is not case-sensitive.
	//
	// TrialComponentName is a required field
	TrialComponentName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateTrialComponentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTrialComponentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTrialComponentInput"}
	if s.DisplayName != nil && len(*s.DisplayName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DisplayName", 1))
	}

	if s.TrialComponentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrialComponentName"))
	}
	if s.TrialComponentName != nil && len(*s.TrialComponentName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrialComponentName", 1))
	}
	if s.InputArtifacts != nil {
		for i, v := range s.InputArtifacts {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "InputArtifacts", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.OutputArtifacts != nil {
		for i, v := range s.OutputArtifacts {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "OutputArtifacts", i), err.(aws.ErrInvalidParams))
			}
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

type CreateTrialComponentOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the trial component.
	TrialComponentArn *string `type:"string"`
}

// String returns the string representation
func (s CreateTrialComponentOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateTrialComponent = "CreateTrialComponent"

// CreateTrialComponentRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Creates a trial component, which is a stage of a machine learning trial.
// A trial is composed of one or more trial components. A trial component can
// be used in multiple trials.
//
// Trial components include pre-processing jobs, training jobs, and batch transform
// jobs.
//
// When you use Amazon SageMaker Studio or the Amazon SageMaker Python SDK,
// all experiments, trials, and trial components are automatically tracked,
// logged, and indexed. When you use the AWS SDK for Python (Boto), you must
// use the logging APIs provided by the SDK.
//
// You can add tags to a trial component and then use the Search API to search
// for the tags.
//
// CreateTrialComponent can only be invoked from within an Amazon SageMaker
// managed environment. This includes Amazon SageMaker training jobs, processing
// jobs, transform jobs, and Amazon SageMaker notebooks. A call to CreateTrialComponent
// from outside one of these environments results in an error.
//
//    // Example sending a request using CreateTrialComponentRequest.
//    req := client.CreateTrialComponentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/CreateTrialComponent
func (c *Client) CreateTrialComponentRequest(input *CreateTrialComponentInput) CreateTrialComponentRequest {
	op := &aws.Operation{
		Name:       opCreateTrialComponent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTrialComponentInput{}
	}

	req := c.newRequest(op, input, &CreateTrialComponentOutput{})
	return CreateTrialComponentRequest{Request: req, Input: input, Copy: c.CreateTrialComponentRequest}
}

// CreateTrialComponentRequest is the request type for the
// CreateTrialComponent API operation.
type CreateTrialComponentRequest struct {
	*aws.Request
	Input *CreateTrialComponentInput
	Copy  func(*CreateTrialComponentInput) CreateTrialComponentRequest
}

// Send marshals and sends the CreateTrialComponent API request.
func (r CreateTrialComponentRequest) Send(ctx context.Context) (*CreateTrialComponentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTrialComponentResponse{
		CreateTrialComponentOutput: r.Request.Data.(*CreateTrialComponentOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTrialComponentResponse is the response type for the
// CreateTrialComponent API operation.
type CreateTrialComponentResponse struct {
	*CreateTrialComponentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTrialComponent request.
func (r *CreateTrialComponentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
