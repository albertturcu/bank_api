package handler

import "restAPI/pkg/domain"

//AppHandler ...
type AppHandler interface {
	UserHandler
	AccountHandler
	AccountTypeHandler
}

type appHandler struct {
	s domain.DService
}

//NewAppHandler ...
func NewAppHandler(s domain.DService) AppHandler {
	return &appHandler{s: s}
}
