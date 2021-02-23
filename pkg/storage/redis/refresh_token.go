package redis

import (
	"errors"
	"log"
	"time"
)

func (s *rdbRepository) SetRefreshToken(userID string, tokenID string, expiresIn time.Duration) error {
	if err := s.rdb.Set(userID, tokenID, expiresIn).Err(); err != nil {
		log.Printf("Could not SET refresh token to redis for userID/tokenID: %s/%s: %v\n", userID, tokenID, err)
		return errors.New("Could not SET refresh token to redis")
	}
	return nil
}

func (s *rdbRepository) DeleteRefreshToken(userID string, previousTokenID string) error {
	if err := s.rdb.Del(userID).Err(); err != nil {
		log.Printf("Could not DEL refresh token to redis for userID/tokenID: %s/%s: %v\n", userID, previousTokenID, err)
		return errors.New("Could not DEL refresh token to redis")
	}
	return nil
}

func (s *rdbRepository) RefreshToken(userID string) (string, error) {
	refreshToken, err := s.rdb.Get(userID).Result()
	if err != nil {
		log.Printf("Could not GET refresh token to redis for userID/tokenID: %s: %v", userID, err)
		return "", errors.New("Could not GET refresh token to redis")
	}
	return refreshToken, nil
}
