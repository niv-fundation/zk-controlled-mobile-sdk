package zk_controlled_mobile_sdk

import (
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
