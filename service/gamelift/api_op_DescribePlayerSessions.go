// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type DescribePlayerSessionsInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the game session to retrieve player sessions for.
	GameSessionId *string `min:"1" type:"string"`

	// The maximum number of results to return. Use this parameter with NextToken
	// to get results as a set of sequential pages. If a player session ID is specified,
	// this parameter is ignored.
	Limit *int64 `min:"1" type:"integer"`

	// Token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this action. To start
	// at the beginning of the result set, do not specify a value. If a player session
	// ID is specified, this parameter is ignored.
	NextToken *string `min:"1" type:"string"`

	// A unique identifier for a player to retrieve player sessions for.
	PlayerId *string `min:"1" type:"string"`

	// A unique identifier for a player session to retrieve.
	PlayerSessionId *string `type:"string"`

	// Player session status to filter results on.
	//
	// Possible player session statuses include the following:
	//
	//    * RESERVED -- The player session request has been received, but the player
	//    has not yet connected to the server process and/or been validated.
	//
	//    * ACTIVE -- The player has been validated by the server process and is
	//    currently connected.
	//
	//    * COMPLETED -- The player connection has been dropped.
	//
	//    * TIMEDOUT -- A player session request was received, but the player did
	//    not connect and/or was not validated within the timeout limit (60 seconds).
	PlayerSessionStatusFilter *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribePlayerSessionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePlayerSessionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePlayerSessionsInput"}
	if s.GameSessionId != nil && len(*s.GameSessionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionId", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.PlayerId != nil && len(*s.PlayerId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlayerId", 1))
	}
	if s.PlayerSessionStatusFilter != nil && len(*s.PlayerSessionStatusFilter) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlayerSessionStatusFilter", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type DescribePlayerSessionsOutput struct {
	_ struct{} `type:"structure"`

	// Token that indicates where to resume retrieving results on the next call
	// to this action. If no token is returned, these results represent the end
	// of the list.
	NextToken *string `min:"1" type:"string"`

	// A collection of objects containing properties for each player session that
	// matches the request.
	PlayerSessions []PlayerSession `type:"list"`
}

// String returns the string representation
func (s DescribePlayerSessionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribePlayerSessions = "DescribePlayerSessions"

// DescribePlayerSessionsRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves properties for one or more player sessions. This action can be
// used in several ways: (1) provide a PlayerSessionId to request properties
// for a specific player session; (2) provide a GameSessionId to request properties
// for all player sessions in the specified game session; (3) provide a PlayerId
// to request properties for all player sessions of a specified player.
//
// To get game session record(s), specify only one of the following: a player
// session ID, a game session ID, or a player ID. You can filter this request
// by player session status. Use the pagination parameters to retrieve results
// as a set of sequential pages. If successful, a PlayerSession object is returned
// for each session matching the request.
//
// Available in Amazon GameLift Local.
//
//    * CreatePlayerSession
//
//    * CreatePlayerSessions
//
//    * DescribePlayerSessions
//
//    * Game session placements StartGameSessionPlacement DescribeGameSessionPlacement
//    StopGameSessionPlacement
//
//    // Example sending a request using DescribePlayerSessionsRequest.
//    req := client.DescribePlayerSessionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribePlayerSessions
func (c *Client) DescribePlayerSessionsRequest(input *DescribePlayerSessionsInput) DescribePlayerSessionsRequest {
	op := &aws.Operation{
		Name:       opDescribePlayerSessions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribePlayerSessionsInput{}
	}

	req := c.newRequest(op, input, &DescribePlayerSessionsOutput{})
	return DescribePlayerSessionsRequest{Request: req, Input: input, Copy: c.DescribePlayerSessionsRequest}
}

// DescribePlayerSessionsRequest is the request type for the
// DescribePlayerSessions API operation.
type DescribePlayerSessionsRequest struct {
	*aws.Request
	Input *DescribePlayerSessionsInput
	Copy  func(*DescribePlayerSessionsInput) DescribePlayerSessionsRequest
}

// Send marshals and sends the DescribePlayerSessions API request.
func (r DescribePlayerSessionsRequest) Send(ctx context.Context) (*DescribePlayerSessionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePlayerSessionsResponse{
		DescribePlayerSessionsOutput: r.Request.Data.(*DescribePlayerSessionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePlayerSessionsResponse is the response type for the
// DescribePlayerSessions API operation.
type DescribePlayerSessionsResponse struct {
	*DescribePlayerSessionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePlayerSessions request.
func (r *DescribePlayerSessionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
