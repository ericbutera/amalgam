package sanitize_test

import (
	"testing"

	"github.com/ericbutera/amalgam/internal/sanitize"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeWhitespace(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty",
			input:    "",
			expected: "",
		},
		{
			name:     "single space",
			input:    " ",
			expected: "",
		},
		{
			name:     "multiple spaces",
			input:    "  ",
			expected: "",
		},
		{
			name:     "leading space",
			input:    " a",
			expected: "a",
		},
		{
			name:     "trailing space",
			input:    "a ",
			expected: "a",
		},
		{
			name:     "random spaces",
			input:    " This         is  a                      test ",
			expected: "This is a test",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, sanitize.NormalizeWhitespace(tc.input))
		})
	}
}
