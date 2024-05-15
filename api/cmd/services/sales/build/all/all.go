package all

import (
	"github.com/ardanlabs/service/api/domain/http/testapi"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(log *logger.Logger, app *web.App) {
	testapi.Routes(log, app)
}
