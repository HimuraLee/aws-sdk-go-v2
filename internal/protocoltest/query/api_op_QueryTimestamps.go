// Code generated by smithy-go-codegen DO NOT EDIT.
package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This test serializes timestamps.
//
//     * Timestamps are serialized as RFC 3339
// date-time values by default.
//
//     * A timestampFormat trait on a member changes
// the format.
//
//     * A timestampFormat trait on the shape targeted by the member
// changes the format.
func (c *Client) QueryTimestamps(ctx context.Context, params *QueryTimestampsInput, optFns ...func(*Options)) (*QueryTimestampsOutput, error) {
	stack := middleware.NewStack("QueryTimestamps", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	stack.Initialize.Add(newServiceMetadataMiddleware_opQueryTimestamps(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "QueryTimestamps",
			Err:           err,
		}
	}
	out := result.(*QueryTimestampsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryTimestampsInput struct {
	NormalFormat *time.Time
	EpochMember  *time.Time
	EpochTarget  *time.Time
}

type QueryTimestampsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func newServiceMetadataMiddleware_opQueryTimestamps(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Query Protocol",
		ServiceID:      "queryprotocol",
		EndpointPrefix: "queryprotocol",
		OperationName:  "QueryTimestamps",
	}
}