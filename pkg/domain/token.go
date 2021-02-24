package domain

import "time"

func (s *dService) SetToken(userID string, tokenID string, expiresIn time.Duration) error {
	return s.rdb.SetToken(userID, tokenID, expiresIn)
}

func (s *dService) DeleteToken(userID string, previousTokenID string) error {
	return s.rdb.DeleteToken(userID, previousTokenID)
}

func (s *dService) GetToken(uuid string) (string, error) {
	return s.rdb.GetToken(uuid)
}
