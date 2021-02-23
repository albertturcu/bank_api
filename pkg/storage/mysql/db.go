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
		os.Getenv("MYSQL_ROOT_USERNAME"),
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	sqlDB, err := sql.Open(os.Getenv("MYSQL_DRIVER"), connectionString)
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
