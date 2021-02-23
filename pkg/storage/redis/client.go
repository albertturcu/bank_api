package redis

import (
	"fmt"
	"log"
	"os"
	"restAPI/pkg/storage/redis/repository"

	"github.com/go-redis/redis"
)

//RdbRepository ...
type RdbRepository interface {
	repository.RefreshToken
}

type rdbRepository struct {
	rdb *redis.Client
}

//NewRepositories ...
func NewRepositories() (RdbRepository, error) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	log.Printf("Connecting to Redis\n")
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping().Result()
	log.Printf("%v, Connected to Redis\n", pong)
	if err != nil {
		return nil, fmt.Errorf("error connecting to redis: %w", err)
	}
	return &rdbRepository{rdb}, nil
}
