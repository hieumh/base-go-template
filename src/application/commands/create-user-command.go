package command

import "github.com/google/uuid"

type CreateUserCommand struct {
	ID    uuid.UUID
	Name  string
	Price float64
	Email string
}
