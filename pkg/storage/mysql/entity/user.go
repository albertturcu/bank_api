package entity

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//User ...
type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Phone    string
	Address  string
	Password string    `gorm:"not null"`
	Accounts []Account `gorm:"foreignKey:user_id;reference:id"`
}

//HashAndSalt ...
func (u *User) HashAndSalt(password string) string {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

//BeforeCreate ...
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	bytePassword := []byte(u.Password)
	fmt.Println(u.Password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

// CheckPassword checks user password
func (u *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
