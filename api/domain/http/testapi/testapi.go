package testapi

import (
	"encoding/json"
	"net/http"
)

func testAPI(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)
}
