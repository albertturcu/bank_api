package mysql

import "restAPI/pkg/storage/mysql/entity"

func (a *dbRepository) CreateAccountType(accountType entity.AccountType) error {
	result := a.db.Create(&accountType)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *dbRepository) GetAllAccountTypes() ([]entity.AccountType, error) {
	var accountTypes []entity.AccountType
	result := a.db.Find(&accountTypes)
	if result.Error != nil {
		return accountTypes, result.Error
	}
	return accountTypes, nil
}
