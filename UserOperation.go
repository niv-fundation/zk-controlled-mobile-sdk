package zk_controlled_mobile_sdk

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"math/big"
	"strings"
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

func (uo UserOperation) print() {
	initCodeHex := hex.EncodeToString(uo.InitCode)
	callDataHex := hex.EncodeToString(uo.CallData)

	fmt.Println("Sender:", uo.Sender)
	fmt.Println("Nonce:", uo.Nonce)
	fmt.Println("InitCode:", initCodeHex)
	fmt.Println("CallData:", callDataHex)
	fmt.Println("AccountGasLimits:", uo.AccountGasLimits)
	fmt.Println("PreVerificationGas:", uo.PreVerificationGas)
	fmt.Println("GasFees:", uo.GasFees)
	fmt.Println("PaymasterAndData:", uo.PaymasterAndData)
	fmt.Println("Signature:", uo.Signature)
}

func ToPackedUserOperation(userOp *UserOperation) PackedUserOperation {
	accountGasLimits := hexutil.MustDecode(userOp.AccountGasLimits)
	gasFees := hexutil.MustDecode(userOp.GasFees)

	var accountGasLimitsBytes [32]byte
	copy(accountGasLimitsBytes[:], accountGasLimits)

	var gasFeesBytes [32]byte
	copy(gasFeesBytes[:], gasFees)

	return PackedUserOperation{
		Sender:             userOp.Sender,
		Nonce:              userOp.Nonce,
		InitCode:           userOp.InitCode,
		CallData:           userOp.CallData,
		AccountGasLimits:   accountGasLimitsBytes,
		PreVerificationGas: userOp.PreVerificationGas,
		GasFees:            gasFeesBytes,
		PaymasterAndData:   userOp.PaymasterAndData,
		Signature:          userOp.Signature,
	}
}

func ToPackedUserOperationArray(userOp *UserOperation) []PackedUserOperation {
	return []PackedUserOperation{ToPackedUserOperation(userOp)}
}

// ComputeAccountGasLimits computes the account gas limits.
func ComputeAccountGasLimits(verificationGasLimit, callGasLimit *big.Int) (*big.Int, error) {
	if verificationGasLimit == nil || callGasLimit == nil {
		return nil, errors.New("verificationGasLimit and callGasLimit cannot be nil")
	}

	shifted := new(big.Int).Lsh(verificationGasLimit, 128)
	accountGasLimits := new(big.Int).Or(shifted, callGasLimit)

	return accountGasLimits, nil
}

// ComputeGasFees computes the gas fees.
func ComputeGasFees(maxPriorityFeePerGas, maxFeePerGas *big.Int) (*big.Int, error) {
	if maxPriorityFeePerGas == nil || maxFeePerGas == nil {
		return nil, errors.New("maxPriorityFeePerGas and maxFeePerGas cannot be nil")
	}

	shifted := new(big.Int).Lsh(maxPriorityFeePerGas, 128)
	gasFees := new(big.Int).Or(shifted, maxFeePerGas)

	return gasFees, nil
}

// GetEmptyPackedUserOperation returns an empty packed user operation.
func GetEmptyPackedUserOperation() (*UserOperation, error) {
	accountGasLimitsInt, err := ComputeAccountGasLimits(VerificationGasLimit, CallGasLimit)
	if err != nil {
		return nil, err
	}
	accountGasLimitsHex := ToBeHex(accountGasLimitsInt, 32)

	gasFeesInt, err := ComputeGasFees(MaxPriorityFeePerGas, MaxFeePerGas)
	if err != nil {
		return nil, err
	}
	gasFeesHex := ToBeHex(gasFeesInt, 32)

	return &UserOperation{
		Sender:             common.HexToAddress("0x0000000000000000000000000000000000000000"),
		Nonce:              big.NewInt(0),
		InitCode:           []byte{},
		CallData:           []byte{},
		AccountGasLimits:   accountGasLimitsHex,
		PreVerificationGas: VerificationGasLimit,
		GasFees:            gasFeesHex,
		PaymasterAndData:   []byte{},
		Signature:          []byte{},
	}, nil
}

func GetPaymasterAndData(paymaster string) string {
	return ToBeHex(MustParseBigInt(paymaster), 20) + "0000000000000000000000000001000000000000000000000000000000001000"
}

// GetDefaultPackedUserOperation returns a default packed user operation for a given account.
func GetDefaultPackedUserOperation(account *Account) (*UserOperation, error) {
	emptyUserOp, err := GetEmptyPackedUserOperation()
	if err != nil {
		return nil, err
	}

	address, err := account.GetAddress()
	if err != nil {
		return nil, err
	}
	emptyUserOp.Sender = address

	nonce, err := account.GetCurrentNonce()
	if err != nil {
		return nil, err
	}
	emptyUserOp.Nonce = nonce

	emptyUserOp.PaymasterAndData = hexutil.MustDecode(GetPaymasterAndData(hexutil.Encode(account.Paymaster.Bytes())))

	return emptyUserOp, nil
}

func GetUserOpHash(client *ethclient.Client, entryPointAddressStr string, userOp *UserOperation) (common.Hash, error) {
	entryPointAddress := common.HexToAddress(entryPointAddressStr)

	entryPointCaller, err := NewEntryPointCaller(entryPointAddress, client)
	if err != nil {
		return common.Hash{}, fmt.Errorf("error creating entry point caller: %v", err)
	}

	packedUserOp := ToPackedUserOperation(userOp)

	hash, err := entryPointCaller.GetUserOpHash(&bind.CallOpts{}, packedUserOp)
	if err != nil {
		return common.Hash{}, fmt.Errorf("error getting user op hash: %v", err)
	}

	return hash, nil
}

func GetInitCode(client *ethclient.Client, factoryAddressStr, nullifier string) ([]byte, common.Address, error) {
	factoryAddress := common.HexToAddress(factoryAddressStr)

	factoryCaller, err := NewSmartAccountFactoryCaller(factoryAddress, client)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("error creating factory caller: %v", err)
	}

	accountFactoryABI, err := abi.JSON(strings.NewReader(SmartAccountFactoryMetaData.ABI))
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to parse ABI: %v", err)
	}

	nullifierBytes := common.HexToHash(nullifier)

	encodedFunctionData, err := accountFactoryABI.Pack("deploySmartAccount", nullifierBytes)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to pack function data: %v", err)
	}

	initCode := append(factoryAddress.Bytes(), encodedFunctionData...)

	predictedAddress, err := factoryCaller.PredictSmartAccountAddress(&bind.CallOpts{}, nullifierBytes)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to predict address: %v", err)
	}

	return initCode, predictedAddress, nil
}

func BuildSendETHUserOperation(client *ethclient.Client, privateKey, eventID, receiver, amount, accountAddress, factoryAddressStr, paymasterAddress string) (*UserOperation, error) {
	nullifier, err := CalculateEventNullifierHex(privateKey, eventID)
	if err != nil {
		return nil, fmt.Errorf("error calculating event nullifier: %v", err)
	}

	var userOp *UserOperation
	if accountAddress == common.HexToAddress("0x0").String() {
		userOp, err = GetEmptyPackedUserOperation()

		initCode, predictedAddress, err := GetInitCode(client, factoryAddressStr, nullifier)
		if err != nil {
			return nil, fmt.Errorf("error getting init code: %v", err)
		}

		userOp.Sender = predictedAddress
		userOp.InitCode = initCode
		userOp.PaymasterAndData = hexutil.MustDecode(GetPaymasterAndData(paymasterAddress))
	} else {
		account := NewAccount(common.HexToAddress(accountAddress), common.HexToAddress(paymasterAddress), client)
		userOp, err = GetDefaultPackedUserOperation(account)
	}

	userOp.CallData, err = GetSendETHData(receiver, amount)

	return userOp, nil
}

func GetSendETHData(receiver, amount string) ([]byte, error) {
	accountABI, err := abi.JSON(strings.NewReader(SmartAccountMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %v", err)
	}

	receiverAddress := common.HexToAddress(receiver)
	amountInt, err := ParseBigInt(amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %v", err)
	}

	encodedFunctionData, err := accountABI.Pack("execute", receiverAddress, amountInt, []byte{})
	if err != nil {
		return nil, fmt.Errorf("failed to pack function data: %v", err)
	}

	return encodedFunctionData, nil
}
