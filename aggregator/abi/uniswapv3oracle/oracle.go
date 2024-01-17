// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv3oracle

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

// Uniswapv3oracleMetaData contains all meta data concerning the Uniswapv3oracle contract.
var Uniswapv3oracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_secondsAgoToStartOfTwap\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_secondsAgoToEndOfTwap\",\"type\":\"uint32\"}],\"name\":\"getHistoricalTwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"getMaxPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"getTimeWeightedAverageTickSafe\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"timeWeightedAverageTick\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"_checkPeriod\",\"type\":\"bool\"}],\"name\":\"getTwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Uniswapv3oracleABI is the input ABI used to generate the binding from.
// Deprecated: Use Uniswapv3oracleMetaData.ABI instead.
var Uniswapv3oracleABI = Uniswapv3oracleMetaData.ABI

// Uniswapv3oracle is an auto generated Go binding around an Ethereum contract.
type Uniswapv3oracle struct {
	Uniswapv3oracleCaller     // Read-only binding to the contract
	Uniswapv3oracleTransactor // Write-only binding to the contract
	Uniswapv3oracleFilterer   // Log filterer for contract events
}

// Uniswapv3oracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswapv3oracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3oracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswapv3oracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3oracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswapv3oracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3oracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswapv3oracleSession struct {
	Contract     *Uniswapv3oracle  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Uniswapv3oracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswapv3oracleCallerSession struct {
	Contract *Uniswapv3oracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Uniswapv3oracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswapv3oracleTransactorSession struct {
	Contract     *Uniswapv3oracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Uniswapv3oracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswapv3oracleRaw struct {
	Contract *Uniswapv3oracle // Generic contract binding to access the raw methods on
}

// Uniswapv3oracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswapv3oracleCallerRaw struct {
	Contract *Uniswapv3oracleCaller // Generic read-only contract binding to access the raw methods on
}

// Uniswapv3oracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswapv3oracleTransactorRaw struct {
	Contract *Uniswapv3oracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapv3oracle creates a new instance of Uniswapv3oracle, bound to a specific deployed contract.
func NewUniswapv3oracle(address common.Address, backend bind.ContractBackend) (*Uniswapv3oracle, error) {
	contract, err := bindUniswapv3oracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3oracle{Uniswapv3oracleCaller: Uniswapv3oracleCaller{contract: contract}, Uniswapv3oracleTransactor: Uniswapv3oracleTransactor{contract: contract}, Uniswapv3oracleFilterer: Uniswapv3oracleFilterer{contract: contract}}, nil
}

// NewUniswapv3oracleCaller creates a new read-only instance of Uniswapv3oracle, bound to a specific deployed contract.
func NewUniswapv3oracleCaller(address common.Address, caller bind.ContractCaller) (*Uniswapv3oracleCaller, error) {
	contract, err := bindUniswapv3oracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3oracleCaller{contract: contract}, nil
}

// NewUniswapv3oracleTransactor creates a new write-only instance of Uniswapv3oracle, bound to a specific deployed contract.
func NewUniswapv3oracleTransactor(address common.Address, transactor bind.ContractTransactor) (*Uniswapv3oracleTransactor, error) {
	contract, err := bindUniswapv3oracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3oracleTransactor{contract: contract}, nil
}

// NewUniswapv3oracleFilterer creates a new log filterer instance of Uniswapv3oracle, bound to a specific deployed contract.
func NewUniswapv3oracleFilterer(address common.Address, filterer bind.ContractFilterer) (*Uniswapv3oracleFilterer, error) {
	contract, err := bindUniswapv3oracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3oracleFilterer{contract: contract}, nil
}

// bindUniswapv3oracle binds a generic wrapper to an already deployed contract.
func bindUniswapv3oracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Uniswapv3oracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv3oracle *Uniswapv3oracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv3oracle.Contract.Uniswapv3oracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv3oracle *Uniswapv3oracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3oracle.Contract.Uniswapv3oracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv3oracle *Uniswapv3oracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv3oracle.Contract.Uniswapv3oracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv3oracle *Uniswapv3oracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv3oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv3oracle *Uniswapv3oracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv3oracle *Uniswapv3oracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv3oracle.Contract.contract.Transact(opts, method, params...)
}

// GetHistoricalTwap is a free data retrieval call binding the contract method 0x4ac78d11.
//
// Solidity: function getHistoricalTwap(address _pool, address _base, address _quote, uint32 _secondsAgoToStartOfTwap, uint32 _secondsAgoToEndOfTwap) view returns(uint256)
func (_Uniswapv3oracle *Uniswapv3oracleCaller) GetHistoricalTwap(opts *bind.CallOpts, _pool common.Address, _base common.Address, _quote common.Address, _secondsAgoToStartOfTwap uint32, _secondsAgoToEndOfTwap uint32) (*big.Int, error) {
	var out []interface{}
	err := _Uniswapv3oracle.contract.Call(opts, &out, "getHistoricalTwap", _pool, _base, _quote, _secondsAgoToStartOfTwap, _secondsAgoToEndOfTwap)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHistoricalTwap is a free data retrieval call binding the contract method 0x4ac78d11.
//
// Solidity: function getHistoricalTwap(address _pool, address _base, address _quote, uint32 _secondsAgoToStartOfTwap, uint32 _secondsAgoToEndOfTwap) view returns(uint256)
func (_Uniswapv3oracle *Uniswapv3oracleSession) GetHistoricalTwap(_pool common.Address, _base common.Address, _quote common.Address, _secondsAgoToStartOfTwap uint32, _secondsAgoToEndOfTwap uint32) (*big.Int, error) {
	return _Uniswapv3oracle.Contract.GetHistoricalTwap(&_Uniswapv3oracle.CallOpts, _pool, _base, _quote, _secondsAgoToStartOfTwap, _secondsAgoToEndOfTwap)
}

// GetHistoricalTwap is a free data retrieval call binding the contract method 0x4ac78d11.
//
// Solidity: function getHistoricalTwap(address _pool, address _base, address _quote, uint32 _secondsAgoToStartOfTwap, uint32 _secondsAgoToEndOfTwap) view returns(uint256)
func (_Uniswapv3oracle *Uniswapv3oracleCallerSession) GetHistoricalTwap(_pool common.Address, _base common.Address, _quote common.Address, _secondsAgoToStartOfTwap uint32, _secondsAgoToEndOfTwap uint32) (*big.Int, error) {
	return _Uniswapv3oracle.Contract.GetHistoricalTwap(&_Uniswapv3oracle.CallOpts, _pool, _base, _quote, _secondsAgoToStartOfTwap, _secondsAgoToEndOfTwap)
}

// GetMaxPeriod is a free data retrieval call binding the contract method 0xde5a6e22.
//
// Solidity: function getMaxPeriod(address _pool) view returns(uint32)
func (_Uniswapv3oracle *Uniswapv3oracleCaller) GetMaxPeriod(opts *bind.CallOpts, _pool common.Address) (uint32, error) {
	var out []interface{}
	err := _Uniswapv3oracle.contract.Call(opts, &out, "getMaxPeriod", _pool)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMaxPeriod is a free data retrieval call binding the contract method 0xde5a6e22.
//
// Solidity: function getMaxPeriod(address _pool) view returns(uint32)
func (_Uniswapv3oracle *Uniswapv3oracleSession) GetMaxPeriod(_pool common.Address) (uint32, error) {
	return _Uniswapv3oracle.Contract.GetMaxPeriod(&_Uniswapv3oracle.CallOpts, _pool)
}

// GetMaxPeriod is a free data retrieval call binding the contract method 0xde5a6e22.
//
// Solidity: function getMaxPeriod(address _pool) view returns(uint32)
func (_Uniswapv3oracle *Uniswapv3oracleCallerSession) GetMaxPeriod(_pool common.Address) (uint32, error) {
	return _Uniswapv3oracle.Contract.GetMaxPeriod(&_Uniswapv3oracle.CallOpts, _pool)
}

// GetTimeWeightedAverageTickSafe is a free data retrieval call binding the contract method 0x4a0a96eb.
//
// Solidity: function getTimeWeightedAverageTickSafe(address _pool, uint32 _period) view returns(int24 timeWeightedAverageTick)
func (_Uniswapv3oracle *Uniswapv3oracleCaller) GetTimeWeightedAverageTickSafe(opts *bind.CallOpts, _pool common.Address, _period uint32) (*big.Int, error) {
	var out []interface{}
	err := _Uniswapv3oracle.contract.Call(opts, &out, "getTimeWeightedAverageTickSafe", _pool, _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeWeightedAverageTickSafe is a free data retrieval call binding the contract method 0x4a0a96eb.
//
// Solidity: function getTimeWeightedAverageTickSafe(address _pool, uint32 _period) view returns(int24 timeWeightedAverageTick)
func (_Uniswapv3oracle *Uniswapv3oracleSession) GetTimeWeightedAverageTickSafe(_pool common.Address, _period uint32) (*big.Int, error) {
	return _Uniswapv3oracle.Contract.GetTimeWeightedAverageTickSafe(&_Uniswapv3oracle.CallOpts, _pool, _period)
}

// GetTimeWeightedAverageTickSafe is a free data retrieval call binding the contract method 0x4a0a96eb.
//
// Solidity: function getTimeWeightedAverageTickSafe(address _pool, uint32 _period) view returns(int24 timeWeightedAverageTick)
func (_Uniswapv3oracle *Uniswapv3oracleCallerSession) GetTimeWeightedAverageTickSafe(_pool common.Address, _period uint32) (*big.Int, error) {
	return _Uniswapv3oracle.Contract.GetTimeWeightedAverageTickSafe(&_Uniswapv3oracle.CallOpts, _pool, _period)
}

// GetTwap is a free data retrieval call binding the contract method 0xcce79bd5.
//
// Solidity: function getTwap(address _pool, address _base, address _quote, uint32 _period, bool _checkPeriod) view returns(uint256)
func (_Uniswapv3oracle *Uniswapv3oracleCaller) GetTwap(opts *bind.CallOpts, _pool common.Address, _base common.Address, _quote common.Address, _period uint32, _checkPeriod bool) (*big.Int, error) {
	var out []interface{}
	err := _Uniswapv3oracle.contract.Call(opts, &out, "getTwap", _pool, _base, _quote, _period, _checkPeriod)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTwap is a free data retrieval call binding the contract method 0xcce79bd5.
//
// Solidity: function getTwap(address _pool, address _base, address _quote, uint32 _period, bool _checkPeriod) view returns(uint256)
func (_Uniswapv3oracle *Uniswapv3oracleSession) GetTwap(_pool common.Address, _base common.Address, _quote common.Address, _period uint32, _checkPeriod bool) (*big.Int, error) {
	return _Uniswapv3oracle.Contract.GetTwap(&_Uniswapv3oracle.CallOpts, _pool, _base, _quote, _period, _checkPeriod)
}

// GetTwap is a free data retrieval call binding the contract method 0xcce79bd5.
//
// Solidity: function getTwap(address _pool, address _base, address _quote, uint32 _period, bool _checkPeriod) view returns(uint256)
func (_Uniswapv3oracle *Uniswapv3oracleCallerSession) GetTwap(_pool common.Address, _base common.Address, _quote common.Address, _period uint32, _checkPeriod bool) (*big.Int, error) {
	return _Uniswapv3oracle.Contract.GetTwap(&_Uniswapv3oracle.CallOpts, _pool, _base, _quote, _period, _checkPeriod)
}
