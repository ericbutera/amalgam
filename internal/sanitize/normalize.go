package sanitize

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`\s+`)

func NormalizeWhitespace(s string) string {
	s = re.ReplaceAllString(s, " ")
	s = strings.Trim(s, " ")
	return s
}
