package utils

import (
	"encoding/hex"
	"errors"
	"math/big"
	"strings"

	"github.com/mr-tron/base58"
	"github.com/shopspring/decimal"
)

func SubgraphHexToHash(s string) (string, error) {
	if !strings.HasPrefix(s, "0x") {
		return "", errors.New("not a hex, should start with 0x")
	}
	buf, err := hex.DecodeString("1220" + strings.TrimPrefix(s, "0x"))
	if err != nil {
		return "", err
	}
	return base58.Encode(buf), nil
}

func SubgraphHashToHex(s string) (string, error) {
	if !strings.HasPrefix(s, "Qm") {
		return "", errors.New("not a valid hash, should start with Qm")
	}
	decoded, err := base58.Decode(s)
	if err != nil {
		return "", err
	}
	return "0x" + strings.TrimPrefix(hex.EncodeToString(decoded), "1220"), nil
}

func ToDecimal(ivalue interface{}, decimals int) (decimal.Decimal, error) {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, err := decimal.NewFromString(value.String())
	if err != nil {
		return decimal.Decimal{}, err
	}
	result := num.Div(mul)

	return result, nil
}

func ToWei(iamount interface{}, decimals int) (*big.Int, error) {
	amount := decimal.NewFromFloat(0)
	var err error
	switch v := iamount.(type) {
	case string:
		amount, err = decimal.NewFromString(v)
		if err != nil {
			return &big.Int{}, err
		}
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei, nil
}
