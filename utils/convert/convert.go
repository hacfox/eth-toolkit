package convert

import (
	"math"
	"math/big"
	"strings"
)

const decimalPrecision = 128

// ToDecimal returns *big.Float format of given decimal string,
// will return nil if input string is empty.
func ToDecimal(valueStr string) *big.Float {
	value, _ := new(big.Float).SetPrec(decimalPrecision).SetString(valueStr)
	return value
}

// BigFloatToString converts big.Float to string.
func BigFloatToString(value *big.Float) string {
	if value == nil {
		return ""
	}

	valueStr := value.Text('f', 32)
	valueStr = strings.TrimRight(valueStr, "0")
	valueStr = strings.TrimSuffix(valueStr, ".")

	return valueStr
}

// RemoveZeros trims trailing zeros of the given value string.
// This function should only be used when the input value is
// promised to be the floating-point string number,
// e.g., when fetching data from database.
func RemoveZeros(valueStr string) string {
	return BigFloatToString(ToDecimal(valueStr))
}

// UintStr2Decimal returns the *big.Float which is rounded quotient valueStr/10^decimals.
func UintStr2Decimal(valueStr string, decimals int) *big.Float {
	return new(big.Float).Quo(ToDecimal(valueStr), big.NewFloat(math.Pow10(decimals)))
}

func ZeroPeddingRight(hex string) string {
	if strings.HasPrefix(hex, "0x") {
		return hex + strings.Repeat("0", 64-2-len(hex))
	}

	return hex + strings.Repeat("0", 64-len(hex))
}

// BalanceToString returns a string balance.
func BalanceToString(value interface{}) string {
	var balance string
	switch value.(type) {
	case string:
		balance, _ = value.(string)
	case float64:
		v := big.NewFloat(value.(float64))
		balance = BigFloatToString(v)
	case *big.Float:
		v, _ := value.(*big.Float)
		balance = BigFloatToString(v)
	default:
		panic("not support this type")
	}
	return balance
}
