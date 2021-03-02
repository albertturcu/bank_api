package web

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//LoginHandler ...
func (*web) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("err")
	}

	tmpl := template.Must(template.ParseFiles(wd + os.Getenv("TEMPLATES_PATH_DOCKER") + "/login.html"))
	tmpl.Execute(w, struct{ Success bool }{true})
}
