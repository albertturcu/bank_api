package mysql

import (
	"fmt"
	"os"
	"restAPI/pkg/storage/mysql/entity"

	"github.com/jinzhu/gorm"
)

//NewConnection ...
func NewConnection() (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_NAME"),
	)

	db, err := gorm.Open(os.Getenv("APP_DB_DRIVER"), connectionString)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	db.AutoMigrate(&entity.User{}, &entity.AccountType{}, &entity.Account{})

	return db, nil
}
