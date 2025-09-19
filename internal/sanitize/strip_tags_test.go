package sanitize_test

import (
	"testing"

	"github.com/ericbutera/amalgam/internal/sanitize"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStripTags(t *testing.T) {
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
			name:     "tags",
			input:    "<p>tags</p>",
			expected: "tags",
		},
		{
			name:     "tags with attributes",
			input:    "<p class='test'>tags</p>",
			expected: "tags",
		},
		{
			name:     "tags with nested",
			input:    "<p><strong>tags</strong></p>",
			expected: "tags",
		},
		{
			name:     "malformed nesting",
			input:    "<p> malformed nesting  </strong></p>",
			expected: "malformed nesting",
		},
		{
			name:     "malformed tags",
			input:    "<p<strong<>malformed</strong<di///></p>tags",
			expected: "malformed tags",
		},
		{
			name:     "happy path",
			input:    "<h1>title</h1><p>content</p>",
			expected: "title content",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res, ok := sanitize.StripTags(tc.input)
			require.True(t, ok)
			assert.Equal(t, tc.expected, res)
		})
	}
}
