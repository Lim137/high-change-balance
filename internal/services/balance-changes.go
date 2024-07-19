package services

import (
	"balance-tracker/internal/api"
	value_format "balance-tracker/pkg/value-format"
	"math/big"
)

type BalanceChange struct {
	Address string
	Change  float64
}

func GetBalanceChanges() (map[string]float64, error) {
	blockNumbers, err := getLastNBlockNumbers(100)
	if err != nil {
		return nil, err
	}
	balanceChanges := make(map[string]float64)
	for _, blockNumber := range blockNumbers {
		block, err := api.GetBlockByNumber(blockNumber)
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
