package mysql

import "restAPI/pkg/storage/mysql/entity"

func (s *dbRepository) CreateAccountType(accountType entity.AccountType) error {
	result := s.db.Create(&accountType)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *dbRepository) GetAllAccountTypes() ([]entity.AccountType, error) {
	var accountTypes []entity.AccountType
	result := s.db.Find(&accountTypes)
	if result.Error != nil {
		return accountTypes, result.Error
	}
	return accountTypes, nil
}
