package user_operations

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func getTestUserOp(t *testing.T) *UserOperationJson {
	return &UserOperationJson{
		Sender:             "0xf7202e2b5209c73462b85e231b9fe09eba05047b",
		Nonce:              "0x0e",
		InitCode:           "0x",
		CallData:           "0xb61d27f6000000000000000000000000e890ef63826e3e073cd0b7cd9b3839ad5b85bf5300000000000000000000000000000000000000000000000000003691d6afc00000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000",
		AccountGasLimits:   "0x000000000000000000000000000c350000000000000000000000000000061a80",
		PreVerificationGas: "0x0c3500",
		GasFees:            "0x3b9aca0000000000000000000000000133f72e19",
		PaymasterAndData:   "0xe8ccf8dd49e297c357fdc1f84f9a6e2fed83c4260000000000000000000000000001000000000000000000000000000000001000",
		Signature:          "0x",
	}
}

func TestEncode(t *testing.T) {
	userOp := getTestUserOp(t)

	expectedEncodedUserOpHex := "0x000000000000000000000000f7202e2b5209c73462b85e231b9fe09eba05047b000000000000000000000000000000000000000000000000000000000000000ec5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47050fc9a13a667bba42694076b988f89e99c41c591706e80bc899e422113307933000000000000000000000000000c350000000000000000000000000000061a8000000000000000000000000000000000000000000000000000000000000c35000000000000000000000000003b9aca0000000000000000000000000133f72e19ece72051698e136a1e4ceefc426f7fb0c37df765f2ad7f4761b548a5a9d7b90b"

	encodedUserOp, err := encode(userOp)
	if err != nil {
		t.Fatal(err)
	}

	encodedUserOpHex := hexutil.Encode(encodedUserOp)

	if expectedEncodedUserOpHex != encodedUserOpHex {
		t.Errorf("Encoded user operation does not match expected value.\nGot: %s\nExp: %s", encodedUserOpHex, expectedEncodedUserOpHex)
	}
}

func TestHash(t *testing.T) {
	userOp := getTestUserOp(t)

	// Expected user operation hash (from your console log)
	expectedUserOpHashHex := "0xff0d4842f8f6addde1a74920c38ae81db39625bcb1958f36ee45551890bdcd07"

	userOpHash, err := Hash(userOp)
	if err != nil {
		t.Fatal(err)
	}

	userOpHashHex := hexutil.Encode(userOpHash)

	if expectedUserOpHashHex != userOpHashHex {
		t.Errorf("User operation hash does not match expected value.\nGot: %s\nExp: %s", userOpHashHex, expectedUserOpHashHex)
	}
}

func TestGetUserOpHash(t *testing.T) {
	userOp := getTestUserOp(t)

	expectedFinalHashHex := "0x5ab03ae5d584d4d028d00d5c7e28b7c47ff7a7f8685c204ab9542586c74c2c9f"

	contractAddress := common.HexToAddress("0x5fbdb2315678afecb367f032d93f642f64180aa3")
	chainID := big.NewInt(31337)

	finalHash, err := GetUserOpHashLocal(userOp, contractAddress, chainID)
	if err != nil {
		t.Fatal(err)
	}

	finalHashHex := hexutil.Encode(finalHash)

	if expectedFinalHashHex != finalHashHex {
		t.Errorf("Final hash does not match expected value.\nGot: %s\nExp: %s", finalHashHex, expectedFinalHashHex)
	}
}
