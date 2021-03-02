package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"restAPI/pkg/domain"
	"restAPI/pkg/http/handler/web"
	"restAPI/pkg/http/router"
	"restAPI/pkg/storage/mysql"
	"restAPI/pkg/storage/redis"

	"github.com/joho/godotenv"
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

	rdb, err := redis.NewRepositories()
	if err != nil {
		panic(err)
	}

	s := domain.NewService(mr, rdb)
	w := web.NewWeb(s)
	r := router.NewWebRouter(w)

	fmt.Printf("Web is running on port %s...\n", os.Getenv("WEB_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("WEB_PORT"), r))
}
