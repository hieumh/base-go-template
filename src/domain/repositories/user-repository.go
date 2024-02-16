package repositories

import (
	"base-go-template/src/domain/entities"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *entities.ValidatedUser) error
	FindById(id uuid.UUID) (*entities.ValidatedUser, error)
	FindAll() ([]*entities.ValidatedUser, error)
	Update(product *entities.ValidatedUser) error
	Delete(id uuid.UUID) error
}
