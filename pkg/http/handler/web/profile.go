package web

import (
	"fmt"
	"html/template"
	"net/http"
)

//ProfileHandler ...
func (web *web) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	user, err := web.s.GetUser("1")
	if err != nil {
		fmt.Println("err")
	}
	tmpl := template.Must(template.ParseFiles("./pkg/http/handler/web/templates/profile.html"))
	tmpl.Execute(w, user)
}
