package user_operations

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendUOToBundler(userOp *UserOperationJson, paymasterAddressStr, entryPointStr string) (string, error) {
	userOperation, err := ConstructUserOperationRequest(userOp, paymasterAddressStr)
	if err != nil {
		return "", fmt.Errorf("error constructing user operation: %v", err)
	}

	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "eth_sendUserOperation",
		"params": []interface{}{
			userOperation,
			entryPointStr,
		},
	}

	// Serialize the request body to JSON
	requestData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request body: %v", err)
	}

	serviceURL := "https://eth-sepolia.g.alchemy.com/v2/-Jm280LvZnniBfiaxZtQa_wL1b_okXCZ"
	resp, err := http.Post(serviceURL, "application/json", bytes.NewReader(requestData))
	if err != nil {
		return "", fmt.Errorf("error sending request to service: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	var response struct {
		Jsonrpc string `json:"jsonrpc"`
		Id      int    `json:"id"`
		Result  string `json:"result,omitempty"`
		Error   *struct {
			Code    int             `json:"code"`
			Message string          `json:"message"`
			Data    json.RawMessage `json:"data,omitempty"`
		} `json:"error,omitempty"`
	}

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling response: %v", err)
	}

	if response.Error != nil {
		var StakeTooLowData struct {
			NeedsStake      json.RawMessage `json:"needsStake"`
			AccessingEntity json.RawMessage `json:"accessingEntity"`
			AccessedAddress json.RawMessage `json:"accessedAddress"`
			AccessedEntity  json.RawMessage `json:"accessedEntity"`
			Slot            json.RawMessage `json:"slot"`
			MinimumStake    json.RawMessage `json:"minimumStake"`
			MinimumUnstake  json.RawMessage `json:"minimumUnstakeDelay"`
		}

		// Try to parse the error data if it exists.
		if len(response.Error.Data) > 0 {
			err = json.Unmarshal(response.Error.Data, &StakeTooLowData)
			if err != nil {
				fmt.Println("Error parsing error data:", err)
				fmt.Println("Needs stake:", hex.EncodeToString(response.Error.Data))
			} else {
				fmt.Println("Needs stake:", hex.EncodeToString(StakeTooLowData.NeedsStake))
				fmt.Println("Accessing entity:", hex.EncodeToString(StakeTooLowData.AccessingEntity))
				fmt.Println("Accessed address:", hex.EncodeToString(StakeTooLowData.AccessedAddress))
				fmt.Println("Accessed entity:", hex.EncodeToString(StakeTooLowData.AccessedEntity))
				fmt.Println("Slot:", hex.EncodeToString(StakeTooLowData.Slot))
				fmt.Println("Minimum stake:", hex.EncodeToString(StakeTooLowData.MinimumStake))
				fmt.Println("Minimum unstake delay:", hex.EncodeToString(StakeTooLowData.MinimumUnstake))
			}
		}

		return "", fmt.Errorf("error from service: %v", response.Error)
	}

	fmt.Println("Transaction hash:", response)

	// Return the transaction hash from the response
	return response.Result, nil
}
