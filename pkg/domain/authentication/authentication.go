package authentication

import (
	"errors"
	"restAPI/pkg/storage/mysql/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	AccessUUID              string
	AccessSecretKey         string
	AccessExpirationMinutes int64
	RefreshUUID             string
	RefreshSecretKey        string
	RefreshExpirationHours  int64
	Issuer                  string
}

// JwtAccessClaim adds email as a claim to the token
type JwtAccessClaim struct {
	ID         uint
	AccessUUID string
	Email      string
	jwt.StandardClaims
}

//JwtRefreshClaim ...
type JwtRefreshClaim struct {
	ID          uint
	RefreshUUID string
	jwt.StandardClaims
}

//ValidateAccessJWT ...
func (j *JwtWrapper) ValidateAccessJWT(signedToken string) (*JwtAccessClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtAccessClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.AccessSecretKey), nil
		},
	)
	if err != nil {
		return &JwtAccessClaim{}, err
	}
	claims, ok := token.Claims.(*JwtAccessClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return claims, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return claims, err
	}
	return claims, nil
}

//ValidateRefreshJWT ...
func (j *JwtWrapper) ValidateRefreshJWT(signedToken string) (*JwtRefreshClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtRefreshClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.RefreshSecretKey), nil
		},
	)
	if err != nil {
		return &JwtRefreshClaim{}, err
	}
	claims, ok := token.Claims.(*JwtRefreshClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return claims, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return claims, err
	}
	return claims, nil
}

//GenerateTokenPair ...
func (j *JwtWrapper) GenerateTokenPair(user entity.User) (map[string]string, error) {
	accessClaims := &JwtAccessClaim{
		ID:         user.ID,
		AccessUUID: j.AccessUUID,
		Email:      user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(j.AccessExpirationMinutes)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	refreshClaims := &JwtRefreshClaim{
		ID:          user.ID,
		RefreshUUID: j.RefreshUUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.RefreshExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	accessToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), accessClaims)
	signedAccessToken, err := accessToken.SignedString([]byte(j.AccessSecretKey))

	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(j.RefreshSecretKey))

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  signedAccessToken,
		"refresh_token": signedRefreshToken}, err
}
