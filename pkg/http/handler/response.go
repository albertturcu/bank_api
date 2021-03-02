package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//RespondWithError ...
func RespondWithError(rw http.ResponseWriter, code int, message error) {
	RespondWithJSON(rw, code, map[string]error{"error": message})
}

//RespondWithJSON ...
func RespondWithJSON(rw http.ResponseWriter, code int, payload interface{}) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	response, err := json.MarshalIndent(payload, "", " ")

	if err != nil {
		fmt.Println(err)
	}
	rw.WriteHeader(code)
	if _, err := rw.Write(response); err != nil {
		fmt.Println(err)
	}
}
