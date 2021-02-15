package handler

import (
	"encoding/json"
	"net/http"
)

//RespondWithError ...
func RespondWithError(rw http.ResponseWriter, code int, message error) {
	RespondWithJSON(rw, code, map[string]error{"error": message})
}

//RespondWithJSON ...
func RespondWithJSON(rw http.ResponseWriter, code int, payload interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(payload)
	rw.WriteHeader(code)
	rw.Write(response)
}
