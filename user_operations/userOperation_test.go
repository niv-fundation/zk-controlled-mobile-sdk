package user_operations

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

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

func TestDecomputeGasFees(t *testing.T) {
	maxPriorityFeePerGas := big.NewInt(256)
	maxFeePerGas := big.NewInt(256)

	gasFeesInt, err := ComputeGasFees(maxPriorityFeePerGas, maxFeePerGas)
	if err != nil {
		t.Fatalf("ComputeGasFees returned an error: %v", err)
	}

	decomputedMaxPriorityFeePerGas, decomputedMaxFeePerGas, err := DecomputeGasFees(gasFeesInt.String())
	if err != nil {
		t.Fatalf("DecomputeGasFees returned an error: %v", err)
	}

	if maxPriorityFeePerGas.Cmp(decomputedMaxPriorityFeePerGas) != 0 {
		t.Errorf("MaxPriorityFeePerGas is %s; expected %s", decomputedMaxPriorityFeePerGas.String(), maxPriorityFeePerGas.String())
	}

	if maxFeePerGas.Cmp(decomputedMaxFeePerGas) != 0 {
		t.Errorf("MaxFeePerGas is %s; expected %s", decomputedMaxFeePerGas.String(), maxFeePerGas.String())
	}
}
