package mapper

import (
	command "base-go-template/src/application/commands"
	"base-go-template/src/domain/entities"
)

func NewUserResultFromEntity(user *entities.ValidatedUser) command.UserResult {
	return command.UserResult{
		Id:    user.Id,
		Name:  user.Name,
		Price: user.Price,
		Email: user.Email,
	}
}
