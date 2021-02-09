package handler

import (
	"encoding/json"
	"net/http"
)

//RespondWithError ...
func RespondWithError(w http.ResponseWriter, code int, message error) {
	RespondWithJSON(w, code, map[string]error{"error": message})
}

//RespondWithJSON ...
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
