package services

import (
	"math/big"

	"balance-tracker/internal/api/getblock"
	value_format "balance-tracker/pkg/value-format"
)

const (
	blockCount = 100
)

type BalanceChange struct {
	Address string
	Change  float64
}

func GetBalanceChanges() (map[string]float64, error) {
	blockNumbers, err := getLastNBlockNumbers(blockCount)
	if err != nil {
		return nil, err
	}

	balanceChanges := make(map[string]float64)
	for _, blockNumber := range blockNumbers {
		block, err := getblock.GetBlockByNumber(blockNumber)
		if err != nil {
			return nil, err
		}
		for _, tx := range block.Transactions {
			value := new(big.Int)
			value.SetString(tx.Value[2:], 16)
			etherValue := value_format.BigValueToEtherFloat64(value)
			balanceChanges[tx.From] -= etherValue
			balanceChanges[tx.To] += etherValue
		}
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
