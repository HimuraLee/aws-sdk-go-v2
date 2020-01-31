// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// The state of an application discovered through Migration Hub import, the
// AWS Agentless Discovery Connector, or the AWS Application Discovery Agent.
type ApplicationState struct {
	_ struct{} `type:"structure"`

	// The configurationId from the Application Discovery Service that uniquely
	// identifies an application.
	ApplicationId *string `min:"1" type:"string"`

	// The current status of an application.
	ApplicationStatus ApplicationStatus `type:"string" enum:"true"`

	// The timestamp when the application status was last updated.
	LastUpdatedTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s ApplicationState) String() string {
	return awsutil.Prettify(s)
}

// An ARN of the AWS cloud resource target receiving the migration (e.g., AMI,
// EC2 instance, RDS instance, etc.).
type CreatedArtifact struct {
	_ struct{} `type:"structure"`

	// A description that can be free-form text to record additional detail about
	// the artifact for clarity or for later reference.
	Description *string `type:"string"`

	// An ARN that uniquely identifies the result of a migration task.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreatedArtifact) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatedArtifact) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatedArtifact"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Object representing the on-premises resource being migrated.
type DiscoveredResource struct {
	_ struct{} `type:"structure"`

	// The configurationId in Application Discovery Service that uniquely identifies
	// the on-premise resource.
	//
	// ConfigurationId is a required field
	ConfigurationId *string `min:"1" type:"string" required:"true"`

	// A description that can be free-form text to record additional detail about
	// the discovered resource for clarity or later reference.
	Description *string `type:"string"`
}

// String returns the string representation
func (s DiscoveredResource) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DiscoveredResource) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DiscoveredResource"}

	if s.ConfigurationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationId"))
	}
	if s.ConfigurationId != nil && len(*s.ConfigurationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigurationId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a migration task in a migration tool.
type MigrationTask struct {
	_ struct{} `type:"structure"`

	// Unique identifier that references the migration task. Do not store personal
	// data in this field.
	MigrationTaskName *string `min:"1" type:"string"`

	// A name that identifies the vendor of the migration tool being used.
	ProgressUpdateStream *string `min:"1" type:"string"`

	// Information about the resource that is being migrated. This data will be
	// used to map the task to a resource in the Application Discovery Service repository.
	ResourceAttributeList []ResourceAttribute `type:"list"`

	// Task object encapsulating task information.
	Task *Task `type:"structure"`

	// The timestamp when the task was gathered.
	UpdateDateTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s MigrationTask) String() string {
	return awsutil.Prettify(s)
}

// MigrationTaskSummary includes MigrationTaskName, ProgressPercent, ProgressUpdateStream,
// Status, and UpdateDateTime for each task.
type MigrationTaskSummary struct {
	_ struct{} `type:"structure"`

	// Unique identifier that references the migration task. Do not store personal
	// data in this field.
	MigrationTaskName *string `min:"1" type:"string"`

	// Indication of the percentage completion of the task.
	ProgressPercent *int64 `type:"integer"`

	// An AWS resource used for access control. It should uniquely identify the
	// migration tool as it is used for all updates made by the tool.
	ProgressUpdateStream *string `min:"1" type:"string"`

	// Status of the task.
	Status Status `type:"string" enum:"true"`

	// Detail information of what is being done within the overall status state.
	StatusDetail *string `type:"string"`

	// The timestamp when the task was gathered.
	UpdateDateTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s MigrationTaskSummary) String() string {
	return awsutil.Prettify(s)
}

// Summary of the AWS resource used for access control that is implicitly linked
// to your AWS account.
type ProgressUpdateStreamSummary struct {
	_ struct{} `type:"structure"`

	// The name of the ProgressUpdateStream. Do not store personal data in this
	// field.
	ProgressUpdateStreamName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ProgressUpdateStreamSummary) String() string {
	return awsutil.Prettify(s)
}

// Attribute associated with a resource.
//
// Note the corresponding format required per type listed below:
//
// IPV4
//
// x.x.x.x
//
// where x is an integer in the range [0,255]
//
// IPV6
//
// y : y : y : y : y : y : y : y
//
// where y is a hexadecimal between 0 and FFFF. [0, FFFF]
//
// MAC_ADDRESS
//
// ^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$
//
// FQDN
//
// ^[^<>{}\\\\/?,=\\p{Cntrl}]{1,256}$
type ResourceAttribute struct {
	_ struct{} `type:"structure"`

	// Type of resource.
	//
	// Type is a required field
	Type ResourceAttributeType `type:"string" required:"true" enum:"true"`

	// Value of the resource type.
	//
	// Value is a required field
	Value *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ResourceAttribute) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResourceAttribute) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResourceAttribute"}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}
	if s.Value != nil && len(*s.Value) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Value", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Task object encapsulating task information.
type Task struct {
	_ struct{} `type:"structure"`

	// Indication of the percentage completion of the task.
	ProgressPercent *int64 `type:"integer"`

	// Status of the task - Not Started, In-Progress, Complete.
	//
	// Status is a required field
	Status Status `type:"string" required:"true" enum:"true"`

	// Details of task status as notified by a migration tool. A tool might use
	// this field to provide clarifying information about the status that is unique
	// to that tool or that explains an error state.
	StatusDetail *string `type:"string"`
}

// String returns the string representation
func (s Task) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Task) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Task"}
	if len(s.Status) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Status"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
