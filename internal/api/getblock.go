package api

import (
	"balance-tracker/internal/domain/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	baseURL = "https://go.getblock.io/"
)

func GetBlockByNumber(number string) (models.Block, error) {
	const op = "api.GetBlockByNumber"
	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf("%s%s", baseURL, apiKey)
	reqBody := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["%s", true],"id":"getblock.io"}`, number)
	resp, err := http.Post(url, "application/json", strings.NewReader(reqBody))
	if err != nil {
		return models.Block{}, errors.New(fmt.Sprintf("функция: %s; не удалось получить ответ от api: %v", op, err))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Block{}, errors.New(fmt.Sprintf("функция: %s; не удалось прочитать ответ от api: %v", op, err))
	}
	var result struct {
		Jsonrpc string       `json:"jsonrpc"`
		ID      string       `json:"id"`
		Result  models.Block `json:"result"`
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return models.Block{}, errors.New(fmt.Sprintf("функция: %s; не удалось распарсить ответ от api: %v", op, err))
	}
	return result.Result, nil
}

func GetLatestBlockNumber() (string, error) {
	const op = "api.GetLatestBlockNumber"
	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf("%s%s", baseURL, apiKey)
	reqBody := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":"getblock.io"}`)
	resp, err := http.Post(url, "application/json", strings.NewReader(reqBody))
	if err != nil {
		return "", errors.New(fmt.Sprintf("функция: %s; не удалось получить ответ от api: %v", op, err))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("функция: %s; не удалось прочитать ответ от api: %v", op, err))
	}
	var result struct {
		Jsonrpc string `json:"jsonrpc"`
		ID      string `json:"id"`
		Result  string `json:"result"`
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return "", errors.New(fmt.Sprintf("функция: %s; не удалось распарсить ответ от api: %v", op, err))
	}
	return result.Result, nil
}
