// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

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

// Erc20MetaData contains all meta data concerning the Erc20 contract.
var Erc20MetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"golemFactory\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_master\",\"type\":\"address\"}],\"name\":\"setMigrationMaster\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"migrationMaster\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenCreationCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\"}],\"name\":\"setMigrationAgent\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"migrationAgent\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fundingEndBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenCreationMin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"funding\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenCreationRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fundingStartBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"create\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_golemFactory\",\"type\":\"address\"},{\"name\":\"_migrationMaster\",\"type\":\"address\"},{\"name\":\"_fundingStartBlock\",\"type\":\"uint256\"},{\"name\":\"_fundingEndBlock\",\"type\":\"uint256\"}],\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Migrate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"}]",
}

// Erc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20MetaData.ABI instead.
var Erc20ABI = Erc20MetaData.ABI

// Erc20 is an auto generated Go binding around an Ethereum contract.
type Erc20 struct {
	Erc20Caller     // Read-only binding to the contract
	Erc20Transactor // Write-only binding to the contract
	Erc20Filterer   // Log filterer for contract events
}

// Erc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20Session struct {
	Contract     *Erc20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20CallerSession struct {
	Contract *Erc20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Erc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20TransactorSession struct {
	Contract     *Erc20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20Raw struct {
	Contract *Erc20 // Generic contract binding to access the raw methods on
}

// Erc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20CallerRaw struct {
	Contract *Erc20Caller // Generic read-only contract binding to access the raw methods on
}

// Erc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20TransactorRaw struct {
	Contract *Erc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20 creates a new instance of Erc20, bound to a specific deployed contract.
func NewErc20(address common.Address, backend bind.ContractBackend) (*Erc20, error) {
	contract, err := bindErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20{Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract}}, nil
}

// NewErc20Caller creates a new read-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Caller(address common.Address, caller bind.ContractCaller) (*Erc20Caller, error) {
	contract, err := bindErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Caller{contract: contract}, nil
}

// NewErc20Transactor creates a new write-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc20Transactor, error) {
	contract, err := bindErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Transactor{contract: contract}, nil
}

// NewErc20Filterer creates a new log filterer instance of Erc20, bound to a specific deployed contract.
func NewErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc20Filterer, error) {
	contract, err := bindErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20Filterer{contract: contract}, nil
}

// bindErc20 binds a generic wrapper to an already deployed contract.
func bindErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.Erc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256)
func (_Erc20 *Erc20Caller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256)
func (_Erc20 *Erc20Session) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256)
func (_Erc20 *Erc20CallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_Erc20 *Erc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_Erc20 *Erc20Session) Decimals() (uint8, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_Erc20 *Erc20CallerSession) Decimals() (uint8, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// Funding is a free data retrieval call binding the contract method 0xcb4c86b7.
//
// Solidity: function funding() returns(bool)
func (_Erc20 *Erc20Caller) Funding(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "funding")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Funding is a free data retrieval call binding the contract method 0xcb4c86b7.
//
// Solidity: function funding() returns(bool)
func (_Erc20 *Erc20Session) Funding() (bool, error) {
	return _Erc20.Contract.Funding(&_Erc20.CallOpts)
}

// Funding is a free data retrieval call binding the contract method 0xcb4c86b7.
//
// Solidity: function funding() returns(bool)
func (_Erc20 *Erc20CallerSession) Funding() (bool, error) {
	return _Erc20.Contract.Funding(&_Erc20.CallOpts)
}

// FundingEndBlock is a free data retrieval call binding the contract method 0x91b43d13.
//
// Solidity: function fundingEndBlock() returns(uint256)
func (_Erc20 *Erc20Caller) FundingEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "fundingEndBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingEndBlock is a free data retrieval call binding the contract method 0x91b43d13.
//
// Solidity: function fundingEndBlock() returns(uint256)
func (_Erc20 *Erc20Session) FundingEndBlock() (*big.Int, error) {
	return _Erc20.Contract.FundingEndBlock(&_Erc20.CallOpts)
}

// FundingEndBlock is a free data retrieval call binding the contract method 0x91b43d13.
//
// Solidity: function fundingEndBlock() returns(uint256)
func (_Erc20 *Erc20CallerSession) FundingEndBlock() (*big.Int, error) {
	return _Erc20.Contract.FundingEndBlock(&_Erc20.CallOpts)
}

// FundingStartBlock is a free data retrieval call binding the contract method 0xd648a647.
//
// Solidity: function fundingStartBlock() returns(uint256)
func (_Erc20 *Erc20Caller) FundingStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "fundingStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingStartBlock is a free data retrieval call binding the contract method 0xd648a647.
//
// Solidity: function fundingStartBlock() returns(uint256)
func (_Erc20 *Erc20Session) FundingStartBlock() (*big.Int, error) {
	return _Erc20.Contract.FundingStartBlock(&_Erc20.CallOpts)
}

// FundingStartBlock is a free data retrieval call binding the contract method 0xd648a647.
//
// Solidity: function fundingStartBlock() returns(uint256)
func (_Erc20 *Erc20CallerSession) FundingStartBlock() (*big.Int, error) {
	return _Erc20.Contract.FundingStartBlock(&_Erc20.CallOpts)
}

// GolemFactory is a free data retrieval call binding the contract method 0x16222950.
//
// Solidity: function golemFactory() returns(address)
func (_Erc20 *Erc20Caller) GolemFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "golemFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GolemFactory is a free data retrieval call binding the contract method 0x16222950.
//
// Solidity: function golemFactory() returns(address)
func (_Erc20 *Erc20Session) GolemFactory() (common.Address, error) {
	return _Erc20.Contract.GolemFactory(&_Erc20.CallOpts)
}

// GolemFactory is a free data retrieval call binding the contract method 0x16222950.
//
// Solidity: function golemFactory() returns(address)
func (_Erc20 *Erc20CallerSession) GolemFactory() (common.Address, error) {
	return _Erc20.Contract.GolemFactory(&_Erc20.CallOpts)
}

// MigrationAgent is a free data retrieval call binding the contract method 0x8328dbcd.
//
// Solidity: function migrationAgent() returns(address)
func (_Erc20 *Erc20Caller) MigrationAgent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "migrationAgent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MigrationAgent is a free data retrieval call binding the contract method 0x8328dbcd.
//
// Solidity: function migrationAgent() returns(address)
func (_Erc20 *Erc20Session) MigrationAgent() (common.Address, error) {
	return _Erc20.Contract.MigrationAgent(&_Erc20.CallOpts)
}

// MigrationAgent is a free data retrieval call binding the contract method 0x8328dbcd.
//
// Solidity: function migrationAgent() returns(address)
func (_Erc20 *Erc20CallerSession) MigrationAgent() (common.Address, error) {
	return _Erc20.Contract.MigrationAgent(&_Erc20.CallOpts)
}

// MigrationMaster is a free data retrieval call binding the contract method 0x676d2e62.
//
// Solidity: function migrationMaster() returns(address)
func (_Erc20 *Erc20Caller) MigrationMaster(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "migrationMaster")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MigrationMaster is a free data retrieval call binding the contract method 0x676d2e62.
//
// Solidity: function migrationMaster() returns(address)
func (_Erc20 *Erc20Session) MigrationMaster() (common.Address, error) {
	return _Erc20.Contract.MigrationMaster(&_Erc20.CallOpts)
}

// MigrationMaster is a free data retrieval call binding the contract method 0x676d2e62.
//
// Solidity: function migrationMaster() returns(address)
func (_Erc20 *Erc20CallerSession) MigrationMaster() (common.Address, error) {
	return _Erc20.Contract.MigrationMaster(&_Erc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_Erc20 *Erc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_Erc20 *Erc20Session) Name() (string, error) {
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_Erc20 *Erc20CallerSession) Name() (string, error) {
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_Erc20 *Erc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_Erc20 *Erc20Session) Symbol() (string, error) {
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_Erc20 *Erc20CallerSession) Symbol() (string, error) {
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
}

// TokenCreationCap is a free data retrieval call binding the contract method 0x6f7920fd.
//
// Solidity: function tokenCreationCap() returns(uint256)
func (_Erc20 *Erc20Caller) TokenCreationCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "tokenCreationCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCreationCap is a free data retrieval call binding the contract method 0x6f7920fd.
//
// Solidity: function tokenCreationCap() returns(uint256)
func (_Erc20 *Erc20Session) TokenCreationCap() (*big.Int, error) {
	return _Erc20.Contract.TokenCreationCap(&_Erc20.CallOpts)
}

// TokenCreationCap is a free data retrieval call binding the contract method 0x6f7920fd.
//
// Solidity: function tokenCreationCap() returns(uint256)
func (_Erc20 *Erc20CallerSession) TokenCreationCap() (*big.Int, error) {
	return _Erc20.Contract.TokenCreationCap(&_Erc20.CallOpts)
}

// TokenCreationMin is a free data retrieval call binding the contract method 0xc039daf6.
//
// Solidity: function tokenCreationMin() returns(uint256)
func (_Erc20 *Erc20Caller) TokenCreationMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "tokenCreationMin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCreationMin is a free data retrieval call binding the contract method 0xc039daf6.
//
// Solidity: function tokenCreationMin() returns(uint256)
func (_Erc20 *Erc20Session) TokenCreationMin() (*big.Int, error) {
	return _Erc20.Contract.TokenCreationMin(&_Erc20.CallOpts)
}

// TokenCreationMin is a free data retrieval call binding the contract method 0xc039daf6.
//
// Solidity: function tokenCreationMin() returns(uint256)
func (_Erc20 *Erc20CallerSession) TokenCreationMin() (*big.Int, error) {
	return _Erc20.Contract.TokenCreationMin(&_Erc20.CallOpts)
}

// TokenCreationRate is a free data retrieval call binding the contract method 0xcf8d652c.
//
// Solidity: function tokenCreationRate() returns(uint256)
func (_Erc20 *Erc20Caller) TokenCreationRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "tokenCreationRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCreationRate is a free data retrieval call binding the contract method 0xcf8d652c.
//
// Solidity: function tokenCreationRate() returns(uint256)
func (_Erc20 *Erc20Session) TokenCreationRate() (*big.Int, error) {
	return _Erc20.Contract.TokenCreationRate(&_Erc20.CallOpts)
}

// TokenCreationRate is a free data retrieval call binding the contract method 0xcf8d652c.
//
// Solidity: function tokenCreationRate() returns(uint256)
func (_Erc20 *Erc20CallerSession) TokenCreationRate() (*big.Int, error) {
	return _Erc20.Contract.TokenCreationRate(&_Erc20.CallOpts)
}

// TotalMigrated is a free data retrieval call binding the contract method 0x95a0f5eb.
//
// Solidity: function totalMigrated() returns(uint256)
func (_Erc20 *Erc20Caller) TotalMigrated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "totalMigrated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMigrated is a free data retrieval call binding the contract method 0x95a0f5eb.
//
// Solidity: function totalMigrated() returns(uint256)
func (_Erc20 *Erc20Session) TotalMigrated() (*big.Int, error) {
	return _Erc20.Contract.TotalMigrated(&_Erc20.CallOpts)
}

// TotalMigrated is a free data retrieval call binding the contract method 0x95a0f5eb.
//
// Solidity: function totalMigrated() returns(uint256)
func (_Erc20 *Erc20CallerSession) TotalMigrated() (*big.Int, error) {
	return _Erc20.Contract.TotalMigrated(&_Erc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_Erc20 *Erc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_Erc20 *Erc20Session) TotalSupply() (*big.Int, error) {
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_Erc20 *Erc20CallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (_Erc20 *Erc20Transactor) Create(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "create")
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (_Erc20 *Erc20Session) Create() (*types.Transaction, error) {
	return _Erc20.Contract.Create(&_Erc20.TransactOpts)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (_Erc20 *Erc20TransactorSession) Create() (*types.Transaction, error) {
	return _Erc20.Contract.Create(&_Erc20.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_Erc20 *Erc20Transactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_Erc20 *Erc20Session) Finalize() (*types.Transaction, error) {
	return _Erc20.Contract.Finalize(&_Erc20.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_Erc20 *Erc20TransactorSession) Finalize() (*types.Transaction, error) {
	return _Erc20.Contract.Finalize(&_Erc20.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _value) returns()
func (_Erc20 *Erc20Transactor) Migrate(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "migrate", _value)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _value) returns()
func (_Erc20 *Erc20Session) Migrate(_value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Migrate(&_Erc20.TransactOpts, _value)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _value) returns()
func (_Erc20 *Erc20TransactorSession) Migrate(_value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Migrate(&_Erc20.TransactOpts, _value)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Erc20 *Erc20Transactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Erc20 *Erc20Session) Refund() (*types.Transaction, error) {
	return _Erc20.Contract.Refund(&_Erc20.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Erc20 *Erc20TransactorSession) Refund() (*types.Transaction, error) {
	return _Erc20.Contract.Refund(&_Erc20.TransactOpts)
}

// SetMigrationAgent is a paid mutator transaction binding the contract method 0x75e2ff65.
//
// Solidity: function setMigrationAgent(address _agent) returns()
func (_Erc20 *Erc20Transactor) SetMigrationAgent(opts *bind.TransactOpts, _agent common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "setMigrationAgent", _agent)
}

// SetMigrationAgent is a paid mutator transaction binding the contract method 0x75e2ff65.
//
// Solidity: function setMigrationAgent(address _agent) returns()
func (_Erc20 *Erc20Session) SetMigrationAgent(_agent common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.SetMigrationAgent(&_Erc20.TransactOpts, _agent)
}

// SetMigrationAgent is a paid mutator transaction binding the contract method 0x75e2ff65.
//
// Solidity: function setMigrationAgent(address _agent) returns()
func (_Erc20 *Erc20TransactorSession) SetMigrationAgent(_agent common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.SetMigrationAgent(&_Erc20.TransactOpts, _agent)
}

// SetMigrationMaster is a paid mutator transaction binding the contract method 0x26316e58.
//
// Solidity: function setMigrationMaster(address _master) returns()
func (_Erc20 *Erc20Transactor) SetMigrationMaster(opts *bind.TransactOpts, _master common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "setMigrationMaster", _master)
}

// SetMigrationMaster is a paid mutator transaction binding the contract method 0x26316e58.
//
// Solidity: function setMigrationMaster(address _master) returns()
func (_Erc20 *Erc20Session) SetMigrationMaster(_master common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.SetMigrationMaster(&_Erc20.TransactOpts, _master)
}

// SetMigrationMaster is a paid mutator transaction binding the contract method 0x26316e58.
//
// Solidity: function setMigrationMaster(address _master) returns()
func (_Erc20 *Erc20TransactorSession) SetMigrationMaster(_master common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.SetMigrationMaster(&_Erc20.TransactOpts, _master)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Erc20 *Erc20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Erc20 *Erc20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Erc20 *Erc20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, _to, _value)
}

// Erc20MigrateIterator is returned from FilterMigrate and is used to iterate over the raw logs and unpacked data for Migrate events raised by the Erc20 contract.
type Erc20MigrateIterator struct {
	Event *Erc20Migrate // Event containing the contract specifics and raw log

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
func (it *Erc20MigrateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Migrate)
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
		it.Event = new(Erc20Migrate)
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
func (it *Erc20MigrateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20MigrateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Migrate represents a Migrate event raised by the Erc20 contract.
type Erc20Migrate struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMigrate is a free log retrieval operation binding the contract event 0x18df02dcc52b9c494f391df09661519c0069bd8540141946280399408205ca1a.
//
// Solidity: event Migrate(address indexed _from, address indexed _to, uint256 _value)
func (_Erc20 *Erc20Filterer) FilterMigrate(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*Erc20MigrateIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Migrate", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20MigrateIterator{contract: _Erc20.contract, event: "Migrate", logs: logs, sub: sub}, nil
}

// WatchMigrate is a free log subscription operation binding the contract event 0x18df02dcc52b9c494f391df09661519c0069bd8540141946280399408205ca1a.
//
// Solidity: event Migrate(address indexed _from, address indexed _to, uint256 _value)
func (_Erc20 *Erc20Filterer) WatchMigrate(opts *bind.WatchOpts, sink chan<- *Erc20Migrate, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Migrate", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Migrate)
				if err := _Erc20.contract.UnpackLog(event, "Migrate", log); err != nil {
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

// ParseMigrate is a log parse operation binding the contract event 0x18df02dcc52b9c494f391df09661519c0069bd8540141946280399408205ca1a.
//
// Solidity: event Migrate(address indexed _from, address indexed _to, uint256 _value)
func (_Erc20 *Erc20Filterer) ParseMigrate(log types.Log) (*Erc20Migrate, error) {
	event := new(Erc20Migrate)
	if err := _Erc20.contract.UnpackLog(event, "Migrate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20RefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the Erc20 contract.
type Erc20RefundIterator struct {
	Event *Erc20Refund // Event containing the contract specifics and raw log

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
func (it *Erc20RefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Refund)
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
		it.Event = new(Erc20Refund)
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
func (it *Erc20RefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20RefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Refund represents a Refund event raised by the Erc20 contract.
type Erc20Refund struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0xbb28353e4598c3b9199101a66e0989549b659a59a54d2c27fbb183f1932c8e6d.
//
// Solidity: event Refund(address indexed _from, uint256 _value)
func (_Erc20 *Erc20Filterer) FilterRefund(opts *bind.FilterOpts, _from []common.Address) (*Erc20RefundIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Refund", _fromRule)
	if err != nil {
		return nil, err
	}
	return &Erc20RefundIterator{contract: _Erc20.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0xbb28353e4598c3b9199101a66e0989549b659a59a54d2c27fbb183f1932c8e6d.
//
// Solidity: event Refund(address indexed _from, uint256 _value)
func (_Erc20 *Erc20Filterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *Erc20Refund, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Refund", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Refund)
				if err := _Erc20.contract.UnpackLog(event, "Refund", log); err != nil {
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

// ParseRefund is a log parse operation binding the contract event 0xbb28353e4598c3b9199101a66e0989549b659a59a54d2c27fbb183f1932c8e6d.
//
// Solidity: event Refund(address indexed _from, uint256 _value)
func (_Erc20 *Erc20Filterer) ParseRefund(log types.Log) (*Erc20Refund, error) {
	event := new(Erc20Refund)
	if err := _Erc20.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20 contract.
type Erc20TransferIterator struct {
	Event *Erc20Transfer // Event containing the contract specifics and raw log

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
func (it *Erc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Transfer)
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
		it.Event = new(Erc20Transfer)
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
func (it *Erc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Transfer represents a Transfer event raised by the Erc20 contract.
type Erc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Erc20 *Erc20Filterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*Erc20TransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20TransferIterator{contract: _Erc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Erc20 *Erc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20Transfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Transfer)
				if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Erc20 *Erc20Filterer) ParseTransfer(log types.Log) (*Erc20Transfer, error) {
	event := new(Erc20Transfer)
	if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
