// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zk_controlled_mobile_sdk

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SmartAccountFactoryMetaData contains all meta data concerning the SmartAccountFactory contract.
var SmartAccountFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"SmartAccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"smartAccountImplementation_\",\"type\":\"address\"}],\"name\":\"__SmartAccountFactory_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nullifier_\",\"type\":\"bytes32\"}],\"name\":\"deploySmartAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nullifier_\",\"type\":\"bytes32\"}],\"name\":\"getSmartAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSmartAccountImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nullifier_\",\"type\":\"bytes32\"}],\"name\":\"predictSmartAccountAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"setSmartAccountImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"smartAccounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// SmartAccountFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SmartAccountFactoryMetaData.ABI instead.
var SmartAccountFactoryABI = SmartAccountFactoryMetaData.ABI

// SmartAccountFactory is an auto generated Go binding around an Ethereum contract.
type SmartAccountFactory struct {
	SmartAccountFactoryCaller     // Read-only binding to the contract
	SmartAccountFactoryTransactor // Write-only binding to the contract
	SmartAccountFactoryFilterer   // Log filterer for contract events
}

// SmartAccountFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartAccountFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartAccountFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartAccountFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartAccountFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartAccountFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartAccountFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartAccountFactorySession struct {
	Contract     *SmartAccountFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SmartAccountFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartAccountFactoryCallerSession struct {
	Contract *SmartAccountFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SmartAccountFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartAccountFactoryTransactorSession struct {
	Contract     *SmartAccountFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SmartAccountFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartAccountFactoryRaw struct {
	Contract *SmartAccountFactory // Generic contract binding to access the raw methods on
}

// SmartAccountFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartAccountFactoryCallerRaw struct {
	Contract *SmartAccountFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SmartAccountFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartAccountFactoryTransactorRaw struct {
	Contract *SmartAccountFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartAccountFactory creates a new instance of SmartAccountFactory, bound to a specific deployed contract.
func NewSmartAccountFactory(address common.Address, backend bind.ContractBackend) (*SmartAccountFactory, error) {
	contract, err := bindSmartAccountFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SmartAccountFactory{SmartAccountFactoryCaller: SmartAccountFactoryCaller{contract: contract}, SmartAccountFactoryTransactor: SmartAccountFactoryTransactor{contract: contract}, SmartAccountFactoryFilterer: SmartAccountFactoryFilterer{contract: contract}}, nil
}

// NewSmartAccountFactoryCaller creates a new read-only instance of SmartAccountFactory, bound to a specific deployed contract.
func NewSmartAccountFactoryCaller(address common.Address, caller bind.ContractCaller) (*SmartAccountFactoryCaller, error) {
	contract, err := bindSmartAccountFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartAccountFactoryCaller{contract: contract}, nil
}

// NewSmartAccountFactoryTransactor creates a new write-only instance of SmartAccountFactory, bound to a specific deployed contract.
func NewSmartAccountFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartAccountFactoryTransactor, error) {
	contract, err := bindSmartAccountFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartAccountFactoryTransactor{contract: contract}, nil
}

// NewSmartAccountFactoryFilterer creates a new log filterer instance of SmartAccountFactory, bound to a specific deployed contract.
func NewSmartAccountFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartAccountFactoryFilterer, error) {
	contract, err := bindSmartAccountFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartAccountFactoryFilterer{contract: contract}, nil
}

// bindSmartAccountFactory binds a generic wrapper to an already deployed contract.
func bindSmartAccountFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SmartAccountFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartAccountFactory *SmartAccountFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SmartAccountFactory.Contract.SmartAccountFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartAccountFactory *SmartAccountFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.SmartAccountFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartAccountFactory *SmartAccountFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.SmartAccountFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartAccountFactory *SmartAccountFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SmartAccountFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartAccountFactory *SmartAccountFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartAccountFactory *SmartAccountFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SmartAccountFactory *SmartAccountFactoryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SmartAccountFactory.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SmartAccountFactory *SmartAccountFactorySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SmartAccountFactory.Contract.UPGRADEINTERFACEVERSION(&_SmartAccountFactory.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SmartAccountFactory *SmartAccountFactoryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SmartAccountFactory.Contract.UPGRADEINTERFACEVERSION(&_SmartAccountFactory.CallOpts)
}

// GetSmartAccount is a free data retrieval call binding the contract method 0x703b48f0.
//
// Solidity: function getSmartAccount(bytes32 nullifier_) view returns(address)
func (_SmartAccountFactory *SmartAccountFactoryCaller) GetSmartAccount(opts *bind.CallOpts, nullifier_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _SmartAccountFactory.contract.Call(opts, &out, "getSmartAccount", nullifier_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSmartAccount is a free data retrieval call binding the contract method 0x703b48f0.
//
// Solidity: function getSmartAccount(bytes32 nullifier_) view returns(address)
func (_SmartAccountFactory *SmartAccountFactorySession) GetSmartAccount(nullifier_ [32]byte) (common.Address, error) {
	return _SmartAccountFactory.Contract.GetSmartAccount(&_SmartAccountFactory.CallOpts, nullifier_)
}

// GetSmartAccount is a free data retrieval call binding the contract method 0x703b48f0.
//
// Solidity: function getSmartAccount(bytes32 nullifier_) view returns(address)
func (_SmartAccountFactory *SmartAccountFactoryCallerSession) GetSmartAccount(nullifier_ [32]byte) (common.Address, error) {
	return _SmartAccountFactory.Contract.GetSmartAccount(&_SmartAccountFactory.CallOpts, nullifier_)
}

// GetSmartAccountImplementation is a free data retrieval call binding the contract method 0x71ea7eda.
//
// Solidity: function getSmartAccountImplementation() view returns(address)
func (_SmartAccountFactory *SmartAccountFactoryCaller) GetSmartAccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SmartAccountFactory.contract.Call(opts, &out, "getSmartAccountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSmartAccountImplementation is a free data retrieval call binding the contract method 0x71ea7eda.
//
// Solidity: function getSmartAccountImplementation() view returns(address)
func (_SmartAccountFactory *SmartAccountFactorySession) GetSmartAccountImplementation() (common.Address, error) {
	return _SmartAccountFactory.Contract.GetSmartAccountImplementation(&_SmartAccountFactory.CallOpts)
}

// GetSmartAccountImplementation is a free data retrieval call binding the contract method 0x71ea7eda.
//
// Solidity: function getSmartAccountImplementation() view returns(address)
func (_SmartAccountFactory *SmartAccountFactoryCallerSession) GetSmartAccountImplementation() (common.Address, error) {
	return _SmartAccountFactory.Contract.GetSmartAccountImplementation(&_SmartAccountFactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SmartAccountFactory *SmartAccountFactoryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SmartAccountFactory.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SmartAccountFactory *SmartAccountFactorySession) Implementation() (common.Address, error) {
	return _SmartAccountFactory.Contract.Implementation(&_SmartAccountFactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SmartAccountFactory *SmartAccountFactoryCallerSession) Implementation() (common.Address, error) {
	return _SmartAccountFactory.Contract.Implementation(&_SmartAccountFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SmartAccountFactory *SmartAccountFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SmartAccountFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SmartAccountFactory *SmartAccountFactorySession) Owner() (common.Address, error) {
	return _SmartAccountFactory.Contract.Owner(&_SmartAccountFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SmartAccountFactory *SmartAccountFactoryCallerSession) Owner() (common.Address, error) {
	return _SmartAccountFactory.Contract.Owner(&_SmartAccountFactory.CallOpts)
}

// PredictSmartAccountAddress is a free data retrieval call binding the contract method 0xbb3ed396.
//
// Solidity: function predictSmartAccountAddress(bytes32 nullifier_) view returns(address)
func (_SmartAccountFactory *SmartAccountFactoryCaller) PredictSmartAccountAddress(opts *bind.CallOpts, nullifier_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _SmartAccountFactory.contract.Call(opts, &out, "predictSmartAccountAddress", nullifier_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PredictSmartAccountAddress is a free data retrieval call binding the contract method 0xbb3ed396.
//
// Solidity: function predictSmartAccountAddress(bytes32 nullifier_) view returns(address)
func (_SmartAccountFactory *SmartAccountFactorySession) PredictSmartAccountAddress(nullifier_ [32]byte) (common.Address, error) {
	return _SmartAccountFactory.Contract.PredictSmartAccountAddress(&_SmartAccountFactory.CallOpts, nullifier_)
}

// PredictSmartAccountAddress is a free data retrieval call binding the contract method 0xbb3ed396.
//
// Solidity: function predictSmartAccountAddress(bytes32 nullifier_) view returns(address)
func (_SmartAccountFactory *SmartAccountFactoryCallerSession) PredictSmartAccountAddress(nullifier_ [32]byte) (common.Address, error) {
	return _SmartAccountFactory.Contract.PredictSmartAccountAddress(&_SmartAccountFactory.CallOpts, nullifier_)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SmartAccountFactory *SmartAccountFactoryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SmartAccountFactory.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SmartAccountFactory *SmartAccountFactorySession) ProxiableUUID() ([32]byte, error) {
	return _SmartAccountFactory.Contract.ProxiableUUID(&_SmartAccountFactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SmartAccountFactory *SmartAccountFactoryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _SmartAccountFactory.Contract.ProxiableUUID(&_SmartAccountFactory.CallOpts)
}

// SmartAccounts is a free data retrieval call binding the contract method 0x17c8ae52.
//
// Solidity: function smartAccounts(bytes32 ) view returns(address)
func (_SmartAccountFactory *SmartAccountFactoryCaller) SmartAccounts(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _SmartAccountFactory.contract.Call(opts, &out, "smartAccounts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SmartAccounts is a free data retrieval call binding the contract method 0x17c8ae52.
//
// Solidity: function smartAccounts(bytes32 ) view returns(address)
func (_SmartAccountFactory *SmartAccountFactorySession) SmartAccounts(arg0 [32]byte) (common.Address, error) {
	return _SmartAccountFactory.Contract.SmartAccounts(&_SmartAccountFactory.CallOpts, arg0)
}

// SmartAccounts is a free data retrieval call binding the contract method 0x17c8ae52.
//
// Solidity: function smartAccounts(bytes32 ) view returns(address)
func (_SmartAccountFactory *SmartAccountFactoryCallerSession) SmartAccounts(arg0 [32]byte) (common.Address, error) {
	return _SmartAccountFactory.Contract.SmartAccounts(&_SmartAccountFactory.CallOpts, arg0)
}

// SmartAccountFactoryInit is a paid mutator transaction binding the contract method 0x6e17a538.
//
// Solidity: function __SmartAccountFactory_init(address smartAccountImplementation_) returns()
func (_SmartAccountFactory *SmartAccountFactoryTransactor) SmartAccountFactoryInit(opts *bind.TransactOpts, smartAccountImplementation_ common.Address) (*types.Transaction, error) {
	return _SmartAccountFactory.contract.Transact(opts, "__SmartAccountFactory_init", smartAccountImplementation_)
}

// SmartAccountFactoryInit is a paid mutator transaction binding the contract method 0x6e17a538.
//
// Solidity: function __SmartAccountFactory_init(address smartAccountImplementation_) returns()
func (_SmartAccountFactory *SmartAccountFactorySession) SmartAccountFactoryInit(smartAccountImplementation_ common.Address) (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.SmartAccountFactoryInit(&_SmartAccountFactory.TransactOpts, smartAccountImplementation_)
}

// SmartAccountFactoryInit is a paid mutator transaction binding the contract method 0x6e17a538.
//
// Solidity: function __SmartAccountFactory_init(address smartAccountImplementation_) returns()
func (_SmartAccountFactory *SmartAccountFactoryTransactorSession) SmartAccountFactoryInit(smartAccountImplementation_ common.Address) (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.SmartAccountFactoryInit(&_SmartAccountFactory.TransactOpts, smartAccountImplementation_)
}

// DeploySmartAccount is a paid mutator transaction binding the contract method 0x9cc3509e.
//
// Solidity: function deploySmartAccount(bytes32 nullifier_) returns(address)
func (_SmartAccountFactory *SmartAccountFactoryTransactor) DeploySmartAccount(opts *bind.TransactOpts, nullifier_ [32]byte) (*types.Transaction, error) {
	return _SmartAccountFactory.contract.Transact(opts, "deploySmartAccount", nullifier_)
}

// DeploySmartAccount is a paid mutator transaction binding the contract method 0x9cc3509e.
//
// Solidity: function deploySmartAccount(bytes32 nullifier_) returns(address)
func (_SmartAccountFactory *SmartAccountFactorySession) DeploySmartAccount(nullifier_ [32]byte) (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.DeploySmartAccount(&_SmartAccountFactory.TransactOpts, nullifier_)
}

// DeploySmartAccount is a paid mutator transaction binding the contract method 0x9cc3509e.
//
// Solidity: function deploySmartAccount(bytes32 nullifier_) returns(address)
func (_SmartAccountFactory *SmartAccountFactoryTransactorSession) DeploySmartAccount(nullifier_ [32]byte) (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.DeploySmartAccount(&_SmartAccountFactory.TransactOpts, nullifier_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SmartAccountFactory *SmartAccountFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartAccountFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SmartAccountFactory *SmartAccountFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.RenounceOwnership(&_SmartAccountFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SmartAccountFactory *SmartAccountFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.RenounceOwnership(&_SmartAccountFactory.TransactOpts)
}

// SetSmartAccountImplementation is a paid mutator transaction binding the contract method 0xceac3e72.
//
// Solidity: function setSmartAccountImplementation(address newImplementation) returns()
func (_SmartAccountFactory *SmartAccountFactoryTransactor) SetSmartAccountImplementation(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _SmartAccountFactory.contract.Transact(opts, "setSmartAccountImplementation", newImplementation)
}

// SetSmartAccountImplementation is a paid mutator transaction binding the contract method 0xceac3e72.
//
// Solidity: function setSmartAccountImplementation(address newImplementation) returns()
func (_SmartAccountFactory *SmartAccountFactorySession) SetSmartAccountImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.SetSmartAccountImplementation(&_SmartAccountFactory.TransactOpts, newImplementation)
}

// SetSmartAccountImplementation is a paid mutator transaction binding the contract method 0xceac3e72.
//
// Solidity: function setSmartAccountImplementation(address newImplementation) returns()
func (_SmartAccountFactory *SmartAccountFactoryTransactorSession) SetSmartAccountImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.SetSmartAccountImplementation(&_SmartAccountFactory.TransactOpts, newImplementation)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SmartAccountFactory *SmartAccountFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SmartAccountFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SmartAccountFactory *SmartAccountFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.TransferOwnership(&_SmartAccountFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SmartAccountFactory *SmartAccountFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.TransferOwnership(&_SmartAccountFactory.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SmartAccountFactory *SmartAccountFactoryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SmartAccountFactory.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SmartAccountFactory *SmartAccountFactorySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.UpgradeToAndCall(&_SmartAccountFactory.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SmartAccountFactory *SmartAccountFactoryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SmartAccountFactory.Contract.UpgradeToAndCall(&_SmartAccountFactory.TransactOpts, newImplementation, data)
}

// SmartAccountFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SmartAccountFactory contract.
type SmartAccountFactoryInitializedIterator struct {
	Event *SmartAccountFactoryInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartAccountFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountFactoryInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartAccountFactoryInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartAccountFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountFactoryInitialized represents a Initialized event raised by the SmartAccountFactory contract.
type SmartAccountFactoryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SmartAccountFactory *SmartAccountFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*SmartAccountFactoryInitializedIterator, error) {

	logs, sub, err := _SmartAccountFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SmartAccountFactoryInitializedIterator{contract: _SmartAccountFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SmartAccountFactory *SmartAccountFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SmartAccountFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _SmartAccountFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountFactoryInitialized)
				if err := _SmartAccountFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SmartAccountFactory *SmartAccountFactoryFilterer) ParseInitialized(log types.Log) (*SmartAccountFactoryInitialized, error) {
	event := new(SmartAccountFactoryInitialized)
	if err := _SmartAccountFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartAccountFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SmartAccountFactory contract.
type SmartAccountFactoryOwnershipTransferredIterator struct {
	Event *SmartAccountFactoryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartAccountFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountFactoryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartAccountFactoryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartAccountFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the SmartAccountFactory contract.
type SmartAccountFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SmartAccountFactory *SmartAccountFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SmartAccountFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SmartAccountFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SmartAccountFactoryOwnershipTransferredIterator{contract: _SmartAccountFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SmartAccountFactory *SmartAccountFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SmartAccountFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SmartAccountFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountFactoryOwnershipTransferred)
				if err := _SmartAccountFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SmartAccountFactory *SmartAccountFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*SmartAccountFactoryOwnershipTransferred, error) {
	event := new(SmartAccountFactoryOwnershipTransferred)
	if err := _SmartAccountFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartAccountFactorySmartAccountDeployedIterator is returned from FilterSmartAccountDeployed and is used to iterate over the raw logs and unpacked data for SmartAccountDeployed events raised by the SmartAccountFactory contract.
type SmartAccountFactorySmartAccountDeployedIterator struct {
	Event *SmartAccountFactorySmartAccountDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartAccountFactorySmartAccountDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountFactorySmartAccountDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartAccountFactorySmartAccountDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartAccountFactorySmartAccountDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountFactorySmartAccountDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountFactorySmartAccountDeployed represents a SmartAccountDeployed event raised by the SmartAccountFactory contract.
type SmartAccountFactorySmartAccountDeployed struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSmartAccountDeployed is a free log retrieval operation binding the contract event 0x6b472ccbb9308c4dcddb80184d3f428e35e4991f03f1fe46a1e034b0c3de639e.
//
// Solidity: event SmartAccountDeployed(address indexed account)
func (_SmartAccountFactory *SmartAccountFactoryFilterer) FilterSmartAccountDeployed(opts *bind.FilterOpts, account []common.Address) (*SmartAccountFactorySmartAccountDeployedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SmartAccountFactory.contract.FilterLogs(opts, "SmartAccountDeployed", accountRule)
	if err != nil {
		return nil, err
	}
	return &SmartAccountFactorySmartAccountDeployedIterator{contract: _SmartAccountFactory.contract, event: "SmartAccountDeployed", logs: logs, sub: sub}, nil
}

// WatchSmartAccountDeployed is a free log subscription operation binding the contract event 0x6b472ccbb9308c4dcddb80184d3f428e35e4991f03f1fe46a1e034b0c3de639e.
//
// Solidity: event SmartAccountDeployed(address indexed account)
func (_SmartAccountFactory *SmartAccountFactoryFilterer) WatchSmartAccountDeployed(opts *bind.WatchOpts, sink chan<- *SmartAccountFactorySmartAccountDeployed, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SmartAccountFactory.contract.WatchLogs(opts, "SmartAccountDeployed", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountFactorySmartAccountDeployed)
				if err := _SmartAccountFactory.contract.UnpackLog(event, "SmartAccountDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSmartAccountDeployed is a log parse operation binding the contract event 0x6b472ccbb9308c4dcddb80184d3f428e35e4991f03f1fe46a1e034b0c3de639e.
//
// Solidity: event SmartAccountDeployed(address indexed account)
func (_SmartAccountFactory *SmartAccountFactoryFilterer) ParseSmartAccountDeployed(log types.Log) (*SmartAccountFactorySmartAccountDeployed, error) {
	event := new(SmartAccountFactorySmartAccountDeployed)
	if err := _SmartAccountFactory.contract.UnpackLog(event, "SmartAccountDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartAccountFactoryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SmartAccountFactory contract.
type SmartAccountFactoryUpgradedIterator struct {
	Event *SmartAccountFactoryUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartAccountFactoryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountFactoryUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartAccountFactoryUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartAccountFactoryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountFactoryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountFactoryUpgraded represents a Upgraded event raised by the SmartAccountFactory contract.
type SmartAccountFactoryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SmartAccountFactory *SmartAccountFactoryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SmartAccountFactoryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SmartAccountFactory.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SmartAccountFactoryUpgradedIterator{contract: _SmartAccountFactory.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SmartAccountFactory *SmartAccountFactoryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SmartAccountFactoryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SmartAccountFactory.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountFactoryUpgraded)
				if err := _SmartAccountFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SmartAccountFactory *SmartAccountFactoryFilterer) ParseUpgraded(log types.Log) (*SmartAccountFactoryUpgraded, error) {
	event := new(SmartAccountFactoryUpgraded)
	if err := _SmartAccountFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
