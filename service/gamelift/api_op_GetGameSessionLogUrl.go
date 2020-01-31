// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type GetGameSessionLogUrlInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the game session to get logs for.
	//
	// GameSessionId is a required field
	GameSessionId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetGameSessionLogUrlInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetGameSessionLogUrlInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetGameSessionLogUrlInput"}

	if s.GameSessionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GameSessionId"))
	}
	if s.GameSessionId != nil && len(*s.GameSessionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type GetGameSessionLogUrlOutput struct {
	_ struct{} `type:"structure"`

	// Location of the requested game session logs, available for download. This
	// URL is valid for 15 minutes, after which S3 will reject any download request
	// using this URL. You can request a new URL any time within the 14-day period
	// that the logs are retained.
	PreSignedUrl *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetGameSessionLogUrlOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetGameSessionLogUrl = "GetGameSessionLogUrl"

// GetGameSessionLogUrlRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves the location of stored game session logs for a specified game session.
// When a game session is terminated, Amazon GameLift automatically stores the
// logs in Amazon S3 and retains them for 14 days. Use this URL to download
// the logs.
//
// See the AWS Service Limits (https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits_gamelift)
// page for maximum log file sizes. Log files that exceed this limit are not
// saved.
//
//    * CreateGameSession
//
//    * DescribeGameSessions
//
//    * DescribeGameSessionDetails
//
//    * SearchGameSessions
//
//    * UpdateGameSession
//
//    * GetGameSessionLogUrl
//
//    * Game session placements StartGameSessionPlacement DescribeGameSessionPlacement
//    StopGameSessionPlacement
//
//    // Example sending a request using GetGameSessionLogUrlRequest.
//    req := client.GetGameSessionLogUrlRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/GetGameSessionLogUrl
func (c *Client) GetGameSessionLogUrlRequest(input *GetGameSessionLogUrlInput) GetGameSessionLogUrlRequest {
	op := &aws.Operation{
		Name:       opGetGameSessionLogUrl,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetGameSessionLogUrlInput{}
	}

	req := c.newRequest(op, input, &GetGameSessionLogUrlOutput{})
	return GetGameSessionLogUrlRequest{Request: req, Input: input, Copy: c.GetGameSessionLogUrlRequest}
}

// GetGameSessionLogUrlRequest is the request type for the
// GetGameSessionLogUrl API operation.
type GetGameSessionLogUrlRequest struct {
	*aws.Request
	Input *GetGameSessionLogUrlInput
	Copy  func(*GetGameSessionLogUrlInput) GetGameSessionLogUrlRequest
}

// Send marshals and sends the GetGameSessionLogUrl API request.
func (r GetGameSessionLogUrlRequest) Send(ctx context.Context) (*GetGameSessionLogUrlResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetGameSessionLogUrlResponse{
		GetGameSessionLogUrlOutput: r.Request.Data.(*GetGameSessionLogUrlOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetGameSessionLogUrlResponse is the response type for the
// GetGameSessionLogUrl API operation.
type GetGameSessionLogUrlResponse struct {
	*GetGameSessionLogUrlOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetGameSessionLogUrl request.
func (r *GetGameSessionLogUrlResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
