package router

import (
	"restAPI/pkg/http/handler"

	"github.com/gorilla/mux"
)

//NewRouter ...
func NewRouter(h handler.AppHandler) *mux.Router {
	r := mux.NewRouter()

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/getOne/{id}", h.GetUser()).Methods("GET")
	userRouter.HandleFunc("/getAll", h.GetUsers()).Methods("GET")
	userRouter.HandleFunc("/deleteOne/{id}", h.DeleteUser()).Methods("DELETE")
	userRouter.HandleFunc("/updateOne/{id}", h.UpdateUser()).Methods("PUT")
	userRouter.HandleFunc("/createOne", h.AddUser()).Methods("POST")

	accountRouter := r.PathPrefix("/account").Subrouter()

	accountRouter.HandleFunc("/getAllAccounts", h.GetAllAccounts()).Methods("GET")
	accountRouter.HandleFunc("/createAccount", h.CreateAccount()).Methods("POST")

	accountRouter.HandleFunc("/withdrawMoney", h.WithdrawMoney()).Methods("POST")
	accountRouter.HandleFunc("/depositMoney", h.DepositMoney()).Methods("POST")
	accountRouter.HandleFunc("/closeAccount/{accountNo}", h.CloseAccount()).Methods("GET")

	accountRouter.HandleFunc("/createAccountType", h.CreateAccountType()).Methods("POST")
	accountRouter.HandleFunc("/getAllAccountTypes", h.GetAllAccountTypes()).Methods("GET")

	return r
}
