package high_change_balance

import (
	httpresponse "balance-tracker/internal/adapters/http"
	"balance-tracker/internal/services"
	"net/http"
)

type BalanceTrackerResponse struct {
	HighChangeAddress string `json:"highChangeAddress"`
}

func BalanceTrackerHandler(w http.ResponseWriter, r *http.Request) {
	const op = "balance_tracker.BalanceTrackerHandler"
	if r.Method != http.MethodGet {
		httpresponse.RespondWithError(w, http.StatusMethodNotAllowed, "метод запроса должен быть GET")
		return
	}
	highChangeAddr, err := services.GetHighChangeAddress()
	if err != nil {
		httpresponse.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpresponse.RespondWithJSON(w, http.StatusOK, BalanceTrackerResponse{HighChangeAddress: highChangeAddr})

}
