package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"restAPI/pkg/http/router"
	"restAPI/pkg/interactor"
	"restAPI/pkg/storage/mysql"

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

	s, err := mysql.NewConnection()
	if err != nil {
		panic(err)
	}

	i := interactor.NewInteractor(s)
	h := i.NewAppHandler()

	r := router.NewRouter(h)

	fmt.Printf("Server is running on port %s...\n", os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), r))
}
