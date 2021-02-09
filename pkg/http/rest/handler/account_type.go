package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"restAPI/pkg/domain"
	"restAPI/pkg/storage/mysql/entity"
)

func createAccountType(a domain.AccountTypeService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var accountType entity.AccountType
		if json.NewDecoder(r.Body).Decode(&accountType) != nil {
			RespondWithError(w, 404, errors.New("Bad request"))
		}
		err := a.CreateAccountType(accountType)
		if err != nil {
			RespondWithError(w, 404, err)
			return
		}
		RespondWithJSON(w, 203, "Account type successfully created!")
	}
}

func getAllAccountTypes(a domain.AccountTypeService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := a.GetAllAccountTypes()

		if err != nil {
			RespondWithError(w, 404, err)
		}
		RespondWithJSON(w, 200, res)
	}
}
