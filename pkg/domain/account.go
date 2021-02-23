package domain

import (
	"restAPI/pkg/storage/mysql/entity"
)

func (s *dService) CreateAccount(account entity.Account) (entity.Account, error) {
	return s.mr.CreateAccount(account)
}

func (s *dService) WithdrawMoney(accountNo string, amount float32) (entity.Account, error) {
	return s.mr.WithdrawMoney(accountNo, amount)
}

func (s *dService) DepositMoney(accountNo string, amount float32) (entity.Account, error) {
	return s.mr.DepositMoney(accountNo, amount)
}

func (s *dService) CloseAccount(accountNo string) error {
	return s.mr.CloseAccount(accountNo)
}

func (s *dService) GetAllAccounts() []entity.Account {
	return s.mr.GetAllAccounts()
}
