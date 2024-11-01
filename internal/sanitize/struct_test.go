package sanitize

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	actual, err := Struct(expected)
	assert.NoError(t, err)
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
			actual, err := Struct(data)
			assert.NoError(t, err)
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
			actual, err := Struct(data)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual.Content)
		})
	}
}
