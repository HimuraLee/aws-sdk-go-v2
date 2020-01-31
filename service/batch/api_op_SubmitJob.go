// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package batch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type SubmitJobInput struct {
	_ struct{} `type:"structure"`

	// The array properties for the submitted job, such as the size of the array.
	// The array size can be between 2 and 10,000. If you specify array properties
	// for a job, it becomes an array job. For more information, see Array Jobs
	// (https://docs.aws.amazon.com/batch/latest/userguide/array_jobs.html) in the
	// AWS Batch User Guide.
	ArrayProperties *ArrayProperties `locationName:"arrayProperties" type:"structure"`

	// A list of container overrides in JSON format that specify the name of a container
	// in the specified job definition and the overrides it should receive. You
	// can override the default command for a container (that is specified in the
	// job definition or the Docker image) with a command override. You can also
	// override existing environment variables (that are specified in the job definition
	// or Docker image) on a container or add new environment variables to it with
	// an environment override.
	ContainerOverrides *ContainerOverrides `locationName:"containerOverrides" type:"structure"`

	// A list of dependencies for the job. A job can depend upon a maximum of 20
	// jobs. You can specify a SEQUENTIAL type dependency without specifying a job
	// ID for array jobs so that each child array job completes sequentially, starting
	// at index 0. You can also specify an N_TO_N type dependency with a job ID
	// for array jobs. In that case, each index child of this job must wait for
	// the corresponding index child of each dependency to complete before it can
	// begin.
	DependsOn []JobDependency `locationName:"dependsOn" type:"list"`

	// The job definition used by this job. This value can be one of name, name:revision,
	// or the Amazon Resource Name (ARN) for the job definition. If name is specified
	// without a revision then the latest active revision is used.
	//
	// JobDefinition is a required field
	JobDefinition *string `locationName:"jobDefinition" type:"string" required:"true"`

	// The name of the job. The first character must be alphanumeric, and up to
	// 128 letters (uppercase and lowercase), numbers, hyphens, and underscores
	// are allowed.
	//
	// JobName is a required field
	JobName *string `locationName:"jobName" type:"string" required:"true"`

	// The job queue into which the job is submitted. You can specify either the
	// name or the Amazon Resource Name (ARN) of the queue.
	//
	// JobQueue is a required field
	JobQueue *string `locationName:"jobQueue" type:"string" required:"true"`

	// A list of node overrides in JSON format that specify the node range to target
	// and the container overrides for that node range.
	NodeOverrides *NodeOverrides `locationName:"nodeOverrides" type:"structure"`

	// Additional parameters passed to the job that replace parameter substitution
	// placeholders that are set in the job definition. Parameters are specified
	// as a key and value pair mapping. Parameters in a SubmitJob request override
	// any corresponding parameter defaults from the job definition.
	Parameters map[string]string `locationName:"parameters" type:"map"`

	// The retry strategy to use for failed jobs from this SubmitJob operation.
	// When a retry strategy is specified here, it overrides the retry strategy
	// defined in the job definition.
	RetryStrategy *RetryStrategy `locationName:"retryStrategy" type:"structure"`

	// The timeout configuration for this SubmitJob operation. You can specify a
	// timeout duration after which AWS Batch terminates your jobs if they have
	// not finished. If a job is terminated due to a timeout, it is not retried.
	// The minimum value for the timeout is 60 seconds. This configuration overrides
	// any timeout configuration specified in the job definition. For array jobs,
	// child jobs have the same timeout configuration as the parent job. For more
	// information, see Job Timeouts (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/job_timeouts.html)
	// in the Amazon Elastic Container Service Developer Guide.
	Timeout *JobTimeout `locationName:"timeout" type:"structure"`
}

// String returns the string representation
func (s SubmitJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SubmitJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SubmitJobInput"}

	if s.JobDefinition == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobDefinition"))
	}

	if s.JobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobName"))
	}

	if s.JobQueue == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobQueue"))
	}
	if s.ContainerOverrides != nil {
		if err := s.ContainerOverrides.Validate(); err != nil {
			invalidParams.AddNested("ContainerOverrides", err.(aws.ErrInvalidParams))
		}
	}
	if s.NodeOverrides != nil {
		if err := s.NodeOverrides.Validate(); err != nil {
			invalidParams.AddNested("NodeOverrides", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SubmitJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ArrayProperties != nil {
		v := s.ArrayProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "arrayProperties", v, metadata)
	}
	if s.ContainerOverrides != nil {
		v := s.ContainerOverrides

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "containerOverrides", v, metadata)
	}
	if s.DependsOn != nil {
		v := s.DependsOn

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "dependsOn", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.JobDefinition != nil {
		v := *s.JobDefinition

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobDefinition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobName != nil {
		v := *s.JobName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobQueue != nil {
		v := *s.JobQueue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobQueue", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NodeOverrides != nil {
		v := s.NodeOverrides

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "nodeOverrides", v, metadata)
	}
	if s.Parameters != nil {
		v := s.Parameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "parameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.RetryStrategy != nil {
		v := s.RetryStrategy

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "retryStrategy", v, metadata)
	}
	if s.Timeout != nil {
		v := s.Timeout

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "timeout", v, metadata)
	}
	return nil
}

type SubmitJobOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the job.
	//
	// JobId is a required field
	JobId *string `locationName:"jobId" type:"string" required:"true"`

	// The name of the job.
	//
	// JobName is a required field
	JobName *string `locationName:"jobName" type:"string" required:"true"`
}

// String returns the string representation
func (s SubmitJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SubmitJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobName != nil {
		v := *s.JobName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opSubmitJob = "SubmitJob"

// SubmitJobRequest returns a request value for making API operation for
// AWS Batch.
//
// Submits an AWS Batch job from a job definition. Parameters specified during
// SubmitJob override parameters defined in the job definition.
//
//    // Example sending a request using SubmitJobRequest.
//    req := client.SubmitJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/SubmitJob
func (c *Client) SubmitJobRequest(input *SubmitJobInput) SubmitJobRequest {
	op := &aws.Operation{
		Name:       opSubmitJob,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/submitjob",
	}

	if input == nil {
		input = &SubmitJobInput{}
	}

	req := c.newRequest(op, input, &SubmitJobOutput{})
	return SubmitJobRequest{Request: req, Input: input, Copy: c.SubmitJobRequest}
}

// SubmitJobRequest is the request type for the
// SubmitJob API operation.
type SubmitJobRequest struct {
	*aws.Request
	Input *SubmitJobInput
	Copy  func(*SubmitJobInput) SubmitJobRequest
}

// Send marshals and sends the SubmitJob API request.
func (r SubmitJobRequest) Send(ctx context.Context) (*SubmitJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SubmitJobResponse{
		SubmitJobOutput: r.Request.Data.(*SubmitJobOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SubmitJobResponse is the response type for the
// SubmitJob API operation.
type SubmitJobResponse struct {
	*SubmitJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SubmitJob request.
func (r *SubmitJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
