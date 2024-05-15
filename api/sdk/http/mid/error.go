package mid

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/app/sdk/mid"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

func Error(log *logger.Logger) web.Middleware {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, r *http.Request) (web.Encoder, error) {

			next := func(ctx context.Context) (mid.Encoder, error) {
				return handler(ctx, r)
			}

			return mid.Errors(ctx, log, next)
		}

		return h
	}

	return m
}
