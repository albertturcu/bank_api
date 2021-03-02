package handler

import (
	"encoding/json"
	"net/http"
	"restAPI/pkg/storage/mysql/entity"

	"github.com/gorilla/mux"
)

type moneyRequest struct {
	AccountNo string
	Amount    float32
}

//AccountHandler ...
type AccountHandler interface {
	GetAllAccounts() func(w http.ResponseWriter, r *http.Request)
	CreateAccount() func(w http.ResponseWriter, r *http.Request)
	WithdrawMoney() func(w http.ResponseWriter, r *http.Request)
	DepositMoney() func(w http.ResponseWriter, r *http.Request)
	CloseAccount() func(w http.ResponseWriter, r *http.Request)
}

func (h *appHandler) GetAllAccounts() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := h.s.GetAllAccounts()

		RespondWithJSON(w, 200, res)
	}
}

func (h *appHandler) CreateAccount() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var account entity.Account
		json.NewDecoder(r.Body).Decode(&account)
		res, err := h.s.CreateAccount(account)
		if err != nil {
			RespondWithError(w, 400, err)
			return
		}
		RespondWithJSON(w, 203, res)
	}
}

func (h *appHandler) WithdrawMoney() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var mr moneyRequest
		json.NewDecoder(r.Body).Decode(&mr)
		acc, err := h.s.WithdrawMoney(mr.AccountNo, mr.Amount)
		if err != nil {
			RespondWithError(w, 400, err)
		}
		RespondWithJSON(w, 200, acc)
	}
}

func (h *appHandler) DepositMoney() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var mr moneyRequest
		json.NewDecoder(r.Body).Decode(&mr)
		acc, err := h.s.DepositMoney(mr.AccountNo, mr.Amount)
		if err != nil {
			RespondWithError(w, 400, err)
		}
		RespondWithJSON(w, 200, acc)
	}
}

func (h *appHandler) CloseAccount() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		accountNo := mux.Vars(r)["accountNo"]
		res := h.s.CloseAccount(accountNo)

		if res != nil {
			RespondWithError(w, 404, res)
		}
		RespondWithJSON(w, 200, "Account closed!")
	}
}
