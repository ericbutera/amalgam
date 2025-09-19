package sanitize

import (
	"strings"

	"golang.org/x/net/html"
)

func StripTags(s string) (string, bool) {
	doc, err := html.Parse(strings.NewReader(s)) // TODO: revisit if err should bubble up; right now it either works or doesn't
	if err != nil {
		return s, false
	}

	var b strings.Builder

	_, ok := appendNode(&b, doc)
	text := b.String()

	return strings.Trim(text, " "), ok
}

func appendNode(b *strings.Builder, node *html.Node) (int, bool) {
	ok := true
	written := 0

	if node.Type == html.TextNode {
		i, err := b.WriteString(node.Data + " ")
		if err != nil {
			ok = false
		}

		written += i
	}

	for next := node.FirstChild; next != nil; next = next.NextSibling {
		w, attempt := appendNode(b, next)
		if !attempt {
			ok = false
		}

		written += w
	}

	return written, ok
}
