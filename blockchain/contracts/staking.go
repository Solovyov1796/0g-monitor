// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// Coin is an auto generated low-level Go binding around an user-defined struct.
type Coin struct {
	Denom  string
	Amount *big.Int
}

// CommissionRates is an auto generated low-level Go binding around an user-defined struct.
type CommissionRates struct {
	Rate          *big.Int
	MaxRate       *big.Int
	MaxChangeRate *big.Int
}

// Description is an auto generated low-level Go binding around an user-defined struct.
type Description struct {
	Moniker         string
	Identity        string
	Website         string
	SecurityContact string
	Details         string
}

// PageRequest is an auto generated low-level Go binding around an user-defined struct.
type PageRequest struct {
	Key        []byte
	Offset     uint64
	Limit      uint64
	CountTotal bool
	Reverse    bool
}

// PageResponse is an auto generated low-level Go binding around an user-defined struct.
type PageResponse struct {
	NextKey []byte
	Total   uint64
}

// Redelegation is an auto generated low-level Go binding around an user-defined struct.
type Redelegation struct {
	DelegatorAddress    string
	ValidatorSrcAddress string
	ValidatorDstAddress string
	Entries             []RedelegationEntry
}

// RedelegationEntry is an auto generated low-level Go binding around an user-defined struct.
type RedelegationEntry struct {
	CreationHeight int64
	CompletionTime int64
	InitialBalance *big.Int
	SharesDst      *big.Int
}

// RedelegationEntryResponse is an auto generated low-level Go binding around an user-defined struct.
type RedelegationEntryResponse struct {
	RedelegationEntry RedelegationEntry
	Balance           *big.Int
}

// RedelegationOutput is an auto generated low-level Go binding around an user-defined struct.
type RedelegationOutput struct {
	DelegatorAddress    string
	ValidatorSrcAddress string
	ValidatorDstAddress string
	Entries             []RedelegationEntry
}

// RedelegationResponse is an auto generated low-level Go binding around an user-defined struct.
type RedelegationResponse struct {
	Redelegation Redelegation
	Entries      []RedelegationEntryResponse
}

// UnbondingDelegationEntry is an auto generated low-level Go binding around an user-defined struct.
type UnbondingDelegationEntry struct {
	CreationHeight          int64
	CompletionTime          int64
	InitialBalance          *big.Int
	Balance                 *big.Int
	UnbondingId             uint64
	UnbondingOnHoldRefCount int64
}

// UnbondingDelegationOutput is an auto generated low-level Go binding around an user-defined struct.
type UnbondingDelegationOutput struct {
	DelegatorAddress string
	ValidatorAddress string
	Entries          []UnbondingDelegationEntry
}

// Validator is an auto generated low-level Go binding around an user-defined struct.
type Validator struct {
	OperatorAddress   string
	ConsensusPubkey   string
	Jailed            bool
	Status            uint8
	Tokens            *big.Int
	DelegatorShares   *big.Int
	Description       string
	UnbondingHeight   int64
	UnbondingTime     int64
	Commission        *big.Int
	MinSelfDelegation *big.Int
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"granter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"methods\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"AllowanceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"granter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"methods\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creationHeight\",\"type\":\"uint256\"}],\"name\":\"CancelUnbondingDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"CreateValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newShares\",\"type\":\"uint256\"}],\"name\":\"Delegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorSrcAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorDstAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"completionTime\",\"type\":\"uint256\"}],\"name\":\"Redelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"granter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"methods\",\"type\":\"string[]\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"completionTime\",\"type\":\"uint256\"}],\"name\":\"Unbond\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"granter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"methods\",\"type\":\"string[]\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationHeight\",\"type\":\"uint256\"}],\"name\":\"cancelUnbondingDelegation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"securityContact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChangeRate\",\"type\":\"uint256\"}],\"internalType\":\"structCommissionRates\",\"name\":\"commissionRates\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pubkey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"createValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"methods\",\"type\":\"string[]\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"}],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin\",\"name\":\"balance\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"methods\",\"type\":\"string[]\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"validatorSrcAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorDstAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redelegate\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"srcValidatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dstValidatorAddress\",\"type\":\"string\"}],\"name\":\"redelegation\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"delegatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorSrcAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorDstAddress\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sharesDst\",\"type\":\"uint256\"}],\"internalType\":\"structRedelegationEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structRedelegationOutput\",\"name\":\"redelegation\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"srcValidatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dstValidatorAddress\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pageRequest\",\"type\":\"tuple\"}],\"name\":\"redelegations\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"delegatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorSrcAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorDstAddress\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sharesDst\",\"type\":\"uint256\"}],\"internalType\":\"structRedelegationEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structRedelegation\",\"name\":\"redelegation\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sharesDst\",\"type\":\"uint256\"}],\"internalType\":\"structRedelegationEntry\",\"name\":\"redelegationEntry\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structRedelegationEntryResponse[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structRedelegationResponse[]\",\"name\":\"response\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"methods\",\"type\":\"string[]\"}],\"name\":\"revoke\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"}],\"name\":\"unbondingDelegation\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"delegatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"unbondingId\",\"type\":\"uint64\"},{\"internalType\":\"int64\",\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\"}],\"internalType\":\"structUnbondingDelegationEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structUnbondingDelegationOutput\",\"name\":\"unbondingDelegation\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"}],\"name\":\"validator\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"operatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"consensusPubkey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"enumBondStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorShares\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"unbondingHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"unbondingTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"}],\"internalType\":\"structValidator\",\"name\":\"validator\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pageRequest\",\"type\":\"tuple\"}],\"name\":\"validators\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"operatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"consensusPubkey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"enumBondStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorShares\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"unbondingHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"unbondingTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xfc08930c.
//
// Solidity: function allowance(address grantee, address granter, string method) view returns(uint256 remaining)
func (_Staking *StakingCaller) Allowance(opts *bind.CallOpts, grantee common.Address, granter common.Address, method string) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "allowance", grantee, granter, method)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xfc08930c.
//
// Solidity: function allowance(address grantee, address granter, string method) view returns(uint256 remaining)
func (_Staking *StakingSession) Allowance(grantee common.Address, granter common.Address, method string) (*big.Int, error) {
	return _Staking.Contract.Allowance(&_Staking.CallOpts, grantee, granter, method)
}

// Allowance is a free data retrieval call binding the contract method 0xfc08930c.
//
// Solidity: function allowance(address grantee, address granter, string method) view returns(uint256 remaining)
func (_Staking *StakingCallerSession) Allowance(grantee common.Address, granter common.Address, method string) (*big.Int, error) {
	return _Staking.Contract.Allowance(&_Staking.CallOpts, grantee, granter, method)
}

// Delegation is a free data retrieval call binding the contract method 0x241774e6.
//
// Solidity: function delegation(address delegatorAddress, string validatorAddress) view returns(uint256 shares, (string,uint256) balance)
func (_Staking *StakingCaller) Delegation(opts *bind.CallOpts, delegatorAddress common.Address, validatorAddress string) (struct {
	Shares  *big.Int
	Balance Coin
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "delegation", delegatorAddress, validatorAddress)

	outstruct := new(struct {
		Shares  *big.Int
		Balance Coin
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Shares = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Balance = *abi.ConvertType(out[1], new(Coin)).(*Coin)

	return *outstruct, err

}

// Delegation is a free data retrieval call binding the contract method 0x241774e6.
//
// Solidity: function delegation(address delegatorAddress, string validatorAddress) view returns(uint256 shares, (string,uint256) balance)
func (_Staking *StakingSession) Delegation(delegatorAddress common.Address, validatorAddress string) (struct {
	Shares  *big.Int
	Balance Coin
}, error) {
	return _Staking.Contract.Delegation(&_Staking.CallOpts, delegatorAddress, validatorAddress)
}

// Delegation is a free data retrieval call binding the contract method 0x241774e6.
//
// Solidity: function delegation(address delegatorAddress, string validatorAddress) view returns(uint256 shares, (string,uint256) balance)
func (_Staking *StakingCallerSession) Delegation(delegatorAddress common.Address, validatorAddress string) (struct {
	Shares  *big.Int
	Balance Coin
}, error) {
	return _Staking.Contract.Delegation(&_Staking.CallOpts, delegatorAddress, validatorAddress)
}

// Redelegation is a free data retrieval call binding the contract method 0x7d9f939c.
//
// Solidity: function redelegation(address delegatorAddress, string srcValidatorAddress, string dstValidatorAddress) view returns((string,string,string,(int64,int64,uint256,uint256)[]) redelegation)
func (_Staking *StakingCaller) Redelegation(opts *bind.CallOpts, delegatorAddress common.Address, srcValidatorAddress string, dstValidatorAddress string) (RedelegationOutput, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "redelegation", delegatorAddress, srcValidatorAddress, dstValidatorAddress)

	if err != nil {
		return *new(RedelegationOutput), err
	}

	out0 := *abi.ConvertType(out[0], new(RedelegationOutput)).(*RedelegationOutput)

	return out0, err

}

// Redelegation is a free data retrieval call binding the contract method 0x7d9f939c.
//
// Solidity: function redelegation(address delegatorAddress, string srcValidatorAddress, string dstValidatorAddress) view returns((string,string,string,(int64,int64,uint256,uint256)[]) redelegation)
func (_Staking *StakingSession) Redelegation(delegatorAddress common.Address, srcValidatorAddress string, dstValidatorAddress string) (RedelegationOutput, error) {
	return _Staking.Contract.Redelegation(&_Staking.CallOpts, delegatorAddress, srcValidatorAddress, dstValidatorAddress)
}

// Redelegation is a free data retrieval call binding the contract method 0x7d9f939c.
//
// Solidity: function redelegation(address delegatorAddress, string srcValidatorAddress, string dstValidatorAddress) view returns((string,string,string,(int64,int64,uint256,uint256)[]) redelegation)
func (_Staking *StakingCallerSession) Redelegation(delegatorAddress common.Address, srcValidatorAddress string, dstValidatorAddress string) (RedelegationOutput, error) {
	return _Staking.Contract.Redelegation(&_Staking.CallOpts, delegatorAddress, srcValidatorAddress, dstValidatorAddress)
}

// Redelegations is a free data retrieval call binding the contract method 0x10a2851c.
//
// Solidity: function redelegations(address delegatorAddress, string srcValidatorAddress, string dstValidatorAddress, (bytes,uint64,uint64,bool,bool) pageRequest) view returns(((string,string,string,(int64,int64,uint256,uint256)[]),((int64,int64,uint256,uint256),uint256)[])[] response, (bytes,uint64) pageResponse)
func (_Staking *StakingCaller) Redelegations(opts *bind.CallOpts, delegatorAddress common.Address, srcValidatorAddress string, dstValidatorAddress string, pageRequest PageRequest) (struct {
	Response     []RedelegationResponse
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "redelegations", delegatorAddress, srcValidatorAddress, dstValidatorAddress, pageRequest)

	outstruct := new(struct {
		Response     []RedelegationResponse
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Response = *abi.ConvertType(out[0], new([]RedelegationResponse)).(*[]RedelegationResponse)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// Redelegations is a free data retrieval call binding the contract method 0x10a2851c.
//
// Solidity: function redelegations(address delegatorAddress, string srcValidatorAddress, string dstValidatorAddress, (bytes,uint64,uint64,bool,bool) pageRequest) view returns(((string,string,string,(int64,int64,uint256,uint256)[]),((int64,int64,uint256,uint256),uint256)[])[] response, (bytes,uint64) pageResponse)
func (_Staking *StakingSession) Redelegations(delegatorAddress common.Address, srcValidatorAddress string, dstValidatorAddress string, pageRequest PageRequest) (struct {
	Response     []RedelegationResponse
	PageResponse PageResponse
}, error) {
	return _Staking.Contract.Redelegations(&_Staking.CallOpts, delegatorAddress, srcValidatorAddress, dstValidatorAddress, pageRequest)
}

// Redelegations is a free data retrieval call binding the contract method 0x10a2851c.
//
// Solidity: function redelegations(address delegatorAddress, string srcValidatorAddress, string dstValidatorAddress, (bytes,uint64,uint64,bool,bool) pageRequest) view returns(((string,string,string,(int64,int64,uint256,uint256)[]),((int64,int64,uint256,uint256),uint256)[])[] response, (bytes,uint64) pageResponse)
func (_Staking *StakingCallerSession) Redelegations(delegatorAddress common.Address, srcValidatorAddress string, dstValidatorAddress string, pageRequest PageRequest) (struct {
	Response     []RedelegationResponse
	PageResponse PageResponse
}, error) {
	return _Staking.Contract.Redelegations(&_Staking.CallOpts, delegatorAddress, srcValidatorAddress, dstValidatorAddress, pageRequest)
}

// UnbondingDelegation is a free data retrieval call binding the contract method 0xa03ffee1.
//
// Solidity: function unbondingDelegation(address delegatorAddress, string validatorAddress) view returns((string,string,(int64,int64,uint256,uint256,uint64,int64)[]) unbondingDelegation)
func (_Staking *StakingCaller) UnbondingDelegation(opts *bind.CallOpts, delegatorAddress common.Address, validatorAddress string) (UnbondingDelegationOutput, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "unbondingDelegation", delegatorAddress, validatorAddress)

	if err != nil {
		return *new(UnbondingDelegationOutput), err
	}

	out0 := *abi.ConvertType(out[0], new(UnbondingDelegationOutput)).(*UnbondingDelegationOutput)

	return out0, err

}

// UnbondingDelegation is a free data retrieval call binding the contract method 0xa03ffee1.
//
// Solidity: function unbondingDelegation(address delegatorAddress, string validatorAddress) view returns((string,string,(int64,int64,uint256,uint256,uint64,int64)[]) unbondingDelegation)
func (_Staking *StakingSession) UnbondingDelegation(delegatorAddress common.Address, validatorAddress string) (UnbondingDelegationOutput, error) {
	return _Staking.Contract.UnbondingDelegation(&_Staking.CallOpts, delegatorAddress, validatorAddress)
}

// UnbondingDelegation is a free data retrieval call binding the contract method 0xa03ffee1.
//
// Solidity: function unbondingDelegation(address delegatorAddress, string validatorAddress) view returns((string,string,(int64,int64,uint256,uint256,uint64,int64)[]) unbondingDelegation)
func (_Staking *StakingCallerSession) UnbondingDelegation(delegatorAddress common.Address, validatorAddress string) (UnbondingDelegationOutput, error) {
	return _Staking.Contract.UnbondingDelegation(&_Staking.CallOpts, delegatorAddress, validatorAddress)
}

// Validator is a free data retrieval call binding the contract method 0x0bc82a17.
//
// Solidity: function validator(string validatorAddress) view returns((string,string,bool,uint8,uint256,uint256,string,int64,int64,uint256,uint256) validator)
func (_Staking *StakingCaller) Validator(opts *bind.CallOpts, validatorAddress string) (Validator, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "validator", validatorAddress)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x0bc82a17.
//
// Solidity: function validator(string validatorAddress) view returns((string,string,bool,uint8,uint256,uint256,string,int64,int64,uint256,uint256) validator)
func (_Staking *StakingSession) Validator(validatorAddress string) (Validator, error) {
	return _Staking.Contract.Validator(&_Staking.CallOpts, validatorAddress)
}

// Validator is a free data retrieval call binding the contract method 0x0bc82a17.
//
// Solidity: function validator(string validatorAddress) view returns((string,string,bool,uint8,uint256,uint256,string,int64,int64,uint256,uint256) validator)
func (_Staking *StakingCallerSession) Validator(validatorAddress string) (Validator, error) {
	return _Staking.Contract.Validator(&_Staking.CallOpts, validatorAddress)
}

// Validators is a free data retrieval call binding the contract method 0x186b2167.
//
// Solidity: function validators(string status, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,string,bool,uint8,uint256,uint256,string,int64,int64,uint256,uint256)[] validators, (bytes,uint64) pageResponse)
func (_Staking *StakingCaller) Validators(opts *bind.CallOpts, status string, pageRequest PageRequest) (struct {
	Validators   []Validator
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "validators", status, pageRequest)

	outstruct := new(struct {
		Validators   []Validator
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]Validator)).(*[]Validator)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0x186b2167.
//
// Solidity: function validators(string status, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,string,bool,uint8,uint256,uint256,string,int64,int64,uint256,uint256)[] validators, (bytes,uint64) pageResponse)
func (_Staking *StakingSession) Validators(status string, pageRequest PageRequest) (struct {
	Validators   []Validator
	PageResponse PageResponse
}, error) {
	return _Staking.Contract.Validators(&_Staking.CallOpts, status, pageRequest)
}

// Validators is a free data retrieval call binding the contract method 0x186b2167.
//
// Solidity: function validators(string status, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,string,bool,uint8,uint256,uint256,string,int64,int64,uint256,uint256)[] validators, (bytes,uint64) pageResponse)
func (_Staking *StakingCallerSession) Validators(status string, pageRequest PageRequest) (struct {
	Validators   []Validator
	PageResponse PageResponse
}, error) {
	return _Staking.Contract.Validators(&_Staking.CallOpts, status, pageRequest)
}

// Approve is a paid mutator transaction binding the contract method 0xb6039895.
//
// Solidity: function approve(address grantee, uint256 amount, string[] methods) returns(bool approved)
func (_Staking *StakingTransactor) Approve(opts *bind.TransactOpts, grantee common.Address, amount *big.Int, methods []string) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "approve", grantee, amount, methods)
}

// Approve is a paid mutator transaction binding the contract method 0xb6039895.
//
// Solidity: function approve(address grantee, uint256 amount, string[] methods) returns(bool approved)
func (_Staking *StakingSession) Approve(grantee common.Address, amount *big.Int, methods []string) (*types.Transaction, error) {
	return _Staking.Contract.Approve(&_Staking.TransactOpts, grantee, amount, methods)
}

// Approve is a paid mutator transaction binding the contract method 0xb6039895.
//
// Solidity: function approve(address grantee, uint256 amount, string[] methods) returns(bool approved)
func (_Staking *StakingTransactorSession) Approve(grantee common.Address, amount *big.Int, methods []string) (*types.Transaction, error) {
	return _Staking.Contract.Approve(&_Staking.TransactOpts, grantee, amount, methods)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x12d58dfe.
//
// Solidity: function cancelUnbondingDelegation(address delegatorAddress, string validatorAddress, uint256 amount, uint256 creationHeight) returns(bool success)
func (_Staking *StakingTransactor) CancelUnbondingDelegation(opts *bind.TransactOpts, delegatorAddress common.Address, validatorAddress string, amount *big.Int, creationHeight *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "cancelUnbondingDelegation", delegatorAddress, validatorAddress, amount, creationHeight)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x12d58dfe.
//
// Solidity: function cancelUnbondingDelegation(address delegatorAddress, string validatorAddress, uint256 amount, uint256 creationHeight) returns(bool success)
func (_Staking *StakingSession) CancelUnbondingDelegation(delegatorAddress common.Address, validatorAddress string, amount *big.Int, creationHeight *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.CancelUnbondingDelegation(&_Staking.TransactOpts, delegatorAddress, validatorAddress, amount, creationHeight)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x12d58dfe.
//
// Solidity: function cancelUnbondingDelegation(address delegatorAddress, string validatorAddress, uint256 amount, uint256 creationHeight) returns(bool success)
func (_Staking *StakingTransactorSession) CancelUnbondingDelegation(delegatorAddress common.Address, validatorAddress string, amount *big.Int, creationHeight *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.CancelUnbondingDelegation(&_Staking.TransactOpts, delegatorAddress, validatorAddress, amount, creationHeight)
}

// CreateValidator is a paid mutator transaction binding the contract method 0x45818645.
//
// Solidity: function createValidator((string,string,string,string,string) description, (uint256,uint256,uint256) commissionRates, uint256 minSelfDelegation, address delegatorAddress, string validatorAddress, string pubkey, uint256 value) returns(bool success)
func (_Staking *StakingTransactor) CreateValidator(opts *bind.TransactOpts, description Description, commissionRates CommissionRates, minSelfDelegation *big.Int, delegatorAddress common.Address, validatorAddress string, pubkey string, value *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "createValidator", description, commissionRates, minSelfDelegation, delegatorAddress, validatorAddress, pubkey, value)
}

// CreateValidator is a paid mutator transaction binding the contract method 0x45818645.
//
// Solidity: function createValidator((string,string,string,string,string) description, (uint256,uint256,uint256) commissionRates, uint256 minSelfDelegation, address delegatorAddress, string validatorAddress, string pubkey, uint256 value) returns(bool success)
func (_Staking *StakingSession) CreateValidator(description Description, commissionRates CommissionRates, minSelfDelegation *big.Int, delegatorAddress common.Address, validatorAddress string, pubkey string, value *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.CreateValidator(&_Staking.TransactOpts, description, commissionRates, minSelfDelegation, delegatorAddress, validatorAddress, pubkey, value)
}

// CreateValidator is a paid mutator transaction binding the contract method 0x45818645.
//
// Solidity: function createValidator((string,string,string,string,string) description, (uint256,uint256,uint256) commissionRates, uint256 minSelfDelegation, address delegatorAddress, string validatorAddress, string pubkey, uint256 value) returns(bool success)
func (_Staking *StakingTransactorSession) CreateValidator(description Description, commissionRates CommissionRates, minSelfDelegation *big.Int, delegatorAddress common.Address, validatorAddress string, pubkey string, value *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.CreateValidator(&_Staking.TransactOpts, description, commissionRates, minSelfDelegation, delegatorAddress, validatorAddress, pubkey, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xf007d286.
//
// Solidity: function decreaseAllowance(address grantee, uint256 amount, string[] methods) returns(bool approved)
func (_Staking *StakingTransactor) DecreaseAllowance(opts *bind.TransactOpts, grantee common.Address, amount *big.Int, methods []string) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "decreaseAllowance", grantee, amount, methods)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xf007d286.
//
// Solidity: function decreaseAllowance(address grantee, uint256 amount, string[] methods) returns(bool approved)
func (_Staking *StakingSession) DecreaseAllowance(grantee common.Address, amount *big.Int, methods []string) (*types.Transaction, error) {
	return _Staking.Contract.DecreaseAllowance(&_Staking.TransactOpts, grantee, amount, methods)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xf007d286.
//
// Solidity: function decreaseAllowance(address grantee, uint256 amount, string[] methods) returns(bool approved)
func (_Staking *StakingTransactorSession) DecreaseAllowance(grantee common.Address, amount *big.Int, methods []string) (*types.Transaction, error) {
	return _Staking.Contract.DecreaseAllowance(&_Staking.TransactOpts, grantee, amount, methods)
}

// Delegate is a paid mutator transaction binding the contract method 0x53266bbb.
//
// Solidity: function delegate(address delegatorAddress, string validatorAddress, uint256 amount) returns(bool success)
func (_Staking *StakingTransactor) Delegate(opts *bind.TransactOpts, delegatorAddress common.Address, validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "delegate", delegatorAddress, validatorAddress, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x53266bbb.
//
// Solidity: function delegate(address delegatorAddress, string validatorAddress, uint256 amount) returns(bool success)
func (_Staking *StakingSession) Delegate(delegatorAddress common.Address, validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, delegatorAddress, validatorAddress, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x53266bbb.
//
// Solidity: function delegate(address delegatorAddress, string validatorAddress, uint256 amount) returns(bool success)
func (_Staking *StakingTransactorSession) Delegate(delegatorAddress common.Address, validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, delegatorAddress, validatorAddress, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0xa386a63c.
//
// Solidity: function increaseAllowance(address grantee, uint256 amount, string[] methods) returns(bool approved)
func (_Staking *StakingTransactor) IncreaseAllowance(opts *bind.TransactOpts, grantee common.Address, amount *big.Int, methods []string) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "increaseAllowance", grantee, amount, methods)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0xa386a63c.
//
// Solidity: function increaseAllowance(address grantee, uint256 amount, string[] methods) returns(bool approved)
func (_Staking *StakingSession) IncreaseAllowance(grantee common.Address, amount *big.Int, methods []string) (*types.Transaction, error) {
	return _Staking.Contract.IncreaseAllowance(&_Staking.TransactOpts, grantee, amount, methods)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0xa386a63c.
//
// Solidity: function increaseAllowance(address grantee, uint256 amount, string[] methods) returns(bool approved)
func (_Staking *StakingTransactorSession) IncreaseAllowance(grantee common.Address, amount *big.Int, methods []string) (*types.Transaction, error) {
	return _Staking.Contract.IncreaseAllowance(&_Staking.TransactOpts, grantee, amount, methods)
}

// Redelegate is a paid mutator transaction binding the contract method 0x54b826f5.
//
// Solidity: function redelegate(address delegatorAddress, string validatorSrcAddress, string validatorDstAddress, uint256 amount) returns(int64 completionTime)
func (_Staking *StakingTransactor) Redelegate(opts *bind.TransactOpts, delegatorAddress common.Address, validatorSrcAddress string, validatorDstAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "redelegate", delegatorAddress, validatorSrcAddress, validatorDstAddress, amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x54b826f5.
//
// Solidity: function redelegate(address delegatorAddress, string validatorSrcAddress, string validatorDstAddress, uint256 amount) returns(int64 completionTime)
func (_Staking *StakingSession) Redelegate(delegatorAddress common.Address, validatorSrcAddress string, validatorDstAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Redelegate(&_Staking.TransactOpts, delegatorAddress, validatorSrcAddress, validatorDstAddress, amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x54b826f5.
//
// Solidity: function redelegate(address delegatorAddress, string validatorSrcAddress, string validatorDstAddress, uint256 amount) returns(int64 completionTime)
func (_Staking *StakingTransactorSession) Redelegate(delegatorAddress common.Address, validatorSrcAddress string, validatorDstAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Redelegate(&_Staking.TransactOpts, delegatorAddress, validatorSrcAddress, validatorDstAddress, amount)
}

// Revoke is a paid mutator transaction binding the contract method 0x61dc5c3b.
//
// Solidity: function revoke(address grantee, string[] methods) returns(bool revoked)
func (_Staking *StakingTransactor) Revoke(opts *bind.TransactOpts, grantee common.Address, methods []string) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "revoke", grantee, methods)
}

// Revoke is a paid mutator transaction binding the contract method 0x61dc5c3b.
//
// Solidity: function revoke(address grantee, string[] methods) returns(bool revoked)
func (_Staking *StakingSession) Revoke(grantee common.Address, methods []string) (*types.Transaction, error) {
	return _Staking.Contract.Revoke(&_Staking.TransactOpts, grantee, methods)
}

// Revoke is a paid mutator transaction binding the contract method 0x61dc5c3b.
//
// Solidity: function revoke(address grantee, string[] methods) returns(bool revoked)
func (_Staking *StakingTransactorSession) Revoke(grantee common.Address, methods []string) (*types.Transaction, error) {
	return _Staking.Contract.Revoke(&_Staking.TransactOpts, grantee, methods)
}

// Undelegate is a paid mutator transaction binding the contract method 0x3edab33c.
//
// Solidity: function undelegate(address delegatorAddress, string validatorAddress, uint256 amount) returns(int64 completionTime)
func (_Staking *StakingTransactor) Undelegate(opts *bind.TransactOpts, delegatorAddress common.Address, validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "undelegate", delegatorAddress, validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x3edab33c.
//
// Solidity: function undelegate(address delegatorAddress, string validatorAddress, uint256 amount) returns(int64 completionTime)
func (_Staking *StakingSession) Undelegate(delegatorAddress common.Address, validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Undelegate(&_Staking.TransactOpts, delegatorAddress, validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x3edab33c.
//
// Solidity: function undelegate(address delegatorAddress, string validatorAddress, uint256 amount) returns(int64 completionTime)
func (_Staking *StakingTransactorSession) Undelegate(delegatorAddress common.Address, validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Undelegate(&_Staking.TransactOpts, delegatorAddress, validatorAddress, amount)
}

// StakingAllowanceChangeIterator is returned from FilterAllowanceChange and is used to iterate over the raw logs and unpacked data for AllowanceChange events raised by the Staking contract.
type StakingAllowanceChangeIterator struct {
	Event *StakingAllowanceChange // Event containing the contract specifics and raw log

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
func (it *StakingAllowanceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingAllowanceChange)
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
		it.Event = new(StakingAllowanceChange)
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
func (it *StakingAllowanceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingAllowanceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingAllowanceChange represents a AllowanceChange event raised by the Staking contract.
type StakingAllowanceChange struct {
	Grantee common.Address
	Granter common.Address
	Methods []string
	Values  []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAllowanceChange is a free log retrieval operation binding the contract event 0x5a22c7e8af595d94a6d652de8e346a8ecdfe49fc2e0e976f33c9fc9358390ea4.
//
// Solidity: event AllowanceChange(address indexed grantee, address indexed granter, string[] methods, uint256[] values)
func (_Staking *StakingFilterer) FilterAllowanceChange(opts *bind.FilterOpts, grantee []common.Address, granter []common.Address) (*StakingAllowanceChangeIterator, error) {

	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}
	var granterRule []interface{}
	for _, granterItem := range granter {
		granterRule = append(granterRule, granterItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "AllowanceChange", granteeRule, granterRule)
	if err != nil {
		return nil, err
	}
	return &StakingAllowanceChangeIterator{contract: _Staking.contract, event: "AllowanceChange", logs: logs, sub: sub}, nil
}

// WatchAllowanceChange is a free log subscription operation binding the contract event 0x5a22c7e8af595d94a6d652de8e346a8ecdfe49fc2e0e976f33c9fc9358390ea4.
//
// Solidity: event AllowanceChange(address indexed grantee, address indexed granter, string[] methods, uint256[] values)
func (_Staking *StakingFilterer) WatchAllowanceChange(opts *bind.WatchOpts, sink chan<- *StakingAllowanceChange, grantee []common.Address, granter []common.Address) (event.Subscription, error) {

	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}
	var granterRule []interface{}
	for _, granterItem := range granter {
		granterRule = append(granterRule, granterItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "AllowanceChange", granteeRule, granterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingAllowanceChange)
				if err := _Staking.contract.UnpackLog(event, "AllowanceChange", log); err != nil {
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

// ParseAllowanceChange is a log parse operation binding the contract event 0x5a22c7e8af595d94a6d652de8e346a8ecdfe49fc2e0e976f33c9fc9358390ea4.
//
// Solidity: event AllowanceChange(address indexed grantee, address indexed granter, string[] methods, uint256[] values)
func (_Staking *StakingFilterer) ParseAllowanceChange(log types.Log) (*StakingAllowanceChange, error) {
	event := new(StakingAllowanceChange)
	if err := _Staking.contract.UnpackLog(event, "AllowanceChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Staking contract.
type StakingApprovalIterator struct {
	Event *StakingApproval // Event containing the contract specifics and raw log

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
func (it *StakingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingApproval)
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
		it.Event = new(StakingApproval)
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
func (it *StakingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingApproval represents a Approval event raised by the Staking contract.
type StakingApproval struct {
	Grantee common.Address
	Granter common.Address
	Methods []string
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xf2638649a77447a76543b3e27c31ee0febe7de7fb20e2b6a781d08207bc7cb8d.
//
// Solidity: event Approval(address indexed grantee, address indexed granter, string[] methods, uint256 value)
func (_Staking *StakingFilterer) FilterApproval(opts *bind.FilterOpts, grantee []common.Address, granter []common.Address) (*StakingApprovalIterator, error) {

	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}
	var granterRule []interface{}
	for _, granterItem := range granter {
		granterRule = append(granterRule, granterItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Approval", granteeRule, granterRule)
	if err != nil {
		return nil, err
	}
	return &StakingApprovalIterator{contract: _Staking.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xf2638649a77447a76543b3e27c31ee0febe7de7fb20e2b6a781d08207bc7cb8d.
//
// Solidity: event Approval(address indexed grantee, address indexed granter, string[] methods, uint256 value)
func (_Staking *StakingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StakingApproval, grantee []common.Address, granter []common.Address) (event.Subscription, error) {

	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}
	var granterRule []interface{}
	for _, granterItem := range granter {
		granterRule = append(granterRule, granterItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Approval", granteeRule, granterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingApproval)
				if err := _Staking.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0xf2638649a77447a76543b3e27c31ee0febe7de7fb20e2b6a781d08207bc7cb8d.
//
// Solidity: event Approval(address indexed grantee, address indexed granter, string[] methods, uint256 value)
func (_Staking *StakingFilterer) ParseApproval(log types.Log) (*StakingApproval, error) {
	event := new(StakingApproval)
	if err := _Staking.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingCancelUnbondingDelegationIterator is returned from FilterCancelUnbondingDelegation and is used to iterate over the raw logs and unpacked data for CancelUnbondingDelegation events raised by the Staking contract.
type StakingCancelUnbondingDelegationIterator struct {
	Event *StakingCancelUnbondingDelegation // Event containing the contract specifics and raw log

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
func (it *StakingCancelUnbondingDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCancelUnbondingDelegation)
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
		it.Event = new(StakingCancelUnbondingDelegation)
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
func (it *StakingCancelUnbondingDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCancelUnbondingDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCancelUnbondingDelegation represents a CancelUnbondingDelegation event raised by the Staking contract.
type StakingCancelUnbondingDelegation struct {
	DelegatorAddress common.Address
	ValidatorAddress common.Address
	Amount           *big.Int
	CreationHeight   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCancelUnbondingDelegation is a free log retrieval operation binding the contract event 0x6dbe2fb6b2613bdd8e3d284a6111592e06c3ab0af846ff89b6688d48f408dbb5.
//
// Solidity: event CancelUnbondingDelegation(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 creationHeight)
func (_Staking *StakingFilterer) FilterCancelUnbondingDelegation(opts *bind.FilterOpts, delegatorAddress []common.Address, validatorAddress []common.Address) (*StakingCancelUnbondingDelegationIterator, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "CancelUnbondingDelegation", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingCancelUnbondingDelegationIterator{contract: _Staking.contract, event: "CancelUnbondingDelegation", logs: logs, sub: sub}, nil
}

// WatchCancelUnbondingDelegation is a free log subscription operation binding the contract event 0x6dbe2fb6b2613bdd8e3d284a6111592e06c3ab0af846ff89b6688d48f408dbb5.
//
// Solidity: event CancelUnbondingDelegation(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 creationHeight)
func (_Staking *StakingFilterer) WatchCancelUnbondingDelegation(opts *bind.WatchOpts, sink chan<- *StakingCancelUnbondingDelegation, delegatorAddress []common.Address, validatorAddress []common.Address) (event.Subscription, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "CancelUnbondingDelegation", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCancelUnbondingDelegation)
				if err := _Staking.contract.UnpackLog(event, "CancelUnbondingDelegation", log); err != nil {
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

// ParseCancelUnbondingDelegation is a log parse operation binding the contract event 0x6dbe2fb6b2613bdd8e3d284a6111592e06c3ab0af846ff89b6688d48f408dbb5.
//
// Solidity: event CancelUnbondingDelegation(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 creationHeight)
func (_Staking *StakingFilterer) ParseCancelUnbondingDelegation(log types.Log) (*StakingCancelUnbondingDelegation, error) {
	event := new(StakingCancelUnbondingDelegation)
	if err := _Staking.contract.UnpackLog(event, "CancelUnbondingDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingCreateValidatorIterator is returned from FilterCreateValidator and is used to iterate over the raw logs and unpacked data for CreateValidator events raised by the Staking contract.
type StakingCreateValidatorIterator struct {
	Event *StakingCreateValidator // Event containing the contract specifics and raw log

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
func (it *StakingCreateValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCreateValidator)
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
		it.Event = new(StakingCreateValidator)
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
func (it *StakingCreateValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCreateValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCreateValidator represents a CreateValidator event raised by the Staking contract.
type StakingCreateValidator struct {
	DelegatorAddress common.Address
	ValidatorAddress common.Address
	Value            *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCreateValidator is a free log retrieval operation binding the contract event 0xc6b9ebca1d8f3f53e580dec2d3e8c7c6152ce1b2157ddb6f22c7868c6a38a0ea.
//
// Solidity: event CreateValidator(address indexed delegatorAddress, address indexed validatorAddress, uint256 value)
func (_Staking *StakingFilterer) FilterCreateValidator(opts *bind.FilterOpts, delegatorAddress []common.Address, validatorAddress []common.Address) (*StakingCreateValidatorIterator, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "CreateValidator", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingCreateValidatorIterator{contract: _Staking.contract, event: "CreateValidator", logs: logs, sub: sub}, nil
}

// WatchCreateValidator is a free log subscription operation binding the contract event 0xc6b9ebca1d8f3f53e580dec2d3e8c7c6152ce1b2157ddb6f22c7868c6a38a0ea.
//
// Solidity: event CreateValidator(address indexed delegatorAddress, address indexed validatorAddress, uint256 value)
func (_Staking *StakingFilterer) WatchCreateValidator(opts *bind.WatchOpts, sink chan<- *StakingCreateValidator, delegatorAddress []common.Address, validatorAddress []common.Address) (event.Subscription, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "CreateValidator", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCreateValidator)
				if err := _Staking.contract.UnpackLog(event, "CreateValidator", log); err != nil {
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

// ParseCreateValidator is a log parse operation binding the contract event 0xc6b9ebca1d8f3f53e580dec2d3e8c7c6152ce1b2157ddb6f22c7868c6a38a0ea.
//
// Solidity: event CreateValidator(address indexed delegatorAddress, address indexed validatorAddress, uint256 value)
func (_Staking *StakingFilterer) ParseCreateValidator(log types.Log) (*StakingCreateValidator, error) {
	event := new(StakingCreateValidator)
	if err := _Staking.contract.UnpackLog(event, "CreateValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDelegateIterator is returned from FilterDelegate and is used to iterate over the raw logs and unpacked data for Delegate events raised by the Staking contract.
type StakingDelegateIterator struct {
	Event *StakingDelegate // Event containing the contract specifics and raw log

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
func (it *StakingDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegate)
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
		it.Event = new(StakingDelegate)
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
func (it *StakingDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegate represents a Delegate event raised by the Staking contract.
type StakingDelegate struct {
	DelegatorAddress common.Address
	ValidatorAddress common.Address
	Amount           *big.Int
	NewShares        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDelegate is a free log retrieval operation binding the contract event 0x500599802164a08023e87ffc3eed0ba3ae60697b3083ba81d046683679d81c6b.
//
// Solidity: event Delegate(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 newShares)
func (_Staking *StakingFilterer) FilterDelegate(opts *bind.FilterOpts, delegatorAddress []common.Address, validatorAddress []common.Address) (*StakingDelegateIterator, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Delegate", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegateIterator{contract: _Staking.contract, event: "Delegate", logs: logs, sub: sub}, nil
}

// WatchDelegate is a free log subscription operation binding the contract event 0x500599802164a08023e87ffc3eed0ba3ae60697b3083ba81d046683679d81c6b.
//
// Solidity: event Delegate(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 newShares)
func (_Staking *StakingFilterer) WatchDelegate(opts *bind.WatchOpts, sink chan<- *StakingDelegate, delegatorAddress []common.Address, validatorAddress []common.Address) (event.Subscription, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Delegate", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegate)
				if err := _Staking.contract.UnpackLog(event, "Delegate", log); err != nil {
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

// ParseDelegate is a log parse operation binding the contract event 0x500599802164a08023e87ffc3eed0ba3ae60697b3083ba81d046683679d81c6b.
//
// Solidity: event Delegate(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 newShares)
func (_Staking *StakingFilterer) ParseDelegate(log types.Log) (*StakingDelegate, error) {
	event := new(StakingDelegate)
	if err := _Staking.contract.UnpackLog(event, "Delegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRedelegateIterator is returned from FilterRedelegate and is used to iterate over the raw logs and unpacked data for Redelegate events raised by the Staking contract.
type StakingRedelegateIterator struct {
	Event *StakingRedelegate // Event containing the contract specifics and raw log

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
func (it *StakingRedelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRedelegate)
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
		it.Event = new(StakingRedelegate)
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
func (it *StakingRedelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRedelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRedelegate represents a Redelegate event raised by the Staking contract.
type StakingRedelegate struct {
	DelegatorAddress    common.Address
	ValidatorSrcAddress common.Address
	ValidatorDstAddress common.Address
	Amount              *big.Int
	CompletionTime      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRedelegate is a free log retrieval operation binding the contract event 0x82b07f2421474f1e3f1e0b34738cb5ffb925273f408e7591d9c803dcae8da657.
//
// Solidity: event Redelegate(address indexed delegatorAddress, address indexed validatorSrcAddress, address indexed validatorDstAddress, uint256 amount, uint256 completionTime)
func (_Staking *StakingFilterer) FilterRedelegate(opts *bind.FilterOpts, delegatorAddress []common.Address, validatorSrcAddress []common.Address, validatorDstAddress []common.Address) (*StakingRedelegateIterator, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorSrcAddressRule []interface{}
	for _, validatorSrcAddressItem := range validatorSrcAddress {
		validatorSrcAddressRule = append(validatorSrcAddressRule, validatorSrcAddressItem)
	}
	var validatorDstAddressRule []interface{}
	for _, validatorDstAddressItem := range validatorDstAddress {
		validatorDstAddressRule = append(validatorDstAddressRule, validatorDstAddressItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Redelegate", delegatorAddressRule, validatorSrcAddressRule, validatorDstAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingRedelegateIterator{contract: _Staking.contract, event: "Redelegate", logs: logs, sub: sub}, nil
}

// WatchRedelegate is a free log subscription operation binding the contract event 0x82b07f2421474f1e3f1e0b34738cb5ffb925273f408e7591d9c803dcae8da657.
//
// Solidity: event Redelegate(address indexed delegatorAddress, address indexed validatorSrcAddress, address indexed validatorDstAddress, uint256 amount, uint256 completionTime)
func (_Staking *StakingFilterer) WatchRedelegate(opts *bind.WatchOpts, sink chan<- *StakingRedelegate, delegatorAddress []common.Address, validatorSrcAddress []common.Address, validatorDstAddress []common.Address) (event.Subscription, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorSrcAddressRule []interface{}
	for _, validatorSrcAddressItem := range validatorSrcAddress {
		validatorSrcAddressRule = append(validatorSrcAddressRule, validatorSrcAddressItem)
	}
	var validatorDstAddressRule []interface{}
	for _, validatorDstAddressItem := range validatorDstAddress {
		validatorDstAddressRule = append(validatorDstAddressRule, validatorDstAddressItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Redelegate", delegatorAddressRule, validatorSrcAddressRule, validatorDstAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRedelegate)
				if err := _Staking.contract.UnpackLog(event, "Redelegate", log); err != nil {
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

// ParseRedelegate is a log parse operation binding the contract event 0x82b07f2421474f1e3f1e0b34738cb5ffb925273f408e7591d9c803dcae8da657.
//
// Solidity: event Redelegate(address indexed delegatorAddress, address indexed validatorSrcAddress, address indexed validatorDstAddress, uint256 amount, uint256 completionTime)
func (_Staking *StakingFilterer) ParseRedelegate(log types.Log) (*StakingRedelegate, error) {
	event := new(StakingRedelegate)
	if err := _Staking.contract.UnpackLog(event, "Redelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the Staking contract.
type StakingRevocationIterator struct {
	Event *StakingRevocation // Event containing the contract specifics and raw log

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
func (it *StakingRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRevocation)
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
		it.Event = new(StakingRevocation)
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
func (it *StakingRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRevocation represents a Revocation event raised by the Staking contract.
type StakingRevocation struct {
	Grantee common.Address
	Granter common.Address
	Methods []string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xb0901d422521d0496e60bfbd8023b219d603a6cb950b43b2fe95043676cb353e.
//
// Solidity: event Revocation(address indexed grantee, address indexed granter, string[] methods)
func (_Staking *StakingFilterer) FilterRevocation(opts *bind.FilterOpts, grantee []common.Address, granter []common.Address) (*StakingRevocationIterator, error) {

	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}
	var granterRule []interface{}
	for _, granterItem := range granter {
		granterRule = append(granterRule, granterItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Revocation", granteeRule, granterRule)
	if err != nil {
		return nil, err
	}
	return &StakingRevocationIterator{contract: _Staking.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xb0901d422521d0496e60bfbd8023b219d603a6cb950b43b2fe95043676cb353e.
//
// Solidity: event Revocation(address indexed grantee, address indexed granter, string[] methods)
func (_Staking *StakingFilterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *StakingRevocation, grantee []common.Address, granter []common.Address) (event.Subscription, error) {

	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}
	var granterRule []interface{}
	for _, granterItem := range granter {
		granterRule = append(granterRule, granterItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Revocation", granteeRule, granterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRevocation)
				if err := _Staking.contract.UnpackLog(event, "Revocation", log); err != nil {
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

// ParseRevocation is a log parse operation binding the contract event 0xb0901d422521d0496e60bfbd8023b219d603a6cb950b43b2fe95043676cb353e.
//
// Solidity: event Revocation(address indexed grantee, address indexed granter, string[] methods)
func (_Staking *StakingFilterer) ParseRevocation(log types.Log) (*StakingRevocation, error) {
	event := new(StakingRevocation)
	if err := _Staking.contract.UnpackLog(event, "Revocation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnbondIterator is returned from FilterUnbond and is used to iterate over the raw logs and unpacked data for Unbond events raised by the Staking contract.
type StakingUnbondIterator struct {
	Event *StakingUnbond // Event containing the contract specifics and raw log

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
func (it *StakingUnbondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnbond)
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
		it.Event = new(StakingUnbond)
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
func (it *StakingUnbondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnbondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnbond represents a Unbond event raised by the Staking contract.
type StakingUnbond struct {
	DelegatorAddress common.Address
	ValidatorAddress common.Address
	Amount           *big.Int
	CompletionTime   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUnbond is a free log retrieval operation binding the contract event 0x4bf8087be3b8a59c2662514df2ed4a3dcaf9ca22f442340cfc05a4e52343d18e.
//
// Solidity: event Unbond(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 completionTime)
func (_Staking *StakingFilterer) FilterUnbond(opts *bind.FilterOpts, delegatorAddress []common.Address, validatorAddress []common.Address) (*StakingUnbondIterator, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unbond", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingUnbondIterator{contract: _Staking.contract, event: "Unbond", logs: logs, sub: sub}, nil
}

// WatchUnbond is a free log subscription operation binding the contract event 0x4bf8087be3b8a59c2662514df2ed4a3dcaf9ca22f442340cfc05a4e52343d18e.
//
// Solidity: event Unbond(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 completionTime)
func (_Staking *StakingFilterer) WatchUnbond(opts *bind.WatchOpts, sink chan<- *StakingUnbond, delegatorAddress []common.Address, validatorAddress []common.Address) (event.Subscription, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unbond", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnbond)
				if err := _Staking.contract.UnpackLog(event, "Unbond", log); err != nil {
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

// ParseUnbond is a log parse operation binding the contract event 0x4bf8087be3b8a59c2662514df2ed4a3dcaf9ca22f442340cfc05a4e52343d18e.
//
// Solidity: event Unbond(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 completionTime)
func (_Staking *StakingFilterer) ParseUnbond(log types.Log) (*StakingUnbond, error) {
	event := new(StakingUnbond)
	if err := _Staking.contract.UnpackLog(event, "Unbond", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
