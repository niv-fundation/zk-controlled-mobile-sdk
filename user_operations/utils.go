package user_operations

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/pkg/errors"
)

func CalculateEventNullifierHex(privateKey, eventID string) (string, error) {
	eventNullifier, err := calculateEventNullifier(privateKey, eventID)
	if err != nil {
		return "", fmt.Errorf("error calculating event nullifier: %v", err)
	}

	eventNullifierHex := hex.EncodeToString(eventNullifier.Bytes())
	if len(eventNullifierHex) < 64 {
		eventNullifierHex = fmt.Sprintf("%0*s", 64, eventNullifierHex)
	}

	return fmt.Sprintf("0x%s", eventNullifierHex), nil
}

func calculateEventNullifier(privateKey, eventID string) (*big.Int, error) {
	secretKey, err := ParseBigInt(privateKey)
	if err != nil {
		return nil, fmt.Errorf("err parsing secret key to big int: %v", err)
	}

	secretKeyHash, err := poseidon.Hash([]*big.Int{secretKey})
	if err != nil {
		return nil, fmt.Errorf("err hashing secret key: %v", err)
	}

	eventIDInt, ok := new(big.Int).SetString(eventID, 0)
	if !ok {
		return nil, fmt.Errorf("err parsing event ID: %v", err)
	}

	eventNullifier, err := poseidon.Hash([]*big.Int{secretKey, secretKeyHash, eventIDInt})
	if err != nil {
		return nil, fmt.Errorf("err hashing event: %v", err)
	}

	return eventNullifier, nil
}

// ToBeHex formats a big.Int into a hex string of the specified byte size.
func ToBeHex(value *big.Int, size int) string {
	if size == 0 {
		return "0x"
	}

	hexStr := value.Text(16)
	expectedLength := size * 2

	if strings.HasPrefix(hexStr, "-") {
		panic("negative numbers are not supported")
	}

	padding := expectedLength - len(hexStr)
	if padding > 0 {
		hexStr = strings.Repeat("0", padding) + hexStr
	}

	return "0x" + hexStr
}

func ParseBigInt(input string) (*big.Int, error) {
	secretKey, success := new(big.Int).SetString(input, 10)
	if success {
		return secretKey, nil
	}

	secretKey, success = new(big.Int).SetString(input, 16)
	if success {
		return secretKey, nil
	}

	if strings.HasPrefix(input, "0x") || strings.HasPrefix(input, "0X") {
		secretKey, success = new(big.Int).SetString(input[2:], 16)
		if success {
			return secretKey, nil
		}
	}

	return nil, errors.New("could not parse string as big.Int in decimal or hexadecimal")
}

func MustParseBigInt(input string) *big.Int {
	secretKey, err := ParseBigInt(input)
	if err != nil {
		panic(err)
	}

	return secretKey
}
