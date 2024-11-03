package user_operations

import "math/big"

var (
	VerificationGasLimit = big.NewInt(800000)
	CallGasLimit         = big.NewInt(400000)
	MaxPriorityFeePerGas = big.NewInt(40000000000)
)
