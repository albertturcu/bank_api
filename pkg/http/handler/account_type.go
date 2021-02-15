package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"restAPI/pkg/storage/mysql/entity"
)

// AccountTypeHandler ...
type AccountTypeHandler interface {
	CreateAccountType() func(w http.ResponseWriter, r *http.Request)
	GetAllAccountTypes() func(w http.ResponseWriter, r *http.Request)
}

func (h *appHandler) CreateAccountType() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var accountType entity.AccountType
		if json.NewDecoder(r.Body).Decode(&accountType) != nil {
			RespondWithError(w, 404, errors.New("Bad request"))
		}
		err := h.s.CreateAccountType(accountType)
		if err != nil {
			RespondWithError(w, 404, err)
			return
		}
		RespondWithJSON(w, 203, "Account type successfully created!")
	}
}

func (h *appHandler) GetAllAccountTypes() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := h.s.GetAllAccountTypes()

		if err != nil {
			RespondWithError(w, 404, err)
		}
		RespondWithJSON(w, 200, res)
	}
}
