package entity

import "gorm.io/gorm"

//User ...
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique;not null"`
	Phone    string
	Address  string
	Accounts []Account `gorm:"foreignKey:user_id;reference:id"`
}
