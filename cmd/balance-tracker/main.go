package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"balance-tracker/internal/app"
	"balance-tracker/internal/logger/sl"
	"balance-tracker/pkg/env"
)

func main() {
	err := env.LoadEnvFile(".env")
	if err != nil {
		log.Fatalf("Error load .env file: %s", err)
	}

	port := env.GetEnv("PORT", "8080")
	environment := env.GetEnv("ENV", "prod")

	logger := sl.GetLogger(environment)
	logger.Info("starting app", slog.Any("port", port), slog.Any("env", environment))

	go app.New(port)

	logger.Info("app started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	logger.Info("app stopped", slog.Any("signal", sign.String()))
}
