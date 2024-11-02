package zk_controlled_mobile_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/constants"
	"github.com/iden3/go-iden3-crypto/poseidon"
)

var (
	VerificationGasLimit = big.NewInt(600000)
	CallGasLimit         = big.NewInt(100000)
	MaxPriorityFeePerGas = big.NewInt(1000000000)
	MaxFeePerGas         = big.NewInt(5000000000)
)

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

func SignRawPoseidon(privateKey, hash string) (*babyjub.Point, *big.Int, error) {
	secretKey, err := ParseBigInt(privateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing secret key to big int: %v", err)
	}

	hashInt, err := ParseBigInt(hash)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing hash to big int: %v", err)
	}

	publicKey := babyjub.NewPoint().Mul(secretKey, babyjub.B8)

	r := new(big.Int).Mod(hashInt, babyjub.SubOrder)
	R8 := babyjub.NewPoint().Mul(r, babyjub.B8)

	hashScalar := new(big.Int).Mod(hashInt, constants.Q)
	hashedMessage, err := poseidon.Hash([]*big.Int{R8.X, R8.Y, publicKey.X, publicKey.Y, hashScalar})
	if err != nil {
		return nil, nil, fmt.Errorf("error hashing message: %v", err)
	}

	part1 := new(big.Int).Mul(hashedMessage, secretKey)
	part2 := new(big.Int).Add(r, part1)
	S := new(big.Int).Mod(part2, babyjub.SubOrder)

	return R8, S, nil
}

func PrintPackedUserOperation(packedUserOp *PackedUserOperation) {
	fmt.Println("Sender:", packedUserOp.Sender.String())
	fmt.Println("Nonce:", packedUserOp.Nonce)
	fmt.Println("InitCode:", hexutil.Encode(packedUserOp.InitCode))
	fmt.Println("CallData:", hexutil.Encode(packedUserOp.CallData))
	fmt.Println("AccountGasLimits:", hexutil.Encode(packedUserOp.AccountGasLimits[:]))
	fmt.Println("PreVerificationGas:", packedUserOp.PreVerificationGas)
	fmt.Println("GasFees:", hexutil.Encode(packedUserOp.GasFees[:]))
	fmt.Println("PaymasterAndData:", hexutil.Encode(packedUserOp.PaymasterAndData))
	fmt.Println("Signature:", hexutil.Encode(packedUserOp.Signature))
}

// UserOperation represents the structure required by eth_sendUserOperation
type UserOperationReq struct {
	Sender                        string `json:"sender"`
	Nonce                         string `json:"nonce"`
	Factory                       string `json:"factory"`
	FactoryData                   string `json:"factoryData"`
	CallData                      string `json:"callData"`
	CallGasLimit                  string `json:"callGasLimit"`
	VerificationGasLimit          string `json:"verificationGasLimit"`
	PreVerificationGas            string `json:"preVerificationGas"`
	MaxFeePerGas                  string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas          string `json:"maxPriorityFeePerGas"`
	Paymaster                     string `json:"paymaster"`
	PaymasterVerificationGasLimit string `json:"paymasterVerificationGasLimit"`
	PaymasterPostOpGasLimit       string `json:"paymasterPostOpGasLimit"`
	PaymasterData                 string `json:"paymasterData"`
	Signature                     string `json:"signature"`
}

func ConstructUserOperation(userOp *UserOperation, paymasterAddressStr string) (UserOperationReq, error) {
	nonce := "0x00"

	if userOp.Nonce.Int64() > 0 {
		nonce = hexutil.Encode(userOp.Nonce.Bytes())
	}

	factory := ""
	factoryData := ""

	if len(userOp.InitCode) > 0 {
		factory = hexutil.Encode(userOp.InitCode[:20])
		factoryData = hexutil.Encode(userOp.InitCode[20:])
	}

	// Ensure all big.Int fields are converted to hex strings with "0x" prefix
	userOperation := UserOperationReq{
		Sender:                        userOp.Sender.Hex(),
		Nonce:                         nonce,
		Factory:                       factory,
		FactoryData:                   factoryData,
		CallData:                      hexutil.Encode(userOp.CallData),
		CallGasLimit:                  hexutil.Encode(CallGasLimit.Bytes()),
		VerificationGasLimit:          hexutil.Encode(VerificationGasLimit.Bytes()),
		PreVerificationGas:            hexutil.Encode(userOp.PreVerificationGas.Bytes()),
		MaxFeePerGas:                  hexutil.Encode(MaxFeePerGas.Bytes()),
		MaxPriorityFeePerGas:          hexutil.Encode(MaxPriorityFeePerGas.Bytes()),
		Paymaster:                     paymasterAddressStr,
		PaymasterVerificationGasLimit: fmt.Sprintf("0x00000000000000000000000000010000"),
		PaymasterPostOpGasLimit:       fmt.Sprintf("0x00000000000000000000000000001000"),
		PaymasterData:                 "0x",
		Signature:                     hexutil.Encode(userOp.Signature),
	}

	return userOperation, nil
}

func EstimateGas(userOp *UserOperation, paymasterAddressStr, entryPointStr string) error {
	userOperation, err := ConstructUserOperation(userOp, paymasterAddressStr)
	if err != nil {
		return fmt.Errorf("error constructing user operation: %v", err)
	}

	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "eth_estimateUserOperationGas",
		"params": []interface{}{
			userOperation,
			entryPointStr,
		},
	}

	// Serialize the request body to JSON
	requestData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("error marshaling request body: %v", err)
	}

	serviceURL := "https://eth-sepolia.g.alchemy.com/v2/-Jm280LvZnniBfiaxZtQa_wL1b_okXCZ"
	resp, err := http.Post(serviceURL, "application/json", bytes.NewReader(requestData))
	if err != nil {
		return fmt.Errorf("error sending request to service: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	var response struct {
		Jsonrpc                          string `json:"jsonrpc"`
		Id                               int    `json:"id"`
		Result                           string `json:"result"`
		PreVerificationGasRes            int    `json:"preVerificationGas"`
		VerificationGasLimitRes          int    `json:"verificationGasLimit"`
		CallGasLimitRes                  int    `json:"callGasLimit"`
		PaymasterVerificationGasLimitRes int    `json:"paymasterVerificationGasLimit"`
		Error                            *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    string `json:"data,omitempty"`
		} `json:"error,omitempty"`
	}

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %v", err)
	}

	if response.Error != nil {
		return fmt.Errorf("error from service: %v", response.Error.Message)
	}

	VerificationGasLimit = big.NewInt(int64(response.VerificationGasLimitRes))
	CallGasLimit = big.NewInt(int64(response.CallGasLimitRes))

	return nil
}

func SendUOToBundler(userOp *UserOperation, paymasterAddressStr, entryPointStr string) (string, error) {
	userOperation, err := ConstructUserOperation(userOp, paymasterAddressStr)
	if err != nil {
		return "", fmt.Errorf("error constructing user operation: %v", err)
	}

	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "eth_sendUserOperation",
		"params": []interface{}{
			userOperation,
			entryPointStr,
		},
	}

	// Serialize the request body to JSON
	requestData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request body: %v", err)
	}

	serviceURL := "https://eth-sepolia.g.alchemy.com/v2/-Jm280LvZnniBfiaxZtQa_wL1b_okXCZ"
	resp, err := http.Post(serviceURL, "application/json", bytes.NewReader(requestData))
	if err != nil {
		return "", fmt.Errorf("error sending request to service: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	var response struct {
		Jsonrpc string `json:"jsonrpc"`
		Id      int    `json:"id"`
		Result  string `json:"result,omitempty"`
		Error   *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error,omitempty"`
	}

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling response: %v", err)
	}

	if response.Error != nil {
		return "", fmt.Errorf("error from service: %v", response.Error.Message)
	}

	fmt.Println("Transaction hash:", response)

	// Return the transaction hash from the response
	return response.Result, nil
}
