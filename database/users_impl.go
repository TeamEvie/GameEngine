package database

import "github.com/go-redis/redis/v9"

func createUser(userId string) (User, error) {
	c := getClient()
	json, err := serialize(baseUser)

	if err != nil {
		return User{}, err
	}

	u := c.Set(ctx, userId, json, 0)

	if u.Err() != nil {
		return User{}, u.Err()
	}

	return baseUser, nil
}

func getUser(userId string) (User, error) {
	c := getClient()
	json := c.Get(ctx, userId)

	if json.Err() != nil {
		if json.Err() == redis.Nil {
			return createUser(userId)
		}
		return User{}, json.Err()
	}

	bytes, err := json.Bytes()

	if err != nil {
		return User{}, err
	}

	user := User{}

	err = deserialize(bytes, &user)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
