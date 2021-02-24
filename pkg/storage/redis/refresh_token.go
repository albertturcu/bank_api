package redis

import (
	"errors"
	"log"
	"time"
)

func (s *rdbRepository) SetToken(tokenID string, userID string, expiresIn time.Duration) error {
	if err := s.rdb.Set(tokenID, userID, expiresIn).Err(); err != nil {
		log.Printf("Could not SET refresh token to redis for userID/tokenID: %s/%s: %v\n", userID, tokenID, err)
		return errors.New("Could not SET refresh token to redis")
	}
	return nil
}

func (s *rdbRepository) DeleteToken(userID string, previousTokenID string) error {
	if err := s.rdb.Del(previousTokenID).Err(); err != nil {
		log.Printf("Could not DEL refresh token to redis for userID/tokenID: %s/%s: %v\n", userID, previousTokenID, err)
		return errors.New("Could not DEL refresh token to redis")
	}
	return nil
}

func (s *rdbRepository) GetToken(uuid string) (string, error) {
	userID, err := s.rdb.Get(uuid).Result()
	if err != nil {
		log.Printf("Could not GET refresh token to redis for userID/tokenID: %s: %v", userID, err)
		return "", errors.New("Could not GET refresh token to redis")
	}
	return userID, nil
}
