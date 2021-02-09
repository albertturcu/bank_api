package handler

import (
	"encoding/json"
	"net/http"
	"restAPI/pkg/domain"
	"restAPI/pkg/storage/mysql/entity"

	"github.com/gorilla/mux"
)

type moneyRequest struct {
	AccountNo string
	Amount    float32
}

func getAllAccounts(a domain.AccountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := a.GetAllAccounts()

		RespondWithJSON(w, 200, res)
	}
}

func createAccount(a domain.AccountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var account entity.Account
		json.NewDecoder(r.Body).Decode(&account)
		res, err := a.CreateAccount(account)

		if err != nil {
			RespondWithError(w, 400, err)
			return
		}
		RespondWithJSON(w, 203, res)
	}
}

func withdrawMoney(a domain.AccountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var mr moneyRequest
		json.NewDecoder(r.Body).Decode(&mr)
		acc, err := a.WithdrawMoney(mr.AccountNo, mr.Amount)
		if err != nil {
			RespondWithError(w, 400, err)
		}
		RespondWithJSON(w, 200, acc)
	}
}

func depositMoney(a domain.AccountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var mr moneyRequest
		json.NewDecoder(r.Body).Decode(&mr)
		acc, err := a.DepositMoney(mr.AccountNo, mr.Amount)
		if err != nil {
			RespondWithError(w, 400, err)
		}
		RespondWithJSON(w, 200, acc)
	}
}

func closeAccount(a domain.AccountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		accountNo := mux.Vars(r)["accountNo"]
		res := a.CloseAccount(accountNo)

		if res != nil {
			RespondWithError(w, 404, res)
		}
		RespondWithJSON(w, 200, "Account closed!")
	}
}
