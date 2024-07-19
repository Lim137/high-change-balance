package services

import (
	"balance-tracker/internal/api"
	value_format "balance-tracker/pkg/value-format"
)

func getLastNBlockNumbers(n int) ([]string, error) {
	currentBlockNumberStr, err := api.GetLatestBlockNumber()
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
	if err != nil {
		return nil, err
	}
	return blockNumbers, nil
}
