package domain

import (
	"restAPI/pkg/storage/mysql/entity"
)

func (s *dService) AddUser(user entity.User) (entity.User, error) {
	return s.mr.AddUser(user)
}

func (s *dService) GetUser(id string) (entity.User, error) {
	return s.mr.GetUser(id)
}

func (s *dService) GetUsers() ([]entity.User, error) {
	return s.mr.GetUsers()
}

func (s *dService) UpdateUser(user entity.User) error {
	return s.mr.UpdateUser(user)
}

func (s *dService) DeleteUser(id string) error {
	return s.mr.DeleteUser(id)
}

func (s *dService) GetUserByEmail(email string) (entity.User, error) {
	return s.mr.GetUserByEmail(email)
}
