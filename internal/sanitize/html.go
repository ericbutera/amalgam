package sanitize

import "github.com/microcosm-cc/bluemonday"

// attempt to sanitize HTML content
//
// quoting https://github.com/microcosm-cc/bluemonday:
// bluemonday takes untrusted user generated content as an input, and will return HTML that has been
// sanitised against an allowlist of approved HTML elements and attributes so that you can safely
// include the content in your web page.
func Html(html string) string {
	return bluemonday.UGCPolicy().Sanitize(html)
}
