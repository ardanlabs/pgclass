package web

import (
	"fmt"
	"io"
	"net/http"
)

// Decoder represents data that can be decoded.
type Decoder interface {
	Decode(data []byte) error
}

// Decode reads the body of an HTTP request and decodes the body into the
// specified data model. If the data model implements the validator interface,
// the method will be called.
func Decode(r *http.Request, v Decoder) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("request: unable to read payload: %w", err)
	}

	if err := v.Decode(data); err != nil {
		return fmt.Errorf("request: encode: %w", err)
	}

	return nil
}
