package services

import (
	"math/big"
	"sync"

	"balance-tracker/internal/api/getblock"
	"balance-tracker/internal/domain/models"
	value_format "balance-tracker/pkg/value-format"
)

const (
	blockCount      = 100
	goroutinesCount = 8
)

type BalanceChange struct {
	Address string
	Change  float64
}

func GetBalanceChanges() (map[string]*float64, error) {
	blockNumbers, err := getLastNBlockNumbers(blockCount)
	if err != nil {
		return nil, err
	}
	blockCh := make(chan models.Block, len(blockNumbers))
	errCh := make(chan error, len(blockNumbers))
	var wg sync.WaitGroup
	balanceChanges := make(map[string]*float64)
	semaphore := make(chan struct{}, goroutinesCount)
	for _, blockNumber := range blockNumbers {
		semaphore <- struct{}{}
		wg.Add(1)
		go getblock.GetBlockByNumber(blockNumber, blockCh, errCh, &wg, semaphore)
	}
	go func() {
		wg.Wait()
		close(blockCh)
		close(errCh)

	}()

	var processWg sync.WaitGroup
	var mu sync.Mutex
	for block := range blockCh {
		processWg.Add(1)
		go processBlock(block, balanceChanges, &mu, &processWg)
	}
	processWg.Wait()
	if len(errCh) > 0 {
		err = <-errCh
		return nil, err
	}
	return balanceChanges, nil
}

func getLastNBlockNumbers(n int) ([]string, error) {
	currentBlockNumberStr, err := getblock.GetLatestBlockNumber()
	if err != nil {
		return nil, err
	}

	currentBlockNumber, err := value_format.GetInt64Value(currentBlockNumberStr)
	if err != nil {
		return nil, err
	}

	blockNumbers := make([]string, 0, n)
	for i := 0; i < n; i++ {
		blockNumber := value_format.Int64ToHexValue(currentBlockNumber - int64(i))
		blockNumbers = append(blockNumbers, blockNumber)
	}

	return blockNumbers, nil
}

func processBlock(block models.Block, balanceChanges map[string]*float64, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, tx := range block.Transactions {
		value := new(big.Int)
		value.SetString(tx.Value[2:], 16)
		etherValue := value_format.BigValueToEtherFloat64(value)
		mu.Lock()
		if _, exists := balanceChanges[tx.From]; !exists {
			balanceChanges[tx.From] = new(float64)
		}
		if _, exists := balanceChanges[tx.To]; !exists {
			balanceChanges[tx.To] = new(float64)
		}
		*balanceChanges[tx.From] -= etherValue
		*balanceChanges[tx.To] += etherValue
		mu.Unlock()
	}
}
