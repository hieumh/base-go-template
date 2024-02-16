package infrastructure

import (
	"base-go-template/src/domain/entities"
	"base-go-template/src/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Create(user *entities.ValidatedUser) error {
	dbUser := ToDBUser(user)

	if err := repo.db.Create(dbUser).Error; err != nil {
		return err
	}

	storedUser, err := repo.FindById(user.Id)
	if err != nil {
		return err
	}

	*user = *storedUser

	return nil
}

func (repo *UserRepository) FindById(id uuid.UUID) (*entities.ValidatedUser, error) {
	var dbUser User
	if err := repo.db.Preload("User").First(&dbUser, id).Error; err != nil {
		return nil, err
	}

	return FromDBUser(&dbUser)
}

func (repo *UserRepository) FindAll() ([]*entities.ValidatedUser, error) {
	var results []*entities.ValidatedUser
	if err := repo.db.Find(&User{}).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (repo *UserRepository) Update(user *entities.ValidatedUser) error {
	dbUser := ToDBUser(user)
	err := repo.db.Model(&User{}).Where("id = ?", dbUser.Id).Updates(dbUser).Error
	if err != nil {
		return err
	}

	storedUser, err := repo.FindById(dbUser.Id)
	if err != nil {
		return err
	}

	*user = *storedUser
	return nil
}

func (repo *UserRepository) Delete(id uuid.UUID) error {
	return repo.db.Delete(&User{}, id).Error
}
