package web

import (
	"html/template"
	"net/http"
)

//LoginHandler ...
func (*web) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl := template.Must(template.ParseFiles("./pkg/http/handler/web/templates/login.html"))
	tmpl.Execute(w, struct{ Success bool }{true})
}
