package getblock

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"balance-tracker/internal/domain/models"
)

const (
	baseURL = "https://go.getblock.io/"
)

func GetBlockByNumber(number string) (models.Block, error) {
	const op = "api.GetBlockByNumber"
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return models.Block{}, fmt.Errorf("%s: API_KEY is empty", op)
	}
	url := fmt.Sprintf("%s%s", baseURL, apiKey)
	reqBody := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["%s", true],"id":"getblock.io"}`, number)
	resp, err := http.Post(url, "application/json", strings.NewReader(reqBody))
	if err != nil {
		return models.Block{}, fmt.Errorf("%s: failed to get response from API: %v", op, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Block{}, fmt.Errorf("%s: failed to read API response: %v", op, err)
	}

	var result struct {
		Jsonrpc string       `json:"jsonrpc"`
		ID      string       `json:"id"`
		Result  models.Block `json:"result"`
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return models.Block{}, fmt.Errorf("%s: failed to unmarshal API response: %v", op, err)
	}

	return result.Result, nil
}

func GetLatestBlockNumber() (string, error) {
	const op = "api.GetLatestBlockNumber"
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("%s: API_KEY is empty", op)
	}
	url := fmt.Sprintf("%s%s", baseURL, apiKey)
	reqBody := `{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":"getblock.io"}`
	resp, err := http.Post(url, "application/json", strings.NewReader(reqBody))

	if err != nil {
		return "", fmt.Errorf("%s: failed to get response from API: %v", op, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("%s: failed to read API response: %v", op, err)
	}

	var result struct {
		Jsonrpc string `json:"jsonrpc"`
		ID      string `json:"id"`
		Result  string `json:"result"`
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("%s: failed to unmarshal API response: %v", op, err)
	}

	return result.Result, nil
}
