package user_operations

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

// TestAccount_GetInitCode tests the GetInitCode function.
// Works only with a local Ethereum client running on http://127.0.0.1:8545
func TestAccount_GetInitCode(t *testing.T) {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Println("Error connecting to the Ethereum client:", err)
		return
	}
	defer client.Close()

	factoryAddressStr := "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9"
	nullifier := "0x0c3d49b8f4939566d014c0751cef7da0bfad2fc9ab52ba583738b8098985f4c4"

	initCode, predictedAddress, err := GetInitCode(client, factoryAddressStr, nullifier)
	if err != nil {
		t.Errorf("Error getting init code: %v", err)
	}

	expectedInitCode := "0xcf7ed3acca5a467e9e704c703e8d87f634fb0fc99cc3509e0c3d49b8f4939566d014c0751cef7da0bfad2fc9ab52ba583738b8098985f4c4"
	expectedPredictedAddress := "0xb85c7dca274328bc69e39e21273BBbAF5776a352"

	initCodeHex := hexutil.Encode(initCode)

	if initCodeHex != expectedInitCode {
		t.Errorf("Init code %s; expected %s", initCode, expectedInitCode)
	}

	if predictedAddress.String() != expectedPredictedAddress {
		t.Errorf("Predicted address %s; expected %s", predictedAddress, expectedPredictedAddress)
	}
}

//func TestGetSendETHData(t *testing.T) {
//	starter := EthereumClient{}
//	client := starter.NewEthereumClient("http://127.0.0.1:8545")
//
//	privateKey := "2683859904373824334341583944441165871856632302722023367041511169550141504364"
//	eventId := "5"
//	receiver := "0xb85c7dca274328bc69e39e21273BBbAF5776a352"
//	amount := "1000000000000000000"
//	accountAddress := "0x0000000000000000000000000000000000000000"
//	factoryAddressStr := "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9"
//	entryPointStr := "0x5FbDB2315678afecb367f032d93F642f64180aa3"
//
//	inputs, err := client.GetSendETHInputs(privateKey, eventId, receiver, amount, accountAddress, factoryAddressStr, entryPointStr)
//	if err != nil {
//		t.Errorf("Error getting send ETH inputs: %v", err)
//	}
//
//	da, err := AuthProofInputFromJSON(inputs)
//	if err != nil {
//		t.Errorf("Error parsing send ETH inputs: %v", err)
//	}
//
//	expectedMessageHash := "101373865665913786847369793212464126103882317415272150032212672803257237916460"
//	expectedSkI := "2683859904373824334341583944441165871856632302722023367041511169550141504364"
//	expectedEventID := "5"
//	expectedSignatureR8x := "2134845251190679957886296794883468566760867366334978115767504948768742331145"
//	expectedSignatureR8y := "3552352080662711047276167584440132768926267183192119239406418413503932103947"
//	expectedSignatureS := "139913670108218062601078696993570022100595631347248102325083436199301545708"
//
//	if da.MessageHash.String() != expectedMessageHash {
//		t.Errorf("Message hash %s; expected %s", da.MessageHash.String(), expectedMessageHash)
//	}
//
//	if da.SkI.String() != expectedSkI {
//		t.Errorf("Secret key %s; expected %s", da.SkI.String(), expectedSkI)
//	}
//
//	if da.EventID.String() != expectedEventID {
//		t.Errorf("Event ID %s; expected %s", da.EventID.String(), expectedEventID)
//	}
//
//	if da.SignatureR8x.String() != expectedSignatureR8x {
//		t.Errorf("Signature R8x %s; expected %s", da.SignatureR8x.String(), expectedSignatureR8x)
//	}
//
//	if da.SignatureR8y.String() != expectedSignatureR8y {
//		t.Errorf("Signature R8y %s; expected %s", da.SignatureR8y.String(), expectedSignatureR8y)
//	}
//
//	if da.SignatureS.String() != expectedSignatureS {
//		t.Errorf("Signature S %s; expected %s", da.SignatureS.String(), expectedSignatureS)
//	}
//}
