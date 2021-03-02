package web

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//ProfileHandler ...
func (web *web) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	user, err := web.s.GetUser("1")
	if err != nil {
		fmt.Println("err")
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("err")
	}

	tmpl := template.Must(template.ParseFiles(wd + os.Getenv("TEMPLATES_PATH_DOCKER") + "/profile.html"))
	tmpl.Execute(w, user)
}
