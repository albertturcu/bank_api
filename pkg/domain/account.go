package domain

import (
	"restAPI/pkg/storage/mysql/entity"
	"restAPI/pkg/storage/mysql/repository"
)

//AccountService ...
type AccountService interface {
	WithdrawMoney(accountNo string, amount float32) (entity.Account, error)
	DepositMoney(accountNo string, amount float32) (entity.Account, error)
	CloseAccount(accountNo string) error
	CreateAccount(account entity.Account) (entity.Account, error)
	GetAllAccounts() []entity.Account
}

type accountService struct {
	r repository.Account
}

//NewAccountService ...
func NewAccountService(r repository.Account) AccountService {
	return &accountService{r: r}
}

func (s *accountService) CreateAccount(account entity.Account) (entity.Account, error) {
	return s.r.CreateAccount(account)
}

func (s *accountService) WithdrawMoney(accountNo string, amount float32) (entity.Account, error) {
	return s.r.WithdrawMoney(accountNo, amount)
}

func (s *accountService) DepositMoney(accountNo string, amount float32) (entity.Account, error) {
	return s.r.DepositMoney(accountNo, amount)
}

func (s *accountService) CloseAccount(accountNo string) error {
	return s.r.CloseAccount(accountNo)
}

func (s *accountService) GetAllAccounts() []entity.Account {
	return s.r.GetAllAccounts()
}
