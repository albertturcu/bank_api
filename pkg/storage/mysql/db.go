package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"restAPI/pkg/storage/mysql/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	sqlDB, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println(err)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	// db.LogMode(true)

	db.AutoMigrate(&entity.User{}, &entity.AccountType{}, &entity.Account{})
	// db.Model(&entity.User{}).AddForeignKey("id", "accounts(account_number)", "RESTRICT", "RESTRICT")
	// var users []entity.User
	// db.Preload(clause.Associations).Find(&users)
	// fmt.Println(users[0])
	// db.Model(&entity.Account{}).AddForeignKey("account_type_id", "accounts(id)", "RESTRICT", "RESTRICT")

	return db, nil
}
