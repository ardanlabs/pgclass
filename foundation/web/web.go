package web

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// Encoder defines behavior that can encode a data model and provide
// the content type for that encoding.
type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

// Handler represents a function that handles a http request within our own
// little mini framework.
type Handler func(context.Context, *http.Request) (Encoder, error)

// Logger represents a function that will be called to add information
// to the logs.
type Logger func(context.Context, string, ...any)

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	*http.ServeMux
	log Logger
	mw  []Middleware
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(log Logger, mw ...Middleware) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		log:      log,
		mw:       mw,
	}
}

// HandleFunc IS MY NEW HANDLER FUNC.
func (app *App) HandleFunc(pattern string, handler Handler, mw ...Middleware) {
	handler = wrapMiddleware(mw, handler)
	handler = wrapMiddleware(app.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := setTraceID(r.Context(), uuid.NewString())

		resp, err := handler(ctx, r)
		if err != nil {
			if err := respondError(ctx, w, err); err != nil {
				app.log(ctx, "respondError", "ERROR", err)
			}
			return
		}

		if err := respond(ctx, w, resp); err != nil {
			app.log(ctx, "respond", "ERROR", err)
		}
	}

	app.ServeMux.HandleFunc(pattern, h)
}
