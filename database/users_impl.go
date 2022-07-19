package database

import (
	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
)

func createUser(userId string) (User, error) {
	_, err := setUser(userId, baseUser)

	if err != nil {
		return User{}, err
	}

	return baseUser, nil
}

func getUser(userId string) (User, error) {
	c := getClient()
	json, err := c.JSONGet(userId, ".")

	if err != nil {
		if err == redis.Nil {
			return createUser(userId)
		}
		return User{}, err
	}

	user := User{}

	err = deserialize(json.([]byte), &user)

	if err != nil {
		return User{}, err
	}

	color.Blue("[database] Deserialized user data(%s): %+v", userId, user)

	return user, nil
}

func setUser(userId string, user User) (User, error) {
	c := getClient()

	_, err := c.JSONSet(userId, ".", user)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
