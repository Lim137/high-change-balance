package services

import (
	"math"
	"sort"
)

func GetHighChangeAddress() (string, error) {
	balanceChanges, err := GetBalanceChanges()
	if err != nil {
		return "", err
	}
	changes := make([]BalanceChange, 0, len(balanceChanges))
	for address, change := range balanceChanges {
		changes = append(changes, BalanceChange{Address: address, Change: change})
	}

	sort.Slice(changes, func(i, j int) bool {
		return math.Abs(changes[i].Change) > math.Abs(changes[j].Change)
	})
	return changes[0].Address, nil
}
