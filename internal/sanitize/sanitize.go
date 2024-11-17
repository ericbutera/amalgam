package sanitize

import (
	"errors"
	"reflect"

	sanitizer "github.com/go-sanitize/sanitize"
)

var ErrInvalidURL = errors.New("invalid URL specified")

// Returns a sanitized copy of the original data.
func Struct[T any](data T) (T, error) {
	s, err := sanitizer.New(
		sanitizer.OptionSanitizerFunc{
			Name:      "url",
			Sanitizer: sanitizeUrl,
		},
		sanitizer.OptionSanitizerFunc{
			Name:      "html",
			Sanitizer: sanitizeHtml,
		},
	)
	if err != nil {
		return data, err
	}
	err = s.Sanitize(&data)
	return data, err
}

func sanitizeUrl(_ sanitizer.Sanitizer, structValue reflect.Value, idx int) error {
	fieldValue := structValue.Field(idx)
	url, err := Url(fieldValue.String())
	if err != nil {
		return ErrInvalidURL
	}
	fieldValue.SetString(url)
	return nil
}

func sanitizeHtml(_ sanitizer.Sanitizer, structValue reflect.Value, idx int) error {
	fieldValue := structValue.Field(idx)
	html := Html(fieldValue.String())
	fieldValue.SetString(html)
	return nil
}
