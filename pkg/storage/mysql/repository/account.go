package repository

import (
	"restAPI/pkg/storage/mysql/entity"
)

//Account ...
type Account interface {
	WithdrawMoney(string, float32) (entity.Account, error)
	DepositMoney(string, float32) (entity.Account, error)
	CloseAccount(string) error
	CreateAccount(entity.Account) (entity.Account, error)
	GetAllAccounts() []entity.Account
}
