package zk_controlled_mobile_sdk

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/niv-fundation/zk-controlled-mobile-sdk/bindings"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/niv-fundation/zk-controlled-mobile-sdk/zkp"
	"github.com/pkg/errors"
	"math/big"

	uo "github.com/niv-fundation/zk-controlled-mobile-sdk/user_operations"
)

type TransactionLog struct {
	To     common.Address `json:"to"`
	Amount string         `json:"amount"`
	Time   uint64         `json:"time"`
}

type EthereumClient struct {
	RPC string
}

func (c *EthereumClient) NewEthereumClient(rpc string) *EthereumClient {
	return &EthereumClient{rpc}
}

func (c *EthereumClient) GetEthAddress(privateKeyStr, chainId string) (string, error) {
	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		return "", fmt.Errorf("error connecting to RPC: %v", err)
	}
	defer client.Close()

	var privateKey *ecdsa.PrivateKey
	privateKey, err = crypto.ToECDSA(crypto.Keccak256([]byte(privateKeyStr)))
	if err != nil {
		return "", fmt.Errorf("error parsing private key: %v", err)
	}

	signerOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, uo.MustParseBigInt(chainId, "chainId in GetEthAddress"))
	if err != nil {
		return "", fmt.Errorf("failed to get keyed transactor: %w", err)
	}

	return signerOpts.From.Hex(), nil
}

func (c *EthereumClient) GetTransactionHistory(contractAddressStr, offset, limit string) (string, error) {
	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		return "", fmt.Errorf("error connecting to RPC: %v", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress(contractAddressStr)

	contractCaller, err := bindings.NewSmartAccountCaller(contractAddress, client)
	if err != nil {
		return "", fmt.Errorf("error creating contract caller: %v", err)
	}

	toBigOffset, success := new(big.Int).SetString(offset, 10)
	if !success {
		return "", fmt.Errorf("error parsing offset to big int: %v", err)
	}
	toBigLimit, success := new(big.Int).SetString(limit, 10)
	if !success {
		return "", fmt.Errorf("error parsing limit to big int: %v", err)
	}

	transactionsCount, err := contractCaller.GetTransactionHistoryLength(&bind.CallOpts{})
	if err != nil {
		return "", fmt.Errorf("error fetching transaction history length: %v", err)
	}

	// Convert big.Int to int64 for calculations
	transactionsCountInt := transactionsCount.Int64()
	offsetInt := toBigOffset.Int64()
	limitInt := toBigLimit.Int64()

	if offsetInt > transactionsCountInt {
		return "[]", nil
	}

	startIndex := transactionsCountInt - offsetInt - limitInt
	if startIndex < 0 {
		startIndex = 0
	}

	newOffset := big.NewInt(startIndex)
	newLimit := big.NewInt(limitInt)

	transactionLogs, err := contractCaller.GetTransactionHistory(&bind.CallOpts{}, newOffset, newLimit)

	if err != nil {
		return "", fmt.Errorf("error fetching transaction history: %v", err)
	}

	var logs []TransactionLog
	for _, log := range transactionLogs {
		ethBalance := new(big.Float).Quo(new(big.Float).SetInt(log.Value), big.NewFloat(1e18))

		logs = append(logs, TransactionLog{
			To:     log.Destination,
			Amount: fmt.Sprintf("%.6f ETH", ethBalance),
			Time:   log.Timestamp.Uint64(),
		})
	}

	if len(logs) == 0 {
		return "[]", nil
	}

	for i, j := 0, len(logs)-1; i < j; i, j = i+1, j-1 {
		logs[i], logs[j] = logs[j], logs[i]
	}

	jsonData, err := json.Marshal(logs)
	if err != nil {
		return "", fmt.Errorf("error marshalling transaction logs to JSON: %v", err)
	}

	return string(jsonData), nil
}

func (c *EthereumClient) GetContractBalance(addressStr string) (string, error) {
	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		return "", fmt.Errorf("error connecting to RPC: %v", err)
	}
	defer client.Close()

	address := common.HexToAddress(addressStr)

	// Get the balance at the latest block
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return "", fmt.Errorf("error retrieving balance: %v", err)
	}

	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	formattedBalance := fmt.Sprintf("%.6f ETH", ethBalance)

	return formattedBalance, nil
}

func (c *EthereumClient) GetUO(privateKeyStr, eventID, receiver, amount, accountAddress, factoryAddressStr, paymasterAddressStr string) (string, error) {
	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		return "", fmt.Errorf("error connecting to RPC: %v", err)
	}
	defer client.Close()

	userOp, err := uo.BuildSendETHUserOperation(client, privateKeyStr, eventID, receiver, amount, accountAddress, factoryAddressStr, paymasterAddressStr)
	if err != nil {
		return "", fmt.Errorf("error building user operation: %v", err)
	}

	baseFeePerGas, err := GetBaseFee(client)
	if err != nil {
		return "", fmt.Errorf("error getting base fee: %v", err)
	}

	fmt.Println("baseFeePerGas: ", baseFeePerGas)

	gasFeesStr, err := uo.ComputeGasFees(uo.MaxPriorityFeePerGas, baseFeePerGas)
	if err != nil {
		return "", fmt.Errorf("error computing gas fees: %v", err)
	}

	userOp.GasFees = hexutil.Encode(gasFeesStr.Bytes())

	if userOp.Nonce == nil {
		userOp.Nonce = big.NewInt(0)
	}

	marshaledUo, err := userOp.MarshalJSON()
	if err != nil {
		return "", fmt.Errorf("error marshalling user operation to JSON: %v", err)
	}

	return string(marshaledUo), nil
}

func (c *EthereumClient) GetSendETHInputs(privateKeyStr, eventID, userOpStr, entryPointSrt string) (string, error) {
	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		return "", fmt.Errorf("error connecting to RPC: %v", err)
	}
	defer client.Close()

	userOp, err := uo.UnmarshalUserOperation(userOpStr)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling user operation: %v", err)
	}

	chainId, err := GetChainID(client)
	if err != nil {
		return "", fmt.Errorf("error getting chain ID: %v", err)
	}

	userOpHash, err := uo.GetUserOpHashLocal(userOp, common.HexToAddress(entryPointSrt), chainId)
	if err != nil {
		return "", fmt.Errorf("error getting user op hash: %v", err)
	}

	userOpHashInt := new(big.Int).SetBytes(userOpHash[:])

	R8, S, err := SignRawPoseidon(privateKeyStr, userOpHashInt.String())
	if err != nil {
		return "", fmt.Errorf("error signing user operation: %v", err)
	}

	inputs := zkp.AuthProofInput{
		SkI:          uo.MustParseBigInt(privateKeyStr, "privateKey in GetSendETHInputs"),
		EventID:      uo.MustParseBigInt(eventID, "eventID in GetSendETHInputs"),
		MessageHash:  userOpHashInt,
		SignatureR8x: R8.X,
		SignatureR8y: R8.Y,
		SignatureS:   S,
	}

	if err := inputs.Validate(); err != nil {
		return "", fmt.Errorf("error validating inputs: %v", err)
	}

	jsonData, err := json.Marshal(inputs)
	if err != nil {
		return "", fmt.Errorf("error marshalling inputs to JSON: %v", err)
	}

	return string(jsonData), nil
}

func GetBaseFee(client *ethclient.Client) (*big.Int, error) {
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error getting block by number: %v", err)
	}

	baseFee := block.BaseFee()
	if baseFee == nil {
		return nil, errors.New("base fee is nil")
	}

	baseFee.Add(baseFee, big.NewInt(1000000000))

	return baseFee, nil
}

func GetChainID(client *ethclient.Client) (*big.Int, error) {
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting chain ID: %v", err)
	}

	return chainID, nil
}

func (c *EthereumClient) SendETH(userOpStr, entryPointStr, paymasterAddressStr, proof string) (string, error) {
	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		return "", fmt.Errorf("error connecting to RPC: %v", err)
	}
	defer client.Close()

	userOp, err := uo.UnmarshalUserOperation(userOpStr)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling user operation: %v", err)
	}

	proofStruct := &zkp.Proof{}
	err = proofStruct.FromJSON(proof)
	if err != nil {
		return "", fmt.Errorf("error parsing proof: %v", err)
	}

	proofPoints, err := proofStruct.ToVerifierHelperProofPoints()
	if err != nil {
		return "", fmt.Errorf("error converting proof to verifier helper proof points: %v", err)
	}

	proofPoints.B[0][0], proofPoints.B[0][1] = proofPoints.B[0][1], proofPoints.B[0][0]
	proofPoints.B[1][0], proofPoints.B[1][1] = proofPoints.B[1][1], proofPoints.B[1][0]

	userOp.Signature, err = zkp.EncodeIdentityProof(proofPoints)
	if err != nil {
		return "", fmt.Errorf("error encoding identity proof: %v", err)
	}

	chainId, err := GetChainID(client)
	if err != nil {
		return "", fmt.Errorf("error getting chain ID: %v", err)
	}

	preSendHash, err := uo.GetUserOpHashLocal(userOp, common.HexToAddress(entryPointStr), chainId)
	if err != nil {
		return "", fmt.Errorf("error getting user op hash: %v", err)
	}

	fmt.Println("preSendHash: ", hexutil.Encode(preSendHash))

	result, err := uo.SendUOToBundler(userOp, paymasterAddressStr, entryPointStr)
	if err != nil {
		return "", err
	}

	return result, nil
}

type UserOperationResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		UserOperation struct {
			Sender                        string `json:"sender"`
			Nonce                         string `json:"nonce"`
			CallData                      string `json:"callData"`
			CallGasLimit                  string `json:"callGasLimit"`
			VerificationGasLimit          string `json:"verificationGasLimit"`
			PreVerificationGas            string `json:"preVerificationGas"`
			MaxPriorityFeePerGas          string `json:"maxPriorityFeePerGas"`
			MaxFeePerGas                  string `json:"maxFeePerGas"`
			Paymaster                     string `json:"paymaster"`
			PaymasterVerificationGasLimit string `json:"paymasterVerificationGasLimit"`
			PaymasterPostOpGasLimit       string `json:"paymasterPostOpGasLimit"`
			PaymasterData                 string `json:"paymasterData"`
			Signature                     string `json:"signature"`
		} `json:"userOperation"`
		EntryPoint      string `json:"entryPoint"`
		BlockNumber     string `json:"blockNumber,omitempty"`
		BlockHash       string `json:"blockHash,omitempty"`
		TransactionHash string `json:"transactionHash,omitempty"`
	} `json:"result"`
}

func (c *EthereumClient) IsUOConfirmed(uoHash string) (bool, error) {
	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		return false, fmt.Errorf("error connecting to RPC: %v", err)
	}
	defer client.Close()

	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "eth_getUserOperationByHash",
		"params": []interface{}{
			uoHash,
		},
	}

	// Serialize the request body to JSON
	requestData, err := json.Marshal(requestBody)
	if err != nil {
		return false, fmt.Errorf("error marshaling request body: %v", err)
	}

	serviceURL := "https://eth-sepolia.g.alchemy.com/v2/TbAcglEXwYO3cAz1Tx-0QARHayvkTEUW"
	resp, err := http.Post(serviceURL, "application/json", bytes.NewReader(requestData))
	if err != nil {
		return false, fmt.Errorf("error sending request to service: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("error reading response: %v", err)
	}

	var response struct {
		Jsonrpc string `json:"jsonrpc"`
		Id      int    `json:"id"`
		Result  struct {
			UserOperation struct {
				Sender                        string `json:"sender"`
				Nonce                         string `json:"nonce"`
				CallData                      string `json:"callData"`
				CallGasLimit                  string `json:"callGasLimit"`
				VerificationGasLimit          string `json:"verificationGasLimit"`
				PreVerificationGas            string `json:"preVerificationGas"`
				MaxPriorityFeePerGas          string `json:"maxPriorityFeePerGas"`
				MaxFeePerGas                  string `json:"maxFeePerGas"`
				Paymaster                     string `json:"paymaster"`
				PaymasterVerificationGasLimit string `json:"paymasterVerificationGasLimit"`
				PaymasterPostOpGasLimit       string `json:"paymasterPostOpGasLimit"`
				PaymasterData                 string `json:"paymasterData"`
				Signature                     string `json:"signature"`
			} `json:"userOperation"`
			EntryPoint      string `json:"entryPoint"`
			BlockNumber     string `json:"blockNumber,omitempty"`
			BlockHash       string `json:"blockHash,omitempty"`
			TransactionHash string `json:"transactionHash,omitempty"`
		} `json:"result"`
		Error *struct {
			Code    int             `json:"code"`
			Message string          `json:"message"`
			Data    json.RawMessage `json:"data,omitempty"`
		} `json:"error,omitempty"`
	}

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return false, fmt.Errorf("error unmarshaling response: %v", err)
	}

	if response.Error != nil {
		return false, fmt.Errorf("error in response: %v", response.Error.Message)
	}

	if response.Result.TransactionHash == "" {
		return false, fmt.Errorf("transaction hash is empty")
	}

	txHashCommon := common.HexToHash(response.Result.TransactionHash)
	receipt, err := client.TransactionReceipt(context.Background(), txHashCommon)
	if err != nil {
		return false, fmt.Errorf("error getting transaction receipt: %v", err)
	}

	if receipt.BlockNumber != nil {
		return true, nil
	}

	return false, nil
}

func (c *EthereumClient) GetPredictedAccountAddress(privateKey, eventID, factoryAddressStr string) (string, error) {
	nullifier, err := uo.CalculateEventNullifierHex(privateKey, eventID)
	if err != nil {
		return "", fmt.Errorf("error calculating event nullifier: %v", err)
	}

	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		return "", fmt.Errorf("error connecting to RPC: %v", err)
	}
	defer client.Close()

	factoryAddress := common.HexToAddress(factoryAddressStr)

	factoryCaller, err := bindings.NewSmartAccountFactoryCaller(factoryAddress, client)
	if err != nil {
		return "", fmt.Errorf("error creating factory caller: %v", err)
	}

	nullifierBytes := common.HexToHash(nullifier)

	address, err := factoryCaller.PredictSmartAccountAddress(&bind.CallOpts{}, nullifierBytes)
	if err != nil {
		return "", fmt.Errorf("error predicting smart account address: %v", err)
	}

	return address.Hex(), nil
}

func (c *EthereumClient) GetAccountDeployed(privateKey, eventID, factoryAddressStr string) (string, error) {
	nullifier, err := uo.CalculateEventNullifierHex(privateKey, eventID)
	if err != nil {
		return "", fmt.Errorf("error calculating event nullifier: %v", err)
	}

	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		return "", fmt.Errorf("error connecting to RPC: %v", err)
	}
	defer client.Close()

	factoryAddress := common.HexToAddress(factoryAddressStr)

	factoryCaller, err := bindings.NewSmartAccountFactoryCaller(factoryAddress, client)
	if err != nil {
		return "", fmt.Errorf("error creating factory caller: %v", err)
	}

	nullifierBytes := common.HexToHash(nullifier)

	address, err := factoryCaller.GetSmartAccount(&bind.CallOpts{}, nullifierBytes)
	if err != nil {
		return "", fmt.Errorf("error predicting smart account address: %v", err)
	}

	return address.Hex(), nil
}
