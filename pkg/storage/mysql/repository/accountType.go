package repository

import "restAPI/pkg/storage/mysql/entity"

//AccountType ...
type AccountType interface {
	CreateAccountType(accountType entity.AccountType) error
	GetAllAccountTypes() ([]entity.AccountType, error)
}
