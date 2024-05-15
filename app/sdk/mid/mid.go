package mid

import "context"

// Encoder represents data that can be encoded.
type Encoder interface {
	Encode() ([]byte, error)
}

// Handler represents an api layer handler function that needs to be called.
type Handler func(context.Context) (Encoder, error)
