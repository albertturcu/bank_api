package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

//TokenHandler ...
type TokenHandler interface {
	SetRefreshToken() func(w http.ResponseWriter, r *http.Request)
	DeleteRefreshToken() func(w http.ResponseWriter, r *http.Request)
	RefreshToken() func(w http.ResponseWriter, r *http.Request)
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

func (h *appHandler) RefreshToken() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		refreshToken, err := h.s.RefreshToken(id)

		if err != nil {
			RespondWithError(w, 400, err)
		}
		RespondWithJSON(w, 200, refreshToken)
	}
}
