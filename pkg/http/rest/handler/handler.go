package handler

import (
	"restAPI/pkg/domain"
	"sync"

	"github.com/gorilla/mux"
)

//NewRouter ...
func NewRouter(u domain.UserService, a domain.AccountService, at domain.AccountTypeService, wg *sync.WaitGroup) *mux.Router {
	r := mux.NewRouter()

	userRouter := r.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/getOne/{id}", getUser(u)).Methods("GET")
	userRouter.HandleFunc("/getAll", getUsers(u)).Methods("GET")
	userRouter.HandleFunc("/deleteOne/{id}", deleteUser(u)).Methods("DELETE")
	userRouter.HandleFunc("/updateOne/{id}", updateUser(u)).Methods("PUT")
	userRouter.HandleFunc("/createOne", addUser(u)).Methods("POST")

	accountRouter := r.PathPrefix("/account").Subrouter()

	accountRouter.HandleFunc("/getAllAccounts", getAllAccounts(a)).Methods("GET")
	accountRouter.HandleFunc("/createAccount", createAccount(a)).Methods("POST")

	accountRouter.HandleFunc("/withdrawMoney", withdrawMoney(a)).Methods("POST")
	accountRouter.HandleFunc("/depositMoney", depositMoney(a)).Methods("POST")
	accountRouter.HandleFunc("/closeAccount/{accountNo}", closeAccount(a)).Methods("GET")

	accountRouter.HandleFunc("/createAccountType", createAccountType(at)).Methods("POST")
	accountRouter.HandleFunc("/getAllAccountTypes", getAllAccountTypes(at)).Methods("GET")

	return r
}
