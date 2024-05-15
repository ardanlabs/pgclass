package all

import (
	"net/http"

	"github.com/ardanlabs/service/api/domain/http/testapi"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(mux *http.ServeMux) {
	testapi.Routes(mux)
}
