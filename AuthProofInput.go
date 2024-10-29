package zk_controlled_mobile_sdk

import (
	"encoding/json"
	"errors"
	"math/big"
)

type AuthProofInput struct {
	SkI          *big.Int `json:"sk_i"`
	EventID      *big.Int `json:"eventID"`
	MessageHash  *big.Int `json:"messageHash"`
	SignatureR8x *big.Int `json:"signatureR8x"`
	SignatureR8y *big.Int `json:"signatureR8y"`
	SignatureS   *big.Int `json:"signatureS"`
}

var Zero = big.NewInt(0)

func (i AuthProofInput) Validate() error {
	if i.SkI.Cmp(Zero) <= 0 {
		return errors.New("invalid sk_i value")
	}
	if i.EventID.Cmp(Zero) <= 0 {
		return errors.New("invalid eventID value")
	}
	if i.MessageHash.Cmp(Zero) <= 0 {
		return errors.New("invalid messageHash value")
	}
	if i.SignatureR8x.Cmp(Zero) <= 0 {
		return errors.New("invalid signatureR8x value")
	}
	if i.SignatureR8y.Cmp(Zero) <= 0 {
		return errors.New("invalid signatureR8y value")
	}
	if i.SignatureS.Cmp(Zero) <= 0 {
		return errors.New("invalid signatureS value")
	}
	return nil
}

// ToMap converts the struct to a map with string values for easy JSON serialization
func (i AuthProofInput) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"sk_i":         i.SkI.String(),
		"eventID":      i.EventID.String(),
		"messageHash":  i.MessageHash.String(),
		"signatureR8x": i.SignatureR8x.String(),
		"signatureR8y": i.SignatureR8y.String(),
		"signatureS":   i.SignatureS.String(),
	}
}

func (i AuthProofInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ToMap())
}

func AuthProofInputFromJSON(data string) (AuthProofInput, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		return AuthProofInput{}, err
	}

	skI, err := ParseBigInt(m["sk_i"].(string))
	if err != nil {
		return AuthProofInput{}, err
	}

	eventID, err := ParseBigInt(m["eventID"].(string))
	if err != nil {
		return AuthProofInput{}, err
	}

	messageHash, err := ParseBigInt(m["messageHash"].(string))
	if err != nil {
		return AuthProofInput{}, err
	}

	signatureR8x, err := ParseBigInt(m["signatureR8x"].(string))
	if err != nil {
		return AuthProofInput{}, err
	}

	signatureR8y, err := ParseBigInt(m["signatureR8y"].(string))
	if err != nil {
		return AuthProofInput{}, err
	}

	signatureS, err := ParseBigInt(m["signatureS"].(string))
	if err != nil {
		return AuthProofInput{}, err
	}

	return AuthProofInput{
		SkI:          skI,
		EventID:      eventID,
		MessageHash:  messageHash,
		SignatureR8x: signatureR8x,
		SignatureR8y: signatureR8y,
		SignatureS:   signatureS,
	}, nil
}
