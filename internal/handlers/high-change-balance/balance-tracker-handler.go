package high_change_balance

import (
	httpresponse "balance-tracker/internal/adapters/http"
	"balance-tracker/internal/logger/sl"
	"balance-tracker/internal/services"
	"net/http"
	"os"
)

type BalanceTrackerResponse struct {
	HighChangeAddress string `json:"highChangeAddress"`
}

func BalanceTrackerHandler(w http.ResponseWriter, r *http.Request) {
	logger := sl.GetLogger(os.Getenv("ENV"))
	if r.Method != http.MethodGet {
		httpresponse.RespondWithError(w, http.StatusMethodNotAllowed, "метод запроса должен быть GET")
		logger.Info("неверный метод запроса")
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
