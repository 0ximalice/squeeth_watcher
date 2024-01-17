// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package opyn

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

// OpynMetaData contains all meta data concerning the Opyn contract.
var OpynMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_shortPowerPerp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wPowerPerp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quoteCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethQuoteCurrencyPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wPowerPerpPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uniPositionManager\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_feeTier\",\"type\":\"uint24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"BurnShort\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"DepositUniPositionToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"FeeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFeeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralPaid\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"MintShort\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldNormFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNormFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastModificationTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NormalizationFactorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"OpenVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pausesLeft\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wPowerPerpAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payoutAmount\",\"type\":\"uint256\"}],\"name\":\"RedeemLong\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vauldId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"RedeemShort\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wPowerPerpRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wPowerPerpBurned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wPowerPerpExcess\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"}],\"name\":\"ReduceDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexForSettlement\",\"type\":\"uint256\"}],\"name\":\"Shutdown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"unpauser\",\"type\":\"address\"}],\"name\":\"UnPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"UpdateOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"WithdrawUniPositionToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FUNDING_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TWAP_PERIOD\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"applyFunding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_powerPerpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"burnPowerPerpAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_wPowerPerpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"burnWPowerPerpAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_uniTokenId\",\"type\":\"uint256\"}],\"name\":\"depositUniPositionToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethQuoteCurrencyPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTier\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"getDenormalizedMark\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"getDenormalizedMarkForFunding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNormalizationFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"getUnscaledIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexForSettlement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutDown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSystemPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"isVaultSafe\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFundingUpdateTimestamp\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPauseTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDebtAmount\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_powerPerpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_uniTokenId\",\"type\":\"uint256\"}],\"name\":\"mintPowerPerpAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_wPowerPerpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_uniTokenId\",\"type\":\"uint256\"}],\"name\":\"mintWPowerPerpAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"normalizationFactor\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pausesLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteCurrency\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_wPerpAmount\",\"type\":\"uint256\"}],\"name\":\"redeemLong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"redeemShort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"reduceDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"reduceDebtShutdown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shortPowerPerp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutDown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unPauseAnyone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unPauseOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"updateOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"NftCollateralId\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"collateralAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint128\",\"name\":\"shortAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wPowerPerp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wPowerPerpPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"withdrawUniPositionToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// OpynABI is the input ABI used to generate the binding from.
// Deprecated: Use OpynMetaData.ABI instead.
var OpynABI = OpynMetaData.ABI

// Opyn is an auto generated Go binding around an Ethereum contract.
type Opyn struct {
	OpynCaller     // Read-only binding to the contract
	OpynTransactor // Write-only binding to the contract
	OpynFilterer   // Log filterer for contract events
}

// OpynCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpynCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpynTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpynTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpynFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpynFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpynSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpynSession struct {
	Contract     *Opyn             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpynCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpynCallerSession struct {
	Contract *OpynCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OpynTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpynTransactorSession struct {
	Contract     *OpynTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpynRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpynRaw struct {
	Contract *Opyn // Generic contract binding to access the raw methods on
}

// OpynCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpynCallerRaw struct {
	Contract *OpynCaller // Generic read-only contract binding to access the raw methods on
}

// OpynTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpynTransactorRaw struct {
	Contract *OpynTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpyn creates a new instance of Opyn, bound to a specific deployed contract.
func NewOpyn(address common.Address, backend bind.ContractBackend) (*Opyn, error) {
	contract, err := bindOpyn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Opyn{OpynCaller: OpynCaller{contract: contract}, OpynTransactor: OpynTransactor{contract: contract}, OpynFilterer: OpynFilterer{contract: contract}}, nil
}

// NewOpynCaller creates a new read-only instance of Opyn, bound to a specific deployed contract.
func NewOpynCaller(address common.Address, caller bind.ContractCaller) (*OpynCaller, error) {
	contract, err := bindOpyn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpynCaller{contract: contract}, nil
}

// NewOpynTransactor creates a new write-only instance of Opyn, bound to a specific deployed contract.
func NewOpynTransactor(address common.Address, transactor bind.ContractTransactor) (*OpynTransactor, error) {
	contract, err := bindOpyn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpynTransactor{contract: contract}, nil
}

// NewOpynFilterer creates a new log filterer instance of Opyn, bound to a specific deployed contract.
func NewOpynFilterer(address common.Address, filterer bind.ContractFilterer) (*OpynFilterer, error) {
	contract, err := bindOpyn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpynFilterer{contract: contract}, nil
}

// bindOpyn binds a generic wrapper to an already deployed contract.
func bindOpyn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OpynMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Opyn *OpynRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Opyn.Contract.OpynCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Opyn *OpynRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opyn.Contract.OpynTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Opyn *OpynRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Opyn.Contract.OpynTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Opyn *OpynCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Opyn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Opyn *OpynTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opyn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Opyn *OpynTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Opyn.Contract.contract.Transact(opts, method, params...)
}

// FUNDINGPERIOD is a free data retrieval call binding the contract method 0xf90c3f27.
//
// Solidity: function FUNDING_PERIOD() view returns(uint256)
func (_Opyn *OpynCaller) FUNDINGPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "FUNDING_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FUNDINGPERIOD is a free data retrieval call binding the contract method 0xf90c3f27.
//
// Solidity: function FUNDING_PERIOD() view returns(uint256)
func (_Opyn *OpynSession) FUNDINGPERIOD() (*big.Int, error) {
	return _Opyn.Contract.FUNDINGPERIOD(&_Opyn.CallOpts)
}

// FUNDINGPERIOD is a free data retrieval call binding the contract method 0xf90c3f27.
//
// Solidity: function FUNDING_PERIOD() view returns(uint256)
func (_Opyn *OpynCallerSession) FUNDINGPERIOD() (*big.Int, error) {
	return _Opyn.Contract.FUNDINGPERIOD(&_Opyn.CallOpts)
}

// TWAPPERIOD is a free data retrieval call binding the contract method 0x7ca25184.
//
// Solidity: function TWAP_PERIOD() view returns(uint32)
func (_Opyn *OpynCaller) TWAPPERIOD(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "TWAP_PERIOD")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TWAPPERIOD is a free data retrieval call binding the contract method 0x7ca25184.
//
// Solidity: function TWAP_PERIOD() view returns(uint32)
func (_Opyn *OpynSession) TWAPPERIOD() (uint32, error) {
	return _Opyn.Contract.TWAPPERIOD(&_Opyn.CallOpts)
}

// TWAPPERIOD is a free data retrieval call binding the contract method 0x7ca25184.
//
// Solidity: function TWAP_PERIOD() view returns(uint32)
func (_Opyn *OpynCallerSession) TWAPPERIOD() (uint32, error) {
	return _Opyn.Contract.TWAPPERIOD(&_Opyn.CallOpts)
}

// EthQuoteCurrencyPool is a free data retrieval call binding the contract method 0x4468c022.
//
// Solidity: function ethQuoteCurrencyPool() view returns(address)
func (_Opyn *OpynCaller) EthQuoteCurrencyPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "ethQuoteCurrencyPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthQuoteCurrencyPool is a free data retrieval call binding the contract method 0x4468c022.
//
// Solidity: function ethQuoteCurrencyPool() view returns(address)
func (_Opyn *OpynSession) EthQuoteCurrencyPool() (common.Address, error) {
	return _Opyn.Contract.EthQuoteCurrencyPool(&_Opyn.CallOpts)
}

// EthQuoteCurrencyPool is a free data retrieval call binding the contract method 0x4468c022.
//
// Solidity: function ethQuoteCurrencyPool() view returns(address)
func (_Opyn *OpynCallerSession) EthQuoteCurrencyPool() (common.Address, error) {
	return _Opyn.Contract.EthQuoteCurrencyPool(&_Opyn.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Opyn *OpynCaller) FeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Opyn *OpynSession) FeeRate() (*big.Int, error) {
	return _Opyn.Contract.FeeRate(&_Opyn.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Opyn *OpynCallerSession) FeeRate() (*big.Int, error) {
	return _Opyn.Contract.FeeRate(&_Opyn.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Opyn *OpynCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Opyn *OpynSession) FeeRecipient() (common.Address, error) {
	return _Opyn.Contract.FeeRecipient(&_Opyn.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Opyn *OpynCallerSession) FeeRecipient() (common.Address, error) {
	return _Opyn.Contract.FeeRecipient(&_Opyn.CallOpts)
}

// FeeTier is a free data retrieval call binding the contract method 0x72f5d98a.
//
// Solidity: function feeTier() view returns(uint24)
func (_Opyn *OpynCaller) FeeTier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "feeTier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeTier is a free data retrieval call binding the contract method 0x72f5d98a.
//
// Solidity: function feeTier() view returns(uint24)
func (_Opyn *OpynSession) FeeTier() (*big.Int, error) {
	return _Opyn.Contract.FeeTier(&_Opyn.CallOpts)
}

// FeeTier is a free data retrieval call binding the contract method 0x72f5d98a.
//
// Solidity: function feeTier() view returns(uint24)
func (_Opyn *OpynCallerSession) FeeTier() (*big.Int, error) {
	return _Opyn.Contract.FeeTier(&_Opyn.CallOpts)
}

// GetDenormalizedMark is a free data retrieval call binding the contract method 0xee3189ff.
//
// Solidity: function getDenormalizedMark(uint32 _period) view returns(uint256)
func (_Opyn *OpynCaller) GetDenormalizedMark(opts *bind.CallOpts, _period uint32) (*big.Int, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "getDenormalizedMark", _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDenormalizedMark is a free data retrieval call binding the contract method 0xee3189ff.
//
// Solidity: function getDenormalizedMark(uint32 _period) view returns(uint256)
func (_Opyn *OpynSession) GetDenormalizedMark(_period uint32) (*big.Int, error) {
	return _Opyn.Contract.GetDenormalizedMark(&_Opyn.CallOpts, _period)
}

// GetDenormalizedMark is a free data retrieval call binding the contract method 0xee3189ff.
//
// Solidity: function getDenormalizedMark(uint32 _period) view returns(uint256)
func (_Opyn *OpynCallerSession) GetDenormalizedMark(_period uint32) (*big.Int, error) {
	return _Opyn.Contract.GetDenormalizedMark(&_Opyn.CallOpts, _period)
}

// GetDenormalizedMarkForFunding is a free data retrieval call binding the contract method 0x377a1936.
//
// Solidity: function getDenormalizedMarkForFunding(uint32 _period) view returns(uint256)
func (_Opyn *OpynCaller) GetDenormalizedMarkForFunding(opts *bind.CallOpts, _period uint32) (*big.Int, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "getDenormalizedMarkForFunding", _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDenormalizedMarkForFunding is a free data retrieval call binding the contract method 0x377a1936.
//
// Solidity: function getDenormalizedMarkForFunding(uint32 _period) view returns(uint256)
func (_Opyn *OpynSession) GetDenormalizedMarkForFunding(_period uint32) (*big.Int, error) {
	return _Opyn.Contract.GetDenormalizedMarkForFunding(&_Opyn.CallOpts, _period)
}

// GetDenormalizedMarkForFunding is a free data retrieval call binding the contract method 0x377a1936.
//
// Solidity: function getDenormalizedMarkForFunding(uint32 _period) view returns(uint256)
func (_Opyn *OpynCallerSession) GetDenormalizedMarkForFunding(_period uint32) (*big.Int, error) {
	return _Opyn.Contract.GetDenormalizedMarkForFunding(&_Opyn.CallOpts, _period)
}

// GetExpectedNormalizationFactor is a free data retrieval call binding the contract method 0x24f5f531.
//
// Solidity: function getExpectedNormalizationFactor() view returns(uint256)
func (_Opyn *OpynCaller) GetExpectedNormalizationFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "getExpectedNormalizationFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpectedNormalizationFactor is a free data retrieval call binding the contract method 0x24f5f531.
//
// Solidity: function getExpectedNormalizationFactor() view returns(uint256)
func (_Opyn *OpynSession) GetExpectedNormalizationFactor() (*big.Int, error) {
	return _Opyn.Contract.GetExpectedNormalizationFactor(&_Opyn.CallOpts)
}

// GetExpectedNormalizationFactor is a free data retrieval call binding the contract method 0x24f5f531.
//
// Solidity: function getExpectedNormalizationFactor() view returns(uint256)
func (_Opyn *OpynCallerSession) GetExpectedNormalizationFactor() (*big.Int, error) {
	return _Opyn.Contract.GetExpectedNormalizationFactor(&_Opyn.CallOpts)
}

// GetIndex is a free data retrieval call binding the contract method 0x8cd21d7c.
//
// Solidity: function getIndex(uint32 _period) view returns(uint256)
func (_Opyn *OpynCaller) GetIndex(opts *bind.CallOpts, _period uint32) (*big.Int, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "getIndex", _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndex is a free data retrieval call binding the contract method 0x8cd21d7c.
//
// Solidity: function getIndex(uint32 _period) view returns(uint256)
func (_Opyn *OpynSession) GetIndex(_period uint32) (*big.Int, error) {
	return _Opyn.Contract.GetIndex(&_Opyn.CallOpts, _period)
}

// GetIndex is a free data retrieval call binding the contract method 0x8cd21d7c.
//
// Solidity: function getIndex(uint32 _period) view returns(uint256)
func (_Opyn *OpynCallerSession) GetIndex(_period uint32) (*big.Int, error) {
	return _Opyn.Contract.GetIndex(&_Opyn.CallOpts, _period)
}

// GetUnscaledIndex is a free data retrieval call binding the contract method 0x15aded83.
//
// Solidity: function getUnscaledIndex(uint32 _period) view returns(uint256)
func (_Opyn *OpynCaller) GetUnscaledIndex(opts *bind.CallOpts, _period uint32) (*big.Int, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "getUnscaledIndex", _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnscaledIndex is a free data retrieval call binding the contract method 0x15aded83.
//
// Solidity: function getUnscaledIndex(uint32 _period) view returns(uint256)
func (_Opyn *OpynSession) GetUnscaledIndex(_period uint32) (*big.Int, error) {
	return _Opyn.Contract.GetUnscaledIndex(&_Opyn.CallOpts, _period)
}

// GetUnscaledIndex is a free data retrieval call binding the contract method 0x15aded83.
//
// Solidity: function getUnscaledIndex(uint32 _period) view returns(uint256)
func (_Opyn *OpynCallerSession) GetUnscaledIndex(_period uint32) (*big.Int, error) {
	return _Opyn.Contract.GetUnscaledIndex(&_Opyn.CallOpts, _period)
}

// IndexForSettlement is a free data retrieval call binding the contract method 0xde4a427a.
//
// Solidity: function indexForSettlement() view returns(uint256)
func (_Opyn *OpynCaller) IndexForSettlement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "indexForSettlement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexForSettlement is a free data retrieval call binding the contract method 0xde4a427a.
//
// Solidity: function indexForSettlement() view returns(uint256)
func (_Opyn *OpynSession) IndexForSettlement() (*big.Int, error) {
	return _Opyn.Contract.IndexForSettlement(&_Opyn.CallOpts)
}

// IndexForSettlement is a free data retrieval call binding the contract method 0xde4a427a.
//
// Solidity: function indexForSettlement() view returns(uint256)
func (_Opyn *OpynCallerSession) IndexForSettlement() (*big.Int, error) {
	return _Opyn.Contract.IndexForSettlement(&_Opyn.CallOpts)
}

// IsShutDown is a free data retrieval call binding the contract method 0xff947525.
//
// Solidity: function isShutDown() view returns(bool)
func (_Opyn *OpynCaller) IsShutDown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "isShutDown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutDown is a free data retrieval call binding the contract method 0xff947525.
//
// Solidity: function isShutDown() view returns(bool)
func (_Opyn *OpynSession) IsShutDown() (bool, error) {
	return _Opyn.Contract.IsShutDown(&_Opyn.CallOpts)
}

// IsShutDown is a free data retrieval call binding the contract method 0xff947525.
//
// Solidity: function isShutDown() view returns(bool)
func (_Opyn *OpynCallerSession) IsShutDown() (bool, error) {
	return _Opyn.Contract.IsShutDown(&_Opyn.CallOpts)
}

// IsSystemPaused is a free data retrieval call binding the contract method 0x7691c4ac.
//
// Solidity: function isSystemPaused() view returns(bool)
func (_Opyn *OpynCaller) IsSystemPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "isSystemPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSystemPaused is a free data retrieval call binding the contract method 0x7691c4ac.
//
// Solidity: function isSystemPaused() view returns(bool)
func (_Opyn *OpynSession) IsSystemPaused() (bool, error) {
	return _Opyn.Contract.IsSystemPaused(&_Opyn.CallOpts)
}

// IsSystemPaused is a free data retrieval call binding the contract method 0x7691c4ac.
//
// Solidity: function isSystemPaused() view returns(bool)
func (_Opyn *OpynCallerSession) IsSystemPaused() (bool, error) {
	return _Opyn.Contract.IsSystemPaused(&_Opyn.CallOpts)
}

// IsVaultSafe is a free data retrieval call binding the contract method 0xa847e674.
//
// Solidity: function isVaultSafe(uint256 _vaultId) view returns(bool)
func (_Opyn *OpynCaller) IsVaultSafe(opts *bind.CallOpts, _vaultId *big.Int) (bool, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "isVaultSafe", _vaultId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVaultSafe is a free data retrieval call binding the contract method 0xa847e674.
//
// Solidity: function isVaultSafe(uint256 _vaultId) view returns(bool)
func (_Opyn *OpynSession) IsVaultSafe(_vaultId *big.Int) (bool, error) {
	return _Opyn.Contract.IsVaultSafe(&_Opyn.CallOpts, _vaultId)
}

// IsVaultSafe is a free data retrieval call binding the contract method 0xa847e674.
//
// Solidity: function isVaultSafe(uint256 _vaultId) view returns(bool)
func (_Opyn *OpynCallerSession) IsVaultSafe(_vaultId *big.Int) (bool, error) {
	return _Opyn.Contract.IsVaultSafe(&_Opyn.CallOpts, _vaultId)
}

// LastFundingUpdateTimestamp is a free data retrieval call binding the contract method 0x4be2822c.
//
// Solidity: function lastFundingUpdateTimestamp() view returns(uint128)
func (_Opyn *OpynCaller) LastFundingUpdateTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "lastFundingUpdateTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFundingUpdateTimestamp is a free data retrieval call binding the contract method 0x4be2822c.
//
// Solidity: function lastFundingUpdateTimestamp() view returns(uint128)
func (_Opyn *OpynSession) LastFundingUpdateTimestamp() (*big.Int, error) {
	return _Opyn.Contract.LastFundingUpdateTimestamp(&_Opyn.CallOpts)
}

// LastFundingUpdateTimestamp is a free data retrieval call binding the contract method 0x4be2822c.
//
// Solidity: function lastFundingUpdateTimestamp() view returns(uint128)
func (_Opyn *OpynCallerSession) LastFundingUpdateTimestamp() (*big.Int, error) {
	return _Opyn.Contract.LastFundingUpdateTimestamp(&_Opyn.CallOpts)
}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_Opyn *OpynCaller) LastPauseTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "lastPauseTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_Opyn *OpynSession) LastPauseTime() (*big.Int, error) {
	return _Opyn.Contract.LastPauseTime(&_Opyn.CallOpts)
}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_Opyn *OpynCallerSession) LastPauseTime() (*big.Int, error) {
	return _Opyn.Contract.LastPauseTime(&_Opyn.CallOpts)
}

// NormalizationFactor is a free data retrieval call binding the contract method 0xd5272584.
//
// Solidity: function normalizationFactor() view returns(uint128)
func (_Opyn *OpynCaller) NormalizationFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "normalizationFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NormalizationFactor is a free data retrieval call binding the contract method 0xd5272584.
//
// Solidity: function normalizationFactor() view returns(uint128)
func (_Opyn *OpynSession) NormalizationFactor() (*big.Int, error) {
	return _Opyn.Contract.NormalizationFactor(&_Opyn.CallOpts)
}

// NormalizationFactor is a free data retrieval call binding the contract method 0xd5272584.
//
// Solidity: function normalizationFactor() view returns(uint128)
func (_Opyn *OpynCallerSession) NormalizationFactor() (*big.Int, error) {
	return _Opyn.Contract.NormalizationFactor(&_Opyn.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Opyn *OpynCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Opyn *OpynSession) Oracle() (common.Address, error) {
	return _Opyn.Contract.Oracle(&_Opyn.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Opyn *OpynCallerSession) Oracle() (common.Address, error) {
	return _Opyn.Contract.Oracle(&_Opyn.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Opyn *OpynCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Opyn *OpynSession) Owner() (common.Address, error) {
	return _Opyn.Contract.Owner(&_Opyn.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Opyn *OpynCallerSession) Owner() (common.Address, error) {
	return _Opyn.Contract.Owner(&_Opyn.CallOpts)
}

// PausesLeft is a free data retrieval call binding the contract method 0x63b38ae4.
//
// Solidity: function pausesLeft() view returns(uint256)
func (_Opyn *OpynCaller) PausesLeft(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "pausesLeft")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PausesLeft is a free data retrieval call binding the contract method 0x63b38ae4.
//
// Solidity: function pausesLeft() view returns(uint256)
func (_Opyn *OpynSession) PausesLeft() (*big.Int, error) {
	return _Opyn.Contract.PausesLeft(&_Opyn.CallOpts)
}

// PausesLeft is a free data retrieval call binding the contract method 0x63b38ae4.
//
// Solidity: function pausesLeft() view returns(uint256)
func (_Opyn *OpynCallerSession) PausesLeft() (*big.Int, error) {
	return _Opyn.Contract.PausesLeft(&_Opyn.CallOpts)
}

// QuoteCurrency is a free data retrieval call binding the contract method 0x82564bca.
//
// Solidity: function quoteCurrency() view returns(address)
func (_Opyn *OpynCaller) QuoteCurrency(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "quoteCurrency")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteCurrency is a free data retrieval call binding the contract method 0x82564bca.
//
// Solidity: function quoteCurrency() view returns(address)
func (_Opyn *OpynSession) QuoteCurrency() (common.Address, error) {
	return _Opyn.Contract.QuoteCurrency(&_Opyn.CallOpts)
}

// QuoteCurrency is a free data retrieval call binding the contract method 0x82564bca.
//
// Solidity: function quoteCurrency() view returns(address)
func (_Opyn *OpynCallerSession) QuoteCurrency() (common.Address, error) {
	return _Opyn.Contract.QuoteCurrency(&_Opyn.CallOpts)
}

// ShortPowerPerp is a free data retrieval call binding the contract method 0x9d4c9442.
//
// Solidity: function shortPowerPerp() view returns(address)
func (_Opyn *OpynCaller) ShortPowerPerp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "shortPowerPerp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShortPowerPerp is a free data retrieval call binding the contract method 0x9d4c9442.
//
// Solidity: function shortPowerPerp() view returns(address)
func (_Opyn *OpynSession) ShortPowerPerp() (common.Address, error) {
	return _Opyn.Contract.ShortPowerPerp(&_Opyn.CallOpts)
}

// ShortPowerPerp is a free data retrieval call binding the contract method 0x9d4c9442.
//
// Solidity: function shortPowerPerp() view returns(address)
func (_Opyn *OpynCallerSession) ShortPowerPerp() (common.Address, error) {
	return _Opyn.Contract.ShortPowerPerp(&_Opyn.CallOpts)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address operator, uint32 NftCollateralId, uint96 collateralAmount, uint128 shortAmount)
func (_Opyn *OpynCaller) Vaults(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Operator         common.Address
	NftCollateralId  uint32
	CollateralAmount *big.Int
	ShortAmount      *big.Int
}, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "vaults", arg0)

	outstruct := new(struct {
		Operator         common.Address
		NftCollateralId  uint32
		CollateralAmount *big.Int
		ShortAmount      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Operator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftCollateralId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.CollateralAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ShortAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address operator, uint32 NftCollateralId, uint96 collateralAmount, uint128 shortAmount)
func (_Opyn *OpynSession) Vaults(arg0 *big.Int) (struct {
	Operator         common.Address
	NftCollateralId  uint32
	CollateralAmount *big.Int
	ShortAmount      *big.Int
}, error) {
	return _Opyn.Contract.Vaults(&_Opyn.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address operator, uint32 NftCollateralId, uint96 collateralAmount, uint128 shortAmount)
func (_Opyn *OpynCallerSession) Vaults(arg0 *big.Int) (struct {
	Operator         common.Address
	NftCollateralId  uint32
	CollateralAmount *big.Int
	ShortAmount      *big.Int
}, error) {
	return _Opyn.Contract.Vaults(&_Opyn.CallOpts, arg0)
}

// WPowerPerp is a free data retrieval call binding the contract method 0x7f07b130.
//
// Solidity: function wPowerPerp() view returns(address)
func (_Opyn *OpynCaller) WPowerPerp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "wPowerPerp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WPowerPerp is a free data retrieval call binding the contract method 0x7f07b130.
//
// Solidity: function wPowerPerp() view returns(address)
func (_Opyn *OpynSession) WPowerPerp() (common.Address, error) {
	return _Opyn.Contract.WPowerPerp(&_Opyn.CallOpts)
}

// WPowerPerp is a free data retrieval call binding the contract method 0x7f07b130.
//
// Solidity: function wPowerPerp() view returns(address)
func (_Opyn *OpynCallerSession) WPowerPerp() (common.Address, error) {
	return _Opyn.Contract.WPowerPerp(&_Opyn.CallOpts)
}

// WPowerPerpPool is a free data retrieval call binding the contract method 0xb707ab99.
//
// Solidity: function wPowerPerpPool() view returns(address)
func (_Opyn *OpynCaller) WPowerPerpPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "wPowerPerpPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WPowerPerpPool is a free data retrieval call binding the contract method 0xb707ab99.
//
// Solidity: function wPowerPerpPool() view returns(address)
func (_Opyn *OpynSession) WPowerPerpPool() (common.Address, error) {
	return _Opyn.Contract.WPowerPerpPool(&_Opyn.CallOpts)
}

// WPowerPerpPool is a free data retrieval call binding the contract method 0xb707ab99.
//
// Solidity: function wPowerPerpPool() view returns(address)
func (_Opyn *OpynCallerSession) WPowerPerpPool() (common.Address, error) {
	return _Opyn.Contract.WPowerPerpPool(&_Opyn.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Opyn *OpynCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Opyn.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Opyn *OpynSession) Weth() (common.Address, error) {
	return _Opyn.Contract.Weth(&_Opyn.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Opyn *OpynCallerSession) Weth() (common.Address, error) {
	return _Opyn.Contract.Weth(&_Opyn.CallOpts)
}

// ApplyFunding is a paid mutator transaction binding the contract method 0x200f4b8d.
//
// Solidity: function applyFunding() returns()
func (_Opyn *OpynTransactor) ApplyFunding(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "applyFunding")
}

// ApplyFunding is a paid mutator transaction binding the contract method 0x200f4b8d.
//
// Solidity: function applyFunding() returns()
func (_Opyn *OpynSession) ApplyFunding() (*types.Transaction, error) {
	return _Opyn.Contract.ApplyFunding(&_Opyn.TransactOpts)
}

// ApplyFunding is a paid mutator transaction binding the contract method 0x200f4b8d.
//
// Solidity: function applyFunding() returns()
func (_Opyn *OpynTransactorSession) ApplyFunding() (*types.Transaction, error) {
	return _Opyn.Contract.ApplyFunding(&_Opyn.TransactOpts)
}

// BurnPowerPerpAmount is a paid mutator transaction binding the contract method 0x4394318d.
//
// Solidity: function burnPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _withdrawAmount) returns(uint256)
func (_Opyn *OpynTransactor) BurnPowerPerpAmount(opts *bind.TransactOpts, _vaultId *big.Int, _powerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "burnPowerPerpAmount", _vaultId, _powerPerpAmount, _withdrawAmount)
}

// BurnPowerPerpAmount is a paid mutator transaction binding the contract method 0x4394318d.
//
// Solidity: function burnPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _withdrawAmount) returns(uint256)
func (_Opyn *OpynSession) BurnPowerPerpAmount(_vaultId *big.Int, _powerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.BurnPowerPerpAmount(&_Opyn.TransactOpts, _vaultId, _powerPerpAmount, _withdrawAmount)
}

// BurnPowerPerpAmount is a paid mutator transaction binding the contract method 0x4394318d.
//
// Solidity: function burnPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _withdrawAmount) returns(uint256)
func (_Opyn *OpynTransactorSession) BurnPowerPerpAmount(_vaultId *big.Int, _powerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.BurnPowerPerpAmount(&_Opyn.TransactOpts, _vaultId, _powerPerpAmount, _withdrawAmount)
}

// BurnWPowerPerpAmount is a paid mutator transaction binding the contract method 0x8632cb03.
//
// Solidity: function burnWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _withdrawAmount) returns()
func (_Opyn *OpynTransactor) BurnWPowerPerpAmount(opts *bind.TransactOpts, _vaultId *big.Int, _wPowerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "burnWPowerPerpAmount", _vaultId, _wPowerPerpAmount, _withdrawAmount)
}

// BurnWPowerPerpAmount is a paid mutator transaction binding the contract method 0x8632cb03.
//
// Solidity: function burnWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _withdrawAmount) returns()
func (_Opyn *OpynSession) BurnWPowerPerpAmount(_vaultId *big.Int, _wPowerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.BurnWPowerPerpAmount(&_Opyn.TransactOpts, _vaultId, _wPowerPerpAmount, _withdrawAmount)
}

// BurnWPowerPerpAmount is a paid mutator transaction binding the contract method 0x8632cb03.
//
// Solidity: function burnWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _withdrawAmount) returns()
func (_Opyn *OpynTransactorSession) BurnWPowerPerpAmount(_vaultId *big.Int, _wPowerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.BurnWPowerPerpAmount(&_Opyn.TransactOpts, _vaultId, _wPowerPerpAmount, _withdrawAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _vaultId) payable returns()
func (_Opyn *OpynTransactor) Deposit(opts *bind.TransactOpts, _vaultId *big.Int) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "deposit", _vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _vaultId) payable returns()
func (_Opyn *OpynSession) Deposit(_vaultId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.Deposit(&_Opyn.TransactOpts, _vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _vaultId) payable returns()
func (_Opyn *OpynTransactorSession) Deposit(_vaultId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.Deposit(&_Opyn.TransactOpts, _vaultId)
}

// DepositUniPositionToken is a paid mutator transaction binding the contract method 0x91b8d34a.
//
// Solidity: function depositUniPositionToken(uint256 _vaultId, uint256 _uniTokenId) returns()
func (_Opyn *OpynTransactor) DepositUniPositionToken(opts *bind.TransactOpts, _vaultId *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "depositUniPositionToken", _vaultId, _uniTokenId)
}

// DepositUniPositionToken is a paid mutator transaction binding the contract method 0x91b8d34a.
//
// Solidity: function depositUniPositionToken(uint256 _vaultId, uint256 _uniTokenId) returns()
func (_Opyn *OpynSession) DepositUniPositionToken(_vaultId *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.DepositUniPositionToken(&_Opyn.TransactOpts, _vaultId, _uniTokenId)
}

// DepositUniPositionToken is a paid mutator transaction binding the contract method 0x91b8d34a.
//
// Solidity: function depositUniPositionToken(uint256 _vaultId, uint256 _uniTokenId) returns()
func (_Opyn *OpynTransactorSession) DepositUniPositionToken(_vaultId *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.DepositUniPositionToken(&_Opyn.TransactOpts, _vaultId, _uniTokenId)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_Opyn *OpynTransactor) Donate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "donate")
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_Opyn *OpynSession) Donate() (*types.Transaction, error) {
	return _Opyn.Contract.Donate(&_Opyn.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_Opyn *OpynTransactorSession) Donate() (*types.Transaction, error) {
	return _Opyn.Contract.Donate(&_Opyn.TransactOpts)
}

// Liquidate is a paid mutator transaction binding the contract method 0xd296d1f1.
//
// Solidity: function liquidate(uint256 _vaultId, uint256 _maxDebtAmount) returns(uint256)
func (_Opyn *OpynTransactor) Liquidate(opts *bind.TransactOpts, _vaultId *big.Int, _maxDebtAmount *big.Int) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "liquidate", _vaultId, _maxDebtAmount)
}

// Liquidate is a paid mutator transaction binding the contract method 0xd296d1f1.
//
// Solidity: function liquidate(uint256 _vaultId, uint256 _maxDebtAmount) returns(uint256)
func (_Opyn *OpynSession) Liquidate(_vaultId *big.Int, _maxDebtAmount *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.Liquidate(&_Opyn.TransactOpts, _vaultId, _maxDebtAmount)
}

// Liquidate is a paid mutator transaction binding the contract method 0xd296d1f1.
//
// Solidity: function liquidate(uint256 _vaultId, uint256 _maxDebtAmount) returns(uint256)
func (_Opyn *OpynTransactorSession) Liquidate(_vaultId *big.Int, _maxDebtAmount *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.Liquidate(&_Opyn.TransactOpts, _vaultId, _maxDebtAmount)
}

// MintPowerPerpAmount is a paid mutator transaction binding the contract method 0x1bf7bf6c.
//
// Solidity: function mintPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _uniTokenId) payable returns(uint256, uint256)
func (_Opyn *OpynTransactor) MintPowerPerpAmount(opts *bind.TransactOpts, _vaultId *big.Int, _powerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "mintPowerPerpAmount", _vaultId, _powerPerpAmount, _uniTokenId)
}

// MintPowerPerpAmount is a paid mutator transaction binding the contract method 0x1bf7bf6c.
//
// Solidity: function mintPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _uniTokenId) payable returns(uint256, uint256)
func (_Opyn *OpynSession) MintPowerPerpAmount(_vaultId *big.Int, _powerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.MintPowerPerpAmount(&_Opyn.TransactOpts, _vaultId, _powerPerpAmount, _uniTokenId)
}

// MintPowerPerpAmount is a paid mutator transaction binding the contract method 0x1bf7bf6c.
//
// Solidity: function mintPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _uniTokenId) payable returns(uint256, uint256)
func (_Opyn *OpynTransactorSession) MintPowerPerpAmount(_vaultId *big.Int, _powerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.MintPowerPerpAmount(&_Opyn.TransactOpts, _vaultId, _powerPerpAmount, _uniTokenId)
}

// MintWPowerPerpAmount is a paid mutator transaction binding the contract method 0x39467918.
//
// Solidity: function mintWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _uniTokenId) payable returns(uint256)
func (_Opyn *OpynTransactor) MintWPowerPerpAmount(opts *bind.TransactOpts, _vaultId *big.Int, _wPowerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "mintWPowerPerpAmount", _vaultId, _wPowerPerpAmount, _uniTokenId)
}

// MintWPowerPerpAmount is a paid mutator transaction binding the contract method 0x39467918.
//
// Solidity: function mintWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _uniTokenId) payable returns(uint256)
func (_Opyn *OpynSession) MintWPowerPerpAmount(_vaultId *big.Int, _wPowerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.MintWPowerPerpAmount(&_Opyn.TransactOpts, _vaultId, _wPowerPerpAmount, _uniTokenId)
}

// MintWPowerPerpAmount is a paid mutator transaction binding the contract method 0x39467918.
//
// Solidity: function mintWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _uniTokenId) payable returns(uint256)
func (_Opyn *OpynTransactorSession) MintWPowerPerpAmount(_vaultId *big.Int, _wPowerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.MintWPowerPerpAmount(&_Opyn.TransactOpts, _vaultId, _wPowerPerpAmount, _uniTokenId)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Opyn *OpynTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Opyn *OpynSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Opyn.Contract.OnERC721Received(&_Opyn.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Opyn *OpynTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Opyn.Contract.OnERC721Received(&_Opyn.TransactOpts, arg0, arg1, arg2, arg3)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Opyn *OpynTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Opyn *OpynSession) Pause() (*types.Transaction, error) {
	return _Opyn.Contract.Pause(&_Opyn.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Opyn *OpynTransactorSession) Pause() (*types.Transaction, error) {
	return _Opyn.Contract.Pause(&_Opyn.TransactOpts)
}

// RedeemLong is a paid mutator transaction binding the contract method 0xac6cd5ef.
//
// Solidity: function redeemLong(uint256 _wPerpAmount) returns()
func (_Opyn *OpynTransactor) RedeemLong(opts *bind.TransactOpts, _wPerpAmount *big.Int) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "redeemLong", _wPerpAmount)
}

// RedeemLong is a paid mutator transaction binding the contract method 0xac6cd5ef.
//
// Solidity: function redeemLong(uint256 _wPerpAmount) returns()
func (_Opyn *OpynSession) RedeemLong(_wPerpAmount *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.RedeemLong(&_Opyn.TransactOpts, _wPerpAmount)
}

// RedeemLong is a paid mutator transaction binding the contract method 0xac6cd5ef.
//
// Solidity: function redeemLong(uint256 _wPerpAmount) returns()
func (_Opyn *OpynTransactorSession) RedeemLong(_wPerpAmount *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.RedeemLong(&_Opyn.TransactOpts, _wPerpAmount)
}

// RedeemShort is a paid mutator transaction binding the contract method 0x97efa942.
//
// Solidity: function redeemShort(uint256 _vaultId) returns()
func (_Opyn *OpynTransactor) RedeemShort(opts *bind.TransactOpts, _vaultId *big.Int) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "redeemShort", _vaultId)
}

// RedeemShort is a paid mutator transaction binding the contract method 0x97efa942.
//
// Solidity: function redeemShort(uint256 _vaultId) returns()
func (_Opyn *OpynSession) RedeemShort(_vaultId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.RedeemShort(&_Opyn.TransactOpts, _vaultId)
}

// RedeemShort is a paid mutator transaction binding the contract method 0x97efa942.
//
// Solidity: function redeemShort(uint256 _vaultId) returns()
func (_Opyn *OpynTransactorSession) RedeemShort(_vaultId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.RedeemShort(&_Opyn.TransactOpts, _vaultId)
}

// ReduceDebt is a paid mutator transaction binding the contract method 0xc9e77ee8.
//
// Solidity: function reduceDebt(uint256 _vaultId) returns()
func (_Opyn *OpynTransactor) ReduceDebt(opts *bind.TransactOpts, _vaultId *big.Int) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "reduceDebt", _vaultId)
}

// ReduceDebt is a paid mutator transaction binding the contract method 0xc9e77ee8.
//
// Solidity: function reduceDebt(uint256 _vaultId) returns()
func (_Opyn *OpynSession) ReduceDebt(_vaultId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.ReduceDebt(&_Opyn.TransactOpts, _vaultId)
}

// ReduceDebt is a paid mutator transaction binding the contract method 0xc9e77ee8.
//
// Solidity: function reduceDebt(uint256 _vaultId) returns()
func (_Opyn *OpynTransactorSession) ReduceDebt(_vaultId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.ReduceDebt(&_Opyn.TransactOpts, _vaultId)
}

// ReduceDebtShutdown is a paid mutator transaction binding the contract method 0xfbfc6bc0.
//
// Solidity: function reduceDebtShutdown(uint256 _vaultId) returns()
func (_Opyn *OpynTransactor) ReduceDebtShutdown(opts *bind.TransactOpts, _vaultId *big.Int) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "reduceDebtShutdown", _vaultId)
}

// ReduceDebtShutdown is a paid mutator transaction binding the contract method 0xfbfc6bc0.
//
// Solidity: function reduceDebtShutdown(uint256 _vaultId) returns()
func (_Opyn *OpynSession) ReduceDebtShutdown(_vaultId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.ReduceDebtShutdown(&_Opyn.TransactOpts, _vaultId)
}

// ReduceDebtShutdown is a paid mutator transaction binding the contract method 0xfbfc6bc0.
//
// Solidity: function reduceDebtShutdown(uint256 _vaultId) returns()
func (_Opyn *OpynTransactorSession) ReduceDebtShutdown(_vaultId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.ReduceDebtShutdown(&_Opyn.TransactOpts, _vaultId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Opyn *OpynTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Opyn *OpynSession) RenounceOwnership() (*types.Transaction, error) {
	return _Opyn.Contract.RenounceOwnership(&_Opyn.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Opyn *OpynTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Opyn.Contract.RenounceOwnership(&_Opyn.TransactOpts)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _newFeeRate) returns()
func (_Opyn *OpynTransactor) SetFeeRate(opts *bind.TransactOpts, _newFeeRate *big.Int) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "setFeeRate", _newFeeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _newFeeRate) returns()
func (_Opyn *OpynSession) SetFeeRate(_newFeeRate *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.SetFeeRate(&_Opyn.TransactOpts, _newFeeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _newFeeRate) returns()
func (_Opyn *OpynTransactorSession) SetFeeRate(_newFeeRate *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.SetFeeRate(&_Opyn.TransactOpts, _newFeeRate)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _newFeeRecipient) returns()
func (_Opyn *OpynTransactor) SetFeeRecipient(opts *bind.TransactOpts, _newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "setFeeRecipient", _newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _newFeeRecipient) returns()
func (_Opyn *OpynSession) SetFeeRecipient(_newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Opyn.Contract.SetFeeRecipient(&_Opyn.TransactOpts, _newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _newFeeRecipient) returns()
func (_Opyn *OpynTransactorSession) SetFeeRecipient(_newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Opyn.Contract.SetFeeRecipient(&_Opyn.TransactOpts, _newFeeRecipient)
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_Opyn *OpynTransactor) ShutDown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "shutDown")
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_Opyn *OpynSession) ShutDown() (*types.Transaction, error) {
	return _Opyn.Contract.ShutDown(&_Opyn.TransactOpts)
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_Opyn *OpynTransactorSession) ShutDown() (*types.Transaction, error) {
	return _Opyn.Contract.ShutDown(&_Opyn.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Opyn *OpynTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Opyn *OpynSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Opyn.Contract.TransferOwnership(&_Opyn.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Opyn *OpynTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Opyn.Contract.TransferOwnership(&_Opyn.TransactOpts, newOwner)
}

// UnPauseAnyone is a paid mutator transaction binding the contract method 0x8146b09f.
//
// Solidity: function unPauseAnyone() returns()
func (_Opyn *OpynTransactor) UnPauseAnyone(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "unPauseAnyone")
}

// UnPauseAnyone is a paid mutator transaction binding the contract method 0x8146b09f.
//
// Solidity: function unPauseAnyone() returns()
func (_Opyn *OpynSession) UnPauseAnyone() (*types.Transaction, error) {
	return _Opyn.Contract.UnPauseAnyone(&_Opyn.TransactOpts)
}

// UnPauseAnyone is a paid mutator transaction binding the contract method 0x8146b09f.
//
// Solidity: function unPauseAnyone() returns()
func (_Opyn *OpynTransactorSession) UnPauseAnyone() (*types.Transaction, error) {
	return _Opyn.Contract.UnPauseAnyone(&_Opyn.TransactOpts)
}

// UnPauseOwner is a paid mutator transaction binding the contract method 0x07633669.
//
// Solidity: function unPauseOwner() returns()
func (_Opyn *OpynTransactor) UnPauseOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "unPauseOwner")
}

// UnPauseOwner is a paid mutator transaction binding the contract method 0x07633669.
//
// Solidity: function unPauseOwner() returns()
func (_Opyn *OpynSession) UnPauseOwner() (*types.Transaction, error) {
	return _Opyn.Contract.UnPauseOwner(&_Opyn.TransactOpts)
}

// UnPauseOwner is a paid mutator transaction binding the contract method 0x07633669.
//
// Solidity: function unPauseOwner() returns()
func (_Opyn *OpynTransactorSession) UnPauseOwner() (*types.Transaction, error) {
	return _Opyn.Contract.UnPauseOwner(&_Opyn.TransactOpts)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xc65a391d.
//
// Solidity: function updateOperator(uint256 _vaultId, address _operator) returns()
func (_Opyn *OpynTransactor) UpdateOperator(opts *bind.TransactOpts, _vaultId *big.Int, _operator common.Address) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "updateOperator", _vaultId, _operator)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xc65a391d.
//
// Solidity: function updateOperator(uint256 _vaultId, address _operator) returns()
func (_Opyn *OpynSession) UpdateOperator(_vaultId *big.Int, _operator common.Address) (*types.Transaction, error) {
	return _Opyn.Contract.UpdateOperator(&_Opyn.TransactOpts, _vaultId, _operator)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xc65a391d.
//
// Solidity: function updateOperator(uint256 _vaultId, address _operator) returns()
func (_Opyn *OpynTransactorSession) UpdateOperator(_vaultId *big.Int, _operator common.Address) (*types.Transaction, error) {
	return _Opyn.Contract.UpdateOperator(&_Opyn.TransactOpts, _vaultId, _operator)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _vaultId, uint256 _amount) returns()
func (_Opyn *OpynTransactor) Withdraw(opts *bind.TransactOpts, _vaultId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "withdraw", _vaultId, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _vaultId, uint256 _amount) returns()
func (_Opyn *OpynSession) Withdraw(_vaultId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.Withdraw(&_Opyn.TransactOpts, _vaultId, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _vaultId, uint256 _amount) returns()
func (_Opyn *OpynTransactorSession) Withdraw(_vaultId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.Withdraw(&_Opyn.TransactOpts, _vaultId, _amount)
}

// WithdrawUniPositionToken is a paid mutator transaction binding the contract method 0x713d517f.
//
// Solidity: function withdrawUniPositionToken(uint256 _vaultId) returns()
func (_Opyn *OpynTransactor) WithdrawUniPositionToken(opts *bind.TransactOpts, _vaultId *big.Int) (*types.Transaction, error) {
	return _Opyn.contract.Transact(opts, "withdrawUniPositionToken", _vaultId)
}

// WithdrawUniPositionToken is a paid mutator transaction binding the contract method 0x713d517f.
//
// Solidity: function withdrawUniPositionToken(uint256 _vaultId) returns()
func (_Opyn *OpynSession) WithdrawUniPositionToken(_vaultId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.WithdrawUniPositionToken(&_Opyn.TransactOpts, _vaultId)
}

// WithdrawUniPositionToken is a paid mutator transaction binding the contract method 0x713d517f.
//
// Solidity: function withdrawUniPositionToken(uint256 _vaultId) returns()
func (_Opyn *OpynTransactorSession) WithdrawUniPositionToken(_vaultId *big.Int) (*types.Transaction, error) {
	return _Opyn.Contract.WithdrawUniPositionToken(&_Opyn.TransactOpts, _vaultId)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Opyn *OpynTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opyn.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Opyn *OpynSession) Receive() (*types.Transaction, error) {
	return _Opyn.Contract.Receive(&_Opyn.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Opyn *OpynTransactorSession) Receive() (*types.Transaction, error) {
	return _Opyn.Contract.Receive(&_Opyn.TransactOpts)
}

// OpynBurnShortIterator is returned from FilterBurnShort and is used to iterate over the raw logs and unpacked data for BurnShort events raised by the Opyn contract.
type OpynBurnShortIterator struct {
	Event *OpynBurnShort // Event containing the contract specifics and raw log

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
func (it *OpynBurnShortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynBurnShort)
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
		it.Event = new(OpynBurnShort)
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
func (it *OpynBurnShortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynBurnShortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynBurnShort represents a BurnShort event raised by the Opyn contract.
type OpynBurnShort struct {
	Sender  common.Address
	Amount  *big.Int
	VaultId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurnShort is a free log retrieval operation binding the contract event 0xea19ffc45b48de6d95594aacff7214dd24595fdb0c60e98c8f0b269058c2f792.
//
// Solidity: event BurnShort(address sender, uint256 amount, uint256 vaultId)
func (_Opyn *OpynFilterer) FilterBurnShort(opts *bind.FilterOpts) (*OpynBurnShortIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "BurnShort")
	if err != nil {
		return nil, err
	}
	return &OpynBurnShortIterator{contract: _Opyn.contract, event: "BurnShort", logs: logs, sub: sub}, nil
}

// WatchBurnShort is a free log subscription operation binding the contract event 0xea19ffc45b48de6d95594aacff7214dd24595fdb0c60e98c8f0b269058c2f792.
//
// Solidity: event BurnShort(address sender, uint256 amount, uint256 vaultId)
func (_Opyn *OpynFilterer) WatchBurnShort(opts *bind.WatchOpts, sink chan<- *OpynBurnShort) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "BurnShort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynBurnShort)
				if err := _Opyn.contract.UnpackLog(event, "BurnShort", log); err != nil {
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

// ParseBurnShort is a log parse operation binding the contract event 0xea19ffc45b48de6d95594aacff7214dd24595fdb0c60e98c8f0b269058c2f792.
//
// Solidity: event BurnShort(address sender, uint256 amount, uint256 vaultId)
func (_Opyn *OpynFilterer) ParseBurnShort(log types.Log) (*OpynBurnShort, error) {
	event := new(OpynBurnShort)
	if err := _Opyn.contract.UnpackLog(event, "BurnShort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynDepositCollateralIterator is returned from FilterDepositCollateral and is used to iterate over the raw logs and unpacked data for DepositCollateral events raised by the Opyn contract.
type OpynDepositCollateralIterator struct {
	Event *OpynDepositCollateral // Event containing the contract specifics and raw log

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
func (it *OpynDepositCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynDepositCollateral)
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
		it.Event = new(OpynDepositCollateral)
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
func (it *OpynDepositCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynDepositCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynDepositCollateral represents a DepositCollateral event raised by the Opyn contract.
type OpynDepositCollateral struct {
	Sender  common.Address
	VaultId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositCollateral is a free log retrieval operation binding the contract event 0x3ca13b7aab12bad7472691fe558faa6b25e99099824a0070a88bd5aa84be610f.
//
// Solidity: event DepositCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Opyn *OpynFilterer) FilterDepositCollateral(opts *bind.FilterOpts) (*OpynDepositCollateralIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "DepositCollateral")
	if err != nil {
		return nil, err
	}
	return &OpynDepositCollateralIterator{contract: _Opyn.contract, event: "DepositCollateral", logs: logs, sub: sub}, nil
}

// WatchDepositCollateral is a free log subscription operation binding the contract event 0x3ca13b7aab12bad7472691fe558faa6b25e99099824a0070a88bd5aa84be610f.
//
// Solidity: event DepositCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Opyn *OpynFilterer) WatchDepositCollateral(opts *bind.WatchOpts, sink chan<- *OpynDepositCollateral) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "DepositCollateral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynDepositCollateral)
				if err := _Opyn.contract.UnpackLog(event, "DepositCollateral", log); err != nil {
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

// ParseDepositCollateral is a log parse operation binding the contract event 0x3ca13b7aab12bad7472691fe558faa6b25e99099824a0070a88bd5aa84be610f.
//
// Solidity: event DepositCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Opyn *OpynFilterer) ParseDepositCollateral(log types.Log) (*OpynDepositCollateral, error) {
	event := new(OpynDepositCollateral)
	if err := _Opyn.contract.UnpackLog(event, "DepositCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynDepositUniPositionTokenIterator is returned from FilterDepositUniPositionToken and is used to iterate over the raw logs and unpacked data for DepositUniPositionToken events raised by the Opyn contract.
type OpynDepositUniPositionTokenIterator struct {
	Event *OpynDepositUniPositionToken // Event containing the contract specifics and raw log

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
func (it *OpynDepositUniPositionTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynDepositUniPositionToken)
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
		it.Event = new(OpynDepositUniPositionToken)
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
func (it *OpynDepositUniPositionTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynDepositUniPositionTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynDepositUniPositionToken represents a DepositUniPositionToken event raised by the Opyn contract.
type OpynDepositUniPositionToken struct {
	Sender  common.Address
	VaultId *big.Int
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositUniPositionToken is a free log retrieval operation binding the contract event 0x3917c2f26ce18614e3aedd1289da672ef6563c5c295f49e9b1697ae0ad315562.
//
// Solidity: event DepositUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Opyn *OpynFilterer) FilterDepositUniPositionToken(opts *bind.FilterOpts) (*OpynDepositUniPositionTokenIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "DepositUniPositionToken")
	if err != nil {
		return nil, err
	}
	return &OpynDepositUniPositionTokenIterator{contract: _Opyn.contract, event: "DepositUniPositionToken", logs: logs, sub: sub}, nil
}

// WatchDepositUniPositionToken is a free log subscription operation binding the contract event 0x3917c2f26ce18614e3aedd1289da672ef6563c5c295f49e9b1697ae0ad315562.
//
// Solidity: event DepositUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Opyn *OpynFilterer) WatchDepositUniPositionToken(opts *bind.WatchOpts, sink chan<- *OpynDepositUniPositionToken) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "DepositUniPositionToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynDepositUniPositionToken)
				if err := _Opyn.contract.UnpackLog(event, "DepositUniPositionToken", log); err != nil {
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

// ParseDepositUniPositionToken is a log parse operation binding the contract event 0x3917c2f26ce18614e3aedd1289da672ef6563c5c295f49e9b1697ae0ad315562.
//
// Solidity: event DepositUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Opyn *OpynFilterer) ParseDepositUniPositionToken(log types.Log) (*OpynDepositUniPositionToken, error) {
	event := new(OpynDepositUniPositionToken)
	if err := _Opyn.contract.UnpackLog(event, "DepositUniPositionToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynFeeRateUpdatedIterator is returned from FilterFeeRateUpdated and is used to iterate over the raw logs and unpacked data for FeeRateUpdated events raised by the Opyn contract.
type OpynFeeRateUpdatedIterator struct {
	Event *OpynFeeRateUpdated // Event containing the contract specifics and raw log

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
func (it *OpynFeeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynFeeRateUpdated)
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
		it.Event = new(OpynFeeRateUpdated)
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
func (it *OpynFeeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynFeeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynFeeRateUpdated represents a FeeRateUpdated event raised by the Opyn contract.
type OpynFeeRateUpdated struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeRateUpdated is a free log retrieval operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFee, uint256 newFee)
func (_Opyn *OpynFilterer) FilterFeeRateUpdated(opts *bind.FilterOpts) (*OpynFeeRateUpdatedIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &OpynFeeRateUpdatedIterator{contract: _Opyn.contract, event: "FeeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRateUpdated is a free log subscription operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFee, uint256 newFee)
func (_Opyn *OpynFilterer) WatchFeeRateUpdated(opts *bind.WatchOpts, sink chan<- *OpynFeeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynFeeRateUpdated)
				if err := _Opyn.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
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

// ParseFeeRateUpdated is a log parse operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFee, uint256 newFee)
func (_Opyn *OpynFilterer) ParseFeeRateUpdated(log types.Log) (*OpynFeeRateUpdated, error) {
	event := new(OpynFeeRateUpdated)
	if err := _Opyn.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynFeeRecipientUpdatedIterator is returned from FilterFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for FeeRecipientUpdated events raised by the Opyn contract.
type OpynFeeRecipientUpdatedIterator struct {
	Event *OpynFeeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *OpynFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynFeeRecipientUpdated)
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
		it.Event = new(OpynFeeRecipientUpdated)
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
func (it *OpynFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynFeeRecipientUpdated represents a FeeRecipientUpdated event raised by the Opyn contract.
type OpynFeeRecipientUpdated struct {
	OldFeeRecipient common.Address
	NewFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientUpdated is a free log retrieval operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address oldFeeRecipient, address newFeeRecipient)
func (_Opyn *OpynFilterer) FilterFeeRecipientUpdated(opts *bind.FilterOpts) (*OpynFeeRecipientUpdatedIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "FeeRecipientUpdated")
	if err != nil {
		return nil, err
	}
	return &OpynFeeRecipientUpdatedIterator{contract: _Opyn.contract, event: "FeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientUpdated is a free log subscription operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address oldFeeRecipient, address newFeeRecipient)
func (_Opyn *OpynFilterer) WatchFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *OpynFeeRecipientUpdated) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "FeeRecipientUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynFeeRecipientUpdated)
				if err := _Opyn.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
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

// ParseFeeRecipientUpdated is a log parse operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address oldFeeRecipient, address newFeeRecipient)
func (_Opyn *OpynFilterer) ParseFeeRecipientUpdated(log types.Log) (*OpynFeeRecipientUpdated, error) {
	event := new(OpynFeeRecipientUpdated)
	if err := _Opyn.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the Opyn contract.
type OpynLiquidateIterator struct {
	Event *OpynLiquidate // Event containing the contract specifics and raw log

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
func (it *OpynLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynLiquidate)
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
		it.Event = new(OpynLiquidate)
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
func (it *OpynLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynLiquidate represents a Liquidate event raised by the Opyn contract.
type OpynLiquidate struct {
	Liquidator     common.Address
	VaultId        *big.Int
	DebtAmount     *big.Int
	CollateralPaid *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0x158ba9ab7bbbd08eeffa4753bad41f4d450e24831d293427308badf3eadd8c76.
//
// Solidity: event Liquidate(address liquidator, uint256 vaultId, uint256 debtAmount, uint256 collateralPaid)
func (_Opyn *OpynFilterer) FilterLiquidate(opts *bind.FilterOpts) (*OpynLiquidateIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "Liquidate")
	if err != nil {
		return nil, err
	}
	return &OpynLiquidateIterator{contract: _Opyn.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0x158ba9ab7bbbd08eeffa4753bad41f4d450e24831d293427308badf3eadd8c76.
//
// Solidity: event Liquidate(address liquidator, uint256 vaultId, uint256 debtAmount, uint256 collateralPaid)
func (_Opyn *OpynFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *OpynLiquidate) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "Liquidate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynLiquidate)
				if err := _Opyn.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0x158ba9ab7bbbd08eeffa4753bad41f4d450e24831d293427308badf3eadd8c76.
//
// Solidity: event Liquidate(address liquidator, uint256 vaultId, uint256 debtAmount, uint256 collateralPaid)
func (_Opyn *OpynFilterer) ParseLiquidate(log types.Log) (*OpynLiquidate, error) {
	event := new(OpynLiquidate)
	if err := _Opyn.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynMintShortIterator is returned from FilterMintShort and is used to iterate over the raw logs and unpacked data for MintShort events raised by the Opyn contract.
type OpynMintShortIterator struct {
	Event *OpynMintShort // Event containing the contract specifics and raw log

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
func (it *OpynMintShortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynMintShort)
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
		it.Event = new(OpynMintShort)
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
func (it *OpynMintShortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynMintShortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynMintShort represents a MintShort event raised by the Opyn contract.
type OpynMintShort struct {
	Sender  common.Address
	Amount  *big.Int
	VaultId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMintShort is a free log retrieval operation binding the contract event 0xb19fa182730a088464dad0e9e0badeb470d0d8d937d854f5caf15c6ad1992c36.
//
// Solidity: event MintShort(address sender, uint256 amount, uint256 vaultId)
func (_Opyn *OpynFilterer) FilterMintShort(opts *bind.FilterOpts) (*OpynMintShortIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "MintShort")
	if err != nil {
		return nil, err
	}
	return &OpynMintShortIterator{contract: _Opyn.contract, event: "MintShort", logs: logs, sub: sub}, nil
}

// WatchMintShort is a free log subscription operation binding the contract event 0xb19fa182730a088464dad0e9e0badeb470d0d8d937d854f5caf15c6ad1992c36.
//
// Solidity: event MintShort(address sender, uint256 amount, uint256 vaultId)
func (_Opyn *OpynFilterer) WatchMintShort(opts *bind.WatchOpts, sink chan<- *OpynMintShort) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "MintShort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynMintShort)
				if err := _Opyn.contract.UnpackLog(event, "MintShort", log); err != nil {
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

// ParseMintShort is a log parse operation binding the contract event 0xb19fa182730a088464dad0e9e0badeb470d0d8d937d854f5caf15c6ad1992c36.
//
// Solidity: event MintShort(address sender, uint256 amount, uint256 vaultId)
func (_Opyn *OpynFilterer) ParseMintShort(log types.Log) (*OpynMintShort, error) {
	event := new(OpynMintShort)
	if err := _Opyn.contract.UnpackLog(event, "MintShort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynNormalizationFactorUpdatedIterator is returned from FilterNormalizationFactorUpdated and is used to iterate over the raw logs and unpacked data for NormalizationFactorUpdated events raised by the Opyn contract.
type OpynNormalizationFactorUpdatedIterator struct {
	Event *OpynNormalizationFactorUpdated // Event containing the contract specifics and raw log

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
func (it *OpynNormalizationFactorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynNormalizationFactorUpdated)
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
		it.Event = new(OpynNormalizationFactorUpdated)
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
func (it *OpynNormalizationFactorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynNormalizationFactorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynNormalizationFactorUpdated represents a NormalizationFactorUpdated event raised by the Opyn contract.
type OpynNormalizationFactorUpdated struct {
	OldNormFactor             *big.Int
	NewNormFactor             *big.Int
	LastModificationTimestamp *big.Int
	Timestamp                 *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterNormalizationFactorUpdated is a free log retrieval operation binding the contract event 0x339e53729b0447795ff69e70a74fed98fc7fef6fe94b7521099b32f0f8de4822.
//
// Solidity: event NormalizationFactorUpdated(uint256 oldNormFactor, uint256 newNormFactor, uint256 lastModificationTimestamp, uint256 timestamp)
func (_Opyn *OpynFilterer) FilterNormalizationFactorUpdated(opts *bind.FilterOpts) (*OpynNormalizationFactorUpdatedIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "NormalizationFactorUpdated")
	if err != nil {
		return nil, err
	}
	return &OpynNormalizationFactorUpdatedIterator{contract: _Opyn.contract, event: "NormalizationFactorUpdated", logs: logs, sub: sub}, nil
}

// WatchNormalizationFactorUpdated is a free log subscription operation binding the contract event 0x339e53729b0447795ff69e70a74fed98fc7fef6fe94b7521099b32f0f8de4822.
//
// Solidity: event NormalizationFactorUpdated(uint256 oldNormFactor, uint256 newNormFactor, uint256 lastModificationTimestamp, uint256 timestamp)
func (_Opyn *OpynFilterer) WatchNormalizationFactorUpdated(opts *bind.WatchOpts, sink chan<- *OpynNormalizationFactorUpdated) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "NormalizationFactorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynNormalizationFactorUpdated)
				if err := _Opyn.contract.UnpackLog(event, "NormalizationFactorUpdated", log); err != nil {
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

// ParseNormalizationFactorUpdated is a log parse operation binding the contract event 0x339e53729b0447795ff69e70a74fed98fc7fef6fe94b7521099b32f0f8de4822.
//
// Solidity: event NormalizationFactorUpdated(uint256 oldNormFactor, uint256 newNormFactor, uint256 lastModificationTimestamp, uint256 timestamp)
func (_Opyn *OpynFilterer) ParseNormalizationFactorUpdated(log types.Log) (*OpynNormalizationFactorUpdated, error) {
	event := new(OpynNormalizationFactorUpdated)
	if err := _Opyn.contract.UnpackLog(event, "NormalizationFactorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynOpenVaultIterator is returned from FilterOpenVault and is used to iterate over the raw logs and unpacked data for OpenVault events raised by the Opyn contract.
type OpynOpenVaultIterator struct {
	Event *OpynOpenVault // Event containing the contract specifics and raw log

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
func (it *OpynOpenVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynOpenVault)
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
		it.Event = new(OpynOpenVault)
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
func (it *OpynOpenVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynOpenVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynOpenVault represents a OpenVault event raised by the Opyn contract.
type OpynOpenVault struct {
	Sender  common.Address
	VaultId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOpenVault is a free log retrieval operation binding the contract event 0x25ff1e0131762176a9084e92880f880f39d6d0e62134607f37e631efe1bad871.
//
// Solidity: event OpenVault(address sender, uint256 vaultId)
func (_Opyn *OpynFilterer) FilterOpenVault(opts *bind.FilterOpts) (*OpynOpenVaultIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "OpenVault")
	if err != nil {
		return nil, err
	}
	return &OpynOpenVaultIterator{contract: _Opyn.contract, event: "OpenVault", logs: logs, sub: sub}, nil
}

// WatchOpenVault is a free log subscription operation binding the contract event 0x25ff1e0131762176a9084e92880f880f39d6d0e62134607f37e631efe1bad871.
//
// Solidity: event OpenVault(address sender, uint256 vaultId)
func (_Opyn *OpynFilterer) WatchOpenVault(opts *bind.WatchOpts, sink chan<- *OpynOpenVault) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "OpenVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynOpenVault)
				if err := _Opyn.contract.UnpackLog(event, "OpenVault", log); err != nil {
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

// ParseOpenVault is a log parse operation binding the contract event 0x25ff1e0131762176a9084e92880f880f39d6d0e62134607f37e631efe1bad871.
//
// Solidity: event OpenVault(address sender, uint256 vaultId)
func (_Opyn *OpynFilterer) ParseOpenVault(log types.Log) (*OpynOpenVault, error) {
	event := new(OpynOpenVault)
	if err := _Opyn.contract.UnpackLog(event, "OpenVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Opyn contract.
type OpynOwnershipTransferredIterator struct {
	Event *OpynOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OpynOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynOwnershipTransferred)
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
		it.Event = new(OpynOwnershipTransferred)
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
func (it *OpynOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynOwnershipTransferred represents a OwnershipTransferred event raised by the Opyn contract.
type OpynOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Opyn *OpynFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OpynOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OpynOwnershipTransferredIterator{contract: _Opyn.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Opyn *OpynFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OpynOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynOwnershipTransferred)
				if err := _Opyn.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Opyn *OpynFilterer) ParseOwnershipTransferred(log types.Log) (*OpynOwnershipTransferred, error) {
	event := new(OpynOwnershipTransferred)
	if err := _Opyn.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Opyn contract.
type OpynPausedIterator struct {
	Event *OpynPaused // Event containing the contract specifics and raw log

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
func (it *OpynPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynPaused)
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
		it.Event = new(OpynPaused)
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
func (it *OpynPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynPaused represents a Paused event raised by the Opyn contract.
type OpynPaused struct {
	PausesLeft *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 pausesLeft)
func (_Opyn *OpynFilterer) FilterPaused(opts *bind.FilterOpts) (*OpynPausedIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OpynPausedIterator{contract: _Opyn.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 pausesLeft)
func (_Opyn *OpynFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OpynPaused) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynPaused)
				if err := _Opyn.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 pausesLeft)
func (_Opyn *OpynFilterer) ParsePaused(log types.Log) (*OpynPaused, error) {
	event := new(OpynPaused)
	if err := _Opyn.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynRedeemLongIterator is returned from FilterRedeemLong and is used to iterate over the raw logs and unpacked data for RedeemLong events raised by the Opyn contract.
type OpynRedeemLongIterator struct {
	Event *OpynRedeemLong // Event containing the contract specifics and raw log

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
func (it *OpynRedeemLongIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynRedeemLong)
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
		it.Event = new(OpynRedeemLong)
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
func (it *OpynRedeemLongIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynRedeemLongIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynRedeemLong represents a RedeemLong event raised by the Opyn contract.
type OpynRedeemLong struct {
	Sender           common.Address
	WPowerPerpAmount *big.Int
	PayoutAmount     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRedeemLong is a free log retrieval operation binding the contract event 0x2131ef4f2f82ca75fe7d2e646ebfa45b6be25e53510c829629c76b641500ec67.
//
// Solidity: event RedeemLong(address sender, uint256 wPowerPerpAmount, uint256 payoutAmount)
func (_Opyn *OpynFilterer) FilterRedeemLong(opts *bind.FilterOpts) (*OpynRedeemLongIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "RedeemLong")
	if err != nil {
		return nil, err
	}
	return &OpynRedeemLongIterator{contract: _Opyn.contract, event: "RedeemLong", logs: logs, sub: sub}, nil
}

// WatchRedeemLong is a free log subscription operation binding the contract event 0x2131ef4f2f82ca75fe7d2e646ebfa45b6be25e53510c829629c76b641500ec67.
//
// Solidity: event RedeemLong(address sender, uint256 wPowerPerpAmount, uint256 payoutAmount)
func (_Opyn *OpynFilterer) WatchRedeemLong(opts *bind.WatchOpts, sink chan<- *OpynRedeemLong) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "RedeemLong")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynRedeemLong)
				if err := _Opyn.contract.UnpackLog(event, "RedeemLong", log); err != nil {
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

// ParseRedeemLong is a log parse operation binding the contract event 0x2131ef4f2f82ca75fe7d2e646ebfa45b6be25e53510c829629c76b641500ec67.
//
// Solidity: event RedeemLong(address sender, uint256 wPowerPerpAmount, uint256 payoutAmount)
func (_Opyn *OpynFilterer) ParseRedeemLong(log types.Log) (*OpynRedeemLong, error) {
	event := new(OpynRedeemLong)
	if err := _Opyn.contract.UnpackLog(event, "RedeemLong", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynRedeemShortIterator is returned from FilterRedeemShort and is used to iterate over the raw logs and unpacked data for RedeemShort events raised by the Opyn contract.
type OpynRedeemShortIterator struct {
	Event *OpynRedeemShort // Event containing the contract specifics and raw log

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
func (it *OpynRedeemShortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynRedeemShort)
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
		it.Event = new(OpynRedeemShort)
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
func (it *OpynRedeemShortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynRedeemShortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynRedeemShort represents a RedeemShort event raised by the Opyn contract.
type OpynRedeemShort struct {
	Sender           common.Address
	VauldId          *big.Int
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRedeemShort is a free log retrieval operation binding the contract event 0x7dff8cdaec6a8d4d1ad32d3c947ed0f0281c3d6456621ef928defae96ec6cddb.
//
// Solidity: event RedeemShort(address sender, uint256 vauldId, uint256 collateralAmount)
func (_Opyn *OpynFilterer) FilterRedeemShort(opts *bind.FilterOpts) (*OpynRedeemShortIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "RedeemShort")
	if err != nil {
		return nil, err
	}
	return &OpynRedeemShortIterator{contract: _Opyn.contract, event: "RedeemShort", logs: logs, sub: sub}, nil
}

// WatchRedeemShort is a free log subscription operation binding the contract event 0x7dff8cdaec6a8d4d1ad32d3c947ed0f0281c3d6456621ef928defae96ec6cddb.
//
// Solidity: event RedeemShort(address sender, uint256 vauldId, uint256 collateralAmount)
func (_Opyn *OpynFilterer) WatchRedeemShort(opts *bind.WatchOpts, sink chan<- *OpynRedeemShort) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "RedeemShort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynRedeemShort)
				if err := _Opyn.contract.UnpackLog(event, "RedeemShort", log); err != nil {
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

// ParseRedeemShort is a log parse operation binding the contract event 0x7dff8cdaec6a8d4d1ad32d3c947ed0f0281c3d6456621ef928defae96ec6cddb.
//
// Solidity: event RedeemShort(address sender, uint256 vauldId, uint256 collateralAmount)
func (_Opyn *OpynFilterer) ParseRedeemShort(log types.Log) (*OpynRedeemShort, error) {
	event := new(OpynRedeemShort)
	if err := _Opyn.contract.UnpackLog(event, "RedeemShort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynReduceDebtIterator is returned from FilterReduceDebt and is used to iterate over the raw logs and unpacked data for ReduceDebt events raised by the Opyn contract.
type OpynReduceDebtIterator struct {
	Event *OpynReduceDebt // Event containing the contract specifics and raw log

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
func (it *OpynReduceDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynReduceDebt)
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
		it.Event = new(OpynReduceDebt)
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
func (it *OpynReduceDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynReduceDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynReduceDebt represents a ReduceDebt event raised by the Opyn contract.
type OpynReduceDebt struct {
	Sender             common.Address
	VaultId            *big.Int
	EthRedeemed        *big.Int
	WPowerPerpRedeemed *big.Int
	WPowerPerpBurned   *big.Int
	WPowerPerpExcess   *big.Int
	Bounty             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReduceDebt is a free log retrieval operation binding the contract event 0xfd0ae2fd36bd955810ae42615bc5ff277c0d0dfcb930f06c9f1777c0ef0752e3.
//
// Solidity: event ReduceDebt(address sender, uint256 vaultId, uint256 ethRedeemed, uint256 wPowerPerpRedeemed, uint256 wPowerPerpBurned, uint256 wPowerPerpExcess, uint256 bounty)
func (_Opyn *OpynFilterer) FilterReduceDebt(opts *bind.FilterOpts) (*OpynReduceDebtIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "ReduceDebt")
	if err != nil {
		return nil, err
	}
	return &OpynReduceDebtIterator{contract: _Opyn.contract, event: "ReduceDebt", logs: logs, sub: sub}, nil
}

// WatchReduceDebt is a free log subscription operation binding the contract event 0xfd0ae2fd36bd955810ae42615bc5ff277c0d0dfcb930f06c9f1777c0ef0752e3.
//
// Solidity: event ReduceDebt(address sender, uint256 vaultId, uint256 ethRedeemed, uint256 wPowerPerpRedeemed, uint256 wPowerPerpBurned, uint256 wPowerPerpExcess, uint256 bounty)
func (_Opyn *OpynFilterer) WatchReduceDebt(opts *bind.WatchOpts, sink chan<- *OpynReduceDebt) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "ReduceDebt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynReduceDebt)
				if err := _Opyn.contract.UnpackLog(event, "ReduceDebt", log); err != nil {
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

// ParseReduceDebt is a log parse operation binding the contract event 0xfd0ae2fd36bd955810ae42615bc5ff277c0d0dfcb930f06c9f1777c0ef0752e3.
//
// Solidity: event ReduceDebt(address sender, uint256 vaultId, uint256 ethRedeemed, uint256 wPowerPerpRedeemed, uint256 wPowerPerpBurned, uint256 wPowerPerpExcess, uint256 bounty)
func (_Opyn *OpynFilterer) ParseReduceDebt(log types.Log) (*OpynReduceDebt, error) {
	event := new(OpynReduceDebt)
	if err := _Opyn.contract.UnpackLog(event, "ReduceDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynShutdownIterator is returned from FilterShutdown and is used to iterate over the raw logs and unpacked data for Shutdown events raised by the Opyn contract.
type OpynShutdownIterator struct {
	Event *OpynShutdown // Event containing the contract specifics and raw log

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
func (it *OpynShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynShutdown)
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
		it.Event = new(OpynShutdown)
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
func (it *OpynShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynShutdown represents a Shutdown event raised by the Opyn contract.
type OpynShutdown struct {
	IndexForSettlement *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterShutdown is a free log retrieval operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 indexForSettlement)
func (_Opyn *OpynFilterer) FilterShutdown(opts *bind.FilterOpts) (*OpynShutdownIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return &OpynShutdownIterator{contract: _Opyn.contract, event: "Shutdown", logs: logs, sub: sub}, nil
}

// WatchShutdown is a free log subscription operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 indexForSettlement)
func (_Opyn *OpynFilterer) WatchShutdown(opts *bind.WatchOpts, sink chan<- *OpynShutdown) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynShutdown)
				if err := _Opyn.contract.UnpackLog(event, "Shutdown", log); err != nil {
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

// ParseShutdown is a log parse operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 indexForSettlement)
func (_Opyn *OpynFilterer) ParseShutdown(log types.Log) (*OpynShutdown, error) {
	event := new(OpynShutdown)
	if err := _Opyn.contract.UnpackLog(event, "Shutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynUnPausedIterator is returned from FilterUnPaused and is used to iterate over the raw logs and unpacked data for UnPaused events raised by the Opyn contract.
type OpynUnPausedIterator struct {
	Event *OpynUnPaused // Event containing the contract specifics and raw log

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
func (it *OpynUnPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynUnPaused)
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
		it.Event = new(OpynUnPaused)
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
func (it *OpynUnPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynUnPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynUnPaused represents a UnPaused event raised by the Opyn contract.
type OpynUnPaused struct {
	Unpauser common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnPaused is a free log retrieval operation binding the contract event 0xff2b959f2bcdb44c7ecb4b16dae055431019d7350607125cfc2b74a06632c90e.
//
// Solidity: event UnPaused(address unpauser)
func (_Opyn *OpynFilterer) FilterUnPaused(opts *bind.FilterOpts) (*OpynUnPausedIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "UnPaused")
	if err != nil {
		return nil, err
	}
	return &OpynUnPausedIterator{contract: _Opyn.contract, event: "UnPaused", logs: logs, sub: sub}, nil
}

// WatchUnPaused is a free log subscription operation binding the contract event 0xff2b959f2bcdb44c7ecb4b16dae055431019d7350607125cfc2b74a06632c90e.
//
// Solidity: event UnPaused(address unpauser)
func (_Opyn *OpynFilterer) WatchUnPaused(opts *bind.WatchOpts, sink chan<- *OpynUnPaused) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "UnPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynUnPaused)
				if err := _Opyn.contract.UnpackLog(event, "UnPaused", log); err != nil {
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

// ParseUnPaused is a log parse operation binding the contract event 0xff2b959f2bcdb44c7ecb4b16dae055431019d7350607125cfc2b74a06632c90e.
//
// Solidity: event UnPaused(address unpauser)
func (_Opyn *OpynFilterer) ParseUnPaused(log types.Log) (*OpynUnPaused, error) {
	event := new(OpynUnPaused)
	if err := _Opyn.contract.UnpackLog(event, "UnPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynUpdateOperatorIterator is returned from FilterUpdateOperator and is used to iterate over the raw logs and unpacked data for UpdateOperator events raised by the Opyn contract.
type OpynUpdateOperatorIterator struct {
	Event *OpynUpdateOperator // Event containing the contract specifics and raw log

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
func (it *OpynUpdateOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynUpdateOperator)
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
		it.Event = new(OpynUpdateOperator)
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
func (it *OpynUpdateOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynUpdateOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynUpdateOperator represents a UpdateOperator event raised by the Opyn contract.
type OpynUpdateOperator struct {
	Sender   common.Address
	VaultId  *big.Int
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateOperator is a free log retrieval operation binding the contract event 0x3137fc9cd2e33c34f86e29c24d81f3c75b0bce639d3c4ed0d31eeff1160a7ff5.
//
// Solidity: event UpdateOperator(address sender, uint256 vaultId, address operator)
func (_Opyn *OpynFilterer) FilterUpdateOperator(opts *bind.FilterOpts) (*OpynUpdateOperatorIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "UpdateOperator")
	if err != nil {
		return nil, err
	}
	return &OpynUpdateOperatorIterator{contract: _Opyn.contract, event: "UpdateOperator", logs: logs, sub: sub}, nil
}

// WatchUpdateOperator is a free log subscription operation binding the contract event 0x3137fc9cd2e33c34f86e29c24d81f3c75b0bce639d3c4ed0d31eeff1160a7ff5.
//
// Solidity: event UpdateOperator(address sender, uint256 vaultId, address operator)
func (_Opyn *OpynFilterer) WatchUpdateOperator(opts *bind.WatchOpts, sink chan<- *OpynUpdateOperator) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "UpdateOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynUpdateOperator)
				if err := _Opyn.contract.UnpackLog(event, "UpdateOperator", log); err != nil {
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

// ParseUpdateOperator is a log parse operation binding the contract event 0x3137fc9cd2e33c34f86e29c24d81f3c75b0bce639d3c4ed0d31eeff1160a7ff5.
//
// Solidity: event UpdateOperator(address sender, uint256 vaultId, address operator)
func (_Opyn *OpynFilterer) ParseUpdateOperator(log types.Log) (*OpynUpdateOperator, error) {
	event := new(OpynUpdateOperator)
	if err := _Opyn.contract.UnpackLog(event, "UpdateOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynWithdrawCollateralIterator is returned from FilterWithdrawCollateral and is used to iterate over the raw logs and unpacked data for WithdrawCollateral events raised by the Opyn contract.
type OpynWithdrawCollateralIterator struct {
	Event *OpynWithdrawCollateral // Event containing the contract specifics and raw log

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
func (it *OpynWithdrawCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynWithdrawCollateral)
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
		it.Event = new(OpynWithdrawCollateral)
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
func (it *OpynWithdrawCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynWithdrawCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynWithdrawCollateral represents a WithdrawCollateral event raised by the Opyn contract.
type OpynWithdrawCollateral struct {
	Sender  common.Address
	VaultId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCollateral is a free log retrieval operation binding the contract event 0x627a692d5a03ab34732c0d2aa319f3ecdebdc4528f383eabcb25441dc0a70cfb.
//
// Solidity: event WithdrawCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Opyn *OpynFilterer) FilterWithdrawCollateral(opts *bind.FilterOpts) (*OpynWithdrawCollateralIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "WithdrawCollateral")
	if err != nil {
		return nil, err
	}
	return &OpynWithdrawCollateralIterator{contract: _Opyn.contract, event: "WithdrawCollateral", logs: logs, sub: sub}, nil
}

// WatchWithdrawCollateral is a free log subscription operation binding the contract event 0x627a692d5a03ab34732c0d2aa319f3ecdebdc4528f383eabcb25441dc0a70cfb.
//
// Solidity: event WithdrawCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Opyn *OpynFilterer) WatchWithdrawCollateral(opts *bind.WatchOpts, sink chan<- *OpynWithdrawCollateral) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "WithdrawCollateral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynWithdrawCollateral)
				if err := _Opyn.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
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

// ParseWithdrawCollateral is a log parse operation binding the contract event 0x627a692d5a03ab34732c0d2aa319f3ecdebdc4528f383eabcb25441dc0a70cfb.
//
// Solidity: event WithdrawCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Opyn *OpynFilterer) ParseWithdrawCollateral(log types.Log) (*OpynWithdrawCollateral, error) {
	event := new(OpynWithdrawCollateral)
	if err := _Opyn.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpynWithdrawUniPositionTokenIterator is returned from FilterWithdrawUniPositionToken and is used to iterate over the raw logs and unpacked data for WithdrawUniPositionToken events raised by the Opyn contract.
type OpynWithdrawUniPositionTokenIterator struct {
	Event *OpynWithdrawUniPositionToken // Event containing the contract specifics and raw log

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
func (it *OpynWithdrawUniPositionTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpynWithdrawUniPositionToken)
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
		it.Event = new(OpynWithdrawUniPositionToken)
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
func (it *OpynWithdrawUniPositionTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpynWithdrawUniPositionTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpynWithdrawUniPositionToken represents a WithdrawUniPositionToken event raised by the Opyn contract.
type OpynWithdrawUniPositionToken struct {
	Sender  common.Address
	VaultId *big.Int
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawUniPositionToken is a free log retrieval operation binding the contract event 0xe59f38fa1264fc25c9f0185eee136eaf810d90b8e7293b342e4037c68720177a.
//
// Solidity: event WithdrawUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Opyn *OpynFilterer) FilterWithdrawUniPositionToken(opts *bind.FilterOpts) (*OpynWithdrawUniPositionTokenIterator, error) {

	logs, sub, err := _Opyn.contract.FilterLogs(opts, "WithdrawUniPositionToken")
	if err != nil {
		return nil, err
	}
	return &OpynWithdrawUniPositionTokenIterator{contract: _Opyn.contract, event: "WithdrawUniPositionToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawUniPositionToken is a free log subscription operation binding the contract event 0xe59f38fa1264fc25c9f0185eee136eaf810d90b8e7293b342e4037c68720177a.
//
// Solidity: event WithdrawUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Opyn *OpynFilterer) WatchWithdrawUniPositionToken(opts *bind.WatchOpts, sink chan<- *OpynWithdrawUniPositionToken) (event.Subscription, error) {

	logs, sub, err := _Opyn.contract.WatchLogs(opts, "WithdrawUniPositionToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpynWithdrawUniPositionToken)
				if err := _Opyn.contract.UnpackLog(event, "WithdrawUniPositionToken", log); err != nil {
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

// ParseWithdrawUniPositionToken is a log parse operation binding the contract event 0xe59f38fa1264fc25c9f0185eee136eaf810d90b8e7293b342e4037c68720177a.
//
// Solidity: event WithdrawUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Opyn *OpynFilterer) ParseWithdrawUniPositionToken(log types.Log) (*OpynWithdrawUniPositionToken, error) {
	event := new(OpynWithdrawUniPositionToken)
	if err := _Opyn.contract.UnpackLog(event, "WithdrawUniPositionToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
