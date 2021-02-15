package service

import "restAPI/pkg/storage/mysql/entity"

//User ...
type User interface {
	AddUser(entity.User) (entity.User, error)
	GetUser(string) (entity.User, error)
	GetUserByEmail(string) (entity.User, error)
	GetUsers() ([]entity.User, error)
	DeleteUser(string) error
	UpdateUser(entity.User) error
}
