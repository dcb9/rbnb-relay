// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake_pool

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
)

// StakePoolMetaData contains all meta data concerning the StakePool contract.
var StakePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"checkAndClaimReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkAndClaimUndelegated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validatorList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountList\",\"type\":\"uint256[]\"}],\"name\":\"delegateVals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getDelegated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getPendingUndelegateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelayerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestInFly\",\"outputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"\",\"type\":\"uint256[3]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalDelegated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeMangerAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorSrc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorDst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validatorList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountList\",\"type\":\"uint256[]\"}],\"name\":\"undelegateVals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawForStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StakePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use StakePoolMetaData.ABI instead.
var StakePoolABI = StakePoolMetaData.ABI

// StakePool is an auto generated Go binding around an Ethereum contract.
type StakePool struct {
	StakePoolCaller     // Read-only binding to the contract
	StakePoolTransactor // Write-only binding to the contract
	StakePoolFilterer   // Log filterer for contract events
}

// StakePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakePoolSession struct {
	Contract     *StakePool        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakePoolCallerSession struct {
	Contract *StakePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StakePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakePoolTransactorSession struct {
	Contract     *StakePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StakePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakePoolRaw struct {
	Contract *StakePool // Generic contract binding to access the raw methods on
}

// StakePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakePoolCallerRaw struct {
	Contract *StakePoolCaller // Generic read-only contract binding to access the raw methods on
}

// StakePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakePoolTransactorRaw struct {
	Contract *StakePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakePool creates a new instance of StakePool, bound to a specific deployed contract.
func NewStakePool(address common.Address, backend bind.ContractBackend) (*StakePool, error) {
	contract, err := bindStakePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakePool{StakePoolCaller: StakePoolCaller{contract: contract}, StakePoolTransactor: StakePoolTransactor{contract: contract}, StakePoolFilterer: StakePoolFilterer{contract: contract}}, nil
}

// NewStakePoolCaller creates a new read-only instance of StakePool, bound to a specific deployed contract.
func NewStakePoolCaller(address common.Address, caller bind.ContractCaller) (*StakePoolCaller, error) {
	contract, err := bindStakePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakePoolCaller{contract: contract}, nil
}

// NewStakePoolTransactor creates a new write-only instance of StakePool, bound to a specific deployed contract.
func NewStakePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*StakePoolTransactor, error) {
	contract, err := bindStakePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakePoolTransactor{contract: contract}, nil
}

// NewStakePoolFilterer creates a new log filterer instance of StakePool, bound to a specific deployed contract.
func NewStakePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*StakePoolFilterer, error) {
	contract, err := bindStakePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakePoolFilterer{contract: contract}, nil
}

// bindStakePool binds a generic wrapper to an already deployed contract.
func bindStakePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakePool *StakePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakePool.Contract.StakePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakePool *StakePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePool.Contract.StakePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakePool *StakePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakePool.Contract.StakePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakePool *StakePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakePool *StakePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakePool *StakePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakePool.Contract.contract.Transact(opts, method, params...)
}

// GetDelegated is a free data retrieval call binding the contract method 0x56f0b57e.
//
// Solidity: function getDelegated(address validator) view returns(uint256)
func (_StakePool *StakePoolCaller) GetDelegated(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "getDelegated", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegated is a free data retrieval call binding the contract method 0x56f0b57e.
//
// Solidity: function getDelegated(address validator) view returns(uint256)
func (_StakePool *StakePoolSession) GetDelegated(validator common.Address) (*big.Int, error) {
	return _StakePool.Contract.GetDelegated(&_StakePool.CallOpts, validator)
}

// GetDelegated is a free data retrieval call binding the contract method 0x56f0b57e.
//
// Solidity: function getDelegated(address validator) view returns(uint256)
func (_StakePool *StakePoolCallerSession) GetDelegated(validator common.Address) (*big.Int, error) {
	return _StakePool.Contract.GetDelegated(&_StakePool.CallOpts, validator)
}

// GetMinDelegation is a free data retrieval call binding the contract method 0x69b635b6.
//
// Solidity: function getMinDelegation() view returns(uint256)
func (_StakePool *StakePoolCaller) GetMinDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "getMinDelegation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinDelegation is a free data retrieval call binding the contract method 0x69b635b6.
//
// Solidity: function getMinDelegation() view returns(uint256)
func (_StakePool *StakePoolSession) GetMinDelegation() (*big.Int, error) {
	return _StakePool.Contract.GetMinDelegation(&_StakePool.CallOpts)
}

// GetMinDelegation is a free data retrieval call binding the contract method 0x69b635b6.
//
// Solidity: function getMinDelegation() view returns(uint256)
func (_StakePool *StakePoolCallerSession) GetMinDelegation() (*big.Int, error) {
	return _StakePool.Contract.GetMinDelegation(&_StakePool.CallOpts)
}

// GetPendingUndelegateTime is a free data retrieval call binding the contract method 0xfec0e154.
//
// Solidity: function getPendingUndelegateTime(address validator) view returns(uint256)
func (_StakePool *StakePoolCaller) GetPendingUndelegateTime(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "getPendingUndelegateTime", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingUndelegateTime is a free data retrieval call binding the contract method 0xfec0e154.
//
// Solidity: function getPendingUndelegateTime(address validator) view returns(uint256)
func (_StakePool *StakePoolSession) GetPendingUndelegateTime(validator common.Address) (*big.Int, error) {
	return _StakePool.Contract.GetPendingUndelegateTime(&_StakePool.CallOpts, validator)
}

// GetPendingUndelegateTime is a free data retrieval call binding the contract method 0xfec0e154.
//
// Solidity: function getPendingUndelegateTime(address validator) view returns(uint256)
func (_StakePool *StakePoolCallerSession) GetPendingUndelegateTime(validator common.Address) (*big.Int, error) {
	return _StakePool.Contract.GetPendingUndelegateTime(&_StakePool.CallOpts, validator)
}

// GetRelayerFee is a free data retrieval call binding the contract method 0xc2117d82.
//
// Solidity: function getRelayerFee() view returns(uint256)
func (_StakePool *StakePoolCaller) GetRelayerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "getRelayerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRelayerFee is a free data retrieval call binding the contract method 0xc2117d82.
//
// Solidity: function getRelayerFee() view returns(uint256)
func (_StakePool *StakePoolSession) GetRelayerFee() (*big.Int, error) {
	return _StakePool.Contract.GetRelayerFee(&_StakePool.CallOpts)
}

// GetRelayerFee is a free data retrieval call binding the contract method 0xc2117d82.
//
// Solidity: function getRelayerFee() view returns(uint256)
func (_StakePool *StakePoolCallerSession) GetRelayerFee() (*big.Int, error) {
	return _StakePool.Contract.GetRelayerFee(&_StakePool.CallOpts)
}

// GetRequestInFly is a free data retrieval call binding the contract method 0x3048ecfa.
//
// Solidity: function getRequestInFly() view returns(uint256[3])
func (_StakePool *StakePoolCaller) GetRequestInFly(opts *bind.CallOpts) ([3]*big.Int, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "getRequestInFly")

	if err != nil {
		return *new([3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([3]*big.Int)).(*[3]*big.Int)

	return out0, err

}

// GetRequestInFly is a free data retrieval call binding the contract method 0x3048ecfa.
//
// Solidity: function getRequestInFly() view returns(uint256[3])
func (_StakePool *StakePoolSession) GetRequestInFly() ([3]*big.Int, error) {
	return _StakePool.Contract.GetRequestInFly(&_StakePool.CallOpts)
}

// GetRequestInFly is a free data retrieval call binding the contract method 0x3048ecfa.
//
// Solidity: function getRequestInFly() view returns(uint256[3])
func (_StakePool *StakePoolCallerSession) GetRequestInFly() ([3]*big.Int, error) {
	return _StakePool.Contract.GetRequestInFly(&_StakePool.CallOpts)
}

// GetTotalDelegated is a free data retrieval call binding the contract method 0xf2976871.
//
// Solidity: function getTotalDelegated() view returns(uint256)
func (_StakePool *StakePoolCaller) GetTotalDelegated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "getTotalDelegated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDelegated is a free data retrieval call binding the contract method 0xf2976871.
//
// Solidity: function getTotalDelegated() view returns(uint256)
func (_StakePool *StakePoolSession) GetTotalDelegated() (*big.Int, error) {
	return _StakePool.Contract.GetTotalDelegated(&_StakePool.CallOpts)
}

// GetTotalDelegated is a free data retrieval call binding the contract method 0xf2976871.
//
// Solidity: function getTotalDelegated() view returns(uint256)
func (_StakePool *StakePoolCallerSession) GetTotalDelegated() (*big.Int, error) {
	return _StakePool.Contract.GetTotalDelegated(&_StakePool.CallOpts)
}

// StakeManagerAddress is a free data retrieval call binding the contract method 0x15f69471.
//
// Solidity: function stakeManagerAddress() view returns(address impl)
func (_StakePool *StakePoolCaller) StakeManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "stakeManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeManagerAddress is a free data retrieval call binding the contract method 0x15f69471.
//
// Solidity: function stakeManagerAddress() view returns(address impl)
func (_StakePool *StakePoolSession) StakeManagerAddress() (common.Address, error) {
	return _StakePool.Contract.StakeManagerAddress(&_StakePool.CallOpts)
}

// StakeManagerAddress is a free data retrieval call binding the contract method 0x15f69471.
//
// Solidity: function stakeManagerAddress() view returns(address impl)
func (_StakePool *StakePoolCallerSession) StakeManagerAddress() (common.Address, error) {
	return _StakePool.Contract.StakeManagerAddress(&_StakePool.CallOpts)
}

// StakingAddress is a free data retrieval call binding the contract method 0xd7b4be24.
//
// Solidity: function stakingAddress() view returns(address impl)
func (_StakePool *StakePoolCaller) StakingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "stakingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingAddress is a free data retrieval call binding the contract method 0xd7b4be24.
//
// Solidity: function stakingAddress() view returns(address impl)
func (_StakePool *StakePoolSession) StakingAddress() (common.Address, error) {
	return _StakePool.Contract.StakingAddress(&_StakePool.CallOpts)
}

// StakingAddress is a free data retrieval call binding the contract method 0xd7b4be24.
//
// Solidity: function stakingAddress() view returns(address impl)
func (_StakePool *StakePoolCallerSession) StakingAddress() (common.Address, error) {
	return _StakePool.Contract.StakingAddress(&_StakePool.CallOpts)
}

// CheckAndClaimReward is a paid mutator transaction binding the contract method 0x28c11e1e.
//
// Solidity: function checkAndClaimReward() returns(uint256)
func (_StakePool *StakePoolTransactor) CheckAndClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "checkAndClaimReward")
}

// CheckAndClaimReward is a paid mutator transaction binding the contract method 0x28c11e1e.
//
// Solidity: function checkAndClaimReward() returns(uint256)
func (_StakePool *StakePoolSession) CheckAndClaimReward() (*types.Transaction, error) {
	return _StakePool.Contract.CheckAndClaimReward(&_StakePool.TransactOpts)
}

// CheckAndClaimReward is a paid mutator transaction binding the contract method 0x28c11e1e.
//
// Solidity: function checkAndClaimReward() returns(uint256)
func (_StakePool *StakePoolTransactorSession) CheckAndClaimReward() (*types.Transaction, error) {
	return _StakePool.Contract.CheckAndClaimReward(&_StakePool.TransactOpts)
}

// CheckAndClaimUndelegated is a paid mutator transaction binding the contract method 0x85770481.
//
// Solidity: function checkAndClaimUndelegated() returns(uint256)
func (_StakePool *StakePoolTransactor) CheckAndClaimUndelegated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "checkAndClaimUndelegated")
}

// CheckAndClaimUndelegated is a paid mutator transaction binding the contract method 0x85770481.
//
// Solidity: function checkAndClaimUndelegated() returns(uint256)
func (_StakePool *StakePoolSession) CheckAndClaimUndelegated() (*types.Transaction, error) {
	return _StakePool.Contract.CheckAndClaimUndelegated(&_StakePool.TransactOpts)
}

// CheckAndClaimUndelegated is a paid mutator transaction binding the contract method 0x85770481.
//
// Solidity: function checkAndClaimUndelegated() returns(uint256)
func (_StakePool *StakePoolTransactorSession) CheckAndClaimUndelegated() (*types.Transaction, error) {
	return _StakePool.Contract.CheckAndClaimUndelegated(&_StakePool.TransactOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address validator, uint256 amount) returns()
func (_StakePool *StakePoolTransactor) Delegate(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "delegate", validator, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address validator, uint256 amount) returns()
func (_StakePool *StakePoolSession) Delegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.Delegate(&_StakePool.TransactOpts, validator, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address validator, uint256 amount) returns()
func (_StakePool *StakePoolTransactorSession) Delegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.Delegate(&_StakePool.TransactOpts, validator, amount)
}

// DelegateVals is a paid mutator transaction binding the contract method 0x4e6c348e.
//
// Solidity: function delegateVals(address[] validatorList, uint256[] amountList) returns()
func (_StakePool *StakePoolTransactor) DelegateVals(opts *bind.TransactOpts, validatorList []common.Address, amountList []*big.Int) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "delegateVals", validatorList, amountList)
}

// DelegateVals is a paid mutator transaction binding the contract method 0x4e6c348e.
//
// Solidity: function delegateVals(address[] validatorList, uint256[] amountList) returns()
func (_StakePool *StakePoolSession) DelegateVals(validatorList []common.Address, amountList []*big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.DelegateVals(&_StakePool.TransactOpts, validatorList, amountList)
}

// DelegateVals is a paid mutator transaction binding the contract method 0x4e6c348e.
//
// Solidity: function delegateVals(address[] validatorList, uint256[] amountList) returns()
func (_StakePool *StakePoolTransactorSession) DelegateVals(validatorList []common.Address, amountList []*big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.DelegateVals(&_StakePool.TransactOpts, validatorList, amountList)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _stakingAddress, address _stakeMangerAddress) returns()
func (_StakePool *StakePoolTransactor) Init(opts *bind.TransactOpts, _stakingAddress common.Address, _stakeMangerAddress common.Address) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "init", _stakingAddress, _stakeMangerAddress)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _stakingAddress, address _stakeMangerAddress) returns()
func (_StakePool *StakePoolSession) Init(_stakingAddress common.Address, _stakeMangerAddress common.Address) (*types.Transaction, error) {
	return _StakePool.Contract.Init(&_StakePool.TransactOpts, _stakingAddress, _stakeMangerAddress)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _stakingAddress, address _stakeMangerAddress) returns()
func (_StakePool *StakePoolTransactorSession) Init(_stakingAddress common.Address, _stakeMangerAddress common.Address) (*types.Transaction, error) {
	return _StakePool.Contract.Init(&_StakePool.TransactOpts, _stakingAddress, _stakeMangerAddress)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address validatorSrc, address validatorDst, uint256 amount) returns()
func (_StakePool *StakePoolTransactor) Redelegate(opts *bind.TransactOpts, validatorSrc common.Address, validatorDst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "redelegate", validatorSrc, validatorDst, amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address validatorSrc, address validatorDst, uint256 amount) returns()
func (_StakePool *StakePoolSession) Redelegate(validatorSrc common.Address, validatorDst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.Redelegate(&_StakePool.TransactOpts, validatorSrc, validatorDst, amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address validatorSrc, address validatorDst, uint256 amount) returns()
func (_StakePool *StakePoolTransactorSession) Redelegate(validatorSrc common.Address, validatorDst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.Redelegate(&_StakePool.TransactOpts, validatorSrc, validatorDst, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_StakePool *StakePoolTransactor) Undelegate(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "undelegate", validator, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_StakePool *StakePoolSession) Undelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.Undelegate(&_StakePool.TransactOpts, validator, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_StakePool *StakePoolTransactorSession) Undelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.Undelegate(&_StakePool.TransactOpts, validator, amount)
}

// UndelegateVals is a paid mutator transaction binding the contract method 0x3dd9fe5c.
//
// Solidity: function undelegateVals(address[] validatorList, uint256[] amountList) returns()
func (_StakePool *StakePoolTransactor) UndelegateVals(opts *bind.TransactOpts, validatorList []common.Address, amountList []*big.Int) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "undelegateVals", validatorList, amountList)
}

// UndelegateVals is a paid mutator transaction binding the contract method 0x3dd9fe5c.
//
// Solidity: function undelegateVals(address[] validatorList, uint256[] amountList) returns()
func (_StakePool *StakePoolSession) UndelegateVals(validatorList []common.Address, amountList []*big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.UndelegateVals(&_StakePool.TransactOpts, validatorList, amountList)
}

// UndelegateVals is a paid mutator transaction binding the contract method 0x3dd9fe5c.
//
// Solidity: function undelegateVals(address[] validatorList, uint256[] amountList) returns()
func (_StakePool *StakePoolTransactorSession) UndelegateVals(validatorList []common.Address, amountList []*big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.UndelegateVals(&_StakePool.TransactOpts, validatorList, amountList)
}

// WithdrawForStaker is a paid mutator transaction binding the contract method 0x71aa38b5.
//
// Solidity: function withdrawForStaker(address staker, uint256 amount) returns()
func (_StakePool *StakePoolTransactor) WithdrawForStaker(opts *bind.TransactOpts, staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "withdrawForStaker", staker, amount)
}

// WithdrawForStaker is a paid mutator transaction binding the contract method 0x71aa38b5.
//
// Solidity: function withdrawForStaker(address staker, uint256 amount) returns()
func (_StakePool *StakePoolSession) WithdrawForStaker(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.WithdrawForStaker(&_StakePool.TransactOpts, staker, amount)
}

// WithdrawForStaker is a paid mutator transaction binding the contract method 0x71aa38b5.
//
// Solidity: function withdrawForStaker(address staker, uint256 amount) returns()
func (_StakePool *StakePoolTransactorSession) WithdrawForStaker(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.WithdrawForStaker(&_StakePool.TransactOpts, staker, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakePool *StakePoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakePool *StakePoolSession) Receive() (*types.Transaction, error) {
	return _StakePool.Contract.Receive(&_StakePool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakePool *StakePoolTransactorSession) Receive() (*types.Transaction, error) {
	return _StakePool.Contract.Receive(&_StakePool.TransactOpts)
}
