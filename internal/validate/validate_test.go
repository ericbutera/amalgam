package validate_test

import (
	"testing"

	"github.com/ericbutera/amalgam/internal/validate"
	"github.com/stretchr/testify/assert"
)

var customMessages = validate.CustomMessages{
	"ID.required":     "The ID field is required.",
	"ID.uuid4":        "The ID must be a valid UUID.",
	"FeedID.required": "The FeedID field is required.",
	"FeedID.uuid4":    "The FeedID must be a valid UUID.",
	"Url.required":    "The URL is required.",
	"Url.url":         "The URL must be a valid URL.",
	"Title.required":  "The title is required.",
	"Title.min":       "The title must be at least 3 characters long.",
}

type TestStruct struct {
	ID     string `validate:"required,uuid4"`
	FeedID string `validate:"required,uuid4"`
	Url    string `validate:"required,url"`
	Title  string `validate:"required"`
}

func TestFeedValidateErrors(t *testing.T) {
	expected := validate.ValidationResult{
		Errors: []validate.ValidationError{
			{
				Field:           "ID",
				Tag:             "required",
				RawMessage:      "Key: 'TestStruct.ID' Error:Field validation for 'ID' failed on the 'required' tag",
				FriendlyMessage: "The ID field is required.",
			},
			{
				Field:           "FeedID",
				Tag:             "required",
				RawMessage:      "Key: 'TestStruct.FeedID' Error:Field validation for 'FeedID' failed on the 'required' tag",
				FriendlyMessage: "The FeedID field is required.",
			},
			{
				Field:           "Url",
				Tag:             "required",
				RawMessage:      "Key: 'TestStruct.Url' Error:Field validation for 'Url' failed on the 'required' tag",
				FriendlyMessage: "The URL is required.",
			},
			{
				Field:           "Title",
				Tag:             "required",
				RawMessage:      "Key: 'TestStruct.Title' Error:Field validation for 'Title' failed on the 'required' tag",
				FriendlyMessage: "The title is required.",
			},
		},
		Ok: false,
	}

	s := TestStruct{}
	actual := validate.Struct(s, customMessages)
	assert.False(t, actual.Ok)
	assert.Len(t, actual.Errors, 4)
	assert.Equal(t, expected, actual)
}

func TestValidateNoErrors(t *testing.T) {
	s := TestStruct{
		ID:     "e5b3b3d7-4f2c-4d4d-8f0a-2d7c8b3d6f6d",
		FeedID: "e5b3b3d7-4f2c-4d4d-8f0a-2d7c8b3d6f6d",
		Url:    "https://example.com",
		Title:  "Title",
	}
	actual := validate.Struct(s, customMessages)
	assert.True(t, actual.Ok)
	assert.Empty(t, actual.Errors)
}
