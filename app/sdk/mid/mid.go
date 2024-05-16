package mid

import "context"

// Encoder defines behavior that can encode a data model and provide
// the content type for that encoding.
type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

// Handler represents an api layer handler function that needs to be called.
type Handler func(context.Context) (Encoder, error)
