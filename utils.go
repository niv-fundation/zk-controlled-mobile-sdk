package zk_controlled_mobile_sdk

import (
	"fmt"
	uo "github.com/niv-fundation/zk-controlled-mobile-sdk/user_operations"
	"math/big"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/constants"
	"github.com/iden3/go-iden3-crypto/poseidon"
)

func SignRawPoseidon(privateKey, hash string) (*babyjub.Point, *big.Int, error) {
	secretKey, err := uo.ParseBigInt(privateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing secret key to big int: %v", err)
	}

	hashInt, err := uo.ParseBigInt(hash)
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
