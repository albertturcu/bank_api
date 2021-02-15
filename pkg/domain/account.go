package domain

import (
	"restAPI/pkg/storage/mysql/entity"
)

func (s *dService) CreateAccount(account entity.Account) (entity.Account, error) {
	return s.r.CreateAccount(account)
}

func (s *dService) WithdrawMoney(accountNo string, amount float32) (entity.Account, error) {
	return s.r.WithdrawMoney(accountNo, amount)
}

func (s *dService) DepositMoney(accountNo string, amount float32) (entity.Account, error) {
	return s.r.DepositMoney(accountNo, amount)
}

func (s *dService) CloseAccount(accountNo string) error {
	return s.r.CloseAccount(accountNo)
}

func (s *dService) GetAllAccounts() []entity.Account {
	return s.r.GetAllAccounts()
}
