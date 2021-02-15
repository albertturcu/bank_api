package entity

import "gorm.io/gorm"

//User ...
type User struct {
	gorm.Model
	Name     string `gorm:"not null;check:name <> ''"`
	Email    string `gorm:"unique;not null;check:email <> ''"`
	Phone    string
	Address  string
	Accounts []Account `gorm:"foreignKey:user_id;reference:id"`
}
