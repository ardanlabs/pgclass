package all

import (
	"github.com/ardanlabs/service/api/domain/http/testapi"
	"github.com/ardanlabs/service/api/domain/http/userapi"
	"github.com/ardanlabs/service/business/domain/userbus"
	"github.com/ardanlabs/service/business/domain/userbus/stores/userdb"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
	"github.com/jmoiron/sqlx"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (a add) Add(log *logger.Logger, db *sqlx.DB, app *web.App) {
	testapi.Routes(log, app)

	userapi.Routes(app, userapi.Config{
		Log:     log,
		UserBus: userbus.NewBusiness(log, userdb.NewStore(log, db)),
	})
}
