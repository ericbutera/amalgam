package sanitize

import (
	"errors"
	"reflect"

	sanitizer "github.com/go-sanitize/sanitize"
)

// Returns a sanitized copy of the original data.
func Struct[T interface{}](data T) (T, error) {
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

func sanitizeUrl(s sanitizer.Sanitizer, structValue reflect.Value, idx int) error {
	fieldValue := structValue.Field(idx)
	url, err := Url(fieldValue.String())
	if err != nil {
		return errors.New("invalid url")
	}
	fieldValue.SetString(url)
	return nil
}

func sanitizeHtml(s sanitizer.Sanitizer, structValue reflect.Value, idx int) error {
	fieldValue := structValue.Field(idx)
	html := Html(fieldValue.String())
	fieldValue.SetString(html)
	return nil
}
