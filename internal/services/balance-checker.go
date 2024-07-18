package services

import (
	"balance-tracker/internal/api"
	"log"
)

func CheckBalance() (string, error) {
	block, err := api.GetLatestBlock()
	if err != nil {
		return "", err
	}
	log.Println(len(block.Transactions))
	//	....
	return block.Number, nil
}
