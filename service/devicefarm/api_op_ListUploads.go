// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the list uploads operation.
type ListUploadsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the project for which you want to list
	// uploads.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// The type of upload.
	//
	// Must be one of the following values:
	//
	//    * ANDROID_APP
	//
	//    * IOS_APP
	//
	//    * WEB_APP
	//
	//    * EXTERNAL_DATA
	//
	//    * APPIUM_JAVA_JUNIT_TEST_PACKAGE
	//
	//    * APPIUM_JAVA_TESTNG_TEST_PACKAGE
	//
	//    * APPIUM_PYTHON_TEST_PACKAGE
	//
	//    * APPIUM_NODE_TEST_PACKAGE
	//
	//    * APPIUM_RUBY_TEST_PACKAGE
	//
	//    * APPIUM_WEB_JAVA_JUNIT_TEST_PACKAGE
	//
	//    * APPIUM_WEB_JAVA_TESTNG_TEST_PACKAGE
	//
	//    * APPIUM_WEB_PYTHON_TEST_PACKAGE
	//
	//    * APPIUM_WEB_NODE_TEST_PACKAGE
	//
	//    * APPIUM_WEB_RUBY_TEST_PACKAGE
	//
	//    * CALABASH_TEST_PACKAGE
	//
	//    * INSTRUMENTATION_TEST_PACKAGE
	//
	//    * UIAUTOMATION_TEST_PACKAGE
	//
	//    * UIAUTOMATOR_TEST_PACKAGE
	//
	//    * XCTEST_TEST_PACKAGE
	//
	//    * XCTEST_UI_TEST_PACKAGE
	//
	//    * APPIUM_JAVA_JUNIT_TEST_SPEC
	//
	//    * APPIUM_JAVA_TESTNG_TEST_SPEC
	//
	//    * APPIUM_PYTHON_TEST_SPEC
	//
	//    * APPIUM_NODE_TEST_SPEC
	//
	//    * APPIUM_RUBY_TEST_SPEC
	//
	//    * APPIUM_WEB_JAVA_JUNIT_TEST_SPEC
	//
	//    * APPIUM_WEB_JAVA_TESTNG_TEST_SPEC
	//
	//    * APPIUM_WEB_PYTHON_TEST_SPEC
	//
	//    * APPIUM_WEB_NODE_TEST_SPEC
	//
	//    * APPIUM_WEB_RUBY_TEST_SPEC
	//
	//    * INSTRUMENTATION_TEST_SPEC
	//
	//    * XCTEST_UI_TEST_SPEC
	Type UploadType `locationName:"type" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListUploadsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListUploadsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListUploadsInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a list uploads request.
type ListUploadsOutput struct {
	_ struct{} `type:"structure"`

	// If the number of items that are returned is significantly large, this is
	// an identifier that is also returned. It can be used in a subsequent call
	// to this operation to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// Information about the uploads.
	Uploads []Upload `locationName:"uploads" type:"list"`
}

// String returns the string representation
func (s ListUploadsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListUploads = "ListUploads"

// ListUploadsRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about uploads, given an AWS Device Farm project ARN.
//
//    // Example sending a request using ListUploadsRequest.
//    req := client.ListUploadsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListUploads
func (c *Client) ListUploadsRequest(input *ListUploadsInput) ListUploadsRequest {
	op := &aws.Operation{
		Name:       opListUploads,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListUploadsInput{}
	}

	req := c.newRequest(op, input, &ListUploadsOutput{})
	return ListUploadsRequest{Request: req, Input: input, Copy: c.ListUploadsRequest}
}

// ListUploadsRequest is the request type for the
// ListUploads API operation.
type ListUploadsRequest struct {
	*aws.Request
	Input *ListUploadsInput
	Copy  func(*ListUploadsInput) ListUploadsRequest
}

// Send marshals and sends the ListUploads API request.
func (r ListUploadsRequest) Send(ctx context.Context) (*ListUploadsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUploadsResponse{
		ListUploadsOutput: r.Request.Data.(*ListUploadsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListUploadsRequestPaginator returns a paginator for ListUploads.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListUploadsRequest(input)
//   p := devicefarm.NewListUploadsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListUploadsPaginator(req ListUploadsRequest) ListUploadsPaginator {
	return ListUploadsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListUploadsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListUploadsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListUploadsPaginator struct {
	aws.Pager
}

func (p *ListUploadsPaginator) CurrentPage() *ListUploadsOutput {
	return p.Pager.CurrentPage().(*ListUploadsOutput)
}

// ListUploadsResponse is the response type for the
// ListUploads API operation.
type ListUploadsResponse struct {
	*ListUploadsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUploads request.
func (r *ListUploadsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
