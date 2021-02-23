package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"restAPI/pkg/domain/authentication"
	"restAPI/pkg/http/handler"
	"strings"
)

//ValidateRequest ...
func ValidateRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			handler.RespondWithError(w, 403, errors.New("No authorizatin header provided"))
			return
		}

		extractedToken := strings.Split(accessToken, "Bearer ")

		if len(extractedToken) == 2 {
			accessToken = strings.TrimSpace(extractedToken[1])
		} else {
			handler.RespondWithError(w, 400, errors.New("Incorrect Format of Authorization Token"))
			return
		}

		jwtWrapper := authentication.JwtWrapper{
			SecretKey: os.Getenv("SECRET_KEY"),
			Issuer:    "AuthService",
		}

		accessClaims, err := jwtWrapper.ValidateJWT(accessToken)
		fmt.Println(accessClaims)
		if err != nil {
			// validate refresh token from redis
			// use refresh token to generateTokenPair if it's valid
			// throw error and request the user a new login flow if it's invalid
			// handler.RespondWithError(w, 401, err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
