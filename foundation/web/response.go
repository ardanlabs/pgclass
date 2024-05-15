package web

import (
	"context"
	"fmt"
	"net/http"
)

func respond(ctx context.Context, w http.ResponseWriter, data Encoder) error {
	b, err := data.Encode()
	if err != nil {
		return fmt.Errorf("respond: encode: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(b); err != nil {
		return fmt.Errorf("respond: write: %w", err)
	}

	return nil
}
