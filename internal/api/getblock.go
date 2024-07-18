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

func GetLatestBlock() (models.Block, error) {
	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf("%s%s", baseURL, apiKey)
	reqBody := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest", true],"id":1}`)
	resp, err := http.Post(url, "application/json", strings.NewReader(reqBody))
	if err != nil {
		//fmt.Println("не удалось получить ответ от api: ", err)
		return models.Block{}, errors.New(fmt.Sprintf("не удалось получить ответ от api: %v", err))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println("не удалось прочитать ответ от api: ", err)
		return models.Block{}, errors.New(fmt.Sprintf("не удалось прочитать ответ от api: %v", err))
	}
	var result struct {
		Jsonrpc string       `json:"jsonrpc"`
		ID      int          `json:"id"`
		Result  models.Block `json:"result"`
	}

	if err = json.Unmarshal(body, &result); err != nil {
		//fmt.Println("не удалось распарсить ответ от api: ", err)
		return models.Block{}, errors.New(fmt.Sprintf("не удалось распарсить ответ от api: %v", err))
	}
	return result.Result, nil
}
