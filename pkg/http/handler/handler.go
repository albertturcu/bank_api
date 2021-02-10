package handler

//AppHandler ...
type AppHandler interface {
	UserHandler
	AccountHandler
	AccountTypeHandler
}
