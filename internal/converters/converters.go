package converters

import "github.com/ericbutera/amalgam/internal/goverter"

type Converter = goverter.Converter

func New() Converter {
	return &goverter.ConverterImpl{}
}
