package database

import "github.com/disgoorg/disgo/json"

func serialize(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func deserialize(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}
