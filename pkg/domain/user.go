package domain

import (
	"restAPI/pkg/storage/mysql/entity"
	"restAPI/pkg/storage/mysql/repository"
)

//UserService ...
type UserService interface {
	AddUser(entity.User) error
	GetUser(string) (entity.User, error)
	GetUsers() ([]entity.User, error)
	DeleteUser(string) error
	UpdateUser(entity.User) error
}

type userService struct {
	r repository.User
}

//NewUserService ...
func NewUserService(r repository.User) UserService {
	return &userService{
		r: r}
}

func (s *userService) AddUser(user entity.User) error {
	return s.r.AddUser(user)
}

func (s *userService) GetUser(id string) (entity.User, error) {
	return s.r.GetUser(id)
}

func (s *userService) GetUsers() ([]entity.User, error) {
	return s.r.GetUsers()
}

func (s *userService) UpdateUser(user entity.User) error {
	return s.r.UpdateUser(user)
}

func (s *userService) DeleteUser(id string) error {
	return s.r.DeleteUser(id)
}
