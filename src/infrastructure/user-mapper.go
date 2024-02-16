package infrastructure

import "base-go-template/src/domain/entities"

func ToDBUser(user *entities.ValidatedUser) *User {
	var _user = &User{
		Name:  user.Name,
		Price: user.Price,
		Email: user.Email,
	}

	_user.Id = user.Id
	return _user
}

func FromDBUser(dbUser *User) (*entities.ValidatedUser, error) {
	var user = &entities.User{
		Name:  dbUser.Name,
		Price: dbUser.Price,
		Email: dbUser.Email,
	}
	user.Id = dbUser.Id

	return entities.NewValidatedUser(user)
}
