package app

import (
	high_change_balance "balance-tracker/internal/handlers/high-change-balance"
	"log"
	"net/http"
)

func New(port string) {
	const op = "app.New"
	http.HandleFunc("/getaddr", high_change_balance.BalanceTrackerHandler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ошибка при запуске сервера: ", err)
	}
}
