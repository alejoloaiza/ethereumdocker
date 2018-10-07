// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package helloworld

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// HelloworldABI is the input ABI used to generate the binding from.
const HelloworldABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_hello\",\"type\":\"string\"}],\"name\":\"Set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Get\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_hello\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// HelloworldBin is the compiled bytecode used for deploying new contracts.
const HelloworldBin = `608060405234801561001057600080fd5b506040516103db3803806103db833981018060405281019080805182019291905050508060009080519060200190610049929190610050565b50506100f5565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061009157805160ff19168380011785556100bf565b828001600101855582156100bf579182015b828111156100be5782518255916020019190600101906100a3565b5b5090506100cc91906100d0565b5090565b6100f291905b808211156100ee5760008160009055506001016100d6565b5090565b90565b6102d7806101046000396000f30060806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806334541fcd14610051578063b1976a02146100ba575b600080fd5b34801561005d57600080fd5b506100b8600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061014a565b005b3480156100c657600080fd5b506100cf610164565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561010f5780820151818401526020810190506100f4565b50505050905090810190601f16801561013c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b8060009080519060200190610160929190610206565b5050565b606060008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101fc5780601f106101d1576101008083540402835291602001916101fc565b820191906000526020600020905b8154815290600101906020018083116101df57829003601f168201915b5050505050905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061024757805160ff1916838001178555610275565b82800160010185558215610275579182015b82811115610274578251825591602001919060010190610259565b5b5090506102829190610286565b5090565b6102a891905b808211156102a457600081600090555060010161028c565b5090565b905600a165627a7a72305820bcbefe69d04ace60fcf80f8a3eb5e3cdeb986240911b669ba7025cefdb95c76a0029`

// DeployHelloworld deploys a new Ethereum contract, binding an instance of Helloworld to it.
func DeployHelloworld(auth *bind.TransactOpts, backend bind.ContractBackend, _hello string) (common.Address, *types.Transaction, *Helloworld, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloworldABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HelloworldBin), backend, _hello)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Helloworld{HelloworldCaller: HelloworldCaller{contract: contract}, HelloworldTransactor: HelloworldTransactor{contract: contract}, HelloworldFilterer: HelloworldFilterer{contract: contract}}, nil
}

// Helloworld is an auto generated Go binding around an Ethereum contract.
type Helloworld struct {
	HelloworldCaller     // Read-only binding to the contract
	HelloworldTransactor // Write-only binding to the contract
	HelloworldFilterer   // Log filterer for contract events
}

// HelloworldCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelloworldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloworldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloworldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloworldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloworldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloworldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloworldSession struct {
	Contract     *Helloworld       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloworldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloworldCallerSession struct {
	Contract *HelloworldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HelloworldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloworldTransactorSession struct {
	Contract     *HelloworldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HelloworldRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelloworldRaw struct {
	Contract *Helloworld // Generic contract binding to access the raw methods on
}

// HelloworldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloworldCallerRaw struct {
	Contract *HelloworldCaller // Generic read-only contract binding to access the raw methods on
}

// HelloworldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloworldTransactorRaw struct {
	Contract *HelloworldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelloworld creates a new instance of Helloworld, bound to a specific deployed contract.
func NewHelloworld(address common.Address, backend bind.ContractBackend) (*Helloworld, error) {
	contract, err := bindHelloworld(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Helloworld{HelloworldCaller: HelloworldCaller{contract: contract}, HelloworldTransactor: HelloworldTransactor{contract: contract}, HelloworldFilterer: HelloworldFilterer{contract: contract}}, nil
}

// NewHelloworldCaller creates a new read-only instance of Helloworld, bound to a specific deployed contract.
func NewHelloworldCaller(address common.Address, caller bind.ContractCaller) (*HelloworldCaller, error) {
	contract, err := bindHelloworld(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloworldCaller{contract: contract}, nil
}

// NewHelloworldTransactor creates a new write-only instance of Helloworld, bound to a specific deployed contract.
func NewHelloworldTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloworldTransactor, error) {
	contract, err := bindHelloworld(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloworldTransactor{contract: contract}, nil
}

// NewHelloworldFilterer creates a new log filterer instance of Helloworld, bound to a specific deployed contract.
func NewHelloworldFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloworldFilterer, error) {
	contract, err := bindHelloworld(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloworldFilterer{contract: contract}, nil
}

// bindHelloworld binds a generic wrapper to an already deployed contract.
func bindHelloworld(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloworldABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helloworld *HelloworldRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Helloworld.Contract.HelloworldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helloworld *HelloworldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helloworld.Contract.HelloworldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helloworld *HelloworldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helloworld.Contract.HelloworldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helloworld *HelloworldCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Helloworld.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helloworld *HelloworldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helloworld.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helloworld *HelloworldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helloworld.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0xb1976a02.
//
// Solidity: function Get() constant returns(string)
func (_Helloworld *HelloworldCaller) Get(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Helloworld.contract.Call(opts, out, "Get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0xb1976a02.
//
// Solidity: function Get() constant returns(string)
func (_Helloworld *HelloworldSession) Get() (string, error) {
	return _Helloworld.Contract.Get(&_Helloworld.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0xb1976a02.
//
// Solidity: function Get() constant returns(string)
func (_Helloworld *HelloworldCallerSession) Get() (string, error) {
	return _Helloworld.Contract.Get(&_Helloworld.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x34541fcd.
//
// Solidity: function Set(_hello string) returns()
func (_Helloworld *HelloworldTransactor) Set(opts *bind.TransactOpts, _hello string) (*types.Transaction, error) {
	return _Helloworld.contract.Transact(opts, "Set", _hello)
}

// Set is a paid mutator transaction binding the contract method 0x34541fcd.
//
// Solidity: function Set(_hello string) returns()
func (_Helloworld *HelloworldSession) Set(_hello string) (*types.Transaction, error) {
	return _Helloworld.Contract.Set(&_Helloworld.TransactOpts, _hello)
}

// Set is a paid mutator transaction binding the contract method 0x34541fcd.
//
// Solidity: function Set(_hello string) returns()
func (_Helloworld *HelloworldTransactorSession) Set(_hello string) (*types.Transaction, error) {
	return _Helloworld.Contract.Set(&_Helloworld.TransactOpts, _hello)
}
