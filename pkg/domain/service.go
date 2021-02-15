package domain

import (
	"restAPI/pkg/domain/service"
	"restAPI/pkg/storage/mysql"
)

//DService ...
type DService interface {
	service.User
	service.Account
	service.AccountType
}

type dService struct {
	r mysql.DBRepository
}

//NewService ...
func NewService(r mysql.DBRepository) DService {
	return &dService{r: r}
}
