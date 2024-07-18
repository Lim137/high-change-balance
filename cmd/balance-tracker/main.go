package main

import (
	"balance-tracker/internal/app"
	"balance-tracker/pkg/env"
	"log"
	"os"
)

func main() {
	err := env.LoadEnvFile(".env")
	if err != nil {
		log.Fatalf("ошибка загрузки .env файла: %s", err)
	}
	port := os.Getenv("PORT")
	app.New(port)
	//value1 := os.Getenv("API_KEY")
	//value2 := os.Getenv("TEST")
	//fmt.Println("API_KEY:", value1)
	//fmt.Println("TEST:", value2)
}
