package zk_controlled_mobile_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
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

func (p *Proof) ToVerifierHelperProofPoints() (*VerifierHelperProofPoints, error) {
	var vhp VerifierHelperProofPoints

	if len(p.PiA) != 2 {
		return nil, errors.New("pi_a must have exactly 2 elements")
	}
	for i := 0; i < 2; i++ {
		vhp.A[i] = MustParseBigInt(p.PiA[i])
	}

	if len(p.PiB) != 2 {
		return nil, errors.New("pi_b must have exactly 2 elements")
	}
	for i := 0; i < 2; i++ {
		if len(p.PiB[i]) != 2 {
			return nil, fmt.Errorf("pi_b[%d] must have exactly 2 elements", i)
		}
		for j := 0; j < 2; j++ {
			vhp.B[i][j] = MustParseBigInt(p.PiB[i][j])
		}
	}

	if len(p.PiC) != 2 {
		return nil, errors.New("pi_c must have exactly 2 elements")
	}
	for i := 0; i < 2; i++ {
		vhp.C[i] = MustParseBigInt(p.PiC[i])
	}

	return &vhp, nil
}

func EncodeIdentityProof(proofPoints *VerifierHelperProofPoints) ([]byte, error) {
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

	identityProof := SmartAccountIdentityProof{
		IdentityProof: *proofPoints,
	}

	// Unmarshal the ABI JSON into abi.Arguments
	var arguments abi.Arguments
	err := json.Unmarshal([]byte(abiJSON), &arguments)
	if err != nil {
		return nil, err
	}

	// Pack the data
	data, err := arguments.Pack(identityProof)
	if err != nil {
		return nil, err
	}

	return data, nil
}
