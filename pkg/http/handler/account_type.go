package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"restAPI/pkg/domain"
	"restAPI/pkg/storage/mysql/entity"
)

//AccountTypeHandler ...
type AccountTypeHandler interface {
	CreateAccountType() func(w http.ResponseWriter, r *http.Request)
	GetAllAccountTypes() func(w http.ResponseWriter, r *http.Request)
}

type accountTypeHandler struct {
	AccountTypeService domain.AccountTypeService
}

//NewAccountTypeHandler ...
func NewAccountTypeHandler(a domain.AccountTypeService) AccountTypeHandler {
	return &accountTypeHandler{AccountTypeService: a}
}

func (h *accountTypeHandler) CreateAccountType() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var accountType entity.AccountType
		if json.NewDecoder(r.Body).Decode(&accountType) != nil {
			RespondWithError(w, 404, errors.New("Bad request"))
		}
		err := h.AccountTypeService.CreateAccountType(accountType)
		if err != nil {
			RespondWithError(w, 404, err)
			return
		}
		RespondWithJSON(w, 203, "Account type successfully created!")
	}
}

func (h *accountTypeHandler) GetAllAccountTypes() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := h.AccountTypeService.GetAllAccountTypes()

		if err != nil {
			RespondWithError(w, 404, err)
		}
		RespondWithJSON(w, 200, res)
	}
}
