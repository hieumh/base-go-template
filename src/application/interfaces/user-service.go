package interfaces

import (
	command "base-go-template/src/application/commands"
	"base-go-template/src/domain/entities"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(userCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error)
	GetAllUser() ([]*entities.ValidatedUser, error)
	FindUserById(id uuid.UUID) (*entities.ValidatedUser, error)
}
