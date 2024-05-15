package testapi

import (
	"io"
	"net/http"

	"github.com/ardanlabs/service/app/domain/testapp"
)

func testPostAPI(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var in testapp.MessageIn
	in.Decode(data)

	out := testapp.TestPost(in)

	outBytes, err := out.Encode()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(outBytes)
}
