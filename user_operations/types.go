package user_operations

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type UserOperation struct {
	Sender             common.Address `json:"sender"`
	Nonce              *big.Int       `json:"nonce"`
	InitCode           []byte         `json:"initCode"`
	CallData           []byte         `json:"callData"`
	AccountGasLimits   string         `json:"accountGasLimits"`
	PreVerificationGas *big.Int       `json:"preVerificationGas"`
	GasFees            string         `json:"gasFees"`
	PaymasterAndData   []byte         `json:"paymasterAndData"`
	Signature          []byte         `json:"signature"`
}

type UserOperationJson struct {
	Sender             string `json:"sender"`
	Nonce              string `json:"nonce"`
	InitCode           string `json:"initCode"`
	CallData           string `json:"callData"`
	AccountGasLimits   string `json:"accountGasLimits"`
	PreVerificationGas string `json:"preVerificationGas"`
	GasFees            string `json:"gasFees"`
	PaymasterAndData   string `json:"paymasterAndData"`
	Signature          string `json:"signature"`
}

type UserOperationReq struct {
	Sender                        string `json:"sender"`
	Nonce                         string `json:"nonce"`
	Factory                       string `json:"factory,omitempty"`
	FactoryData                   string `json:"factoryData,omitempty"`
	CallData                      string `json:"callData"`
	CallGasLimit                  string `json:"callGasLimit"`
	VerificationGasLimit          string `json:"verificationGasLimit"`
	PreVerificationGas            string `json:"preVerificationGas"`
	MaxFeePerGas                  string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas          string `json:"maxPriorityFeePerGas"`
	PaymasterVerificationGasLimit string `json:"paymasterVerificationGasLimit"`
	PaymasterPostOpGasLimit       string `json:"paymasterPostOpGasLimit"`
	Signature                     string `json:"signature"`
	Paymaster                     string `json:"paymaster"`
	PaymasterData                 string `json:"paymasterData"`
}
