package value_format

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func TrimValuePrefix(value string) string {
	return strings.TrimPrefix(value, "0x")
}

func GetInt64Value(value string) (int64, error) {
	const op = "value-format.GetInt64Value"
	valueInt, err := strconv.ParseInt(TrimValuePrefix(value), 16, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("функция: %s; ошибка при преобразовании адреса: %v", op, err.Error()))
	}
	return valueInt, nil
}

func Int64ToHexValue(value int64) string {
	return fmt.Sprintf("0x%x", value)
}

func BigValueToEtherFloat64(value *big.Int) float64 {
	etherValue := new(big.Float).Quo(new(big.Float).SetInt(value), big.NewFloat(1e18))
	etherFloat64, _ := etherValue.Float64()
	return etherFloat64
}
