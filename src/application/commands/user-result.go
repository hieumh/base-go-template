package command

import "github.com/google/uuid"

type UserResult struct {
	Id    uuid.UUID
	Name  string
	Price float64
	Email string
}
