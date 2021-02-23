package service

import "time"

//Token ...
type Token interface {
	SetRefreshToken(userID string, tokenID string, expiresIn time.Duration) error
	DeleteRefreshToken(userID string, previousTokenID string) error
	RefreshToken(userID string) (string, error)
}
