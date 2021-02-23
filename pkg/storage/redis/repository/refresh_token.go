package repository

import "time"

//RefreshToken ...
type RefreshToken interface {
	RefreshToken(userID string) (string, error)
	SetRefreshToken(userID string, tokenID string, expiresIn time.Duration) error
	DeleteRefreshToken(userID string, previousTokenID string) error
}
