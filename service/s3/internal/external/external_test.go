// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package external_test

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	svcExternal "github.com/aws/aws-sdk-go-v2/service/s3/internal/external"
)

// UseARNRegionProvider Assertions
var (
	_ svcExternal.UseARNRegionProvider = &external.EnvConfig{}
	_ svcExternal.UseARNRegionProvider = &external.SharedConfig{}
)
