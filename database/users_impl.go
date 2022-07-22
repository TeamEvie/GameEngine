package database

import (
	"eviecoin/items"
	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
	"strings"
)

func NewSummonFrom(reason string) string {
	reason = strings.Replace(reason, " ", "_", -1)
	reason = strings.ToLower(reason)
	return "me:" + reason
}

func addItem(userId string, item items.RawItemData) (User, error) {
	user, err := getUser(userId)

	if err != nil {
		return User{}, err
	}

	user.Inventory = append(user.Inventory, item)

	_, err = setUser(userId, user)

	if err != nil {
		return User{}, err
	}

	_ = writeToLedger(newTransaction(userId, 1, item.Id, NewSummonFrom("addItem")))

	return user, nil
}

func addBalance(userId string, amount int64, from string) (User, error) {
	user, err := getUser(userId)

	if err != nil {
		return User{}, err
	}

	user.Balance += amount

	_, err = setUser(userId, user)

	if err != nil {
		return User{}, err
	}

	var typeOfTransaction string
	if amount < 0 {
		typeOfTransaction = "debit"
	} else {
		typeOfTransaction = "credit"
	}

	_ = writeToLedger(newTransaction(userId, amount, typeOfTransaction, from))

	return user, nil
}

func createUser(userId string) (User, error) {
	color.Blue("[database] Creating a new user(%s)", userId)

	_, err := setUser(userId, baseUser)

	if err != nil {
		return User{}, err
	}

	return baseUser, nil
}

func getUser(userId string) (User, error) {
	c := getClient()
	json, err := c.JSONGet(formatUserKey(userId), ".")

	if err != nil {
		if err == redis.Nil {
			return createUser(userId)
		}
		return User{}, err
	}

	user := User{}

	err = deserializeUser(json.([]byte), &user)

	defer func(userId string, user User) {
		color.Blue("[database] Patching user(%s) in case there is missing required fields...", userId)
		_, err := setUser(userId, user)
		if err != nil {
			color.Red("[database] Failed to update user(%s) to redis", userId)
		}
	}(userId, user)

	if err != nil {
		return User{}, err
	}

	color.Blue("[database] Deserialized user data(%s): %+v", userId, user)

	return user, nil
}

func setUser(userId string, user User) (User, error) {
	c := getClient()

	_, err := c.JSONSet(formatUserKey(userId), ".", user)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
