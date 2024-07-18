package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Couldn't write response data: "+err.Error())
	}
}

func RespondWithError(w http.ResponseWriter, code int, message string) {

	type errRespond struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, code, errRespond{Error: message})
}
