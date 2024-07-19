package handlers

import (
	"net/http"

	httpresponse "balance-tracker/internal/adapters/http"
	"balance-tracker/internal/logger/sl"
	"balance-tracker/internal/services"
	"balance-tracker/pkg/env"
)

type BalanceTrackerResponse struct {
	HighChangeAddress string `json:"highChangeAddress"`
}

func BalanceTrackerHandler(w http.ResponseWriter, r *http.Request) {
	environment := env.GetEnv("ENV", "prod")
	logger := sl.GetLogger(environment)
	if r.Method != http.MethodGet {
		httpresponse.RespondWithError(w, http.StatusMethodNotAllowed, "request method must be GET")
		logger.Info("invalid request method")
		return
	}
	highChangeAddr, err := services.GetHighChangeAddress()
	if err != nil {
		httpresponse.RespondWithError(w, http.StatusInternalServerError, err.Error())
		logger.Error(err.Error())
		return
	}
	httpresponse.RespondWithJSON(w, http.StatusOK, BalanceTrackerResponse{HighChangeAddress: highChangeAddr})
}
