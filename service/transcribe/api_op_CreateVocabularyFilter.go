// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateVocabularyFilterInput struct {
	_ struct{} `type:"structure"`

	// The language code of the words in the vocabulary filter. All words in the
	// filter must be in the same language. The vocabulary filter can only be used
	// with transcription jobs in the specified language.
	//
	// LanguageCode is a required field
	LanguageCode LanguageCode `type:"string" required:"true" enum:"true"`

	// The Amazon S3 location of a text file used as input to create the vocabulary
	// filter. Only use characters from the character set defined for custom vocabularies.
	// For a list of character sets, see Character Sets for Custom Vocabularies
	// (https://docs.aws.amazon.com/transcribe/latest/dg/how-vocabulary.html#charsets).
	//
	// The specified file must be less than 50 KB of UTF-8 characters.
	//
	// If you provide the location of a list of words in the VocabularyFilterFileUri
	// parameter, you can't use the Words parameter.
	VocabularyFilterFileUri *string `min:"1" type:"string"`

	// The vocabulary filter name. The name must be unique within the account that
	// contains it.
	//
	// VocabularyFilterName is a required field
	VocabularyFilterName *string `min:"1" type:"string" required:"true"`

	// The words to use in the vocabulary filter. Only use characters from the character
	// set defined for custom vocabularies. For a list of character sets, see Character
	// Sets for Custom Vocabularies (https://docs.aws.amazon.com/transcribe/latest/dg/how-vocabulary.html#charsets).
	//
	// If you provide a list of words in the Words parameter, you can't use the
	// VocabularyFilterFileUri parameter.
	Words []string `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateVocabularyFilterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVocabularyFilterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVocabularyFilterInput"}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}
	if s.VocabularyFilterFileUri != nil && len(*s.VocabularyFilterFileUri) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VocabularyFilterFileUri", 1))
	}

	if s.VocabularyFilterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VocabularyFilterName"))
	}
	if s.VocabularyFilterName != nil && len(*s.VocabularyFilterName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VocabularyFilterName", 1))
	}
	if s.Words != nil && len(s.Words) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Words", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateVocabularyFilterOutput struct {
	_ struct{} `type:"structure"`

	// The language code of the words in the collection.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// The date and time that the vocabulary filter was modified.
	LastModifiedTime *time.Time `type:"timestamp"`

	// The name of the vocabulary filter.
	VocabularyFilterName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateVocabularyFilterOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateVocabularyFilter = "CreateVocabularyFilter"

// CreateVocabularyFilterRequest returns a request value for making API operation for
// Amazon Transcribe Service.
//
// Creates a new vocabulary filter that you can use to filter words, such as
// profane words, from the output of a transcription job.
//
//    // Example sending a request using CreateVocabularyFilterRequest.
//    req := client.CreateVocabularyFilterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/CreateVocabularyFilter
func (c *Client) CreateVocabularyFilterRequest(input *CreateVocabularyFilterInput) CreateVocabularyFilterRequest {
	op := &aws.Operation{
		Name:       opCreateVocabularyFilter,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVocabularyFilterInput{}
	}

	req := c.newRequest(op, input, &CreateVocabularyFilterOutput{})
	return CreateVocabularyFilterRequest{Request: req, Input: input, Copy: c.CreateVocabularyFilterRequest}
}

// CreateVocabularyFilterRequest is the request type for the
// CreateVocabularyFilter API operation.
type CreateVocabularyFilterRequest struct {
	*aws.Request
	Input *CreateVocabularyFilterInput
	Copy  func(*CreateVocabularyFilterInput) CreateVocabularyFilterRequest
}

// Send marshals and sends the CreateVocabularyFilter API request.
func (r CreateVocabularyFilterRequest) Send(ctx context.Context) (*CreateVocabularyFilterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVocabularyFilterResponse{
		CreateVocabularyFilterOutput: r.Request.Data.(*CreateVocabularyFilterOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVocabularyFilterResponse is the response type for the
// CreateVocabularyFilter API operation.
type CreateVocabularyFilterResponse struct {
	*CreateVocabularyFilterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVocabularyFilter request.
func (r *CreateVocabularyFilterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
