package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"restAPI/pkg/storage/mysql/entity"
	"restAPI/pkg/storage/mysql/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DBRepository ...
type DBRepository interface {
	repository.User
	repository.Account
	repository.AccountType
}

//Repositories ...
type dbRepository struct {
	db *gorm.DB
}

//NewRepositories ...
func NewRepositories() (DBRepository, error) {
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

	db.AutoMigrate(&entity.User{}, &entity.AccountType{}, &entity.Account{})

	return &dbRepository{
		db: db,
	}, nil
}

// //NewConnection ...
// func NewConnection() (*gorm.DB, error) {
// 	connectionString := fmt.Sprintf(
// 		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
// 		os.Getenv("APP_DB_USERNAME"),
// 		os.Getenv("APP_DB_PASSWORD"),
// 		os.Getenv("APP_DB_HOST"),
// 		os.Getenv("APP_DB_PORT"),
// 		os.Getenv("APP_DB_NAME"),
// 	)

// 	sqlDB, err := sql.Open("mysql", connectionString)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	db, err := gorm.Open(mysql.New(mysql.Config{
// 		Conn: sqlDB,
// 	}), &gorm.Config{})

// 	if err != nil {
// 		return nil, err
// 	}

// 	db.AutoMigrate(&entity.User{}, &entity.AccountType{}, &entity.Account{})

// 	return db, nil
// }
