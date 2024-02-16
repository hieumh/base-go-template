package request

import command "base-go-template/src/application/commands"

type CreateUserRequest struct {
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
	Email string  `json:"Email"`
}

func (req *CreateUserRequest) ToCreateUserCommand() (*command.CreateUserCommand, error) {
	return &command.CreateUserCommand{
		Name:  req.Name,
		Price: req.Price,
		Email: req.Email,
	}, nil
}
