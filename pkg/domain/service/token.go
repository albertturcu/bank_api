package service

import "time"

//Token ...
type Token interface {
	SetToken(userID string, tokenID string, expiresIn time.Duration) error
	DeleteToken(userID string, previousTokenID string) error
	GetToken(uuid string) (string, error)
}
