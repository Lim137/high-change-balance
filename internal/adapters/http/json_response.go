package http

import (
	"encoding/json"
	"net/http"

	"balance-tracker/internal/logger/sl"
	"balance-tracker/pkg/env"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	environment := env.GetEnv("ENV", "prod")
	logger := sl.GetLogger(environment)
	if err != nil {
		logger.Error("failed to marshal JSON response: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err = w.Write(data); err != nil {
		logger.Error("failed to write response: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}
