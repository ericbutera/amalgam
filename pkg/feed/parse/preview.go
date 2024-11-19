package parse

import (
	"log/slog"
	"strings"

	"github.com/ericbutera/amalgam/internal/sanitize"
)

const DefaultPreviewLength = 1000

// creates a plaintext preview string from an HTML string
func Preview(s string) string {
	return PreviewWithLength(s, DefaultPreviewLength)
}

func PreviewWithLength(s string, length int) string {
	s, ok := sanitize.StripTags(s)
	if !ok {
		slog.Warn("parse: failed to strip tags from preview")
	}

	if len(s) > length {
		s = truncateWithoutBreakingWords(s, length)
	}

	s = sanitize.NormalizeWhitespace(s)

	return s
}

func truncateWithoutBreakingWords(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	words := strings.Fields(s[:maxLength+1])
	if len(words) > 1 && len(strings.Join(words[:len(words)-1], " ")) <= maxLength {
		return strings.Join(words[:len(words)-1], " ")
	}
	return strings.Join(words, " ")
}
