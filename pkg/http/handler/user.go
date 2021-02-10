package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"restAPI/pkg/domain"
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
}

type userHandler struct {
	UserService domain.UserService
}

//NewUserHandler ...
func NewUserHandler(u domain.UserService) UserHandler {
	return &userHandler{UserService: u}
}

func (h *userHandler) GetUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		user, err := h.UserService.GetUser(id)
		if err != nil {
			RespondWithError(w, 404, errors.New("the requested user does not exist"))
		}
		RespondWithJSON(w, 200, user)
	}
}

func (h *userHandler) GetUsers() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := h.UserService.GetUsers()
		if err != nil {
			RespondWithError(w, 400, err)
		}
		RespondWithJSON(w, 200, users)
	}
}

func (h *userHandler) AddUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user entity.User

		if json.NewDecoder(r.Body).Decode(&user) != nil {
			RespondWithError(w, 404, errors.New("Bad request"))
		}
		res := h.UserService.AddUser(user)
		if res != nil {
			RespondWithError(w, 404, res)
		}
	}
}

func (h *userHandler) DeleteUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Deleting endpoint")
	}
}

func (h *userHandler) UpdateUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Updating endpoint")
	}
}
