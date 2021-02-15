package mysql

import "restAPI/pkg/storage/mysql/entity"

func (s *dbRepository) CreateAccount(account entity.Account) (entity.Account, error) {
	result := s.db.Create(&account)
	if result.Error != nil {
		return entity.Account{}, result.Error
	}
	s.db.Preload("AccountType").Find(&account)

	return account, nil
}

func (s *dbRepository) GetAllAccounts() []entity.Account {
	accounts := []entity.Account{}
	result := s.db.Preload("AccountType").Find(&accounts)

	if result.Error != nil {
		return accounts
	}
	return accounts
}

func (s *dbRepository) WithdrawMoney(accountNo string, amount float32) (entity.Account, error) {
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

func (s *dbRepository) DepositMoney(accountNo string, amount float32) (entity.Account, error) {
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

func (s *dbRepository) CloseAccount(accountNo string) error { return nil }
