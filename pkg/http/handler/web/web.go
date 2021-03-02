package web

import (
	"net/http"
	"restAPI/pkg/domain"
)

//Web ...
type Web interface {
	LoginHandler(w http.ResponseWriter, r *http.Request)
	ProfileHandler(w http.ResponseWriter, r *http.Request)
}

type web struct {
	s domain.DService
}

//NewWeb ...
func NewWeb(s domain.DService) Web {
	return &web{s: s}
}
