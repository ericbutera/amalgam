package converters

import "github.com/ericbutera/amalgam/internal/goverter"

func New() goverter.Converter {
	return &goverter.ConverterImpl{}
}
