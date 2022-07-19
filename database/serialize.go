package database

import (
	"encoding/json"
)

func deserializeUser(b []byte, v interface{}) error {
	baseJson, err := json.Marshal(baseUser)

	if err != nil {
		return err
	}

	err = json.Unmarshal(b, v)

	for k, v := range baseUser.(map[string]interface{}) {
		if _, ok := v.(map[string]interface{}); ok {
			continue
		}
	}

	if err != nil {
		return err
	}

	return nil
}
