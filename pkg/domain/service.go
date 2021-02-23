package domain

import (
	"restAPI/pkg/domain/service"
	"restAPI/pkg/storage/mysql"
	"restAPI/pkg/storage/redis"
)

//DService ...
type DService interface {
	service.User
	service.Account
	service.AccountType
	service.Token
}

type dService struct {
	mr  mysql.DBRepository
	rdb redis.RdbRepository
}

//NewService ...
func NewService(mr mysql.DBRepository, rdb redis.RdbRepository) DService {
	return &dService{
		mr:  mr,
		rdb: rdb,
	}
}
