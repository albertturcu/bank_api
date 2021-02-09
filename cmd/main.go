package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"restAPI/pkg/domain"
	"restAPI/pkg/http/rest/handler"
	"restAPI/pkg/storage/mysql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	var wg sync.WaitGroup

	s, err := mysql.NewRepositories()
	if err != nil {
		panic(err)
	}

	u := domain.NewUserService(s)
	a := domain.NewAccountService(s)
	at := domain.NewAccountTypeService(s)

	r := handler.NewRouter(u, a, at, &wg)

	fmt.Printf("Server is running on port %s...\n", os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), r))
}
