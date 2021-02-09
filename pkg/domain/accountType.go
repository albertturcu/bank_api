package domain

import (
	"restAPI/pkg/storage/mysql/entity"
	"restAPI/pkg/storage/mysql/repository"
)

//AccountTypeService ...
type AccountTypeService interface {
	CreateAccountType(accountType entity.AccountType) error
	GetAllAccountTypes() ([]entity.AccountType, error)
}

type accountTypeService struct {
	r repository.AccountType
}

//NewAccountTypeService ...
func NewAccountTypeService(r repository.AccountType) AccountTypeService {
	return &accountTypeService{r: r}
}

func (a *accountTypeService) CreateAccountType(accountType entity.AccountType) error {
	return a.r.CreateAccountType(accountType)
}

func (a *accountTypeService) GetAllAccountTypes() ([]entity.AccountType, error) {
	return a.r.GetAllAccountTypes()
}
