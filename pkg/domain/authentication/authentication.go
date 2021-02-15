package authentication

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// JwtClaim adds email as a claim to the token
type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

//ValidateJWT ...
func (j *JwtWrapper) ValidateJWT(signedToken string) (*JwtClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return &JwtClaim{}, err
	}

	claims, ok := token.Claims.(*JwtClaim)

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

//GenerateJWT ...
func (j *JwtWrapper) GenerateJWT(email string) (string, error) {
	// Do stuff here
	claims := &JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		return signedToken, err
	}
	return signedToken, err
}
