package zkp

import (
	"testing"
)

func TestEncodeIdentityProof(t *testing.T) {
	// Verified data
	proofData := `{
        "pi_a": [
            "296935625294623977750917323254780655270224064827329102045876697782036738222",
            "14410858609693465947919850925523220400172503798905469833355730539170969286384"
        ],
        "pi_b": [
            [
                "17324811497758163581271558486792198140894105686338810980308828049979436039667",
                "4437811339423783930708345295353662002609579183197921807004666970131088647355"
            ],
            [
                "2271745139522372768254608728341121867836804521421999334595687843339028033188",
                "353152366440981541410371385254798304196963469739506168358664382663516961666"
            ]
        ],
        "pi_c": [
            "13785186540492072226227548468301452067153491525642891500491048816265240393571",
            "8623940764607207732046785735585322919939006678477394903042844051087573740141"
        ],
        "protocol": "groth16"
    }`

	// Parse the Proof
	var proof Proof
	if err := proof.FromJSON(proofData); err != nil {
		t.Fatalf("Error parsing JSON: %v", err)
	}

	// Convert to VerifierHelperProofPoints
	vhp, err := proof.ToVerifierHelperProofPoints()
	if err != nil {
		t.Fatalf("Error converting to VerifierHelperProofPoints: %v", err)
	}

	// Encode the identity proof
	encodedData, err := EncodeIdentityProof(vhp)
	if err != nil {
		t.Fatalf("Error encoding proof: %v", err)
	}

	// Expected output
	expectedHex := "0x00a80f427a1d59a10b569312cc4184bf3e5babe2c3b1840483974c3d630030ae1fdc41c5eae90106cc4b15932ad9f4f0dba9dc2662d27eee4c694fbd876df6f0264d7eedab7b1e18c05347b1ffbcdf380706333766a58cd07c2495aa8b5dfdf309cfb65328664a6bcba2b5f67375ab0c4d76970d6e8d027ecd516944d445a4bb0505c31e716f6babb81e4ed43731ab23bed17f44a451ef3592af0df2ed0e6ea400c7e08d4c1da072c26aca11084f45ab25398107f0baea56074b70208eb33b821e7a239b0b2dba5cfa8b8d25e94fa8d77c25e1c1eef407072536009bbd3637631310fa3d45bf4ef8253e2a9aa6248a65ae34eb28046c70cb0bd68b734de56a6d"

	// Compare the encoded data
	if expectedHex != encodedData {
		t.Errorf("Encoded data does not match expected output.\nGot:      %x\nExpected: %x", encodedData, expectedHex)
	}
}
