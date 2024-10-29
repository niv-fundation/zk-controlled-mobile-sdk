package zk_controlled_mobile_sdk

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
	"testing"
)

func TestToBeHex(t *testing.T) {
	testCase1, success := new(big.Int).SetString("a513E6E4b8f2a923D98304ec87F64353C4D5C853", 16)
	if !success {
		t.Errorf("Error converting string to big.Int")
	}

	tests := []struct {
		value    *big.Int
		size     int
		expected string
	}{
		{big.NewInt(0), 32, "0x" + strings.Repeat("0", 64)},
		{big.NewInt(1), 32, "0x" + strings.Repeat("0", 63) + "1"},
		{new(big.Int).SetBytes([]byte{0xFF}), 32, "0x" + strings.Repeat("0", 62) + "ff"},
		{testCase1, 32, "0x000000000000000000000000a513e6e4b8f2a923d98304ec87f64353c4d5c853"},
		{new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1)), 32, "0x" + strings.Repeat("f", 64)},
	}

	for _, tt := range tests {
		result := ToBeHex(tt.value, tt.size)
		if result != tt.expected {
			t.Errorf("ToBeHex(%v, %d) = %s; expected %s", tt.value, tt.size, result, tt.expected)
		}
	}
}

func TestComputeAccountGasLimits(t *testing.T) {
	verificationGasLimit := big.NewInt(2777216)
	callGasLimit := big.NewInt(2777216)

	expected := new(big.Int).Or(
		new(big.Int).Lsh(verificationGasLimit, 128),
		callGasLimit,
	)

	result, err := ComputeAccountGasLimits(verificationGasLimit, callGasLimit)
	if err != nil {
		t.Fatalf("ComputeAccountGasLimits returned an error: %v", err)
	}

	if result.Cmp(expected) != 0 {
		t.Errorf("ComputeAccountGasLimits result %s; expected %s", result.String(), expected.String())
	}
}

func TestComputeAccountGasLimits_NilInput(t *testing.T) {
	_, err := ComputeAccountGasLimits(nil, big.NewInt(1))
	if err == nil {
		t.Error("Expected error when verificationGasLimit is nil")
	}
}

func TestComputeGasFees(t *testing.T) {
	maxPriorityFeePerGas := big.NewInt(256)
	maxFeePerGas := big.NewInt(256)

	expected := new(big.Int).Or(
		new(big.Int).Lsh(maxPriorityFeePerGas, 128),
		maxFeePerGas,
	)

	result, err := ComputeGasFees(maxPriorityFeePerGas, maxFeePerGas)
	if err != nil {
		t.Fatalf("ComputeGasFees returned an error: %v", err)
	}

	if result.Cmp(expected) != 0 {
		t.Errorf("ComputeGasFees result %s; expected %s", result.String(), expected.String())
	}
}

func TestComputeGasFees_NilInput(t *testing.T) {
	_, err := ComputeGasFees(nil, big.NewInt(1))
	if err == nil {
		t.Error("Expected error when maxPriorityFeePerGas is nil")
	}
}

func TestGetEmptyPackedUserOperation(t *testing.T) {
	userOp, err := GetEmptyPackedUserOperation()
	if err != nil {
		t.Fatalf("GetEmptyPackedUserOperation returned an error: %v", err)
	}

	// Check Sender
	expectedSender := common.HexToAddress("0x0000000000000000000000000000000000000000")
	if userOp.Sender != expectedSender {
		t.Errorf("Sender is %s; expected %s", userOp.Sender.Hex(), expectedSender.Hex())
	}

	// Check Nonce
	if userOp.Nonce.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Nonce is %s; expected %s", userOp.Nonce.String(), "0")
	}

	// Check AccountGasLimits
	verificationGasLimit := big.NewInt(2777216)
	callGasLimit := big.NewInt(2777216)
	accountGasLimitsInt, _ := ComputeAccountGasLimits(verificationGasLimit, callGasLimit)
	expectedAccountGasLimits := ToBeHex(accountGasLimitsInt, 32)
	if userOp.AccountGasLimits != expectedAccountGasLimits {
		t.Errorf("AccountGasLimits is %s; expected %s", userOp.AccountGasLimits, expectedAccountGasLimits)
	}

	// Check GasFees
	maxPriorityFeePerGas := big.NewInt(256)
	maxFeePerGas := big.NewInt(256)
	gasFeesInt, _ := ComputeGasFees(maxPriorityFeePerGas, maxFeePerGas)
	expectedGasFees := ToBeHex(gasFeesInt, 32)
	if userOp.GasFees != expectedGasFees {
		t.Errorf("GasFees is %s; expected %s", userOp.GasFees, expectedGasFees)
	}

	// Check PreVerificationGas
	if userOp.PreVerificationGas.Cmp(verificationGasLimit) != 0 {
		t.Errorf("PreVerificationGas is %s; expected %s", userOp.PreVerificationGas.String(), verificationGasLimit.String())
	}
}

func TestToBeHex_EdgeCases(t *testing.T) {
	value := big.NewInt(0)
	size := 0
	expected := "0x"
	result := ToBeHex(value, size)
	if result != expected {
		t.Errorf("ToBeHex(%v, %d) = %s; expected %s", value, size, result, expected)
	}
}

func TestSignRawPoseidon(t *testing.T) {
	testPK := "18586133768512220936620570745912940619677854269274689475585506675881198879027"
	testMessage := "0xbea17b6896df76e4389d5ea75eaffc6c0362801995ad8695356a5d50965d69e6"

	expectedR8X := "13989537675992952592749254734422527013515207516029614061867549975534361158287"
	expectedR8Y := "12183525468408496816700163568705869208055039206371131047872828002085851855615"
	expectedS := "19232498072797748515268330878939422128145160512284888529599501925137143411"

	R8, S, err := SignRawPoseidon(testPK, testMessage)
	if err != nil {
		t.Fatalf("SignRawPoseidon returned an error: %v", err)
	}

	if R8.X.String() != expectedR8X {
		t.Errorf("R8.X is %s; expected %s", R8.X.String(), expectedR8X)
	}

	if R8.Y.String() != expectedR8Y {
		t.Errorf("R8.Y is %s; expected %s", R8.Y.String(), expectedR8Y)
	}

	if S.String() != expectedS {
		t.Errorf("S is %s; expected %s", S.String(), expectedS)
	}
}
