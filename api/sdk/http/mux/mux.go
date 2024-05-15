// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"net/http"
)

// RouteAdder defines behavior that sets the routes to bind for an instance
// of the service.
type RouteAdder interface {
	Add(*http.ServeMux)
}

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(routeAdder RouteAdder) *http.ServeMux {
	mux := http.NewServeMux()

	routeAdder.Add(mux)

	return mux
}
