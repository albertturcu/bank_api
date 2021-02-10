package repository

import (
	"restAPI/pkg/storage/mysql/entity"

	"github.com/jinzhu/gorm"
)

//User ...
type User interface {
	GetUser(id string) (entity.User, error)
	GetUsers() ([]entity.User, error)
	AddUser(user entity.User) error
	DeleteUser(id string) error
	UpdateUser(user entity.User) error
}

type user struct {
	db *gorm.DB
}

//NewUserRepository ...
func NewUserRepository(db *gorm.DB) User {
	return &user{db}
}

//GetUser ...
func (s *user) GetUser(id string) (entity.User, error) {
	user := entity.User{}
	result := s.db.Preload("Accounts").Preload("Accounts.AccountType").Find(&user)

	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

//GetUsers ...
func (s *user) GetUsers() ([]entity.User, error) {
	users := []entity.User{}
	result := s.db.Preload("Accounts").Preload("Accounts.AccountType").Find(&users)

	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

//AddUser ...
func (s *user) AddUser(user entity.User) error {
	result := s.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

//DeleteUser ...
func (s *user) DeleteUser(id string) error {
	if err := s.db.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ...
func (s *user) UpdateUser(user entity.User) error {

	err := s.db.Model(&user).Where("id=?", user.ID).Updates(entity.User{Name: user.Name, Email: user.Email, Phone: user.Phone, Address: user.Address}).Error
	if err != nil {
		return err
	}
	return nil
}
