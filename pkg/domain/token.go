package domain

import "time"

func (s *dService) SetRefreshToken(userID string, tokenID string, expiresIn time.Duration) error {
	return s.rdb.SetRefreshToken(userID, tokenID, expiresIn)
}
func (s *dService) DeleteRefreshToken(userID string, previousTokenID string) error {
	return s.rdb.DeleteRefreshToken(userID, previousTokenID)
}

func (s *dService) RefreshToken(userID string) (string, error) {
	return s.rdb.RefreshToken(userID)
}
