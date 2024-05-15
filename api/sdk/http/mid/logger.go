package mid

import (
	"context"
	"net/http"
	"time"

	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

func Logger(log *logger.Logger) web.Middleware {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, r *http.Request) (web.Encoder, error) {
			now := time.Now()

			log.Info(ctx, "request started", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

			resp, err := handler(ctx, r)

			log.Info(ctx, "request completed", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr,
				"since", time.Since(now).String())

			return resp, err
		}

		return h
	}

	return m
}
