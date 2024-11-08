package sanitize_test

import (
	"testing"

	"github.com/ericbutera/amalgam/internal/sanitize"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	Url     string `san:"url"`
	Content string `san:"html"`
	Title   string `validate:"trim"`
}

func TestHappyPath(t *testing.T) {
	expected := TestStruct{
		Url:     "https://example.com",
		Content: "<p>Content</p>",
		Title:   "Title",
	}
	actual, err := sanitize.Struct(expected)
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUrl(t *testing.T) {
	tt := []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "Valid URL",
			url:      "https://example.com",
			expected: "https://example.com",
		},
		{
			name:     "URL with trailing slashes",
			url:      "https://example.com/////",
			expected: "https://example.com/",
		},
		{
			name:     "URL with query parameters",
			url:      "https://example.com?query=param",
			expected: "https://example.com?query=param",
		},
		{
			name:     "caps",
			url:      "HTTPS://EXAMPLE.COM",
			expected: "https://example.com",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			data := TestStruct{
				Url: tc.url,
			}
			actual, err := sanitize.Struct(data)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, actual.Url)
		})
	}
}

func TestHtml(t *testing.T) {
	tt := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "denied tag",
			content:  "<script>Content</script>",
			expected: "",
		},
		{
			name:     "allowed tag",
			content:  "<p>Content</p>",
			expected: "<p>Content</p>",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			data := TestStruct{
				Content: tc.content,
			}
			actual, err := sanitize.Struct(data)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, actual.Content)
		})
	}
}
