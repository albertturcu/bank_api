package repository

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"
// 	"testing"

// 	_ "github.com/go-sql-driver/mysql"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var testUserRepository User

// func TestMain(m *testing.M) {
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
// 		log.Fatal(err)
// 	}
// 	testUserRepository = NewUserRepository(db)

// 	os.Exit(m.Run())
// }
