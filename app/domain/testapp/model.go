package testapp

import "encoding/json"

type MessageIn struct {
	Value string `json:"value"`
}

// Decode implments the decoder interface.
func (app *MessageIn) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

type MessageOut struct {
	Value string `json:"value"`
}

// Encode implments the encoder interface.
func (app MessageOut) Encode() ([]byte, error) {
	return json.Marshal(app)
}
