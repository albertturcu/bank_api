package repository

import "time"

//Token ...
type Token interface {
	GetToken(userID string) (string, error)
	SetToken(userID string, tokenID string, expiresIn time.Duration) error
	DeleteToken(userID string, previousTokenID string) error
}
