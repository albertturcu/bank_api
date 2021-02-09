package entity

import "github.com/jinzhu/gorm"

//User ...
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique;not null"`
	Phone    string
	Address  string
	Accounts []Account `gorm:"OnDelete:CASCADE"`
}
