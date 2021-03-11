package utils

import (
	"encoding/hex"
	"strings"

	"github.com/mr-tron/base58"
)

func SubgraphHexToHash(s string) (string, error) {
	buf, err := hex.DecodeString("1220" + strings.TrimPrefix(s, "0x"))
	if err != nil {
		return "", err
	}
	return base58.Encode(buf), nil
}

func SubgraphHashToHex(s string) (string, error) {
	decoded, err := base58.Decode(s)
	if err != nil {
		return "", err
	}
	return "0x" + strings.TrimPrefix(hex.EncodeToString(decoded), "1220"), nil
}
