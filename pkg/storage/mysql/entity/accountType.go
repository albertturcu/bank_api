package entity

import "github.com/jinzhu/gorm"

//AccountType ...
type AccountType struct {
	gorm.Model
	Type string `gorm:"unique;not null"`
}
