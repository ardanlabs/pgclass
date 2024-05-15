package testapi

import (
	"net/http"

	"github.com/ardanlabs/service/foundation/logger"
)

// Routes adds specific routes for this group.
func Routes(log *logger.Logger, mux *http.ServeMux) {
	api := testapi{
		Log: log,
	}

	mux.HandleFunc("POST /test", api.testPostAPI)
}
