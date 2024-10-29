package zk_controlled_mobile_sdk

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/constants"
	"github.com/iden3/go-iden3-crypto/poseidon"
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
