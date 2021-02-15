package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"restAPI/pkg/domain"
	"restAPI/pkg/http/handler"
	"restAPI/pkg/http/router"
	"restAPI/pkg/storage/mysql"

	"github.com/joho/godotenv"
	_ "gorm.io/driver/mysql"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	mr, err := mysql.NewRepositories()
	if err != nil {
		panic(err)
	}

	s := domain.NewService(mr)
	h := handler.NewAppHandler(s)
	r := router.NewRouter(h)

	fmt.Printf("Server is running on port %s...\n", os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), r))
}
