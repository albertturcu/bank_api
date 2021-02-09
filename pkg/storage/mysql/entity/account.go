package entity

import "github.com/jinzhu/gorm"

//Account ...
type Account struct {
	gorm.Model
	AccountNumber string `gorm:"unique"`
	Balance       float32
	UserID        uint `gorm:"not null"`
	AccountTypeID int  `gorm:"not null;unique"`
	AccountType   AccountType
}
