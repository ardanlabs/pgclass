package testapi

import "net/http"

// Routes adds specific routes for this group.
func Routes(mux *http.ServeMux) {
	mux.HandleFunc("POST /test", testPostAPI)
}
