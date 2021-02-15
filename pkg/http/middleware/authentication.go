package middleware

import (
	"errors"
	"net/http"
	"os"
	"restAPI/pkg/domain/authentication"
	"restAPI/pkg/http/handler"
	"strings"
)

//ValidateRequest ...
func ValidateRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientToken := r.Header.Get("Authorization")
		if clientToken == "" {
			handler.RespondWithError(w, 403, errors.New("No authorizatin header provided"))
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			handler.RespondWithError(w, 400, errors.New("Incorrect Format of Authorization Token"))
			return
		}

		jwtWrapper := authentication.JwtWrapper{
			SecretKey: os.Getenv("SECRET_KEY"),
			Issuer:    "AuthService",
		}

		_, err := jwtWrapper.ValidateJWT(clientToken)

		if err != nil {
			handler.RespondWithError(w, 401, err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
