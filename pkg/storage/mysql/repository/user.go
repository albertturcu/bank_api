package repository

import (
	"restAPI/pkg/storage/mysql/entity"
)

//User ...
type User interface {
	GetUser(id string) (entity.User, error)
	GetUsers() ([]entity.User, error)
	AddUser(user entity.User) (entity.User, error)
	DeleteUser(id string) error
	UpdateUser(user entity.User) error
	GetUserByEmail(email string) (entity.User, error)
}
