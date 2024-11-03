package user_operations

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func ConstructUserOperationRequest(userOp *UserOperationJson, paymasterAddressStr string) (UserOperationReq, error) {
	maxPriorityFeePerGas, maxFeePerGas, err := DecomputeGasFees(userOp.GasFees)
	if err != nil {
		panic(err)
	}

	userOperation := UserOperationReq{
		Sender:                        userOp.Sender,
		Nonce:                         userOp.Nonce,
		CallData:                      userOp.CallData,
		CallGasLimit:                  hexutil.Encode(CallGasLimit.Bytes()),
		VerificationGasLimit:          hexutil.Encode(VerificationGasLimit.Bytes()),
		PreVerificationGas:            userOp.PreVerificationGas,
		MaxFeePerGas:                  hexutil.Encode(maxFeePerGas.Bytes()),
		MaxPriorityFeePerGas:          hexutil.Encode(maxPriorityFeePerGas.Bytes()),
		Paymaster:                     paymasterAddressStr,
		PaymasterVerificationGasLimit: fmt.Sprintf("0x00000000000000000000000000010000"),
		PaymasterPostOpGasLimit:       fmt.Sprintf("0x00000000000000000000000000001000"),
		PaymasterData:                 "0x",
		Signature:                     userOp.Signature,
	}

	initCode := hexutil.MustDecode(userOp.InitCode)
	if len(initCode) > 0 {
		userOperation.Factory = hexutil.Encode(initCode[:20])
		userOperation.FactoryData = hexutil.Encode(initCode[20:])
	}

	return userOperation, nil
}
