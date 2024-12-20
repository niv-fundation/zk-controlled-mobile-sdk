package user_operations

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/niv-fundation/zk-controlled-mobile-sdk/bindings"
	"math/big"
)

// Account represents the smart account.
type Account struct {
	Address         common.Address
	Paymaster       common.Address
	AccountInstance *bindings.SmartAccount
	Client          *ethclient.Client
}

func NewAccount(address, paymaster common.Address, client *ethclient.Client) *Account {
	contractInstance, err := bindings.NewSmartAccount(address, client)
	if err != nil {
		panic(err)
	}

	return &Account{
		Address:         address,
		Paymaster:       paymaster,
		Client:          client,
		AccountInstance: contractInstance,
	}
}

// GetAddress returns the address of the smart account.
func (sa *Account) GetAddress() (common.Address, error) {
	// Implement the logic to get the address
	return sa.Address, nil
}

// GetCurrentNonce returns the current nonce of the smart account.
func (sa *Account) GetCurrentNonce() (*big.Int, error) {
	nonce, err := sa.AccountInstance.GetCurrentNonce(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return nonce, nil
}
