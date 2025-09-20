package parse_test

import (
	"testing"

	"github.com/ericbutera/amalgam/pkg/feed/parse"
	"github.com/stretchr/testify/assert"
)

func TestPreview(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    string
		expected string
		length   int
	}{
		{
			name:     "empty",
			input:    "",
			expected: "",
			length:   parse.DefaultPreviewLength,
		},
		{
			name:     "html",
			input:    "<p><strong>html</strong></p>",
			expected: "html",
			length:   parse.DefaultPreviewLength,
		},
		{
			name:     "test break words",
			input:    "This is a test of the break words function. It should break at the first space after 70 characters.",
			expected: "This is a test of the break words",
			length:   42,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := parse.PreviewWithLength(tc.input, tc.length)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
