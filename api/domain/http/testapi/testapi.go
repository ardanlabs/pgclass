package testapi

import (
	"io"
	"net/http"

	"github.com/ardanlabs/service/app/domain/testapp"
	"github.com/ardanlabs/service/foundation/logger"
)

type testapi struct {
	Log *logger.Logger
}

func (api *testapi) testPostAPI(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	api.Log.Info(ctx, "READALL", "Data", string(data))

	var in testapp.MessageIn
	if err := in.Decode(data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.Log.Info(ctx, "READALL", "IN", in)

	out := testapp.TestPost(in)

	outBytes, err := out.Encode()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(outBytes)
}
