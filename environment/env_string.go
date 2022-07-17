package environment

import "os"

func GetStringNullPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("environment variable " + key + " is not set")
	}
	return value
}
