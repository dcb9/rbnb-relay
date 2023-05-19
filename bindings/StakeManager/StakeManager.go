// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake_manager

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

// StakeManagerMetaData contains all meta data concerning the StakeManager contract.
var StakeManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"era\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ExecuteNewEra\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalId\",\"type\":\"bytes32\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"govDelegated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"localDelegated\",\"type\":\"uint256\"}],\"name\":\"RepairDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"era\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"Settle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakePool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rTokenAmount\",\"type\":\"uint256\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakePool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeIndex\",\"type\":\"uint256\"}],\"name\":\"Unstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakePool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"unstakeIndexList\",\"type\":\"uint256[]\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakePool\",\"type\":\"address\"}],\"name\":\"addStakePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"addVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEra\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegatedDiffLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegatedOfValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eraOffset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"eraRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eraSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBondedPools\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"pools\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakeRelayerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getUnstakeIndexListOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"unstakeIndexList\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnstakeRelayerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"getValidatorsOf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proposalId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_initialVoters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_initialThreshold\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unbondingDuration\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_initialThreshold\",\"type\":\"uint256\"}],\"name\":\"initMultisig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestEra\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"latestRewardTimestampOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakePoolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_govDelegated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unbond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pendingDelegate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalRTokenSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_era\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_latestRewardtimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_undistributedReward\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_era\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_poolAddressList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_newRewardList\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_latestRewardTimestampList\",\"type\":\"uint256[]\"}],\"name\":\"newEra\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextUnstakeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingDelegateOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingUndelegateOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolInfoOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"era\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unbond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"active\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"enumMultisig.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"_yesVotes\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"_yesVotesTotal\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rateChangeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_srcValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dstValidator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"removeVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakePool\",\"type\":\"address\"}],\"name\":\"rmStakePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"rmValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unstakeFeeCommission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_protocolFeeCommission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unbondingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rateChangeLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_eraSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_eraOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transferGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_delegatedDiffLimit\",\"type\":\"uint256\"}],\"name\":\"setParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakePoolAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"}],\"name\":\"stakeWithPool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalProtocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRTokenSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unbondingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"undistributedRewardOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rTokenAmount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unstakeAtIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"era\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeFeeCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakePoolAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rTokenAmount\",\"type\":\"uint256\"}],\"name\":\"unstakeWithPool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawRelayerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"withdrawWithPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeManagerMetaData.ABI instead.
var StakeManagerABI = StakeManagerMetaData.ABI

// StakeManager is an auto generated Go binding around an Ethereum contract.
type StakeManager struct {
	StakeManagerCaller     // Read-only binding to the contract
	StakeManagerTransactor // Write-only binding to the contract
	StakeManagerFilterer   // Log filterer for contract events
}

// StakeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeManagerSession struct {
	Contract     *StakeManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeManagerCallerSession struct {
	Contract *StakeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StakeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeManagerTransactorSession struct {
	Contract     *StakeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeManagerRaw struct {
	Contract *StakeManager // Generic contract binding to access the raw methods on
}

// StakeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeManagerCallerRaw struct {
	Contract *StakeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeManagerTransactorRaw struct {
	Contract *StakeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeManager creates a new instance of StakeManager, bound to a specific deployed contract.
func NewStakeManager(address common.Address, backend bind.ContractBackend) (*StakeManager, error) {
	contract, err := bindStakeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeManager{StakeManagerCaller: StakeManagerCaller{contract: contract}, StakeManagerTransactor: StakeManagerTransactor{contract: contract}, StakeManagerFilterer: StakeManagerFilterer{contract: contract}}, nil
}

// NewStakeManagerCaller creates a new read-only instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerCaller(address common.Address, caller bind.ContractCaller) (*StakeManagerCaller, error) {
	contract, err := bindStakeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeManagerCaller{contract: contract}, nil
}

// NewStakeManagerTransactor creates a new write-only instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeManagerTransactor, error) {
	contract, err := bindStakeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeManagerTransactor{contract: contract}, nil
}

// NewStakeManagerFilterer creates a new log filterer instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeManagerFilterer, error) {
	contract, err := bindStakeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeManagerFilterer{contract: contract}, nil
}

// bindStakeManager binds a generic wrapper to an already deployed contract.
func bindStakeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeManager *StakeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeManager.Contract.StakeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeManager *StakeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeManager *StakeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeManager *StakeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeManager *StakeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeManager *StakeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeManager.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_StakeManager *StakeManagerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_StakeManager *StakeManagerSession) Admin() (common.Address, error) {
	return _StakeManager.Contract.Admin(&_StakeManager.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_StakeManager *StakeManagerCallerSession) Admin() (common.Address, error) {
	return _StakeManager.Contract.Admin(&_StakeManager.CallOpts)
}

// CurrentEra is a free data retrieval call binding the contract method 0x973628f6.
//
// Solidity: function currentEra() view returns(uint256)
func (_StakeManager *StakeManagerCaller) CurrentEra(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "currentEra")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEra is a free data retrieval call binding the contract method 0x973628f6.
//
// Solidity: function currentEra() view returns(uint256)
func (_StakeManager *StakeManagerSession) CurrentEra() (*big.Int, error) {
	return _StakeManager.Contract.CurrentEra(&_StakeManager.CallOpts)
}

// CurrentEra is a free data retrieval call binding the contract method 0x973628f6.
//
// Solidity: function currentEra() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) CurrentEra() (*big.Int, error) {
	return _StakeManager.Contract.CurrentEra(&_StakeManager.CallOpts)
}

// DelegatedDiffLimit is a free data retrieval call binding the contract method 0x24ccfeec.
//
// Solidity: function delegatedDiffLimit() view returns(uint256)
func (_StakeManager *StakeManagerCaller) DelegatedDiffLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "delegatedDiffLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedDiffLimit is a free data retrieval call binding the contract method 0x24ccfeec.
//
// Solidity: function delegatedDiffLimit() view returns(uint256)
func (_StakeManager *StakeManagerSession) DelegatedDiffLimit() (*big.Int, error) {
	return _StakeManager.Contract.DelegatedDiffLimit(&_StakeManager.CallOpts)
}

// DelegatedDiffLimit is a free data retrieval call binding the contract method 0x24ccfeec.
//
// Solidity: function delegatedDiffLimit() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) DelegatedDiffLimit() (*big.Int, error) {
	return _StakeManager.Contract.DelegatedDiffLimit(&_StakeManager.CallOpts)
}

// DelegatedOfValidator is a free data retrieval call binding the contract method 0x33160634.
//
// Solidity: function delegatedOfValidator(address , address ) view returns(uint256)
func (_StakeManager *StakeManagerCaller) DelegatedOfValidator(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "delegatedOfValidator", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedOfValidator is a free data retrieval call binding the contract method 0x33160634.
//
// Solidity: function delegatedOfValidator(address , address ) view returns(uint256)
func (_StakeManager *StakeManagerSession) DelegatedOfValidator(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.DelegatedOfValidator(&_StakeManager.CallOpts, arg0, arg1)
}

// DelegatedOfValidator is a free data retrieval call binding the contract method 0x33160634.
//
// Solidity: function delegatedOfValidator(address , address ) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) DelegatedOfValidator(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.DelegatedOfValidator(&_StakeManager.CallOpts, arg0, arg1)
}

// EraOffset is a free data retrieval call binding the contract method 0xc8c20263.
//
// Solidity: function eraOffset() view returns(uint256)
func (_StakeManager *StakeManagerCaller) EraOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "eraOffset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EraOffset is a free data retrieval call binding the contract method 0xc8c20263.
//
// Solidity: function eraOffset() view returns(uint256)
func (_StakeManager *StakeManagerSession) EraOffset() (*big.Int, error) {
	return _StakeManager.Contract.EraOffset(&_StakeManager.CallOpts)
}

// EraOffset is a free data retrieval call binding the contract method 0xc8c20263.
//
// Solidity: function eraOffset() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) EraOffset() (*big.Int, error) {
	return _StakeManager.Contract.EraOffset(&_StakeManager.CallOpts)
}

// EraRate is a free data retrieval call binding the contract method 0xeb8ad76e.
//
// Solidity: function eraRate(uint256 ) view returns(uint256)
func (_StakeManager *StakeManagerCaller) EraRate(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "eraRate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EraRate is a free data retrieval call binding the contract method 0xeb8ad76e.
//
// Solidity: function eraRate(uint256 ) view returns(uint256)
func (_StakeManager *StakeManagerSession) EraRate(arg0 *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.EraRate(&_StakeManager.CallOpts, arg0)
}

// EraRate is a free data retrieval call binding the contract method 0xeb8ad76e.
//
// Solidity: function eraRate(uint256 ) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) EraRate(arg0 *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.EraRate(&_StakeManager.CallOpts, arg0)
}

// EraSeconds is a free data retrieval call binding the contract method 0xe81f1553.
//
// Solidity: function eraSeconds() view returns(uint256)
func (_StakeManager *StakeManagerCaller) EraSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "eraSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EraSeconds is a free data retrieval call binding the contract method 0xe81f1553.
//
// Solidity: function eraSeconds() view returns(uint256)
func (_StakeManager *StakeManagerSession) EraSeconds() (*big.Int, error) {
	return _StakeManager.Contract.EraSeconds(&_StakeManager.CallOpts)
}

// EraSeconds is a free data retrieval call binding the contract method 0xe81f1553.
//
// Solidity: function eraSeconds() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) EraSeconds() (*big.Int, error) {
	return _StakeManager.Contract.EraSeconds(&_StakeManager.CallOpts)
}

// GetBondedPools is a free data retrieval call binding the contract method 0x58af8bf0.
//
// Solidity: function getBondedPools() view returns(address[] pools)
func (_StakeManager *StakeManagerCaller) GetBondedPools(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getBondedPools")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBondedPools is a free data retrieval call binding the contract method 0x58af8bf0.
//
// Solidity: function getBondedPools() view returns(address[] pools)
func (_StakeManager *StakeManagerSession) GetBondedPools() ([]common.Address, error) {
	return _StakeManager.Contract.GetBondedPools(&_StakeManager.CallOpts)
}

// GetBondedPools is a free data retrieval call binding the contract method 0x58af8bf0.
//
// Solidity: function getBondedPools() view returns(address[] pools)
func (_StakeManager *StakeManagerCallerSession) GetBondedPools() ([]common.Address, error) {
	return _StakeManager.Contract.GetBondedPools(&_StakeManager.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_StakeManager *StakeManagerCaller) GetRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_StakeManager *StakeManagerSession) GetRate() (*big.Int, error) {
	return _StakeManager.Contract.GetRate(&_StakeManager.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) GetRate() (*big.Int, error) {
	return _StakeManager.Contract.GetRate(&_StakeManager.CallOpts)
}

// GetStakeRelayerFee is a free data retrieval call binding the contract method 0x95859bfb.
//
// Solidity: function getStakeRelayerFee() view returns(uint256)
func (_StakeManager *StakeManagerCaller) GetStakeRelayerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getStakeRelayerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeRelayerFee is a free data retrieval call binding the contract method 0x95859bfb.
//
// Solidity: function getStakeRelayerFee() view returns(uint256)
func (_StakeManager *StakeManagerSession) GetStakeRelayerFee() (*big.Int, error) {
	return _StakeManager.Contract.GetStakeRelayerFee(&_StakeManager.CallOpts)
}

// GetStakeRelayerFee is a free data retrieval call binding the contract method 0x95859bfb.
//
// Solidity: function getStakeRelayerFee() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) GetStakeRelayerFee() (*big.Int, error) {
	return _StakeManager.Contract.GetStakeRelayerFee(&_StakeManager.CallOpts)
}

// GetUnstakeIndexListOf is a free data retrieval call binding the contract method 0x615acc36.
//
// Solidity: function getUnstakeIndexListOf(address _staker) view returns(uint256[] unstakeIndexList)
func (_StakeManager *StakeManagerCaller) GetUnstakeIndexListOf(opts *bind.CallOpts, _staker common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getUnstakeIndexListOf", _staker)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUnstakeIndexListOf is a free data retrieval call binding the contract method 0x615acc36.
//
// Solidity: function getUnstakeIndexListOf(address _staker) view returns(uint256[] unstakeIndexList)
func (_StakeManager *StakeManagerSession) GetUnstakeIndexListOf(_staker common.Address) ([]*big.Int, error) {
	return _StakeManager.Contract.GetUnstakeIndexListOf(&_StakeManager.CallOpts, _staker)
}

// GetUnstakeIndexListOf is a free data retrieval call binding the contract method 0x615acc36.
//
// Solidity: function getUnstakeIndexListOf(address _staker) view returns(uint256[] unstakeIndexList)
func (_StakeManager *StakeManagerCallerSession) GetUnstakeIndexListOf(_staker common.Address) ([]*big.Int, error) {
	return _StakeManager.Contract.GetUnstakeIndexListOf(&_StakeManager.CallOpts, _staker)
}

// GetUnstakeRelayerFee is a free data retrieval call binding the contract method 0x1ed79e09.
//
// Solidity: function getUnstakeRelayerFee() view returns(uint256)
func (_StakeManager *StakeManagerCaller) GetUnstakeRelayerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getUnstakeRelayerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnstakeRelayerFee is a free data retrieval call binding the contract method 0x1ed79e09.
//
// Solidity: function getUnstakeRelayerFee() view returns(uint256)
func (_StakeManager *StakeManagerSession) GetUnstakeRelayerFee() (*big.Int, error) {
	return _StakeManager.Contract.GetUnstakeRelayerFee(&_StakeManager.CallOpts)
}

// GetUnstakeRelayerFee is a free data retrieval call binding the contract method 0x1ed79e09.
//
// Solidity: function getUnstakeRelayerFee() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) GetUnstakeRelayerFee() (*big.Int, error) {
	return _StakeManager.Contract.GetUnstakeRelayerFee(&_StakeManager.CallOpts)
}

// GetValidatorsOf is a free data retrieval call binding the contract method 0x301b97b0.
//
// Solidity: function getValidatorsOf(address _poolAddress) view returns(address[] validators)
func (_StakeManager *StakeManagerCaller) GetValidatorsOf(opts *bind.CallOpts, _poolAddress common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getValidatorsOf", _poolAddress)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidatorsOf is a free data retrieval call binding the contract method 0x301b97b0.
//
// Solidity: function getValidatorsOf(address _poolAddress) view returns(address[] validators)
func (_StakeManager *StakeManagerSession) GetValidatorsOf(_poolAddress common.Address) ([]common.Address, error) {
	return _StakeManager.Contract.GetValidatorsOf(&_StakeManager.CallOpts, _poolAddress)
}

// GetValidatorsOf is a free data retrieval call binding the contract method 0x301b97b0.
//
// Solidity: function getValidatorsOf(address _poolAddress) view returns(address[] validators)
func (_StakeManager *StakeManagerCallerSession) GetValidatorsOf(_poolAddress common.Address) ([]common.Address, error) {
	return _StakeManager.Contract.GetValidatorsOf(&_StakeManager.CallOpts, _poolAddress)
}

// GetVoterIndex is a free data retrieval call binding the contract method 0x7941743a.
//
// Solidity: function getVoterIndex(address _voter) view returns(uint256)
func (_StakeManager *StakeManagerCaller) GetVoterIndex(opts *bind.CallOpts, _voter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getVoterIndex", _voter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVoterIndex is a free data retrieval call binding the contract method 0x7941743a.
//
// Solidity: function getVoterIndex(address _voter) view returns(uint256)
func (_StakeManager *StakeManagerSession) GetVoterIndex(_voter common.Address) (*big.Int, error) {
	return _StakeManager.Contract.GetVoterIndex(&_StakeManager.CallOpts, _voter)
}

// GetVoterIndex is a free data retrieval call binding the contract method 0x7941743a.
//
// Solidity: function getVoterIndex(address _voter) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) GetVoterIndex(_voter common.Address) (*big.Int, error) {
	return _StakeManager.Contract.GetVoterIndex(&_StakeManager.CallOpts, _voter)
}

// HasVoted is a free data retrieval call binding the contract method 0xaadc3b72.
//
// Solidity: function hasVoted(bytes32 _proposalId, address _voter) view returns(bool)
func (_StakeManager *StakeManagerCaller) HasVoted(opts *bind.CallOpts, _proposalId [32]byte, _voter common.Address) (bool, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "hasVoted", _proposalId, _voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0xaadc3b72.
//
// Solidity: function hasVoted(bytes32 _proposalId, address _voter) view returns(bool)
func (_StakeManager *StakeManagerSession) HasVoted(_proposalId [32]byte, _voter common.Address) (bool, error) {
	return _StakeManager.Contract.HasVoted(&_StakeManager.CallOpts, _proposalId, _voter)
}

// HasVoted is a free data retrieval call binding the contract method 0xaadc3b72.
//
// Solidity: function hasVoted(bytes32 _proposalId, address _voter) view returns(bool)
func (_StakeManager *StakeManagerCallerSession) HasVoted(_proposalId [32]byte, _voter common.Address) (bool, error) {
	return _StakeManager.Contract.HasVoted(&_StakeManager.CallOpts, _proposalId, _voter)
}

// LatestEra is a free data retrieval call binding the contract method 0x3f6f5f32.
//
// Solidity: function latestEra() view returns(uint256)
func (_StakeManager *StakeManagerCaller) LatestEra(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "latestEra")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestEra is a free data retrieval call binding the contract method 0x3f6f5f32.
//
// Solidity: function latestEra() view returns(uint256)
func (_StakeManager *StakeManagerSession) LatestEra() (*big.Int, error) {
	return _StakeManager.Contract.LatestEra(&_StakeManager.CallOpts)
}

// LatestEra is a free data retrieval call binding the contract method 0x3f6f5f32.
//
// Solidity: function latestEra() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) LatestEra() (*big.Int, error) {
	return _StakeManager.Contract.LatestEra(&_StakeManager.CallOpts)
}

// LatestRewardTimestampOf is a free data retrieval call binding the contract method 0x98140d2f.
//
// Solidity: function latestRewardTimestampOf(address ) view returns(uint256)
func (_StakeManager *StakeManagerCaller) LatestRewardTimestampOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "latestRewardTimestampOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRewardTimestampOf is a free data retrieval call binding the contract method 0x98140d2f.
//
// Solidity: function latestRewardTimestampOf(address ) view returns(uint256)
func (_StakeManager *StakeManagerSession) LatestRewardTimestampOf(arg0 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.LatestRewardTimestampOf(&_StakeManager.CallOpts, arg0)
}

// LatestRewardTimestampOf is a free data retrieval call binding the contract method 0x98140d2f.
//
// Solidity: function latestRewardTimestampOf(address ) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) LatestRewardTimestampOf(arg0 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.LatestRewardTimestampOf(&_StakeManager.CallOpts, arg0)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_StakeManager *StakeManagerCaller) MinStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "minStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_StakeManager *StakeManagerSession) MinStakeAmount() (*big.Int, error) {
	return _StakeManager.Contract.MinStakeAmount(&_StakeManager.CallOpts)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) MinStakeAmount() (*big.Int, error) {
	return _StakeManager.Contract.MinStakeAmount(&_StakeManager.CallOpts)
}

// NextUnstakeIndex is a free data retrieval call binding the contract method 0x3bea9ee3.
//
// Solidity: function nextUnstakeIndex() view returns(uint256)
func (_StakeManager *StakeManagerCaller) NextUnstakeIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "nextUnstakeIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextUnstakeIndex is a free data retrieval call binding the contract method 0x3bea9ee3.
//
// Solidity: function nextUnstakeIndex() view returns(uint256)
func (_StakeManager *StakeManagerSession) NextUnstakeIndex() (*big.Int, error) {
	return _StakeManager.Contract.NextUnstakeIndex(&_StakeManager.CallOpts)
}

// NextUnstakeIndex is a free data retrieval call binding the contract method 0x3bea9ee3.
//
// Solidity: function nextUnstakeIndex() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) NextUnstakeIndex() (*big.Int, error) {
	return _StakeManager.Contract.NextUnstakeIndex(&_StakeManager.CallOpts)
}

// PendingDelegateOf is a free data retrieval call binding the contract method 0xa070d0e0.
//
// Solidity: function pendingDelegateOf(address ) view returns(uint256)
func (_StakeManager *StakeManagerCaller) PendingDelegateOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "pendingDelegateOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingDelegateOf is a free data retrieval call binding the contract method 0xa070d0e0.
//
// Solidity: function pendingDelegateOf(address ) view returns(uint256)
func (_StakeManager *StakeManagerSession) PendingDelegateOf(arg0 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.PendingDelegateOf(&_StakeManager.CallOpts, arg0)
}

// PendingDelegateOf is a free data retrieval call binding the contract method 0xa070d0e0.
//
// Solidity: function pendingDelegateOf(address ) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) PendingDelegateOf(arg0 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.PendingDelegateOf(&_StakeManager.CallOpts, arg0)
}

// PendingUndelegateOf is a free data retrieval call binding the contract method 0xe8bdddd6.
//
// Solidity: function pendingUndelegateOf(address ) view returns(uint256)
func (_StakeManager *StakeManagerCaller) PendingUndelegateOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "pendingUndelegateOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingUndelegateOf is a free data retrieval call binding the contract method 0xe8bdddd6.
//
// Solidity: function pendingUndelegateOf(address ) view returns(uint256)
func (_StakeManager *StakeManagerSession) PendingUndelegateOf(arg0 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.PendingUndelegateOf(&_StakeManager.CallOpts, arg0)
}

// PendingUndelegateOf is a free data retrieval call binding the contract method 0xe8bdddd6.
//
// Solidity: function pendingUndelegateOf(address ) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) PendingUndelegateOf(arg0 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.PendingUndelegateOf(&_StakeManager.CallOpts, arg0)
}

// PoolInfoOf is a free data retrieval call binding the contract method 0x008b5dd2.
//
// Solidity: function poolInfoOf(address ) view returns(uint256 era, uint256 bond, uint256 unbond, uint256 active)
func (_StakeManager *StakeManagerCaller) PoolInfoOf(opts *bind.CallOpts, arg0 common.Address) (struct {
	Era    *big.Int
	Bond   *big.Int
	Unbond *big.Int
	Active *big.Int
}, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "poolInfoOf", arg0)

	outstruct := new(struct {
		Era    *big.Int
		Bond   *big.Int
		Unbond *big.Int
		Active *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Era = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Bond = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Unbond = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfoOf is a free data retrieval call binding the contract method 0x008b5dd2.
//
// Solidity: function poolInfoOf(address ) view returns(uint256 era, uint256 bond, uint256 unbond, uint256 active)
func (_StakeManager *StakeManagerSession) PoolInfoOf(arg0 common.Address) (struct {
	Era    *big.Int
	Bond   *big.Int
	Unbond *big.Int
	Active *big.Int
}, error) {
	return _StakeManager.Contract.PoolInfoOf(&_StakeManager.CallOpts, arg0)
}

// PoolInfoOf is a free data retrieval call binding the contract method 0x008b5dd2.
//
// Solidity: function poolInfoOf(address ) view returns(uint256 era, uint256 bond, uint256 unbond, uint256 active)
func (_StakeManager *StakeManagerCallerSession) PoolInfoOf(arg0 common.Address) (struct {
	Era    *big.Int
	Bond   *big.Int
	Unbond *big.Int
	Active *big.Int
}, error) {
	return _StakeManager.Contract.PoolInfoOf(&_StakeManager.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(uint8 _status, uint16 _yesVotes, uint8 _yesVotesTotal)
func (_StakeManager *StakeManagerCaller) Proposals(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status        uint8
	YesVotes      uint16
	YesVotesTotal uint8
}, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Status        uint8
		YesVotes      uint16
		YesVotesTotal uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.YesVotes = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.YesVotesTotal = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(uint8 _status, uint16 _yesVotes, uint8 _yesVotesTotal)
func (_StakeManager *StakeManagerSession) Proposals(arg0 [32]byte) (struct {
	Status        uint8
	YesVotes      uint16
	YesVotesTotal uint8
}, error) {
	return _StakeManager.Contract.Proposals(&_StakeManager.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(uint8 _status, uint16 _yesVotes, uint8 _yesVotesTotal)
func (_StakeManager *StakeManagerCallerSession) Proposals(arg0 [32]byte) (struct {
	Status        uint8
	YesVotes      uint16
	YesVotesTotal uint8
}, error) {
	return _StakeManager.Contract.Proposals(&_StakeManager.CallOpts, arg0)
}

// ProtocolFeeCommission is a free data retrieval call binding the contract method 0x19301c26.
//
// Solidity: function protocolFeeCommission() view returns(uint256)
func (_StakeManager *StakeManagerCaller) ProtocolFeeCommission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "protocolFeeCommission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeCommission is a free data retrieval call binding the contract method 0x19301c26.
//
// Solidity: function protocolFeeCommission() view returns(uint256)
func (_StakeManager *StakeManagerSession) ProtocolFeeCommission() (*big.Int, error) {
	return _StakeManager.Contract.ProtocolFeeCommission(&_StakeManager.CallOpts)
}

// ProtocolFeeCommission is a free data retrieval call binding the contract method 0x19301c26.
//
// Solidity: function protocolFeeCommission() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) ProtocolFeeCommission() (*big.Int, error) {
	return _StakeManager.Contract.ProtocolFeeCommission(&_StakeManager.CallOpts)
}

// RTokenAddress is a free data retrieval call binding the contract method 0x35381358.
//
// Solidity: function rTokenAddress() view returns(address)
func (_StakeManager *StakeManagerCaller) RTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "rTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RTokenAddress is a free data retrieval call binding the contract method 0x35381358.
//
// Solidity: function rTokenAddress() view returns(address)
func (_StakeManager *StakeManagerSession) RTokenAddress() (common.Address, error) {
	return _StakeManager.Contract.RTokenAddress(&_StakeManager.CallOpts)
}

// RTokenAddress is a free data retrieval call binding the contract method 0x35381358.
//
// Solidity: function rTokenAddress() view returns(address)
func (_StakeManager *StakeManagerCallerSession) RTokenAddress() (common.Address, error) {
	return _StakeManager.Contract.RTokenAddress(&_StakeManager.CallOpts)
}

// RateChangeLimit is a free data retrieval call binding the contract method 0xc0152b71.
//
// Solidity: function rateChangeLimit() view returns(uint256)
func (_StakeManager *StakeManagerCaller) RateChangeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "rateChangeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RateChangeLimit is a free data retrieval call binding the contract method 0xc0152b71.
//
// Solidity: function rateChangeLimit() view returns(uint256)
func (_StakeManager *StakeManagerSession) RateChangeLimit() (*big.Int, error) {
	return _StakeManager.Contract.RateChangeLimit(&_StakeManager.CallOpts)
}

// RateChangeLimit is a free data retrieval call binding the contract method 0xc0152b71.
//
// Solidity: function rateChangeLimit() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) RateChangeLimit() (*big.Int, error) {
	return _StakeManager.Contract.RateChangeLimit(&_StakeManager.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_StakeManager *StakeManagerCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_StakeManager *StakeManagerSession) Threshold() (uint8, error) {
	return _StakeManager.Contract.Threshold(&_StakeManager.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_StakeManager *StakeManagerCallerSession) Threshold() (uint8, error) {
	return _StakeManager.Contract.Threshold(&_StakeManager.CallOpts)
}

// TotalProtocolFee is a free data retrieval call binding the contract method 0x88611f35.
//
// Solidity: function totalProtocolFee() view returns(uint256)
func (_StakeManager *StakeManagerCaller) TotalProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "totalProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProtocolFee is a free data retrieval call binding the contract method 0x88611f35.
//
// Solidity: function totalProtocolFee() view returns(uint256)
func (_StakeManager *StakeManagerSession) TotalProtocolFee() (*big.Int, error) {
	return _StakeManager.Contract.TotalProtocolFee(&_StakeManager.CallOpts)
}

// TotalProtocolFee is a free data retrieval call binding the contract method 0x88611f35.
//
// Solidity: function totalProtocolFee() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) TotalProtocolFee() (*big.Int, error) {
	return _StakeManager.Contract.TotalProtocolFee(&_StakeManager.CallOpts)
}

// TotalRTokenSupply is a free data retrieval call binding the contract method 0x7a7c27c0.
//
// Solidity: function totalRTokenSupply() view returns(uint256)
func (_StakeManager *StakeManagerCaller) TotalRTokenSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "totalRTokenSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRTokenSupply is a free data retrieval call binding the contract method 0x7a7c27c0.
//
// Solidity: function totalRTokenSupply() view returns(uint256)
func (_StakeManager *StakeManagerSession) TotalRTokenSupply() (*big.Int, error) {
	return _StakeManager.Contract.TotalRTokenSupply(&_StakeManager.CallOpts)
}

// TotalRTokenSupply is a free data retrieval call binding the contract method 0x7a7c27c0.
//
// Solidity: function totalRTokenSupply() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) TotalRTokenSupply() (*big.Int, error) {
	return _StakeManager.Contract.TotalRTokenSupply(&_StakeManager.CallOpts)
}

// TransferGas is a free data retrieval call binding the contract method 0xfa03f797.
//
// Solidity: function transferGas() view returns(uint256)
func (_StakeManager *StakeManagerCaller) TransferGas(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "transferGas")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferGas is a free data retrieval call binding the contract method 0xfa03f797.
//
// Solidity: function transferGas() view returns(uint256)
func (_StakeManager *StakeManagerSession) TransferGas() (*big.Int, error) {
	return _StakeManager.Contract.TransferGas(&_StakeManager.CallOpts)
}

// TransferGas is a free data retrieval call binding the contract method 0xfa03f797.
//
// Solidity: function transferGas() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) TransferGas() (*big.Int, error) {
	return _StakeManager.Contract.TransferGas(&_StakeManager.CallOpts)
}

// UnbondingDuration is a free data retrieval call binding the contract method 0xccf6802a.
//
// Solidity: function unbondingDuration() view returns(uint256)
func (_StakeManager *StakeManagerCaller) UnbondingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "unbondingDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnbondingDuration is a free data retrieval call binding the contract method 0xccf6802a.
//
// Solidity: function unbondingDuration() view returns(uint256)
func (_StakeManager *StakeManagerSession) UnbondingDuration() (*big.Int, error) {
	return _StakeManager.Contract.UnbondingDuration(&_StakeManager.CallOpts)
}

// UnbondingDuration is a free data retrieval call binding the contract method 0xccf6802a.
//
// Solidity: function unbondingDuration() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) UnbondingDuration() (*big.Int, error) {
	return _StakeManager.Contract.UnbondingDuration(&_StakeManager.CallOpts)
}

// UndistributedRewardOf is a free data retrieval call binding the contract method 0x07731a3b.
//
// Solidity: function undistributedRewardOf(address ) view returns(uint256)
func (_StakeManager *StakeManagerCaller) UndistributedRewardOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "undistributedRewardOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UndistributedRewardOf is a free data retrieval call binding the contract method 0x07731a3b.
//
// Solidity: function undistributedRewardOf(address ) view returns(uint256)
func (_StakeManager *StakeManagerSession) UndistributedRewardOf(arg0 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.UndistributedRewardOf(&_StakeManager.CallOpts, arg0)
}

// UndistributedRewardOf is a free data retrieval call binding the contract method 0x07731a3b.
//
// Solidity: function undistributedRewardOf(address ) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) UndistributedRewardOf(arg0 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.UndistributedRewardOf(&_StakeManager.CallOpts, arg0)
}

// UnstakeAtIndex is a free data retrieval call binding the contract method 0x6e436c6e.
//
// Solidity: function unstakeAtIndex(uint256 ) view returns(uint256 era, address pool, address receiver, uint256 amount)
func (_StakeManager *StakeManagerCaller) UnstakeAtIndex(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Era      *big.Int
	Pool     common.Address
	Receiver common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "unstakeAtIndex", arg0)

	outstruct := new(struct {
		Era      *big.Int
		Pool     common.Address
		Receiver common.Address
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Era = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Pool = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Receiver = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UnstakeAtIndex is a free data retrieval call binding the contract method 0x6e436c6e.
//
// Solidity: function unstakeAtIndex(uint256 ) view returns(uint256 era, address pool, address receiver, uint256 amount)
func (_StakeManager *StakeManagerSession) UnstakeAtIndex(arg0 *big.Int) (struct {
	Era      *big.Int
	Pool     common.Address
	Receiver common.Address
	Amount   *big.Int
}, error) {
	return _StakeManager.Contract.UnstakeAtIndex(&_StakeManager.CallOpts, arg0)
}

// UnstakeAtIndex is a free data retrieval call binding the contract method 0x6e436c6e.
//
// Solidity: function unstakeAtIndex(uint256 ) view returns(uint256 era, address pool, address receiver, uint256 amount)
func (_StakeManager *StakeManagerCallerSession) UnstakeAtIndex(arg0 *big.Int) (struct {
	Era      *big.Int
	Pool     common.Address
	Receiver common.Address
	Amount   *big.Int
}, error) {
	return _StakeManager.Contract.UnstakeAtIndex(&_StakeManager.CallOpts, arg0)
}

// UnstakeFeeCommission is a free data retrieval call binding the contract method 0xe1bb5897.
//
// Solidity: function unstakeFeeCommission() view returns(uint256)
func (_StakeManager *StakeManagerCaller) UnstakeFeeCommission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "unstakeFeeCommission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnstakeFeeCommission is a free data retrieval call binding the contract method 0xe1bb5897.
//
// Solidity: function unstakeFeeCommission() view returns(uint256)
func (_StakeManager *StakeManagerSession) UnstakeFeeCommission() (*big.Int, error) {
	return _StakeManager.Contract.UnstakeFeeCommission(&_StakeManager.CallOpts)
}

// UnstakeFeeCommission is a free data retrieval call binding the contract method 0xe1bb5897.
//
// Solidity: function unstakeFeeCommission() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) UnstakeFeeCommission() (*big.Int, error) {
	return _StakeManager.Contract.UnstakeFeeCommission(&_StakeManager.CallOpts)
}

// AddStakePool is a paid mutator transaction binding the contract method 0x0f772a1d.
//
// Solidity: function addStakePool(address _stakePool) returns()
func (_StakeManager *StakeManagerTransactor) AddStakePool(opts *bind.TransactOpts, _stakePool common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "addStakePool", _stakePool)
}

// AddStakePool is a paid mutator transaction binding the contract method 0x0f772a1d.
//
// Solidity: function addStakePool(address _stakePool) returns()
func (_StakeManager *StakeManagerSession) AddStakePool(_stakePool common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.AddStakePool(&_StakeManager.TransactOpts, _stakePool)
}

// AddStakePool is a paid mutator transaction binding the contract method 0x0f772a1d.
//
// Solidity: function addStakePool(address _stakePool) returns()
func (_StakeManager *StakeManagerTransactorSession) AddStakePool(_stakePool common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.AddStakePool(&_StakeManager.TransactOpts, _stakePool)
}

// AddValidator is a paid mutator transaction binding the contract method 0xb1819c2b.
//
// Solidity: function addValidator(address _stakePool, address _validator) returns()
func (_StakeManager *StakeManagerTransactor) AddValidator(opts *bind.TransactOpts, _stakePool common.Address, _validator common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "addValidator", _stakePool, _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0xb1819c2b.
//
// Solidity: function addValidator(address _stakePool, address _validator) returns()
func (_StakeManager *StakeManagerSession) AddValidator(_stakePool common.Address, _validator common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.AddValidator(&_StakeManager.TransactOpts, _stakePool, _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0xb1819c2b.
//
// Solidity: function addValidator(address _stakePool, address _validator) returns()
func (_StakeManager *StakeManagerTransactorSession) AddValidator(_stakePool common.Address, _validator common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.AddValidator(&_StakeManager.TransactOpts, _stakePool, _validator)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(address _voter) returns()
func (_StakeManager *StakeManagerTransactor) AddVoter(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "addVoter", _voter)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(address _voter) returns()
func (_StakeManager *StakeManagerSession) AddVoter(_voter common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.AddVoter(&_StakeManager.TransactOpts, _voter)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(address _voter) returns()
func (_StakeManager *StakeManagerTransactorSession) AddVoter(_voter common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.AddVoter(&_StakeManager.TransactOpts, _voter)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _newThreshold) returns()
func (_StakeManager *StakeManagerTransactor) ChangeThreshold(opts *bind.TransactOpts, _newThreshold *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "changeThreshold", _newThreshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _newThreshold) returns()
func (_StakeManager *StakeManagerSession) ChangeThreshold(_newThreshold *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ChangeThreshold(&_StakeManager.TransactOpts, _newThreshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _newThreshold) returns()
func (_StakeManager *StakeManagerTransactorSession) ChangeThreshold(_newThreshold *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ChangeThreshold(&_StakeManager.TransactOpts, _newThreshold)
}

// Init is a paid mutator transaction binding the contract method 0xec8690c8.
//
// Solidity: function init(address[] _initialVoters, uint256 _initialThreshold, address _rTokenAddress, uint256 _unbondingDuration) returns()
func (_StakeManager *StakeManagerTransactor) Init(opts *bind.TransactOpts, _initialVoters []common.Address, _initialThreshold *big.Int, _rTokenAddress common.Address, _unbondingDuration *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "init", _initialVoters, _initialThreshold, _rTokenAddress, _unbondingDuration)
}

// Init is a paid mutator transaction binding the contract method 0xec8690c8.
//
// Solidity: function init(address[] _initialVoters, uint256 _initialThreshold, address _rTokenAddress, uint256 _unbondingDuration) returns()
func (_StakeManager *StakeManagerSession) Init(_initialVoters []common.Address, _initialThreshold *big.Int, _rTokenAddress common.Address, _unbondingDuration *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Init(&_StakeManager.TransactOpts, _initialVoters, _initialThreshold, _rTokenAddress, _unbondingDuration)
}

// Init is a paid mutator transaction binding the contract method 0xec8690c8.
//
// Solidity: function init(address[] _initialVoters, uint256 _initialThreshold, address _rTokenAddress, uint256 _unbondingDuration) returns()
func (_StakeManager *StakeManagerTransactorSession) Init(_initialVoters []common.Address, _initialThreshold *big.Int, _rTokenAddress common.Address, _unbondingDuration *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Init(&_StakeManager.TransactOpts, _initialVoters, _initialThreshold, _rTokenAddress, _unbondingDuration)
}

// InitMultisig is a paid mutator transaction binding the contract method 0x88a09d91.
//
// Solidity: function initMultisig(address[] _voters, uint256 _initialThreshold) returns()
func (_StakeManager *StakeManagerTransactor) InitMultisig(opts *bind.TransactOpts, _voters []common.Address, _initialThreshold *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "initMultisig", _voters, _initialThreshold)
}

// InitMultisig is a paid mutator transaction binding the contract method 0x88a09d91.
//
// Solidity: function initMultisig(address[] _voters, uint256 _initialThreshold) returns()
func (_StakeManager *StakeManagerSession) InitMultisig(_voters []common.Address, _initialThreshold *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.InitMultisig(&_StakeManager.TransactOpts, _voters, _initialThreshold)
}

// InitMultisig is a paid mutator transaction binding the contract method 0x88a09d91.
//
// Solidity: function initMultisig(address[] _voters, uint256 _initialThreshold) returns()
func (_StakeManager *StakeManagerTransactorSession) InitMultisig(_voters []common.Address, _initialThreshold *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.InitMultisig(&_StakeManager.TransactOpts, _voters, _initialThreshold)
}

// Migrate is a paid mutator transaction binding the contract method 0x197513b0.
//
// Solidity: function migrate(address _stakePoolAddress, address _validator, uint256 _govDelegated, uint256 _bond, uint256 _unbond, uint256 _pendingDelegate, uint256 _rate, uint256 _totalRTokenSupply, uint256 _totalProtocolFee, uint256 _era, uint256 _latestRewardtimestamp, uint256 _undistributedReward) returns()
func (_StakeManager *StakeManagerTransactor) Migrate(opts *bind.TransactOpts, _stakePoolAddress common.Address, _validator common.Address, _govDelegated *big.Int, _bond *big.Int, _unbond *big.Int, _pendingDelegate *big.Int, _rate *big.Int, _totalRTokenSupply *big.Int, _totalProtocolFee *big.Int, _era *big.Int, _latestRewardtimestamp *big.Int, _undistributedReward *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "migrate", _stakePoolAddress, _validator, _govDelegated, _bond, _unbond, _pendingDelegate, _rate, _totalRTokenSupply, _totalProtocolFee, _era, _latestRewardtimestamp, _undistributedReward)
}

// Migrate is a paid mutator transaction binding the contract method 0x197513b0.
//
// Solidity: function migrate(address _stakePoolAddress, address _validator, uint256 _govDelegated, uint256 _bond, uint256 _unbond, uint256 _pendingDelegate, uint256 _rate, uint256 _totalRTokenSupply, uint256 _totalProtocolFee, uint256 _era, uint256 _latestRewardtimestamp, uint256 _undistributedReward) returns()
func (_StakeManager *StakeManagerSession) Migrate(_stakePoolAddress common.Address, _validator common.Address, _govDelegated *big.Int, _bond *big.Int, _unbond *big.Int, _pendingDelegate *big.Int, _rate *big.Int, _totalRTokenSupply *big.Int, _totalProtocolFee *big.Int, _era *big.Int, _latestRewardtimestamp *big.Int, _undistributedReward *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Migrate(&_StakeManager.TransactOpts, _stakePoolAddress, _validator, _govDelegated, _bond, _unbond, _pendingDelegate, _rate, _totalRTokenSupply, _totalProtocolFee, _era, _latestRewardtimestamp, _undistributedReward)
}

// Migrate is a paid mutator transaction binding the contract method 0x197513b0.
//
// Solidity: function migrate(address _stakePoolAddress, address _validator, uint256 _govDelegated, uint256 _bond, uint256 _unbond, uint256 _pendingDelegate, uint256 _rate, uint256 _totalRTokenSupply, uint256 _totalProtocolFee, uint256 _era, uint256 _latestRewardtimestamp, uint256 _undistributedReward) returns()
func (_StakeManager *StakeManagerTransactorSession) Migrate(_stakePoolAddress common.Address, _validator common.Address, _govDelegated *big.Int, _bond *big.Int, _unbond *big.Int, _pendingDelegate *big.Int, _rate *big.Int, _totalRTokenSupply *big.Int, _totalProtocolFee *big.Int, _era *big.Int, _latestRewardtimestamp *big.Int, _undistributedReward *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Migrate(&_StakeManager.TransactOpts, _stakePoolAddress, _validator, _govDelegated, _bond, _unbond, _pendingDelegate, _rate, _totalRTokenSupply, _totalProtocolFee, _era, _latestRewardtimestamp, _undistributedReward)
}

// NewEra is a paid mutator transaction binding the contract method 0xfb1930c5.
//
// Solidity: function newEra(uint256 _era, address[] _poolAddressList, uint256[] _newRewardList, uint256[] _latestRewardTimestampList) returns()
func (_StakeManager *StakeManagerTransactor) NewEra(opts *bind.TransactOpts, _era *big.Int, _poolAddressList []common.Address, _newRewardList []*big.Int, _latestRewardTimestampList []*big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "newEra", _era, _poolAddressList, _newRewardList, _latestRewardTimestampList)
}

// NewEra is a paid mutator transaction binding the contract method 0xfb1930c5.
//
// Solidity: function newEra(uint256 _era, address[] _poolAddressList, uint256[] _newRewardList, uint256[] _latestRewardTimestampList) returns()
func (_StakeManager *StakeManagerSession) NewEra(_era *big.Int, _poolAddressList []common.Address, _newRewardList []*big.Int, _latestRewardTimestampList []*big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.NewEra(&_StakeManager.TransactOpts, _era, _poolAddressList, _newRewardList, _latestRewardTimestampList)
}

// NewEra is a paid mutator transaction binding the contract method 0xfb1930c5.
//
// Solidity: function newEra(uint256 _era, address[] _poolAddressList, uint256[] _newRewardList, uint256[] _latestRewardTimestampList) returns()
func (_StakeManager *StakeManagerTransactorSession) NewEra(_era *big.Int, _poolAddressList []common.Address, _newRewardList []*big.Int, _latestRewardTimestampList []*big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.NewEra(&_StakeManager.TransactOpts, _era, _poolAddressList, _newRewardList, _latestRewardTimestampList)
}

// Redelegate is a paid mutator transaction binding the contract method 0x77fc27fa.
//
// Solidity: function redelegate(address _stakePool, address _srcValidator, address _dstValidator, uint256 _amount) returns()
func (_StakeManager *StakeManagerTransactor) Redelegate(opts *bind.TransactOpts, _stakePool common.Address, _srcValidator common.Address, _dstValidator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "redelegate", _stakePool, _srcValidator, _dstValidator, _amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x77fc27fa.
//
// Solidity: function redelegate(address _stakePool, address _srcValidator, address _dstValidator, uint256 _amount) returns()
func (_StakeManager *StakeManagerSession) Redelegate(_stakePool common.Address, _srcValidator common.Address, _dstValidator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Redelegate(&_StakeManager.TransactOpts, _stakePool, _srcValidator, _dstValidator, _amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x77fc27fa.
//
// Solidity: function redelegate(address _stakePool, address _srcValidator, address _dstValidator, uint256 _amount) returns()
func (_StakeManager *StakeManagerTransactorSession) Redelegate(_stakePool common.Address, _srcValidator common.Address, _dstValidator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Redelegate(&_StakeManager.TransactOpts, _stakePool, _srcValidator, _dstValidator, _amount)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(address _voter) returns()
func (_StakeManager *StakeManagerTransactor) RemoveVoter(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "removeVoter", _voter)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(address _voter) returns()
func (_StakeManager *StakeManagerSession) RemoveVoter(_voter common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.RemoveVoter(&_StakeManager.TransactOpts, _voter)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(address _voter) returns()
func (_StakeManager *StakeManagerTransactorSession) RemoveVoter(_voter common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.RemoveVoter(&_StakeManager.TransactOpts, _voter)
}

// RmStakePool is a paid mutator transaction binding the contract method 0x58c76ca8.
//
// Solidity: function rmStakePool(address _stakePool) returns()
func (_StakeManager *StakeManagerTransactor) RmStakePool(opts *bind.TransactOpts, _stakePool common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "rmStakePool", _stakePool)
}

// RmStakePool is a paid mutator transaction binding the contract method 0x58c76ca8.
//
// Solidity: function rmStakePool(address _stakePool) returns()
func (_StakeManager *StakeManagerSession) RmStakePool(_stakePool common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.RmStakePool(&_StakeManager.TransactOpts, _stakePool)
}

// RmStakePool is a paid mutator transaction binding the contract method 0x58c76ca8.
//
// Solidity: function rmStakePool(address _stakePool) returns()
func (_StakeManager *StakeManagerTransactorSession) RmStakePool(_stakePool common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.RmStakePool(&_StakeManager.TransactOpts, _stakePool)
}

// RmValidator is a paid mutator transaction binding the contract method 0x89a407c9.
//
// Solidity: function rmValidator(address _stakePool, address _validator) returns()
func (_StakeManager *StakeManagerTransactor) RmValidator(opts *bind.TransactOpts, _stakePool common.Address, _validator common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "rmValidator", _stakePool, _validator)
}

// RmValidator is a paid mutator transaction binding the contract method 0x89a407c9.
//
// Solidity: function rmValidator(address _stakePool, address _validator) returns()
func (_StakeManager *StakeManagerSession) RmValidator(_stakePool common.Address, _validator common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.RmValidator(&_StakeManager.TransactOpts, _stakePool, _validator)
}

// RmValidator is a paid mutator transaction binding the contract method 0x89a407c9.
//
// Solidity: function rmValidator(address _stakePool, address _validator) returns()
func (_StakeManager *StakeManagerTransactorSession) RmValidator(_stakePool common.Address, _validator common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.RmValidator(&_StakeManager.TransactOpts, _stakePool, _validator)
}

// SetParams is a paid mutator transaction binding the contract method 0xa6f41a40.
//
// Solidity: function setParams(uint256 _unstakeFeeCommission, uint256 _protocolFeeCommission, uint256 _minStakeAmount, uint256 _unbondingDuration, uint256 _rateChangeLimit, uint256 _eraSeconds, uint256 _eraOffset, uint256 _transferGas, uint256 _delegatedDiffLimit) returns()
func (_StakeManager *StakeManagerTransactor) SetParams(opts *bind.TransactOpts, _unstakeFeeCommission *big.Int, _protocolFeeCommission *big.Int, _minStakeAmount *big.Int, _unbondingDuration *big.Int, _rateChangeLimit *big.Int, _eraSeconds *big.Int, _eraOffset *big.Int, _transferGas *big.Int, _delegatedDiffLimit *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "setParams", _unstakeFeeCommission, _protocolFeeCommission, _minStakeAmount, _unbondingDuration, _rateChangeLimit, _eraSeconds, _eraOffset, _transferGas, _delegatedDiffLimit)
}

// SetParams is a paid mutator transaction binding the contract method 0xa6f41a40.
//
// Solidity: function setParams(uint256 _unstakeFeeCommission, uint256 _protocolFeeCommission, uint256 _minStakeAmount, uint256 _unbondingDuration, uint256 _rateChangeLimit, uint256 _eraSeconds, uint256 _eraOffset, uint256 _transferGas, uint256 _delegatedDiffLimit) returns()
func (_StakeManager *StakeManagerSession) SetParams(_unstakeFeeCommission *big.Int, _protocolFeeCommission *big.Int, _minStakeAmount *big.Int, _unbondingDuration *big.Int, _rateChangeLimit *big.Int, _eraSeconds *big.Int, _eraOffset *big.Int, _transferGas *big.Int, _delegatedDiffLimit *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetParams(&_StakeManager.TransactOpts, _unstakeFeeCommission, _protocolFeeCommission, _minStakeAmount, _unbondingDuration, _rateChangeLimit, _eraSeconds, _eraOffset, _transferGas, _delegatedDiffLimit)
}

// SetParams is a paid mutator transaction binding the contract method 0xa6f41a40.
//
// Solidity: function setParams(uint256 _unstakeFeeCommission, uint256 _protocolFeeCommission, uint256 _minStakeAmount, uint256 _unbondingDuration, uint256 _rateChangeLimit, uint256 _eraSeconds, uint256 _eraOffset, uint256 _transferGas, uint256 _delegatedDiffLimit) returns()
func (_StakeManager *StakeManagerTransactorSession) SetParams(_unstakeFeeCommission *big.Int, _protocolFeeCommission *big.Int, _minStakeAmount *big.Int, _unbondingDuration *big.Int, _rateChangeLimit *big.Int, _eraSeconds *big.Int, _eraOffset *big.Int, _transferGas *big.Int, _delegatedDiffLimit *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetParams(&_StakeManager.TransactOpts, _unstakeFeeCommission, _protocolFeeCommission, _minStakeAmount, _unbondingDuration, _rateChangeLimit, _eraSeconds, _eraOffset, _transferGas, _delegatedDiffLimit)
}

// Settle is a paid mutator transaction binding the contract method 0x6a256b29.
//
// Solidity: function settle(address _poolAddress) returns()
func (_StakeManager *StakeManagerTransactor) Settle(opts *bind.TransactOpts, _poolAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "settle", _poolAddress)
}

// Settle is a paid mutator transaction binding the contract method 0x6a256b29.
//
// Solidity: function settle(address _poolAddress) returns()
func (_StakeManager *StakeManagerSession) Settle(_poolAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.Settle(&_StakeManager.TransactOpts, _poolAddress)
}

// Settle is a paid mutator transaction binding the contract method 0x6a256b29.
//
// Solidity: function settle(address _poolAddress) returns()
func (_StakeManager *StakeManagerTransactorSession) Settle(_poolAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.Settle(&_StakeManager.TransactOpts, _poolAddress)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _stakeAmount) payable returns()
func (_StakeManager *StakeManagerTransactor) Stake(opts *bind.TransactOpts, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "stake", _stakeAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _stakeAmount) payable returns()
func (_StakeManager *StakeManagerSession) Stake(_stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Stake(&_StakeManager.TransactOpts, _stakeAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _stakeAmount) payable returns()
func (_StakeManager *StakeManagerTransactorSession) Stake(_stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Stake(&_StakeManager.TransactOpts, _stakeAmount)
}

// StakeWithPool is a paid mutator transaction binding the contract method 0x1525be32.
//
// Solidity: function stakeWithPool(address _stakePoolAddress, uint256 _stakeAmount) payable returns()
func (_StakeManager *StakeManagerTransactor) StakeWithPool(opts *bind.TransactOpts, _stakePoolAddress common.Address, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "stakeWithPool", _stakePoolAddress, _stakeAmount)
}

// StakeWithPool is a paid mutator transaction binding the contract method 0x1525be32.
//
// Solidity: function stakeWithPool(address _stakePoolAddress, uint256 _stakeAmount) payable returns()
func (_StakeManager *StakeManagerSession) StakeWithPool(_stakePoolAddress common.Address, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeWithPool(&_StakeManager.TransactOpts, _stakePoolAddress, _stakeAmount)
}

// StakeWithPool is a paid mutator transaction binding the contract method 0x1525be32.
//
// Solidity: function stakeWithPool(address _stakePoolAddress, uint256 _stakeAmount) payable returns()
func (_StakeManager *StakeManagerTransactorSession) StakeWithPool(_stakePoolAddress common.Address, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeWithPool(&_StakeManager.TransactOpts, _stakePoolAddress, _stakeAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_StakeManager *StakeManagerTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_StakeManager *StakeManagerSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.TransferOwnership(&_StakeManager.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_StakeManager *StakeManagerTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.TransferOwnership(&_StakeManager.TransactOpts, _newOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _rTokenAmount) payable returns()
func (_StakeManager *StakeManagerTransactor) Unstake(opts *bind.TransactOpts, _rTokenAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "unstake", _rTokenAmount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _rTokenAmount) payable returns()
func (_StakeManager *StakeManagerSession) Unstake(_rTokenAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Unstake(&_StakeManager.TransactOpts, _rTokenAmount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _rTokenAmount) payable returns()
func (_StakeManager *StakeManagerTransactorSession) Unstake(_rTokenAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Unstake(&_StakeManager.TransactOpts, _rTokenAmount)
}

// UnstakeWithPool is a paid mutator transaction binding the contract method 0xb608b458.
//
// Solidity: function unstakeWithPool(address _stakePoolAddress, uint256 _rTokenAmount) payable returns()
func (_StakeManager *StakeManagerTransactor) UnstakeWithPool(opts *bind.TransactOpts, _stakePoolAddress common.Address, _rTokenAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "unstakeWithPool", _stakePoolAddress, _rTokenAmount)
}

// UnstakeWithPool is a paid mutator transaction binding the contract method 0xb608b458.
//
// Solidity: function unstakeWithPool(address _stakePoolAddress, uint256 _rTokenAmount) payable returns()
func (_StakeManager *StakeManagerSession) UnstakeWithPool(_stakePoolAddress common.Address, _rTokenAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UnstakeWithPool(&_StakeManager.TransactOpts, _stakePoolAddress, _rTokenAmount)
}

// UnstakeWithPool is a paid mutator transaction binding the contract method 0xb608b458.
//
// Solidity: function unstakeWithPool(address _stakePoolAddress, uint256 _rTokenAmount) payable returns()
func (_StakeManager *StakeManagerTransactorSession) UnstakeWithPool(_stakePoolAddress common.Address, _rTokenAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UnstakeWithPool(&_StakeManager.TransactOpts, _stakePoolAddress, _rTokenAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StakeManager *StakeManagerTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StakeManager *StakeManagerSession) Withdraw() (*types.Transaction, error) {
	return _StakeManager.Contract.Withdraw(&_StakeManager.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StakeManager *StakeManagerTransactorSession) Withdraw() (*types.Transaction, error) {
	return _StakeManager.Contract.Withdraw(&_StakeManager.TransactOpts)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0x668fb6dc.
//
// Solidity: function withdrawProtocolFee(address _to) returns()
func (_StakeManager *StakeManagerTransactor) WithdrawProtocolFee(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdrawProtocolFee", _to)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0x668fb6dc.
//
// Solidity: function withdrawProtocolFee(address _to) returns()
func (_StakeManager *StakeManagerSession) WithdrawProtocolFee(_to common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawProtocolFee(&_StakeManager.TransactOpts, _to)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0x668fb6dc.
//
// Solidity: function withdrawProtocolFee(address _to) returns()
func (_StakeManager *StakeManagerTransactorSession) WithdrawProtocolFee(_to common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawProtocolFee(&_StakeManager.TransactOpts, _to)
}

// WithdrawRelayerFee is a paid mutator transaction binding the contract method 0x3b9b1822.
//
// Solidity: function withdrawRelayerFee(address _to) returns()
func (_StakeManager *StakeManagerTransactor) WithdrawRelayerFee(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdrawRelayerFee", _to)
}

// WithdrawRelayerFee is a paid mutator transaction binding the contract method 0x3b9b1822.
//
// Solidity: function withdrawRelayerFee(address _to) returns()
func (_StakeManager *StakeManagerSession) WithdrawRelayerFee(_to common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawRelayerFee(&_StakeManager.TransactOpts, _to)
}

// WithdrawRelayerFee is a paid mutator transaction binding the contract method 0x3b9b1822.
//
// Solidity: function withdrawRelayerFee(address _to) returns()
func (_StakeManager *StakeManagerTransactorSession) WithdrawRelayerFee(_to common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawRelayerFee(&_StakeManager.TransactOpts, _to)
}

// WithdrawWithPool is a paid mutator transaction binding the contract method 0xf737abac.
//
// Solidity: function withdrawWithPool(address _poolAddress) returns()
func (_StakeManager *StakeManagerTransactor) WithdrawWithPool(opts *bind.TransactOpts, _poolAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdrawWithPool", _poolAddress)
}

// WithdrawWithPool is a paid mutator transaction binding the contract method 0xf737abac.
//
// Solidity: function withdrawWithPool(address _poolAddress) returns()
func (_StakeManager *StakeManagerSession) WithdrawWithPool(_poolAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawWithPool(&_StakeManager.TransactOpts, _poolAddress)
}

// WithdrawWithPool is a paid mutator transaction binding the contract method 0xf737abac.
//
// Solidity: function withdrawWithPool(address _poolAddress) returns()
func (_StakeManager *StakeManagerTransactorSession) WithdrawWithPool(_poolAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawWithPool(&_StakeManager.TransactOpts, _poolAddress)
}

// StakeManagerExecuteNewEraIterator is returned from FilterExecuteNewEra and is used to iterate over the raw logs and unpacked data for ExecuteNewEra events raised by the StakeManager contract.
type StakeManagerExecuteNewEraIterator struct {
	Event *StakeManagerExecuteNewEra // Event containing the contract specifics and raw log

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
func (it *StakeManagerExecuteNewEraIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerExecuteNewEra)
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
		it.Event = new(StakeManagerExecuteNewEra)
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
func (it *StakeManagerExecuteNewEraIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerExecuteNewEraIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerExecuteNewEra represents a ExecuteNewEra event raised by the StakeManager contract.
type StakeManagerExecuteNewEra struct {
	Era  *big.Int
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExecuteNewEra is a free log retrieval operation binding the contract event 0x02105621fc31aa3ac04a9845beacd54c700e2ab23ff8acdd755dfd878ae61f02.
//
// Solidity: event ExecuteNewEra(uint256 indexed era, uint256 rate)
func (_StakeManager *StakeManagerFilterer) FilterExecuteNewEra(opts *bind.FilterOpts, era []*big.Int) (*StakeManagerExecuteNewEraIterator, error) {

	var eraRule []interface{}
	for _, eraItem := range era {
		eraRule = append(eraRule, eraItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "ExecuteNewEra", eraRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerExecuteNewEraIterator{contract: _StakeManager.contract, event: "ExecuteNewEra", logs: logs, sub: sub}, nil
}

// WatchExecuteNewEra is a free log subscription operation binding the contract event 0x02105621fc31aa3ac04a9845beacd54c700e2ab23ff8acdd755dfd878ae61f02.
//
// Solidity: event ExecuteNewEra(uint256 indexed era, uint256 rate)
func (_StakeManager *StakeManagerFilterer) WatchExecuteNewEra(opts *bind.WatchOpts, sink chan<- *StakeManagerExecuteNewEra, era []*big.Int) (event.Subscription, error) {

	var eraRule []interface{}
	for _, eraItem := range era {
		eraRule = append(eraRule, eraItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "ExecuteNewEra", eraRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerExecuteNewEra)
				if err := _StakeManager.contract.UnpackLog(event, "ExecuteNewEra", log); err != nil {
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

// ParseExecuteNewEra is a log parse operation binding the contract event 0x02105621fc31aa3ac04a9845beacd54c700e2ab23ff8acdd755dfd878ae61f02.
//
// Solidity: event ExecuteNewEra(uint256 indexed era, uint256 rate)
func (_StakeManager *StakeManagerFilterer) ParseExecuteNewEra(log types.Log) (*StakeManagerExecuteNewEra, error) {
	event := new(StakeManagerExecuteNewEra)
	if err := _StakeManager.contract.UnpackLog(event, "ExecuteNewEra", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the StakeManager contract.
type StakeManagerProposalExecutedIterator struct {
	Event *StakeManagerProposalExecuted // Event containing the contract specifics and raw log

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
func (it *StakeManagerProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerProposalExecuted)
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
		it.Event = new(StakeManagerProposalExecuted)
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
func (it *StakeManagerProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerProposalExecuted represents a ProposalExecuted event raised by the StakeManager contract.
type StakeManagerProposalExecuted struct {
	ProposalId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalId)
func (_StakeManager *StakeManagerFilterer) FilterProposalExecuted(opts *bind.FilterOpts, proposalId [][32]byte) (*StakeManagerProposalExecutedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerProposalExecutedIterator{contract: _StakeManager.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalId)
func (_StakeManager *StakeManagerFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *StakeManagerProposalExecuted, proposalId [][32]byte) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerProposalExecuted)
				if err := _StakeManager.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalId)
func (_StakeManager *StakeManagerFilterer) ParseProposalExecuted(log types.Log) (*StakeManagerProposalExecuted, error) {
	event := new(StakeManagerProposalExecuted)
	if err := _StakeManager.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerRepairDelegatedIterator is returned from FilterRepairDelegated and is used to iterate over the raw logs and unpacked data for RepairDelegated events raised by the StakeManager contract.
type StakeManagerRepairDelegatedIterator struct {
	Event *StakeManagerRepairDelegated // Event containing the contract specifics and raw log

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
func (it *StakeManagerRepairDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerRepairDelegated)
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
		it.Event = new(StakeManagerRepairDelegated)
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
func (it *StakeManagerRepairDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerRepairDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerRepairDelegated represents a RepairDelegated event raised by the StakeManager contract.
type StakeManagerRepairDelegated struct {
	GovDelegated   *big.Int
	LocalDelegated *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepairDelegated is a free log retrieval operation binding the contract event 0x5eb59c12a32152528b0ae50929843b3125858e1189a5867d0a34458ff728a14c.
//
// Solidity: event RepairDelegated(uint256 govDelegated, uint256 localDelegated)
func (_StakeManager *StakeManagerFilterer) FilterRepairDelegated(opts *bind.FilterOpts) (*StakeManagerRepairDelegatedIterator, error) {

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "RepairDelegated")
	if err != nil {
		return nil, err
	}
	return &StakeManagerRepairDelegatedIterator{contract: _StakeManager.contract, event: "RepairDelegated", logs: logs, sub: sub}, nil
}

// WatchRepairDelegated is a free log subscription operation binding the contract event 0x5eb59c12a32152528b0ae50929843b3125858e1189a5867d0a34458ff728a14c.
//
// Solidity: event RepairDelegated(uint256 govDelegated, uint256 localDelegated)
func (_StakeManager *StakeManagerFilterer) WatchRepairDelegated(opts *bind.WatchOpts, sink chan<- *StakeManagerRepairDelegated) (event.Subscription, error) {

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "RepairDelegated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerRepairDelegated)
				if err := _StakeManager.contract.UnpackLog(event, "RepairDelegated", log); err != nil {
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

// ParseRepairDelegated is a log parse operation binding the contract event 0x5eb59c12a32152528b0ae50929843b3125858e1189a5867d0a34458ff728a14c.
//
// Solidity: event RepairDelegated(uint256 govDelegated, uint256 localDelegated)
func (_StakeManager *StakeManagerFilterer) ParseRepairDelegated(log types.Log) (*StakeManagerRepairDelegated, error) {
	event := new(StakeManagerRepairDelegated)
	if err := _StakeManager.contract.UnpackLog(event, "RepairDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerSettleIterator is returned from FilterSettle and is used to iterate over the raw logs and unpacked data for Settle events raised by the StakeManager contract.
type StakeManagerSettleIterator struct {
	Event *StakeManagerSettle // Event containing the contract specifics and raw log

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
func (it *StakeManagerSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerSettle)
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
		it.Event = new(StakeManagerSettle)
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
func (it *StakeManagerSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerSettle represents a Settle event raised by the StakeManager contract.
type StakeManagerSettle struct {
	Era         *big.Int
	PoolAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSettle is a free log retrieval operation binding the contract event 0x5a208f0e8f17e3d4544100afcf65662217cc8fb06e47d9016a83e21a1f96a791.
//
// Solidity: event Settle(uint256 indexed era, address indexed poolAddress)
func (_StakeManager *StakeManagerFilterer) FilterSettle(opts *bind.FilterOpts, era []*big.Int, poolAddress []common.Address) (*StakeManagerSettleIterator, error) {

	var eraRule []interface{}
	for _, eraItem := range era {
		eraRule = append(eraRule, eraItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Settle", eraRule, poolAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerSettleIterator{contract: _StakeManager.contract, event: "Settle", logs: logs, sub: sub}, nil
}

// WatchSettle is a free log subscription operation binding the contract event 0x5a208f0e8f17e3d4544100afcf65662217cc8fb06e47d9016a83e21a1f96a791.
//
// Solidity: event Settle(uint256 indexed era, address indexed poolAddress)
func (_StakeManager *StakeManagerFilterer) WatchSettle(opts *bind.WatchOpts, sink chan<- *StakeManagerSettle, era []*big.Int, poolAddress []common.Address) (event.Subscription, error) {

	var eraRule []interface{}
	for _, eraItem := range era {
		eraRule = append(eraRule, eraItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Settle", eraRule, poolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerSettle)
				if err := _StakeManager.contract.UnpackLog(event, "Settle", log); err != nil {
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

// ParseSettle is a log parse operation binding the contract event 0x5a208f0e8f17e3d4544100afcf65662217cc8fb06e47d9016a83e21a1f96a791.
//
// Solidity: event Settle(uint256 indexed era, address indexed poolAddress)
func (_StakeManager *StakeManagerFilterer) ParseSettle(log types.Log) (*StakeManagerSettle, error) {
	event := new(StakeManagerSettle)
	if err := _StakeManager.contract.UnpackLog(event, "Settle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the StakeManager contract.
type StakeManagerStakeIterator struct {
	Event *StakeManagerStake // Event containing the contract specifics and raw log

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
func (it *StakeManagerStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerStake)
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
		it.Event = new(StakeManagerStake)
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
func (it *StakeManagerStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerStake represents a Stake event raised by the StakeManager contract.
type StakeManagerStake struct {
	Staker       common.Address
	StakePool    common.Address
	TokenAmount  *big.Int
	RTokenAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0x63602d0ecc7b3a0ef7ff1a116e23056662d64280355ba8031b6d0d767c4b4458.
//
// Solidity: event Stake(address staker, address stakePool, uint256 tokenAmount, uint256 rTokenAmount)
func (_StakeManager *StakeManagerFilterer) FilterStake(opts *bind.FilterOpts) (*StakeManagerStakeIterator, error) {

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Stake")
	if err != nil {
		return nil, err
	}
	return &StakeManagerStakeIterator{contract: _StakeManager.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0x63602d0ecc7b3a0ef7ff1a116e23056662d64280355ba8031b6d0d767c4b4458.
//
// Solidity: event Stake(address staker, address stakePool, uint256 tokenAmount, uint256 rTokenAmount)
func (_StakeManager *StakeManagerFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *StakeManagerStake) (event.Subscription, error) {

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Stake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerStake)
				if err := _StakeManager.contract.UnpackLog(event, "Stake", log); err != nil {
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

// ParseStake is a log parse operation binding the contract event 0x63602d0ecc7b3a0ef7ff1a116e23056662d64280355ba8031b6d0d767c4b4458.
//
// Solidity: event Stake(address staker, address stakePool, uint256 tokenAmount, uint256 rTokenAmount)
func (_StakeManager *StakeManagerFilterer) ParseStake(log types.Log) (*StakeManagerStake, error) {
	event := new(StakeManagerStake)
	if err := _StakeManager.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerUnstakeIterator is returned from FilterUnstake and is used to iterate over the raw logs and unpacked data for Unstake events raised by the StakeManager contract.
type StakeManagerUnstakeIterator struct {
	Event *StakeManagerUnstake // Event containing the contract specifics and raw log

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
func (it *StakeManagerUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerUnstake)
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
		it.Event = new(StakeManagerUnstake)
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
func (it *StakeManagerUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerUnstake represents a Unstake event raised by the StakeManager contract.
type StakeManagerUnstake struct {
	Staker       common.Address
	StakePool    common.Address
	TokenAmount  *big.Int
	RTokenAmount *big.Int
	BurnAmount   *big.Int
	UnstakeIndex *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnstake is a free log retrieval operation binding the contract event 0x4e5916c8cf4042e78e6a44e60bf6b48e8e969ce5abf8d4227613ddc3454015ea.
//
// Solidity: event Unstake(address staker, address stakePool, uint256 tokenAmount, uint256 rTokenAmount, uint256 burnAmount, uint256 unstakeIndex)
func (_StakeManager *StakeManagerFilterer) FilterUnstake(opts *bind.FilterOpts) (*StakeManagerUnstakeIterator, error) {

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Unstake")
	if err != nil {
		return nil, err
	}
	return &StakeManagerUnstakeIterator{contract: _StakeManager.contract, event: "Unstake", logs: logs, sub: sub}, nil
}

// WatchUnstake is a free log subscription operation binding the contract event 0x4e5916c8cf4042e78e6a44e60bf6b48e8e969ce5abf8d4227613ddc3454015ea.
//
// Solidity: event Unstake(address staker, address stakePool, uint256 tokenAmount, uint256 rTokenAmount, uint256 burnAmount, uint256 unstakeIndex)
func (_StakeManager *StakeManagerFilterer) WatchUnstake(opts *bind.WatchOpts, sink chan<- *StakeManagerUnstake) (event.Subscription, error) {

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Unstake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerUnstake)
				if err := _StakeManager.contract.UnpackLog(event, "Unstake", log); err != nil {
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

// ParseUnstake is a log parse operation binding the contract event 0x4e5916c8cf4042e78e6a44e60bf6b48e8e969ce5abf8d4227613ddc3454015ea.
//
// Solidity: event Unstake(address staker, address stakePool, uint256 tokenAmount, uint256 rTokenAmount, uint256 burnAmount, uint256 unstakeIndex)
func (_StakeManager *StakeManagerFilterer) ParseUnstake(log types.Log) (*StakeManagerUnstake, error) {
	event := new(StakeManagerUnstake)
	if err := _StakeManager.contract.UnpackLog(event, "Unstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the StakeManager contract.
type StakeManagerWithdrawIterator struct {
	Event *StakeManagerWithdraw // Event containing the contract specifics and raw log

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
func (it *StakeManagerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerWithdraw)
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
		it.Event = new(StakeManagerWithdraw)
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
func (it *StakeManagerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerWithdraw represents a Withdraw event raised by the StakeManager contract.
type StakeManagerWithdraw struct {
	Staker           common.Address
	StakePool        common.Address
	TokenAmount      *big.Int
	UnstakeIndexList []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xe4b7499d334dcb3a4338114f8df473bb4444d9cace993f8d2eb779921f074dd3.
//
// Solidity: event Withdraw(address staker, address stakePool, uint256 tokenAmount, uint256[] unstakeIndexList)
func (_StakeManager *StakeManagerFilterer) FilterWithdraw(opts *bind.FilterOpts) (*StakeManagerWithdrawIterator, error) {

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &StakeManagerWithdrawIterator{contract: _StakeManager.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xe4b7499d334dcb3a4338114f8df473bb4444d9cace993f8d2eb779921f074dd3.
//
// Solidity: event Withdraw(address staker, address stakePool, uint256 tokenAmount, uint256[] unstakeIndexList)
func (_StakeManager *StakeManagerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *StakeManagerWithdraw) (event.Subscription, error) {

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerWithdraw)
				if err := _StakeManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xe4b7499d334dcb3a4338114f8df473bb4444d9cace993f8d2eb779921f074dd3.
//
// Solidity: event Withdraw(address staker, address stakePool, uint256 tokenAmount, uint256[] unstakeIndexList)
func (_StakeManager *StakeManagerFilterer) ParseWithdraw(log types.Log) (*StakeManagerWithdraw, error) {
	event := new(StakeManagerWithdraw)
	if err := _StakeManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
