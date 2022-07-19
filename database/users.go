package database

type User struct {
	Balance int64 `json:"balance"`
}

type users struct {
	GetUser func(userId string) (User, error)
}

var baseUser = User{
	Balance: 0,
}
