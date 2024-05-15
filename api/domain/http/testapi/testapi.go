package testapi

import (
	"context"
	"io"
	"net/http"

	"github.com/ardanlabs/service/app/domain/testapp"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

type testapi struct {
	Log *logger.Logger
}

func (api *testapi) newTestPostAPI(ctx context.Context, r *http.Request) (web.Encoder, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	api.Log.Info(ctx, "READALL", "Data", string(data))

	var in testapp.MessageIn
	if err := in.Decode(data); err != nil {
		return nil, err
	}

	api.Log.Info(ctx, "READALL", "IN", in)

	out := testapp.TestPost(in)

	return out, nil
}
