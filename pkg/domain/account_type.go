package domain

import (
	"restAPI/pkg/storage/mysql/entity"
)

func (s *dService) CreateAccountType(accountType entity.AccountType) error {
	return s.r.CreateAccountType(accountType)
}

func (s *dService) GetAllAccountTypes() ([]entity.AccountType, error) {
	return s.r.GetAllAccountTypes()
}
