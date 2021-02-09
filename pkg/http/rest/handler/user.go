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

func getUser(u domain.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		user, err := u.GetUser(id)
		if err != nil {
			RespondWithError(w, 404, errors.New("the requested user does not exist"))
		}
		RespondWithJSON(w, 200, user)
	}
}

func getUsers(u domain.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := u.GetUsers()
		if err != nil {
			RespondWithError(w, 400, err)
		}
		RespondWithJSON(w, 200, users)
	}
}

func addUser(u domain.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user entity.User

		if json.NewDecoder(r.Body).Decode(&user) != nil {
			RespondWithError(w, 404, errors.New("Bad request"))
		}
		res := u.AddUser(user)
		if res != nil {
			RespondWithError(w, 404, res)
		}
	}
}

func deleteUser(d domain.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Deleting endpoint")
	}
}
func updateUser(u domain.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Updating endpoint")
	}
}
