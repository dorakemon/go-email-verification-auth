package models

import "errors"

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var (
	Users = []User{
		{
			Name:     "admin",
			Password: "password",
			Email:    "admin@example.com",
		},
		{
			Name:     "testuser",
			Password: "password",
			Email:    "test@example.com",
		},
	}
)

func AddUser(name string, password string, email string) User {
	newUser := User{
		Name:     name,
		Password: password,
		Email:    email,
	}

	Users = append(Users, newUser)

	return newUser
}

func FindUserByEmail(email string) (User, error) {
	for _, user := range Users {
		if user.Email == email {
			return user, nil
		}
	}
	return User{}, errors.New("cannot find user")
}
