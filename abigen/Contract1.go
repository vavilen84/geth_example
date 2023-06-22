// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigen

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

// Contract1MetaData contains all meta data concerning the Contract1 contract.
var Contract1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Contract1ABI is the input ABI used to generate the binding from.
// Deprecated: Use Contract1MetaData.ABI instead.
var Contract1ABI = Contract1MetaData.ABI

// Contract1 is an auto generated Go binding around an Ethereum contract.
type Contract1 struct {
	Contract1Caller     // Read-only binding to the contract
	Contract1Transactor // Write-only binding to the contract
	Contract1Filterer   // Log filterer for contract events
}

// Contract1Caller is an auto generated read-only Go binding around an Ethereum contract.
type Contract1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Contract1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Contract1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Contract1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Contract1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Contract1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Contract1Session struct {
	Contract     *Contract1        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Contract1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Contract1CallerSession struct {
	Contract *Contract1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Contract1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Contract1TransactorSession struct {
	Contract     *Contract1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Contract1Raw is an auto generated low-level Go binding around an Ethereum contract.
type Contract1Raw struct {
	Contract *Contract1 // Generic contract binding to access the raw methods on
}

// Contract1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Contract1CallerRaw struct {
	Contract *Contract1Caller // Generic read-only contract binding to access the raw methods on
}

// Contract1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Contract1TransactorRaw struct {
	Contract *Contract1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewContract1 creates a new instance of Contract1, bound to a specific deployed contract.
func NewContract1(address common.Address, backend bind.ContractBackend) (*Contract1, error) {
	contract, err := bindContract1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract1{Contract1Caller: Contract1Caller{contract: contract}, Contract1Transactor: Contract1Transactor{contract: contract}, Contract1Filterer: Contract1Filterer{contract: contract}}, nil
}

// NewContract1Caller creates a new read-only instance of Contract1, bound to a specific deployed contract.
func NewContract1Caller(address common.Address, caller bind.ContractCaller) (*Contract1Caller, error) {
	contract, err := bindContract1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Contract1Caller{contract: contract}, nil
}

// NewContract1Transactor creates a new write-only instance of Contract1, bound to a specific deployed contract.
func NewContract1Transactor(address common.Address, transactor bind.ContractTransactor) (*Contract1Transactor, error) {
	contract, err := bindContract1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Contract1Transactor{contract: contract}, nil
}

// NewContract1Filterer creates a new log filterer instance of Contract1, bound to a specific deployed contract.
func NewContract1Filterer(address common.Address, filterer bind.ContractFilterer) (*Contract1Filterer, error) {
	contract, err := bindContract1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Contract1Filterer{contract: contract}, nil
}

// bindContract1 binds a generic wrapper to an already deployed contract.
func bindContract1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Contract1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract1 *Contract1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract1.Contract.Contract1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract1 *Contract1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract1.Contract.Contract1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract1 *Contract1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract1.Contract.Contract1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract1 *Contract1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract1 *Contract1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract1 *Contract1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract1.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract1 *Contract1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract1 *Contract1Session) Owner() (common.Address, error) {
	return _Contract1.Contract.Owner(&_Contract1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract1 *Contract1CallerSession) Owner() (common.Address, error) {
	return _Contract1.Contract.Owner(&_Contract1.CallOpts)
}
