package repository

import (
	"restAPI/pkg/storage/mysql/entity"

	"github.com/jinzhu/gorm"
)

//AccountType ...
type AccountType interface {
	CreateAccountType(accountType entity.AccountType) error
	GetAllAccountTypes() ([]entity.AccountType, error)
}

type accountType struct {
	db *gorm.DB
}

//NewAccountTypeRepository ...
func NewAccountTypeRepository(db *gorm.DB) AccountType {
	return &accountType{db}
}

func (a *accountType) CreateAccountType(accountType entity.AccountType) error {
	result := a.db.Create(&accountType)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *accountType) GetAllAccountTypes() ([]entity.AccountType, error) {
	var accountTypes []entity.AccountType
	result := a.db.Find(&accountTypes)
	if result.Error != nil {
		return accountTypes, result.Error
	}
	return accountTypes, nil
}
