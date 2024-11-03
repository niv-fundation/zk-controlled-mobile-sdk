// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// SmartAccountIdentityProof is an auto generated low-level Go binding around an user-defined struct.
type SmartAccountIdentityProof struct {
	IdentityProof VerifierHelperProofPoints
}

// SmartAccountTransactionLog is an auto generated low-level Go binding around an user-defined struct.
type SmartAccountTransactionLog struct {
	Destination common.Address
	Timestamp   *big.Int
	Value       *big.Int
	Data        []byte
}

// VerifierHelperProofPoints is an auto generated low-level Go binding around an user-defined struct.
type VerifierHelperProofPoints struct {
	A [2]*big.Int
	B [2][2]*big.Int
	C [2]*big.Int
}

// SmartAccountMetaData contains all meta data concerning the SmartAccount contract.
var SmartAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"entryPoint_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"identityAuthVerifier_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotFromEntryPoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotFromEntryPointOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotFromThis\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SessionAccountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ENTRY_POINT\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IDENTITY_AUTH_VERIFIER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nullifier_\",\"type\":\"bytes32\"}],\"name\":\"__SmartAccount_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"decodeIdentityProof\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerifierHelper.ProofPoints\",\"name\":\"identityProof\",\"type\":\"tuple\"}],\"internalType\":\"structSmartAccount.IdentityProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerifierHelper.ProofPoints\",\"name\":\"identityProof\",\"type\":\"tuple\"}],\"internalType\":\"structSmartAccount.IdentityProof\",\"name\":\"proof_\",\"type\":\"tuple\"}],\"name\":\"encodeIdentityProof\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"functionData_\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getTransactionHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSmartAccount.TransactionLog[]\",\"name\":\"list_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransactionHistoryLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"history\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nullifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sessionAccounts\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"setSessionAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SmartAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use SmartAccountMetaData.ABI instead.
var SmartAccountABI = SmartAccountMetaData.ABI

// SmartAccount is an auto generated Go binding around an Ethereum contract.
type SmartAccount struct {
	SmartAccountCaller     // Read-only binding to the contract
	SmartAccountTransactor // Write-only binding to the contract
	SmartAccountFilterer   // Log filterer for contract events
}

// SmartAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartAccountSession struct {
	Contract     *SmartAccount     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartAccountCallerSession struct {
	Contract *SmartAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SmartAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartAccountTransactorSession struct {
	Contract     *SmartAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SmartAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartAccountRaw struct {
	Contract *SmartAccount // Generic contract binding to access the raw methods on
}

// SmartAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartAccountCallerRaw struct {
	Contract *SmartAccountCaller // Generic read-only contract binding to access the raw methods on
}

// SmartAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartAccountTransactorRaw struct {
	Contract *SmartAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartAccount creates a new instance of SmartAccount, bound to a specific deployed contract.
func NewSmartAccount(address common.Address, backend bind.ContractBackend) (*SmartAccount, error) {
	contract, err := bindSmartAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SmartAccount{SmartAccountCaller: SmartAccountCaller{contract: contract}, SmartAccountTransactor: SmartAccountTransactor{contract: contract}, SmartAccountFilterer: SmartAccountFilterer{contract: contract}}, nil
}

// NewSmartAccountCaller creates a new read-only instance of SmartAccount, bound to a specific deployed contract.
func NewSmartAccountCaller(address common.Address, caller bind.ContractCaller) (*SmartAccountCaller, error) {
	contract, err := bindSmartAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartAccountCaller{contract: contract}, nil
}

// NewSmartAccountTransactor creates a new write-only instance of SmartAccount, bound to a specific deployed contract.
func NewSmartAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartAccountTransactor, error) {
	contract, err := bindSmartAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartAccountTransactor{contract: contract}, nil
}

// NewSmartAccountFilterer creates a new log filterer instance of SmartAccount, bound to a specific deployed contract.
func NewSmartAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartAccountFilterer, error) {
	contract, err := bindSmartAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartAccountFilterer{contract: contract}, nil
}

// bindSmartAccount binds a generic wrapper to an already deployed contract.
func bindSmartAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SmartAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartAccount *SmartAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SmartAccount.Contract.SmartAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartAccount *SmartAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartAccount.Contract.SmartAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartAccount *SmartAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartAccount.Contract.SmartAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartAccount *SmartAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SmartAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartAccount *SmartAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartAccount *SmartAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartAccount.Contract.contract.Transact(opts, method, params...)
}

// ENTRYPOINT is a free data retrieval call binding the contract method 0x94430fa5.
//
// Solidity: function ENTRY_POINT() view returns(address)
func (_SmartAccount *SmartAccountCaller) ENTRYPOINT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "ENTRY_POINT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ENTRYPOINT is a free data retrieval call binding the contract method 0x94430fa5.
//
// Solidity: function ENTRY_POINT() view returns(address)
func (_SmartAccount *SmartAccountSession) ENTRYPOINT() (common.Address, error) {
	return _SmartAccount.Contract.ENTRYPOINT(&_SmartAccount.CallOpts)
}

// ENTRYPOINT is a free data retrieval call binding the contract method 0x94430fa5.
//
// Solidity: function ENTRY_POINT() view returns(address)
func (_SmartAccount *SmartAccountCallerSession) ENTRYPOINT() (common.Address, error) {
	return _SmartAccount.Contract.ENTRYPOINT(&_SmartAccount.CallOpts)
}

// IDENTITYAUTHVERIFIER is a free data retrieval call binding the contract method 0x966619d0.
//
// Solidity: function IDENTITY_AUTH_VERIFIER() view returns(address)
func (_SmartAccount *SmartAccountCaller) IDENTITYAUTHVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "IDENTITY_AUTH_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IDENTITYAUTHVERIFIER is a free data retrieval call binding the contract method 0x966619d0.
//
// Solidity: function IDENTITY_AUTH_VERIFIER() view returns(address)
func (_SmartAccount *SmartAccountSession) IDENTITYAUTHVERIFIER() (common.Address, error) {
	return _SmartAccount.Contract.IDENTITYAUTHVERIFIER(&_SmartAccount.CallOpts)
}

// IDENTITYAUTHVERIFIER is a free data retrieval call binding the contract method 0x966619d0.
//
// Solidity: function IDENTITY_AUTH_VERIFIER() view returns(address)
func (_SmartAccount *SmartAccountCallerSession) IDENTITYAUTHVERIFIER() (common.Address, error) {
	return _SmartAccount.Contract.IDENTITYAUTHVERIFIER(&_SmartAccount.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SmartAccount *SmartAccountCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SmartAccount *SmartAccountSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SmartAccount.Contract.UPGRADEINTERFACEVERSION(&_SmartAccount.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SmartAccount *SmartAccountCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SmartAccount.Contract.UPGRADEINTERFACEVERSION(&_SmartAccount.CallOpts)
}

// DecodeIdentityProof is a free data retrieval call binding the contract method 0xf6661101.
//
// Solidity: function decodeIdentityProof(bytes data_) pure returns(((uint256[2],uint256[2][2],uint256[2])))
func (_SmartAccount *SmartAccountCaller) DecodeIdentityProof(opts *bind.CallOpts, data_ []byte) (SmartAccountIdentityProof, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "decodeIdentityProof", data_)

	if err != nil {
		return *new(SmartAccountIdentityProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmartAccountIdentityProof)).(*SmartAccountIdentityProof)

	return out0, err

}

// DecodeIdentityProof is a free data retrieval call binding the contract method 0xf6661101.
//
// Solidity: function decodeIdentityProof(bytes data_) pure returns(((uint256[2],uint256[2][2],uint256[2])))
func (_SmartAccount *SmartAccountSession) DecodeIdentityProof(data_ []byte) (SmartAccountIdentityProof, error) {
	return _SmartAccount.Contract.DecodeIdentityProof(&_SmartAccount.CallOpts, data_)
}

// DecodeIdentityProof is a free data retrieval call binding the contract method 0xf6661101.
//
// Solidity: function decodeIdentityProof(bytes data_) pure returns(((uint256[2],uint256[2][2],uint256[2])))
func (_SmartAccount *SmartAccountCallerSession) DecodeIdentityProof(data_ []byte) (SmartAccountIdentityProof, error) {
	return _SmartAccount.Contract.DecodeIdentityProof(&_SmartAccount.CallOpts, data_)
}

// EncodeIdentityProof is a free data retrieval call binding the contract method 0xaef59cf6.
//
// Solidity: function encodeIdentityProof(((uint256[2],uint256[2][2],uint256[2])) proof_) pure returns(bytes)
func (_SmartAccount *SmartAccountCaller) EncodeIdentityProof(opts *bind.CallOpts, proof_ SmartAccountIdentityProof) ([]byte, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "encodeIdentityProof", proof_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeIdentityProof is a free data retrieval call binding the contract method 0xaef59cf6.
//
// Solidity: function encodeIdentityProof(((uint256[2],uint256[2][2],uint256[2])) proof_) pure returns(bytes)
func (_SmartAccount *SmartAccountSession) EncodeIdentityProof(proof_ SmartAccountIdentityProof) ([]byte, error) {
	return _SmartAccount.Contract.EncodeIdentityProof(&_SmartAccount.CallOpts, proof_)
}

// EncodeIdentityProof is a free data retrieval call binding the contract method 0xaef59cf6.
//
// Solidity: function encodeIdentityProof(((uint256[2],uint256[2][2],uint256[2])) proof_) pure returns(bytes)
func (_SmartAccount *SmartAccountCallerSession) EncodeIdentityProof(proof_ SmartAccountIdentityProof) ([]byte, error) {
	return _SmartAccount.Contract.EncodeIdentityProof(&_SmartAccount.CallOpts, proof_)
}

// GetCurrentNonce is a free data retrieval call binding the contract method 0x3a60c386.
//
// Solidity: function getCurrentNonce() view returns(uint256)
func (_SmartAccount *SmartAccountCaller) GetCurrentNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "getCurrentNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentNonce is a free data retrieval call binding the contract method 0x3a60c386.
//
// Solidity: function getCurrentNonce() view returns(uint256)
func (_SmartAccount *SmartAccountSession) GetCurrentNonce() (*big.Int, error) {
	return _SmartAccount.Contract.GetCurrentNonce(&_SmartAccount.CallOpts)
}

// GetCurrentNonce is a free data retrieval call binding the contract method 0x3a60c386.
//
// Solidity: function getCurrentNonce() view returns(uint256)
func (_SmartAccount *SmartAccountCallerSession) GetCurrentNonce() (*big.Int, error) {
	return _SmartAccount.Contract.GetCurrentNonce(&_SmartAccount.CallOpts)
}

// GetTransactionHistory is a free data retrieval call binding the contract method 0xa39c7a9c.
//
// Solidity: function getTransactionHistory(uint256 offset_, uint256 limit_) view returns((address,uint256,uint256,bytes)[] list_)
func (_SmartAccount *SmartAccountCaller) GetTransactionHistory(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]SmartAccountTransactionLog, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "getTransactionHistory", offset_, limit_)

	if err != nil {
		return *new([]SmartAccountTransactionLog), err
	}

	out0 := *abi.ConvertType(out[0], new([]SmartAccountTransactionLog)).(*[]SmartAccountTransactionLog)

	return out0, err

}

// GetTransactionHistory is a free data retrieval call binding the contract method 0xa39c7a9c.
//
// Solidity: function getTransactionHistory(uint256 offset_, uint256 limit_) view returns((address,uint256,uint256,bytes)[] list_)
func (_SmartAccount *SmartAccountSession) GetTransactionHistory(offset_ *big.Int, limit_ *big.Int) ([]SmartAccountTransactionLog, error) {
	return _SmartAccount.Contract.GetTransactionHistory(&_SmartAccount.CallOpts, offset_, limit_)
}

// GetTransactionHistory is a free data retrieval call binding the contract method 0xa39c7a9c.
//
// Solidity: function getTransactionHistory(uint256 offset_, uint256 limit_) view returns((address,uint256,uint256,bytes)[] list_)
func (_SmartAccount *SmartAccountCallerSession) GetTransactionHistory(offset_ *big.Int, limit_ *big.Int) ([]SmartAccountTransactionLog, error) {
	return _SmartAccount.Contract.GetTransactionHistory(&_SmartAccount.CallOpts, offset_, limit_)
}

// GetTransactionHistoryLength is a free data retrieval call binding the contract method 0x9f71c3c9.
//
// Solidity: function getTransactionHistoryLength() view returns(uint256)
func (_SmartAccount *SmartAccountCaller) GetTransactionHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "getTransactionHistoryLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransactionHistoryLength is a free data retrieval call binding the contract method 0x9f71c3c9.
//
// Solidity: function getTransactionHistoryLength() view returns(uint256)
func (_SmartAccount *SmartAccountSession) GetTransactionHistoryLength() (*big.Int, error) {
	return _SmartAccount.Contract.GetTransactionHistoryLength(&_SmartAccount.CallOpts)
}

// GetTransactionHistoryLength is a free data retrieval call binding the contract method 0x9f71c3c9.
//
// Solidity: function getTransactionHistoryLength() view returns(uint256)
func (_SmartAccount *SmartAccountCallerSession) GetTransactionHistoryLength() (*big.Int, error) {
	return _SmartAccount.Contract.GetTransactionHistoryLength(&_SmartAccount.CallOpts)
}

// History is a free data retrieval call binding the contract method 0xa7a38f0b.
//
// Solidity: function history(uint256 ) view returns(address destination, uint256 timestamp, uint256 value, bytes data)
func (_SmartAccount *SmartAccountCaller) History(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Timestamp   *big.Int
	Value       *big.Int
	Data        []byte
}, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "history", arg0)

	outstruct := new(struct {
		Destination common.Address
		Timestamp   *big.Int
		Value       *big.Int
		Data        []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Destination = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// History is a free data retrieval call binding the contract method 0xa7a38f0b.
//
// Solidity: function history(uint256 ) view returns(address destination, uint256 timestamp, uint256 value, bytes data)
func (_SmartAccount *SmartAccountSession) History(arg0 *big.Int) (struct {
	Destination common.Address
	Timestamp   *big.Int
	Value       *big.Int
	Data        []byte
}, error) {
	return _SmartAccount.Contract.History(&_SmartAccount.CallOpts, arg0)
}

// History is a free data retrieval call binding the contract method 0xa7a38f0b.
//
// Solidity: function history(uint256 ) view returns(address destination, uint256 timestamp, uint256 value, bytes data)
func (_SmartAccount *SmartAccountCallerSession) History(arg0 *big.Int) (struct {
	Destination common.Address
	Timestamp   *big.Int
	Value       *big.Int
	Data        []byte
}, error) {
	return _SmartAccount.Contract.History(&_SmartAccount.CallOpts, arg0)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SmartAccount *SmartAccountCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SmartAccount *SmartAccountSession) Implementation() (common.Address, error) {
	return _SmartAccount.Contract.Implementation(&_SmartAccount.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SmartAccount *SmartAccountCallerSession) Implementation() (common.Address, error) {
	return _SmartAccount.Contract.Implementation(&_SmartAccount.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_SmartAccount *SmartAccountCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_SmartAccount *SmartAccountSession) Nonces(owner common.Address) (*big.Int, error) {
	return _SmartAccount.Contract.Nonces(&_SmartAccount.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_SmartAccount *SmartAccountCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _SmartAccount.Contract.Nonces(&_SmartAccount.CallOpts, owner)
}

// Nullifier is a free data retrieval call binding the contract method 0x43bf6cf0.
//
// Solidity: function nullifier() view returns(bytes32)
func (_SmartAccount *SmartAccountCaller) Nullifier(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "nullifier")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Nullifier is a free data retrieval call binding the contract method 0x43bf6cf0.
//
// Solidity: function nullifier() view returns(bytes32)
func (_SmartAccount *SmartAccountSession) Nullifier() ([32]byte, error) {
	return _SmartAccount.Contract.Nullifier(&_SmartAccount.CallOpts)
}

// Nullifier is a free data retrieval call binding the contract method 0x43bf6cf0.
//
// Solidity: function nullifier() view returns(bytes32)
func (_SmartAccount *SmartAccountCallerSession) Nullifier() ([32]byte, error) {
	return _SmartAccount.Contract.Nullifier(&_SmartAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SmartAccount *SmartAccountCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SmartAccount *SmartAccountSession) ProxiableUUID() ([32]byte, error) {
	return _SmartAccount.Contract.ProxiableUUID(&_SmartAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SmartAccount *SmartAccountCallerSession) ProxiableUUID() ([32]byte, error) {
	return _SmartAccount.Contract.ProxiableUUID(&_SmartAccount.CallOpts)
}

// SessionAccounts is a free data retrieval call binding the contract method 0x2511d0ad.
//
// Solidity: function sessionAccounts(address ) view returns(uint48)
func (_SmartAccount *SmartAccountCaller) SessionAccounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "sessionAccounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SessionAccounts is a free data retrieval call binding the contract method 0x2511d0ad.
//
// Solidity: function sessionAccounts(address ) view returns(uint48)
func (_SmartAccount *SmartAccountSession) SessionAccounts(arg0 common.Address) (*big.Int, error) {
	return _SmartAccount.Contract.SessionAccounts(&_SmartAccount.CallOpts, arg0)
}

// SessionAccounts is a free data retrieval call binding the contract method 0x2511d0ad.
//
// Solidity: function sessionAccounts(address ) view returns(uint48)
func (_SmartAccount *SmartAccountCallerSession) SessionAccounts(arg0 common.Address) (*big.Int, error) {
	return _SmartAccount.Contract.SessionAccounts(&_SmartAccount.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SmartAccount *SmartAccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SmartAccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SmartAccount *SmartAccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SmartAccount.Contract.SupportsInterface(&_SmartAccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SmartAccount *SmartAccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SmartAccount.Contract.SupportsInterface(&_SmartAccount.CallOpts, interfaceId)
}

// SmartAccountInit is a paid mutator transaction binding the contract method 0x62a7e336.
//
// Solidity: function __SmartAccount_init(bytes32 nullifier_) returns()
func (_SmartAccount *SmartAccountTransactor) SmartAccountInit(opts *bind.TransactOpts, nullifier_ [32]byte) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "__SmartAccount_init", nullifier_)
}

// SmartAccountInit is a paid mutator transaction binding the contract method 0x62a7e336.
//
// Solidity: function __SmartAccount_init(bytes32 nullifier_) returns()
func (_SmartAccount *SmartAccountSession) SmartAccountInit(nullifier_ [32]byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.SmartAccountInit(&_SmartAccount.TransactOpts, nullifier_)
}

// SmartAccountInit is a paid mutator transaction binding the contract method 0x62a7e336.
//
// Solidity: function __SmartAccount_init(bytes32 nullifier_) returns()
func (_SmartAccount *SmartAccountTransactorSession) SmartAccountInit(nullifier_ [32]byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.SmartAccountInit(&_SmartAccount.TransactOpts, nullifier_)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address destination_, uint256 value_, bytes functionData_) returns()
func (_SmartAccount *SmartAccountTransactor) Execute(opts *bind.TransactOpts, destination_ common.Address, value_ *big.Int, functionData_ []byte) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "execute", destination_, value_, functionData_)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address destination_, uint256 value_, bytes functionData_) returns()
func (_SmartAccount *SmartAccountSession) Execute(destination_ common.Address, value_ *big.Int, functionData_ []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.Execute(&_SmartAccount.TransactOpts, destination_, value_, functionData_)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address destination_, uint256 value_, bytes functionData_) returns()
func (_SmartAccount *SmartAccountTransactorSession) Execute(destination_ common.Address, value_ *big.Int, functionData_ []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.Execute(&_SmartAccount.TransactOpts, destination_, value_, functionData_)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_SmartAccount *SmartAccountTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_SmartAccount *SmartAccountSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.OnERC1155BatchReceived(&_SmartAccount.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_SmartAccount *SmartAccountTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.OnERC1155BatchReceived(&_SmartAccount.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_SmartAccount *SmartAccountTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_SmartAccount *SmartAccountSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.OnERC1155Received(&_SmartAccount.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_SmartAccount *SmartAccountTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.OnERC1155Received(&_SmartAccount.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SetSessionAccount is a paid mutator transaction binding the contract method 0xa5515878.
//
// Solidity: function setSessionAccount(address candidate_, bytes signature_) returns()
func (_SmartAccount *SmartAccountTransactor) SetSessionAccount(opts *bind.TransactOpts, candidate_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "setSessionAccount", candidate_, signature_)
}

// SetSessionAccount is a paid mutator transaction binding the contract method 0xa5515878.
//
// Solidity: function setSessionAccount(address candidate_, bytes signature_) returns()
func (_SmartAccount *SmartAccountSession) SetSessionAccount(candidate_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.SetSessionAccount(&_SmartAccount.TransactOpts, candidate_, signature_)
}

// SetSessionAccount is a paid mutator transaction binding the contract method 0xa5515878.
//
// Solidity: function setSessionAccount(address candidate_, bytes signature_) returns()
func (_SmartAccount *SmartAccountTransactorSession) SetSessionAccount(candidate_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.SetSessionAccount(&_SmartAccount.TransactOpts, candidate_, signature_)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SmartAccount *SmartAccountTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SmartAccount *SmartAccountSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.UpgradeToAndCall(&_SmartAccount.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SmartAccount *SmartAccountTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SmartAccount.Contract.UpgradeToAndCall(&_SmartAccount.TransactOpts, newImplementation, data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_SmartAccount *SmartAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _SmartAccount.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_SmartAccount *SmartAccountSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _SmartAccount.Contract.ValidateUserOp(&_SmartAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_SmartAccount *SmartAccountTransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _SmartAccount.Contract.ValidateUserOp(&_SmartAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SmartAccount *SmartAccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartAccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SmartAccount *SmartAccountSession) Receive() (*types.Transaction, error) {
	return _SmartAccount.Contract.Receive(&_SmartAccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SmartAccount *SmartAccountTransactorSession) Receive() (*types.Transaction, error) {
	return _SmartAccount.Contract.Receive(&_SmartAccount.TransactOpts)
}

// SmartAccountExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the SmartAccount contract.
type SmartAccountExecutedIterator struct {
	Event *SmartAccountExecuted // Event containing the contract specifics and raw log

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
func (it *SmartAccountExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountExecuted)
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
		it.Event = new(SmartAccountExecuted)
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
func (it *SmartAccountExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountExecuted represents a Executed event raised by the SmartAccount contract.
type SmartAccountExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_SmartAccount *SmartAccountFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*SmartAccountExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _SmartAccount.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &SmartAccountExecutedIterator{contract: _SmartAccount.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_SmartAccount *SmartAccountFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *SmartAccountExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _SmartAccount.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountExecuted)
				if err := _SmartAccount.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_SmartAccount *SmartAccountFilterer) ParseExecuted(log types.Log) (*SmartAccountExecuted, error) {
	event := new(SmartAccountExecuted)
	if err := _SmartAccount.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartAccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SmartAccount contract.
type SmartAccountInitializedIterator struct {
	Event *SmartAccountInitialized // Event containing the contract specifics and raw log

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
func (it *SmartAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountInitialized)
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
		it.Event = new(SmartAccountInitialized)
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
func (it *SmartAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountInitialized represents a Initialized event raised by the SmartAccount contract.
type SmartAccountInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SmartAccount *SmartAccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*SmartAccountInitializedIterator, error) {

	logs, sub, err := _SmartAccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SmartAccountInitializedIterator{contract: _SmartAccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SmartAccount *SmartAccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SmartAccountInitialized) (event.Subscription, error) {

	logs, sub, err := _SmartAccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountInitialized)
				if err := _SmartAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SmartAccount *SmartAccountFilterer) ParseInitialized(log types.Log) (*SmartAccountInitialized, error) {
	event := new(SmartAccountInitialized)
	if err := _SmartAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartAccountSessionAccountSetIterator is returned from FilterSessionAccountSet and is used to iterate over the raw logs and unpacked data for SessionAccountSet events raised by the SmartAccount contract.
type SmartAccountSessionAccountSetIterator struct {
	Event *SmartAccountSessionAccountSet // Event containing the contract specifics and raw log

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
func (it *SmartAccountSessionAccountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountSessionAccountSet)
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
		it.Event = new(SmartAccountSessionAccountSet)
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
func (it *SmartAccountSessionAccountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountSessionAccountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountSessionAccountSet represents a SessionAccountSet event raised by the SmartAccount contract.
type SmartAccountSessionAccountSet struct {
	Account   common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSessionAccountSet is a free log retrieval operation binding the contract event 0xf6d323c5e751dbe442743a74622a237eef71145a93957026bf2ecc6f7f6d2ac3.
//
// Solidity: event SessionAccountSet(address indexed account, uint256 timestamp)
func (_SmartAccount *SmartAccountFilterer) FilterSessionAccountSet(opts *bind.FilterOpts, account []common.Address) (*SmartAccountSessionAccountSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SmartAccount.contract.FilterLogs(opts, "SessionAccountSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &SmartAccountSessionAccountSetIterator{contract: _SmartAccount.contract, event: "SessionAccountSet", logs: logs, sub: sub}, nil
}

// WatchSessionAccountSet is a free log subscription operation binding the contract event 0xf6d323c5e751dbe442743a74622a237eef71145a93957026bf2ecc6f7f6d2ac3.
//
// Solidity: event SessionAccountSet(address indexed account, uint256 timestamp)
func (_SmartAccount *SmartAccountFilterer) WatchSessionAccountSet(opts *bind.WatchOpts, sink chan<- *SmartAccountSessionAccountSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SmartAccount.contract.WatchLogs(opts, "SessionAccountSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountSessionAccountSet)
				if err := _SmartAccount.contract.UnpackLog(event, "SessionAccountSet", log); err != nil {
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

// ParseSessionAccountSet is a log parse operation binding the contract event 0xf6d323c5e751dbe442743a74622a237eef71145a93957026bf2ecc6f7f6d2ac3.
//
// Solidity: event SessionAccountSet(address indexed account, uint256 timestamp)
func (_SmartAccount *SmartAccountFilterer) ParseSessionAccountSet(log types.Log) (*SmartAccountSessionAccountSet, error) {
	event := new(SmartAccountSessionAccountSet)
	if err := _SmartAccount.contract.UnpackLog(event, "SessionAccountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartAccountUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SmartAccount contract.
type SmartAccountUpgradedIterator struct {
	Event *SmartAccountUpgraded // Event containing the contract specifics and raw log

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
func (it *SmartAccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartAccountUpgraded)
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
		it.Event = new(SmartAccountUpgraded)
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
func (it *SmartAccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartAccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartAccountUpgraded represents a Upgraded event raised by the SmartAccount contract.
type SmartAccountUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SmartAccount *SmartAccountFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SmartAccountUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SmartAccount.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SmartAccountUpgradedIterator{contract: _SmartAccount.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SmartAccount *SmartAccountFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SmartAccountUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SmartAccount.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartAccountUpgraded)
				if err := _SmartAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_SmartAccount *SmartAccountFilterer) ParseUpgraded(log types.Log) (*SmartAccountUpgraded, error) {
	event := new(SmartAccountUpgraded)
	if err := _SmartAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
