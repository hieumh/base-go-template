package entities

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	Id    uuid.UUID
	Name  string
	Price float64
	Email string
}

func (user *User) validate() error {
	if user.Name == "" || user.Price <= 0 {
		return errors.New("Invalid user details")
	}

	return nil
}

func NewUser(name string, price float64, email string) *User {
	return &User{
		Id:    uuid.New(),
		Name:  name,
		Price: price,
		Email: email,
	}
}
