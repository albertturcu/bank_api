package service

import "restAPI/pkg/storage/mysql/entity"

//Account ...
type Account interface {
	WithdrawMoney(accountNo string, amount float32) (entity.Account, error)
	DepositMoney(accountNo string, amount float32) (entity.Account, error)
	CloseAccount(accountNo string) error
	CreateAccount(account entity.Account) (entity.Account, error)
	GetAllAccounts() []entity.Account
}
