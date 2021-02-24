package handler

import (
	"errors"
	"net/http"
	"os"
	"restAPI/pkg/domain/authentication"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

//ContextKey ...
type ContextKey string

const contextTokenKey ContextKey = "tokenString"

//TokenHandler ...
type TokenHandler interface {
	SetRefreshToken() func(w http.ResponseWriter, r *http.Request)
	DeleteRefreshToken() func(w http.ResponseWriter, r *http.Request)
	GetToken() func(w http.ResponseWriter, r *http.Request)
}

func (h *appHandler) SetRefreshToken() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		RespondWithJSON(w, 200, "Success")
	}
}
func (h *appHandler) DeleteRefreshToken() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		RespondWithJSON(w, 200, "Success")
	}
}

func (h *appHandler) GetToken() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtWrapper := authentication.JwtWrapper{
			AccessUUID:              uuid.NewV4().String(),
			AccessSecretKey:         os.Getenv("ACCESS_SECRET_KEY"),
			AccessExpirationMinutes: 15,
			RefreshUUID:             uuid.NewV4().String(),
			RefreshSecretKey:        os.Getenv("REFRESH_SECRET_KEY"),
			RefreshExpirationHours:  7 * 24,
			Issuer:                  "AuthService",
		}
		token := r.Header.Get("Authorization")
		if token == "" {
			RespondWithError(w, 403, errors.New("No authorizatin header provided"))
			return
		}

		extractedToken := strings.Split(token, "Bearer ")

		if len(extractedToken) == 2 {
			token = strings.TrimSpace(extractedToken[1])
		} else {
			RespondWithError(w, 400, errors.New("Incorrect Format of Authorization Token"))
			return
		}

		refreshClaims, err := jwtWrapper.ValidateRefreshJWT(token)
		if err != nil {
			RespondWithError(w, 400, err)
			return
		}
		user, err := h.s.GetUser(strconv.FormatUint(uint64(refreshClaims.ID), 10))
		if _, err := h.s.GetToken(refreshClaims.RefreshUUID); err != nil {
			RespondWithError(w, 400, err)
			return
		}

		ok := h.s.DeleteToken(strconv.FormatUint(uint64(user.ID), 10), refreshClaims.RefreshUUID)
		if ok != nil {
			RespondWithError(w, 400, ok)
			return
		}

		signedTokenPair, err := jwtWrapper.GenerateTokenPair(user)

		if err != nil {
			RespondWithError(w, 400, err)
			return
		}

		ok = h.s.SetToken(jwtWrapper.RefreshUUID, strconv.FormatUint(uint64(user.ID), 10), time.Duration(jwtWrapper.RefreshExpirationHours)*time.Hour)
		if ok != nil {
			RespondWithError(w, 400, ok)
			return
		}

		ok = h.s.SetToken(jwtWrapper.AccessUUID, strconv.FormatUint(uint64(user.ID), 10), time.Duration(jwtWrapper.AccessExpirationMinutes)*time.Minute)
		if ok != nil {
			RespondWithError(w, 400, ok)
			return
		}

		RespondWithJSON(w, 200, signedTokenPair)
	}
}
