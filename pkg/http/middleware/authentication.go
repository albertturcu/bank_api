package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"restAPI/pkg/domain"
	"restAPI/pkg/domain/authentication"
	"restAPI/pkg/http/handler"
	"strings"
)

//Middleware ...
type Middleware interface {
	ValidateRequest(next http.Handler) http.Handler
	ExtractToken(next http.Handler) http.Handler
}

type middleware struct {
	s domain.DService
}

//NewMiddleware ...
func NewMiddleware(s domain.DService) Middleware {
	return &middleware{s: s}
}

//ContextKey ...
type ContextKey string

const contextTokenKey ContextKey = "tokenString"

//ValidateRequest ...
func (m *middleware) ValidateRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtWrapper := authentication.JwtWrapper{
			AccessSecretKey: os.Getenv("ACCESS_SECRET_KEY"),
			Issuer:          "AuthService",
		}
		token := r.Context().Value(contextTokenKey)
		tokenString, ok := token.(string)
		if !ok {
			handler.RespondWithError(w, 401, errors.New("Invalid token type"))
			return
		}

		tokenClaims, err := jwtWrapper.ValidateAccessJWT(tokenString)

		if err != nil {
			handler.RespondWithError(w, 401, err)
			return
		}

		if _, ok := m.s.GetToken(tokenClaims.AccessUUID); ok != nil {
			handler.RespondWithError(w, 401, ok)
			return
		}
		next.ServeHTTP(w, r)
	})
}

//ExtractToken ...
func (m *middleware) ExtractToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			handler.RespondWithError(w, 403, errors.New("No authorizatin header provided"))
			return
		}

		extractedToken := strings.Split(accessToken, "Bearer ")

		if len(extractedToken) == 2 {
			accessToken = strings.TrimSpace(extractedToken[1])
			ctx = context.WithValue(r.Context(), contextTokenKey, accessToken)
		} else {
			handler.RespondWithError(w, 400, errors.New("Incorrect Format of Authorization Token"))
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
