package user_operations

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/iden3/go-iden3-crypto/keccak256"
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

func calldataKeccak(data string) string {
	return hexutil.Encode(keccak256.Hash(hexutil.MustDecode(data)))
}

func encode(userOp *UserOperationJson) ([]byte, error) {
	nonce := userOp.Nonce
	hashInitCode := calldataKeccak(userOp.InitCode)
	hashCallData := calldataKeccak(userOp.CallData)
	accountGasLimits := userOp.AccountGasLimits
	preVerificationGas := userOp.PreVerificationGas
	gasFees := userOp.GasFees
	hashPaymasterAndData := calldataKeccak(userOp.PaymasterAndData)

	// Define ABI types
	senderType, _ := abi.NewType("address", "", nil)
	uint256Type, _ := abi.NewType("uint256", "", nil)
	bytes32Type, _ := abi.NewType("bytes32", "", nil)

	// Prepare arguments for ABI encoding
	arguments := abi.Arguments{
		{Name: "sender", Type: senderType},
		{Name: "nonce", Type: uint256Type},
		{Name: "hashInitCode", Type: bytes32Type},
		{Name: "hashCallData", Type: bytes32Type},
		{Name: "accountGasLimits", Type: uint256Type},
		{Name: "preVerificationGas", Type: uint256Type},
		{Name: "gasFees", Type: uint256Type},
		{Name: "hashPaymasterAndData", Type: bytes32Type},
	}

	// ABI-encode the data
	packedData, err := arguments.Pack(
		common.HexToAddress(userOp.Sender),
		MustParseBigInt(nonce),
		common.HexToHash(hashInitCode),
		common.HexToHash(hashCallData),
		MustParseBigInt(accountGasLimits),
		MustParseBigInt(preVerificationGas),
		MustParseBigInt(gasFees),
		common.HexToHash(hashPaymasterAndData),
	)
	if err != nil {
		return nil, err
	}
	return packedData, nil
}

func Hash(userOp *UserOperationJson) ([]byte, error) {
	encoded, err := encode(userOp)
	if err != nil {
		return nil, err
	}
	h := keccak256.Hash(encoded)
	return h[:], nil
}

func GetUserOpHashLocal(userOp *UserOperationJson, contractAddress common.Address, chainID *big.Int) ([]byte, error) {
	userOpHash, err := Hash(userOp)
	if err != nil {
		return nil, err
	}
	addressType, _ := abi.NewType("address", "", nil)
	uint256Type, _ := abi.NewType("uint256", "", nil)
	bytes32Type, _ := abi.NewType("bytes32", "", nil)

	arguments := abi.Arguments{
		{Name: "userOpHash", Type: bytes32Type},
		{Name: "contractAddress", Type: addressType},
		{Name: "chainID", Type: uint256Type},
	}

	// ABI-encode the data
	packedData, err := arguments.Pack(
		common.BytesToHash(userOpHash),
		contractAddress,
		chainID,
	)
	if err != nil {
		return nil, err
	}
	h := keccak256.Hash(packedData)
	return h[:], nil
}
