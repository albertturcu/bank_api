package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"restAPI/pkg/domain/authentication"
	"restAPI/pkg/storage/mysql/entity"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

//UserHandler ...
type UserHandler interface {
	AddUser() func(w http.ResponseWriter, r *http.Request)
	GetUser() func(w http.ResponseWriter, r *http.Request)
	GetUsers() func(w http.ResponseWriter, r *http.Request)
	DeleteUser() func(w http.ResponseWriter, r *http.Request)
	UpdateUser() func(w http.ResponseWriter, r *http.Request)
	Login() func(w http.ResponseWriter, r *http.Request)
}

//LoginResponse ...
type LoginResponse struct {
	Email        string `json:"Email"`
	AccessToken  string `json:"Access Token"`
	RefreshToken string `json:"Refresh Token"`
}

//LoginPayload ...
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *appHandler) GetUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		user, err := h.s.GetUser(id)
		if err != nil {
			RespondWithError(w, 404, err)
			return
		}
		RespondWithJSON(w, 200, user)
	}
}

func (h *appHandler) GetUsers() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := h.s.GetUsers()
		if err != nil {
			RespondWithError(w, 400, err)
			return
		}
		RespondWithJSON(w, 200, users)
	}
}

func (h *appHandler) AddUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user entity.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			RespondWithError(w, 404, err)
			return
		}
		user, err := h.s.AddUser(user)
		if err != nil {
			RespondWithError(w, 404, err)
			return
		}
		RespondWithJSON(w, 203, user)
	}
}

func (h *appHandler) DeleteUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Deleting endpoint")
	}
}

func (h *appHandler) UpdateUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Updating endpoint")
	}
}

func (h *appHandler) Login() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload LoginPayload

		jwtWrapper := authentication.JwtWrapper{
			AccessUUID:              uuid.NewV4().String(),
			AccessSecretKey:         os.Getenv("ACCESS_SECRET_KEY"),
			AccessExpirationMinutes: 15,
			RefreshUUID:             uuid.NewV4().String(),
			RefreshSecretKey:        os.Getenv("REFRESH_SECRET_KEY"),
			RefreshExpirationHours:  7 * 24,
			Issuer:                  "AuthService",
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			RespondWithError(w, 400, err)
		}

		expectedUser, err := h.s.GetUserByEmail(payload.Email)
		valid := expectedUser.CheckPassword(payload.Password)

		if err != nil {
			RespondWithError(w, 400, err)
			return
		}

		if valid != nil {
			RespondWithError(w, 400, valid)
			return
		}
		signedTokenPair, err := jwtWrapper.GenerateTokenPair(expectedUser)
		if err != nil {
			RespondWithError(w, 400, err)
			return
		}

		loginRespone := LoginResponse{
			Email:        payload.Email,
			AccessToken:  signedTokenPair["access_token"],
			RefreshToken: signedTokenPair["refresh_token"],
		}
		ok := h.s.SetToken(jwtWrapper.RefreshUUID, strconv.FormatUint(uint64(expectedUser.ID), 10), time.Duration(jwtWrapper.RefreshExpirationHours)*time.Hour)
		if ok != nil {
			RespondWithError(w, 400, err)
			return
		}

		ok = h.s.SetToken(jwtWrapper.AccessUUID, strconv.FormatUint(uint64(expectedUser.ID), 10), time.Duration(jwtWrapper.AccessExpirationMinutes)*time.Minute)
		if ok != nil {
			RespondWithError(w, 400, err)
			return
		}

		RespondWithJSON(w, 200, loginRespone)
	}
}
