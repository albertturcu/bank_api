package mysql

import (
	"restAPI/pkg/storage/mysql/entity"
)

//GetUser ...
func (s *dbRepository) GetUser(id string) (entity.User, error) {
	user := entity.User{}
	result := s.db.Preload("Accounts").Preload("Accounts.AccountType").Find(&user)

	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

//GetUserByEmail ...
func (s *dbRepository) GetUserByEmail(email string) (entity.User, error) {
	user := entity.User{}

	result := s.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

//GetUsers ...
func (s *dbRepository) GetUsers() ([]entity.User, error) {
	users := []entity.User{}
	result := s.db.Preload("Accounts").Preload("Accounts.AccountType").Find(&users)

	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

//AddUser ...
func (s *dbRepository) AddUser(user entity.User) (entity.User, error) {
	result := s.db.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

//DeleteUser ...
func (s *dbRepository) DeleteUser(id string) error {
	if err := s.db.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ...
func (s *dbRepository) UpdateUser(user entity.User) error {

	err := s.db.Model(&user).Where("id=?", user.ID).Updates(entity.User{Name: user.Name, Email: user.Email, Phone: user.Phone, Address: user.Address}).Error
	if err != nil {
		return err
	}
	return nil
}
