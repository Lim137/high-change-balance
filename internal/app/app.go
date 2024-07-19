package app

import (
	"log"
	"net/http"

	"balance-tracker/internal/app/middleware"
	high_change_balance "balance-tracker/internal/handlers/high-change-balance"
)

func New(port string) {
	const op = "app.New"
	http.Handle("/getaddr", middleware.CORS(http.HandlerFunc(high_change_balance.BalanceTrackerHandler)))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("%s: error when starting the server: %v", op, err)
	}
}
