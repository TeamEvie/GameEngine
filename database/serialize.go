package database

import (
	"encoding/json"
	"reflect"
)

func deserializeUser(b []byte, v interface{}) error {
	err := json.Unmarshal(b, v)

	if err != nil {
		return err
	}

	baseReflect := reflect.ValueOf(baseUser)
	userReflect := reflect.ValueOf(v)

	for i := 0; i < baseReflect.NumField(); i++ {
		fieldName := baseReflect.Type().Field(i).Name
		indirect := reflect.Indirect(userReflect)
		if indirect.FieldByName(fieldName).IsZero() {
			indirect.FieldByName(fieldName).Set(baseReflect.FieldByName(fieldName))
		}
	}

	return nil
}
