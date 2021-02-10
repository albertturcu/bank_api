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

//AccountHandler ...
type AccountHandler interface {
	GetAllAccounts() func(w http.ResponseWriter, r *http.Request)
	CreateAccount() func(w http.ResponseWriter, r *http.Request)
	WithdrawMoney() func(w http.ResponseWriter, r *http.Request)
	DepositMoney() func(w http.ResponseWriter, r *http.Request)
	CloseAccount() func(w http.ResponseWriter, r *http.Request)
}

type accountHandler struct {
	AccountService domain.AccountService
}

//NewAccountHandler ...
func NewAccountHandler(a domain.AccountService) AccountHandler {
	return &accountHandler{AccountService: a}
}

func (h *accountHandler) GetAllAccounts() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := h.AccountService.GetAllAccounts()

		RespondWithJSON(w, 200, res)
	}
}

func (h *accountHandler) CreateAccount() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var account entity.Account
		json.NewDecoder(r.Body).Decode(&account)
		res, err := h.AccountService.CreateAccount(account)

		if err != nil {
			RespondWithError(w, 400, err)
			return
		}
		RespondWithJSON(w, 203, res)
	}
}

func (h *accountHandler) WithdrawMoney() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var mr moneyRequest
		json.NewDecoder(r.Body).Decode(&mr)
		acc, err := h.AccountService.WithdrawMoney(mr.AccountNo, mr.Amount)
		if err != nil {
			RespondWithError(w, 400, err)
		}
		RespondWithJSON(w, 200, acc)
	}
}

func (h *accountHandler) DepositMoney() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var mr moneyRequest
		json.NewDecoder(r.Body).Decode(&mr)
		acc, err := h.AccountService.DepositMoney(mr.AccountNo, mr.Amount)
		if err != nil {
			RespondWithError(w, 400, err)
		}
		RespondWithJSON(w, 200, acc)
	}
}

func (h *accountHandler) CloseAccount() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		accountNo := mux.Vars(r)["accountNo"]
		res := h.AccountService.CloseAccount(accountNo)

		if res != nil {
			RespondWithError(w, 404, res)
		}
		RespondWithJSON(w, 200, "Account closed!")
	}
}
