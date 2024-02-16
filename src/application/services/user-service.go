package services

import (
	command "base-go-template/src/application/commands"
	"base-go-template/src/application/mapper"
	"base-go-template/src/domain/entities"
	"base-go-template/src/domain/repositories"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (service *UserService) CreateUser(userCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	newUser := entities.NewUser(
		userCommand.Name, userCommand.Price, userCommand.Email,
	)

	validatedUser, err := entities.NewValidatedUser(newUser)
	if err != nil {
		return nil, err
	}

	err = service.userRepository.Create(validatedUser)
	if err != nil {
		return nil, err
	}

	var result command.CreateUserCommandResult
	result.Result = mapper.NewUserResultFromEntity(validatedUser)

	return &result, nil
}

func (service *UserService) GetAllUser() ([]*entities.ValidatedUser, error) {
	return service.userRepository.FindAll()
}

func (service *UserService) FindUserById(id uuid.UUID) (*entities.ValidatedUser, error) {
	return service.userRepository.FindById(id)
}
