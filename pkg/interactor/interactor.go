package interactor

import (
	"restAPI/pkg/http/handler"
	"restAPI/pkg/storage/mysql/repository"

	"github.com/jinzhu/gorm"
)

//Interactor ...
type Interactor interface {
	NewAppHandler() handler.AppHandler
	NewAccountRepository() repository.Account
	NewAccountTypeRepository() repository.AccountType
	NewUserRepository() repository.User
	NewAccountHandler() handler.AccountHandler
	NewAccountTypeHandler() handler.AccountTypeHandler
	NewUserHandler() handler.UserHandler
}

type interactor struct {
	Conn *gorm.DB
}

//NewInteractor ...
func NewInteractor(Conn *gorm.DB) Interactor {
	return &interactor{Conn}
}

type appHandler struct {
	handler.UserHandler
	handler.AccountHandler
	handler.AccountTypeHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.UserHandler = i.NewUserHandler()
	appHandler.AccountHandler = i.NewAccountHandler()
	appHandler.AccountTypeHandler = i.NewAccountTypeHandler()
	return appHandler
}

func (i *interactor) NewUserRepository() repository.User {
	return repository.NewUserRepository(i.Conn)
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserRepository())
}

func (i *interactor) NewAccountRepository() repository.Account {
	return repository.NewAccountRepository(i.Conn)
}

func (i *interactor) NewAccountHandler() handler.AccountHandler {
	return handler.NewAccountHandler(i.NewAccountRepository())
}

func (i *interactor) NewAccountTypeRepository() repository.AccountType {
	return repository.NewAccountTypeRepository(i.Conn)
}

func (i *interactor) NewAccountTypeHandler() handler.AccountTypeHandler {
	return handler.NewAccountTypeHandler(i.NewAccountTypeRepository())
}
