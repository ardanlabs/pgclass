package testapi

import (
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(log *logger.Logger, app *web.App) {
	api := testapi{
		Log: log,
	}

	app.HandleFunc("POST /test", api.newTestPostAPI)
}
