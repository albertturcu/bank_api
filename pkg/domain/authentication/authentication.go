package authentication

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey              string
	Issuer                 string
	AccessExpirationHours  int64
	RefreshExpirationHours int64
}

// JwtAccessClaim adds email as a claim to the token
type JwtAccessClaim struct {
	ID    uint
	Email string
	jwt.StandardClaims
}

//JwtRefreshClaim ...
type JwtRefreshClaim struct {
	ID uint
	jwt.StandardClaims
}

//ValidateJWT ...
func (j *JwtWrapper) ValidateJWT(signedToken string) (*JwtAccessClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtAccessClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
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

//GenerateTokenPair ...
func (j *JwtWrapper) GenerateTokenPair(email string, uid uint) (map[string]string, error) {
	accessClaims := &JwtAccessClaim{
		Email: email,
		ID:    uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.AccessExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	refreshClaims := &JwtRefreshClaim{
		ID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.RefreshExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	accessToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), accessClaims)

	signedAccessToken, err := accessToken.SignedString([]byte(j.SecretKey))

	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), refreshClaims)

	signedRefreshToken, err := refreshToken.SignedString([]byte(j.SecretKey))

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  signedAccessToken,
		"refresh_token": signedRefreshToken}, err
}
