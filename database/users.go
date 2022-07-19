package database

type User struct {
	Balance int64 `json:"balance"`
}

type users struct {
	GetUser func(userId string) (User, error)
	SetUser func(userId string, user User) (User, error)
}

var baseUser = User{
	Balance: 10,
}
