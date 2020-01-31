// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codebuildiface provides an interface to enable mocking the AWS CodeBuild service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codebuildiface

import (
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
)

// ClientAPI provides an interface to enable mocking the
// codebuild.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS CodeBuild.
//    func myFunc(svc codebuildiface.ClientAPI) bool {
//        // Make svc.BatchDeleteBuilds request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := codebuild.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        codebuildiface.ClientPI
//    }
//    func (m *mockClientClient) BatchDeleteBuilds(input *codebuild.BatchDeleteBuildsInput) (*codebuild.BatchDeleteBuildsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	BatchDeleteBuildsRequest(*codebuild.BatchDeleteBuildsInput) codebuild.BatchDeleteBuildsRequest

	BatchGetBuildsRequest(*codebuild.BatchGetBuildsInput) codebuild.BatchGetBuildsRequest

	BatchGetProjectsRequest(*codebuild.BatchGetProjectsInput) codebuild.BatchGetProjectsRequest

	BatchGetReportGroupsRequest(*codebuild.BatchGetReportGroupsInput) codebuild.BatchGetReportGroupsRequest

	BatchGetReportsRequest(*codebuild.BatchGetReportsInput) codebuild.BatchGetReportsRequest

	CreateProjectRequest(*codebuild.CreateProjectInput) codebuild.CreateProjectRequest

	CreateReportGroupRequest(*codebuild.CreateReportGroupInput) codebuild.CreateReportGroupRequest

	CreateWebhookRequest(*codebuild.CreateWebhookInput) codebuild.CreateWebhookRequest

	DeleteProjectRequest(*codebuild.DeleteProjectInput) codebuild.DeleteProjectRequest

	DeleteReportRequest(*codebuild.DeleteReportInput) codebuild.DeleteReportRequest

	DeleteReportGroupRequest(*codebuild.DeleteReportGroupInput) codebuild.DeleteReportGroupRequest

	DeleteResourcePolicyRequest(*codebuild.DeleteResourcePolicyInput) codebuild.DeleteResourcePolicyRequest

	DeleteSourceCredentialsRequest(*codebuild.DeleteSourceCredentialsInput) codebuild.DeleteSourceCredentialsRequest

	DeleteWebhookRequest(*codebuild.DeleteWebhookInput) codebuild.DeleteWebhookRequest

	DescribeTestCasesRequest(*codebuild.DescribeTestCasesInput) codebuild.DescribeTestCasesRequest

	GetResourcePolicyRequest(*codebuild.GetResourcePolicyInput) codebuild.GetResourcePolicyRequest

	ImportSourceCredentialsRequest(*codebuild.ImportSourceCredentialsInput) codebuild.ImportSourceCredentialsRequest

	InvalidateProjectCacheRequest(*codebuild.InvalidateProjectCacheInput) codebuild.InvalidateProjectCacheRequest

	ListBuildsRequest(*codebuild.ListBuildsInput) codebuild.ListBuildsRequest

	ListBuildsForProjectRequest(*codebuild.ListBuildsForProjectInput) codebuild.ListBuildsForProjectRequest

	ListCuratedEnvironmentImagesRequest(*codebuild.ListCuratedEnvironmentImagesInput) codebuild.ListCuratedEnvironmentImagesRequest

	ListProjectsRequest(*codebuild.ListProjectsInput) codebuild.ListProjectsRequest

	ListReportGroupsRequest(*codebuild.ListReportGroupsInput) codebuild.ListReportGroupsRequest

	ListReportsRequest(*codebuild.ListReportsInput) codebuild.ListReportsRequest

	ListReportsForReportGroupRequest(*codebuild.ListReportsForReportGroupInput) codebuild.ListReportsForReportGroupRequest

	ListSharedProjectsRequest(*codebuild.ListSharedProjectsInput) codebuild.ListSharedProjectsRequest

	ListSharedReportGroupsRequest(*codebuild.ListSharedReportGroupsInput) codebuild.ListSharedReportGroupsRequest

	ListSourceCredentialsRequest(*codebuild.ListSourceCredentialsInput) codebuild.ListSourceCredentialsRequest

	PutResourcePolicyRequest(*codebuild.PutResourcePolicyInput) codebuild.PutResourcePolicyRequest

	StartBuildRequest(*codebuild.StartBuildInput) codebuild.StartBuildRequest

	StopBuildRequest(*codebuild.StopBuildInput) codebuild.StopBuildRequest

	UpdateProjectRequest(*codebuild.UpdateProjectInput) codebuild.UpdateProjectRequest

	UpdateReportGroupRequest(*codebuild.UpdateReportGroupInput) codebuild.UpdateReportGroupRequest

	UpdateWebhookRequest(*codebuild.UpdateWebhookInput) codebuild.UpdateWebhookRequest
}

var _ ClientAPI = (*codebuild.Client)(nil)
