// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package workspacesiface provides an interface to enable mocking the Amazon WorkSpaces service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package workspacesiface

import (
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
)

// ClientAPI provides an interface to enable mocking the
// workspaces.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon WorkSpaces.
//    func myFunc(svc workspacesiface.ClientAPI) bool {
//        // Make svc.AssociateIpGroups request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := workspaces.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        workspacesiface.ClientPI
//    }
//    func (m *mockClientClient) AssociateIpGroups(input *workspaces.AssociateIpGroupsInput) (*workspaces.AssociateIpGroupsOutput, error) {
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
	AssociateIpGroupsRequest(*workspaces.AssociateIpGroupsInput) workspaces.AssociateIpGroupsRequest

	AuthorizeIpRulesRequest(*workspaces.AuthorizeIpRulesInput) workspaces.AuthorizeIpRulesRequest

	CopyWorkspaceImageRequest(*workspaces.CopyWorkspaceImageInput) workspaces.CopyWorkspaceImageRequest

	CreateIpGroupRequest(*workspaces.CreateIpGroupInput) workspaces.CreateIpGroupRequest

	CreateTagsRequest(*workspaces.CreateTagsInput) workspaces.CreateTagsRequest

	CreateWorkspacesRequest(*workspaces.CreateWorkspacesInput) workspaces.CreateWorkspacesRequest

	DeleteIpGroupRequest(*workspaces.DeleteIpGroupInput) workspaces.DeleteIpGroupRequest

	DeleteTagsRequest(*workspaces.DeleteTagsInput) workspaces.DeleteTagsRequest

	DeleteWorkspaceImageRequest(*workspaces.DeleteWorkspaceImageInput) workspaces.DeleteWorkspaceImageRequest

	DeregisterWorkspaceDirectoryRequest(*workspaces.DeregisterWorkspaceDirectoryInput) workspaces.DeregisterWorkspaceDirectoryRequest

	DescribeAccountRequest(*workspaces.DescribeAccountInput) workspaces.DescribeAccountRequest

	DescribeAccountModificationsRequest(*workspaces.DescribeAccountModificationsInput) workspaces.DescribeAccountModificationsRequest

	DescribeClientPropertiesRequest(*workspaces.DescribeClientPropertiesInput) workspaces.DescribeClientPropertiesRequest

	DescribeIpGroupsRequest(*workspaces.DescribeIpGroupsInput) workspaces.DescribeIpGroupsRequest

	DescribeTagsRequest(*workspaces.DescribeTagsInput) workspaces.DescribeTagsRequest

	DescribeWorkspaceBundlesRequest(*workspaces.DescribeWorkspaceBundlesInput) workspaces.DescribeWorkspaceBundlesRequest

	DescribeWorkspaceDirectoriesRequest(*workspaces.DescribeWorkspaceDirectoriesInput) workspaces.DescribeWorkspaceDirectoriesRequest

	DescribeWorkspaceImagesRequest(*workspaces.DescribeWorkspaceImagesInput) workspaces.DescribeWorkspaceImagesRequest

	DescribeWorkspaceSnapshotsRequest(*workspaces.DescribeWorkspaceSnapshotsInput) workspaces.DescribeWorkspaceSnapshotsRequest

	DescribeWorkspacesRequest(*workspaces.DescribeWorkspacesInput) workspaces.DescribeWorkspacesRequest

	DescribeWorkspacesConnectionStatusRequest(*workspaces.DescribeWorkspacesConnectionStatusInput) workspaces.DescribeWorkspacesConnectionStatusRequest

	DisassociateIpGroupsRequest(*workspaces.DisassociateIpGroupsInput) workspaces.DisassociateIpGroupsRequest

	ImportWorkspaceImageRequest(*workspaces.ImportWorkspaceImageInput) workspaces.ImportWorkspaceImageRequest

	ListAvailableManagementCidrRangesRequest(*workspaces.ListAvailableManagementCidrRangesInput) workspaces.ListAvailableManagementCidrRangesRequest

	MigrateWorkspaceRequest(*workspaces.MigrateWorkspaceInput) workspaces.MigrateWorkspaceRequest

	ModifyAccountRequest(*workspaces.ModifyAccountInput) workspaces.ModifyAccountRequest

	ModifyClientPropertiesRequest(*workspaces.ModifyClientPropertiesInput) workspaces.ModifyClientPropertiesRequest

	ModifySelfservicePermissionsRequest(*workspaces.ModifySelfservicePermissionsInput) workspaces.ModifySelfservicePermissionsRequest

	ModifyWorkspaceAccessPropertiesRequest(*workspaces.ModifyWorkspaceAccessPropertiesInput) workspaces.ModifyWorkspaceAccessPropertiesRequest

	ModifyWorkspaceCreationPropertiesRequest(*workspaces.ModifyWorkspaceCreationPropertiesInput) workspaces.ModifyWorkspaceCreationPropertiesRequest

	ModifyWorkspacePropertiesRequest(*workspaces.ModifyWorkspacePropertiesInput) workspaces.ModifyWorkspacePropertiesRequest

	ModifyWorkspaceStateRequest(*workspaces.ModifyWorkspaceStateInput) workspaces.ModifyWorkspaceStateRequest

	RebootWorkspacesRequest(*workspaces.RebootWorkspacesInput) workspaces.RebootWorkspacesRequest

	RebuildWorkspacesRequest(*workspaces.RebuildWorkspacesInput) workspaces.RebuildWorkspacesRequest

	RegisterWorkspaceDirectoryRequest(*workspaces.RegisterWorkspaceDirectoryInput) workspaces.RegisterWorkspaceDirectoryRequest

	RestoreWorkspaceRequest(*workspaces.RestoreWorkspaceInput) workspaces.RestoreWorkspaceRequest

	RevokeIpRulesRequest(*workspaces.RevokeIpRulesInput) workspaces.RevokeIpRulesRequest

	StartWorkspacesRequest(*workspaces.StartWorkspacesInput) workspaces.StartWorkspacesRequest

	StopWorkspacesRequest(*workspaces.StopWorkspacesInput) workspaces.StopWorkspacesRequest

	TerminateWorkspacesRequest(*workspaces.TerminateWorkspacesInput) workspaces.TerminateWorkspacesRequest

	UpdateRulesOfIpGroupRequest(*workspaces.UpdateRulesOfIpGroupInput) workspaces.UpdateRulesOfIpGroupRequest
}

var _ ClientAPI = (*workspaces.Client)(nil)
