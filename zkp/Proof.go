package zkp

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/niv-fundation/zk-controlled-mobile-sdk/bindings"

	uo "github.com/niv-fundation/zk-controlled-mobile-sdk/user_operations"
)

type Proof struct {
	PiA      []string   `json:"pi_a"`
	PiB      [][]string `json:"pi_b"`
	PiC      []string   `json:"pi_c"`
	Protocol string     `json:"protocol"`
}

func (p *Proof) FromJSON(jsonString string) error {
	return json.Unmarshal([]byte(jsonString), p)
}

type ZkProof struct {
	Proof      Proof    `json:"proof"`
	PubSignals []string `json:"pub_signals"`
}

func (p *Proof) ToVerifierHelperProofPoints() (*bindings.VerifierHelperProofPoints, error) {
	var vhp bindings.VerifierHelperProofPoints

	for i := 0; i < 2; i++ {
		vhp.A[i] = uo.MustParseBigInt(p.PiA[i], "p.PiA[i] in ToVerifierHelperProofPoints")
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			vhp.B[i][j] = uo.MustParseBigInt(p.PiB[i][j], "p.PiB[i][j] in ToVerifierHelperProofPoints")
		}
	}

	for i := 0; i < 2; i++ {
		vhp.C[i] = uo.MustParseBigInt(p.PiC[i], "p.PiC[i] in ToVerifierHelperProofPoints")
	}

	return &vhp, nil
}

func EncodeIdentityProof(proofPoints *bindings.VerifierHelperProofPoints) (string, error) {
	// Define the ABI of the custom type
	abiJSON := `[{
       "components": [
         {
           "components": [
             {
               "internalType": "uint256[2]",
               "name": "a",
               "type": "uint256[2]"
             },
             {
               "internalType": "uint256[2][2]",
               "name": "b",
               "type": "uint256[2][2]"
             },
             {
               "internalType": "uint256[2]",
               "name": "c",
               "type": "uint256[2]"
             }
           ],
           "internalType": "struct VerifierHelper.ProofPoints",
           "name": "identityProof",
           "type": "tuple"
         }
       ],
       "internalType": "struct SmartAccount.IdentityProof",
       "name": "proof_",
       "type": "tuple"
     }]`

	identityProof := bindings.SmartAccountIdentityProof{
		IdentityProof: *proofPoints,
	}

	// Unmarshal the ABI JSON into abi.Arguments
	var arguments abi.Arguments
	err := json.Unmarshal([]byte(abiJSON), &arguments)
	if err != nil {
		return "", err
	}

	// Pack the data
	data, err := arguments.Pack(identityProof)
	if err != nil {
		return "", err
	}

	return hexutil.Encode(data), nil
}
