package redis

import "encoding/json"

type Mail struct {
	From string `json:"from"`
	To   string `json:"to"`
	Body string `json:"body"`
}

func (i Mail) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}
