package environment

import (
	"os"
	"strconv"
)

func GetInt(key string) (int, error) {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return 0, err
	}
	return value, nil
}

func GetIntWithDefault(key string, defaultValue int) int {
	value, err := GetInt(key)
	if err != nil {
		return defaultValue
	}
	return value
}
