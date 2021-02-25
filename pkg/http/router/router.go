package router

import (
	"restAPI/pkg/http/handler"
	"restAPI/pkg/http/middleware"

	"github.com/gorilla/mux"
)

//NewRouter ...
func NewRouter(h handler.AppHandler, m middleware.Middleware) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login", h.Login()).Methods("POST")
	r.HandleFunc("/register", h.AddUser()).Methods("POST")
	r.HandleFunc("/refreshToken", h.GetToken()).Methods("GET")

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.Use(m.ExtractToken, m.ValidateRequest)

	userRouter.HandleFunc("/getOne/{id}", h.GetUser()).Methods("GET")
	userRouter.HandleFunc("/getAll", h.GetUsers()).Methods("GET")
	userRouter.HandleFunc("/deleteOne/{id}", h.DeleteUser()).Methods("DELETE")
	userRouter.HandleFunc("/updateOne/{id}", h.UpdateUser()).Methods("PUT")

	accountRouter := r.PathPrefix("/account").Subrouter()
	userRouter.Use(m.ExtractToken, m.ValidateRequest)

	accountRouter.HandleFunc("/getAllAccounts", h.GetAllAccounts()).Methods("GET")
	accountRouter.HandleFunc("/createAccount", h.CreateAccount()).Methods("POST")
	accountRouter.HandleFunc("/withdrawMoney", h.WithdrawMoney()).Methods("POST")
	accountRouter.HandleFunc("/depositMoney", h.DepositMoney()).Methods("POST")
	accountRouter.HandleFunc("/closeAccount/{accountNo}", h.CloseAccount()).Methods("GET")

	accountRouter.HandleFunc("/createAccountType", h.CreateAccountType()).Methods("POST")
	accountRouter.HandleFunc("/getAllAccountTypes", h.GetAllAccountTypes()).Methods("GET")

	return r
}
