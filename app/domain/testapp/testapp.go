package testapp

import (
	"fmt"
	"math/rand"

	"github.com/ardanlabs/service/app/sdk/errs"
)

func TestPost(in MessageIn) (MessageOut, error) {
	if n := rand.Intn(100); n%2 == 0 {
		return MessageOut{}, errs.Newf(errs.FailedPrecondition, "My APP Error")
	}

	mo := MessageOut{
		Value: fmt.Sprintf("ECHO: %s", in.Value),
	}

	return mo, nil
}
