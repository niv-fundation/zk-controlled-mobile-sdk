package zk_controlled_mobile_sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/pkg/errors"
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

func (c *EthereumClient) GetTransactionHistory(contractAddressStr, offset, limit string) (string, error) {
	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		return "", fmt.Errorf("error connecting to RPC: %v", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress(contractAddressStr)

	contractCaller, err := NewSmartAccountCaller(contractAddress, client)
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

	transactionLogs, err := contractCaller.GetTransactionHistory(&bind.CallOpts{}, toBigOffset, toBigLimit)

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

func (c *EthereumClient) GetPredictedAccountAddress(privateKey, eventID, factoryAddressStr string) (string, error) {
	nullifier, err := CalculateEventNullifierHex(privateKey, eventID)
	if err != nil {
		return "", fmt.Errorf("error calculating event nullifier: %v", err)
	}

	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		return "", fmt.Errorf("error connecting to RPC: %v", err)
	}
	defer client.Close()

	factoryAddress := common.HexToAddress(factoryAddressStr)

	factoryCaller, err := NewSmartAccountFactoryCaller(factoryAddress, client)
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
