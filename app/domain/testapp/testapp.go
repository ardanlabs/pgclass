package testapp

import "fmt"

func TestPost(in MessageIn) MessageOut {
	return MessageOut{
		Value: fmt.Sprintf("ECHO: %s", in.Value),
	}
}
