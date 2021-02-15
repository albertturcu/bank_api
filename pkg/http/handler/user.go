package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"restAPI/pkg/storage/mysql/entity"
	"unsafe"

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
		fmt.Println(unsafe.Sizeof(user))
		if json.NewDecoder(r.Body).Decode(&user) != nil {
			RespondWithError(w, 404, errors.New("Bad request"))
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
