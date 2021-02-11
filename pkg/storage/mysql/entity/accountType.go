package entity

import "gorm.io/gorm"

//AccountType ...
type AccountType struct {
	gorm.Model
	Type string `gorm:"unique;not null"`
}
