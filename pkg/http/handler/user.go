package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"restAPI/pkg/domain/authentication"
	"restAPI/pkg/storage/mysql/entity"

	"github.com/gorilla/mux"
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
	Email string `json:"Email"`
	Token string `json:"Auth Bearer"`
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
			RespondWithError(w, 404, errors.New("the requested user does not exist"))
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

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			RespondWithError(w, 400, err)
		}

		expectedUser, err := h.s.GetUserByEmail(payload.Email)
		valid := expectedUser.CheckPassword(payload.Password)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			RespondWithError(w, 400, errors.New("Email not found"))
			return
		}

		if valid != nil {
			w.WriteHeader(http.StatusUnauthorized)
			RespondWithError(w, 400, errors.New("Incorrect Password"))
			return
		}

		jwtWrapper := authentication.JwtWrapper{
			SecretKey:       os.Getenv("SECRET_KEY"),
			Issuer:          "AuthService",
			ExpirationHours: 24,
		}

		signedToken, err := jwtWrapper.GenerateJWT(payload.Email)
		if err != nil {
			fmt.Println(err)
			RespondWithError(w, 400, err)
		}

		loginRespone := LoginResponse{
			Email: payload.Email,
			Token: signedToken,
		}
		RespondWithJSON(w, 200, loginRespone)
	}
}
