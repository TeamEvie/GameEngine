package database

import "time"

type User struct {
	Balance           int64     `json:"balance"`
	LastRedeemedDaily time.Time `json:"last_redeemed_daily"`
}

type users struct {
	GetUser    func(userId string) (User, error)
	SetUser    func(userId string, user User) (User, error)
	AddBalance func(userId string, amount int64, from string) (User, error)
}

var creationDate = time.Date(2021, 2, 6, 9, 28, 4, 0, time.UTC)

var baseUser = User{
	Balance:           10,
	LastRedeemedDaily: creationDate,
}
