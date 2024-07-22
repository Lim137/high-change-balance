package getblock

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"balance-tracker/internal/domain/models"
)

const (
	baseURL = "https://go.getblock.io/"
)

func GetBlockByNumber(number string, ch chan models.Block, errCh chan error, wg *sync.WaitGroup, semaphore chan struct{}) {
	defer wg.Done()
	defer func() { <-semaphore }()
	const op = "api.GetBlockByNumber"
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		errCh <- fmt.Errorf("%s: API_KEY is empty", op)
		return
	}
	url := fmt.Sprintf("%s%s", baseURL, apiKey)
	reqBody := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["%s", true],"id":"getblock.io"}`, number)
	resp, err := http.Post(url, "application/json", strings.NewReader(reqBody))
	if err != nil {
		errCh <- fmt.Errorf("%s: failed to get response from API: %v", op, err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errCh <- fmt.Errorf("%s: failed to read API response: %v", op, err)
		return
	}

	var result struct {
		Jsonrpc string       `json:"jsonrpc"`
		ID      string       `json:"id"`
		Result  models.Block `json:"result"`
	}
	if err = json.Unmarshal(body, &result); err != nil {
		errCh <- fmt.Errorf("%s: failed to unmarshal API response: %v", op, err)
		return
	}
	ch <- result.Result

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
