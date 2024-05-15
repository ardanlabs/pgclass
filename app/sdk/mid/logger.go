package mid

import (
	"context"
	"fmt"
	"time"

	"github.com/ardanlabs/service/foundation/logger"
)

// Logger writes information about the request to the logs.
func Logger(ctx context.Context, log *logger.Logger, path string, rawQuery string, method string, remoteAddr string, next Handler) (Encoder, error) {
	now := time.Now()

	if rawQuery != "" {
		path = fmt.Sprintf("%s?%s", path, rawQuery)
	}

	log.Info(ctx, "request started", "method", method, "path", path, "remoteaddr", remoteAddr)

	resp, err := next(ctx)

	log.Info(ctx, "request completed", "method", method, "path", path, "remoteaddr", remoteAddr,
		"since", time.Since(now).String())

	return resp, err
}
