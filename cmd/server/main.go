package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"restAPI/pkg/domain"
	"restAPI/pkg/http/handler"
	"restAPI/pkg/http/middleware"
	"restAPI/pkg/http/router"
	"restAPI/pkg/storage/mysql"
	"restAPI/pkg/storage/redis"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	_ "gorm.io/driver/mysql"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	c := cors.New(cors.Options{
		AllowedHeaders:   []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowedOrigins:   []string{"http://localhost:8081"},
		AllowCredentials: true,
	})

	mr, err := mysql.NewRepositories()
	if err != nil {
		panic(err)
	}

	rdb, err := redis.NewRepositories()
	if err != nil {
		panic(err)
	}

	s := domain.NewService(mr, rdb)
	m := middleware.NewMiddleware(s)
	h := handler.NewAppHandler(s)
	r := router.NewRouter(h, m)

	fmt.Printf("Server is running on port %s...\n", os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), c.Handler(r)))
}
