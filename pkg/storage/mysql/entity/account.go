package entity

import "gorm.io/gorm"

//Account ...
type Account struct {
	gorm.Model
	AccountNumber string `gorm:"unique"`
	Balance       float32
	UserID        uint
	AccountTypeID int         `gorm:"not null;unique"`
	AccountType   AccountType `gorm:"foreignKey:id;reference:account_type_id"`
}
