package testapi

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/app/domain/testapp"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

type testapi struct {
	Log *logger.Logger
}

func (api *testapi) newTestPostAPI(ctx context.Context, r *http.Request) (web.Encoder, error) {
	var in testapp.MessageIn
	if err := web.Decode(r, &in); err != nil {
		return nil, err
	}

	api.Log.Info(ctx, "READALL", "IN", in)

	out := testapp.TestPost(in)

	return out, nil
}
