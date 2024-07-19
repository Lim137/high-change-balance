package main

import (
	"balance-tracker/internal/app"
	"balance-tracker/internal/logger/sl"
	"balance-tracker/pkg/env"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := env.LoadEnvFile(".env")
	if err != nil {
		log.Fatalf("ошибка загрузки .env файла: %s", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	projEnv := os.Getenv("ENV")
	if projEnv == "" {
		projEnv = "prod"
	}
	logger := sl.GetLogger(projEnv)
	logger.Info("starting app", slog.Any("port", port), slog.Any("env", os.Getenv("ENV")))
	go app.New(port)
	logger.Info("app started")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop
	logger.Info("app stopped", slog.Any("signal", sign.String()))
}
