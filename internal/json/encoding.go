package json

import (
	encoding "encoding/json"
	"io"

	errs "github.com/configservice/internal/error"
)

func DecodeBody(v interface{}, r io.Reader) errs.Error {
	err := encoding.NewDecoder(r).Decode(v)
	if err != nil {
		return errs.NewDecoderError(err)
	}

	return nil
}
