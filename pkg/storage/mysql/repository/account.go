package repository

import (
	"restAPI/pkg/storage/mysql/entity"

	"gorm.io/gorm"
)

//Account ...
type Account interface {
	WithdrawMoney(string, float32) (entity.Account, error)
	DepositMoney(string, float32) (entity.Account, error)
	CloseAccount(string) error
	CreateAccount(entity.Account) (entity.Account, error)
	GetAllAccounts() []entity.Account
}

type account struct {
	db *gorm.DB
}

//NewAccountRepository ...
func NewAccountRepository(db *gorm.DB) Account {
	return &account{db}
}

func (s *account) CreateAccount(account entity.Account) (entity.Account, error) {
	result := s.db.Create(&account)
	if result.Error != nil {
		return entity.Account{}, result.Error
	}
	s.db.Preload("AccountType").Find(&account)

	return account, nil
}

func (s *account) GetAllAccounts() []entity.Account {
	accounts := []entity.Account{}
	result := s.db.Preload("AccountType").Find(&accounts)

	if result.Error != nil {
		return accounts
	}
	return accounts
}

func (s *account) WithdrawMoney(accountNo string, amount float32) (entity.Account, error) {
	account := entity.Account{}
	tx := s.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return account, err
	}

	if err := tx.Preload("AccountType").Where("account_number = ?", accountNo).Find(&account).Error; err != nil {
		tx.Rollback()
		return account, err
	}
	account.Balance -= amount

	if err := tx.Save(&account).Error; err != nil {
		tx.Rollback()
		return account, err
	}

	return account, tx.Commit().Error
}

func (s *account) DepositMoney(accountNo string, amount float32) (entity.Account, error) {
	account := entity.Account{}
	tx := s.db.Begin()

	if err := tx.Preload("AccountType").Where("account_number = ?", accountNo).Find(&account).Error; err != nil {
		tx.Rollback()
		return account, err
	}
	account.Balance += amount

	if err := tx.Save(&account).Error; err != nil {
		tx.Rollback()
		return account, err
	}

	return account, tx.Commit().Error
}

func (s *account) CloseAccount(accountNo string) error { return nil }
