// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StakingPool

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StakingPoolABI is the input ABI used to generate the binding from.
const StakingPoolABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"name\":\"LogRevoke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"LogSetTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"name\":\"LogStake\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"weightOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StakingPoolBin is the compiled bytecode used for deploying new contracts.
var StakingPoolBin = "0x608060405234801561001057600080fd5b5060405161199a38038061199a8339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160046000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050611886806101146000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80639d76ea58116100715780639d76ea58146101e0578063a2e80c5e1461022a578063a694fc3a14610248578063ad7a672f14610276578063dd4bc10114610294578063fc0c546a146102ec576100a9565b806320c5429b146100ae57806326a4e8d2146100dc57806327e235e31461012057806342cde4e814610178578063570ca73514610196575b600080fd5b6100da600480360360208110156100c457600080fd5b8101908080359060200190929190505050610336565b005b61011e600480360360208110156100f257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105e0565b005b6101626004803603602081101561013657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107ac565b6040518082815260200191505060405180910390f35b6101806107c4565b6040518082815260200191505060405180910390f35b61019e6107d5565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101e86107fa565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610232610820565b6040518082815260200191505060405180910390f35b6102746004803603602081101561025e57600080fd5b8101908080359060200190929190505050610826565b005b61027e610b6a565b6040518082815260200191505060405180910390f35b6102d6600480360360208110156102aa57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b70565b6040518082815260200191505060405180910390f35b6102f4610bb9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548111156103eb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5374616b696e6720696e73756666696369656e7420746f207265766f6b652e0081525060200191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561049457600080fd5b505af11580156104a8573d6000803e3d6000fd5b505050506040513d60208110156104be57600080fd5b81019080805190602001909291905050505061051b3382600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205403610bdf565b806006600082825403925050819055507fdec5ce19f447be2b1838432cb938cf95c3c6675bfbc389ec94f1c2f79956a4573382600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390a150565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106a2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d75737420626520746865206f70657261746f722e000000000000000000000081525060200191505060405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fa664206ac1d008004105250bc16103eeae43c3550ab4dca3c187db27b62bd3ed600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60036020528060005260406000206000915090505481565b60006107d0600a610dc5565b905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60055481565b600081141561089d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f416d6f756e74206d757374206e6f74207a65726f2e000000000000000000000081525060200191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15801561097a57600080fd5b505af115801561098e573d6000803e3d6000fd5b505050506040513d60208110156109a457600080fd5b810190808051906020019092919050505050600073ffffffffffffffffffffffffffffffffffffffff16600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610a5957610a543382610f0b565b610aa5565b610aa43382600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401610bdf565b5b806006600082825401925050819055507fa0d1fa51224fb08c7c0b48d481448ad6a31c945d1a68cac8758cac3a92d5001c3382600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390a150565b60065481565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff16600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610ce1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f5374616b657220646f6573206e6f742065786973742e0000000000000000000081525060200191505060405180910390fd5b6000610cec836111cf565b90506000600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050610d5d8284836112f4565b15610dab5782600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610dbf565b610db4846113fa565b610dbe8484610f0b565b5b50505050565b6000600554821115610dd75760055491505b6000809050600060046000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008090505b84811015610f0057600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483019250600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150806001019050610e47565b508192505050919050565b600073ffffffffffffffffffffffffffffffffffffffff16600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461100c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f5374616b657220616c7265616479206578697374732e0000000000000000000081525060200191505060405180910390fd5b6000611017826116c1565b905081600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600560008154809291906001019190505550505050565b600080600190505b600173ffffffffffffffffffffffffffffffffffffffff16600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146112e95761127483826117b9565b1561128257809150506112ef565b600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506111d7565b60009150505b919050565b6000600173ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480611370575082600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410155b80156113f15750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614806113f05750600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483115b5b90509392505050565b600073ffffffffffffffffffffffffffffffffffffffff16600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156114fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f5374616b657220646f6573206e6f742065786973742e0000000000000000000081525060200191505060405180910390fd5b6000611507826111cf565b9050600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600560008154809291906001900391905055505050565b600080600190505b6001156117af5761173a8184600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166112f4565b1561174857809150506117b4565b600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506116c9565b809150505b919050565b60008273ffffffffffffffffffffffffffffffffffffffff16600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161490509291505056fea265627a7a7231582095110a2526972e8623fd615d9dc84b6ace4461c90a108009590bef0515452f1e64736f6c63430005110032"

// DeployStakingPool deploys a new Ethereum contract, binding an instance of StakingPool to it.
func DeployStakingPool(auth *bind.TransactOpts, backend bind.ContractBackend, _operator common.Address) (common.Address, *types.Transaction, *StakingPool, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingPoolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StakingPoolBin), backend, _operator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingPool{StakingPoolCaller: StakingPoolCaller{contract: contract}, StakingPoolTransactor: StakingPoolTransactor{contract: contract}, StakingPoolFilterer: StakingPoolFilterer{contract: contract}}, nil
}

// StakingPool is an auto generated Go binding around an Ethereum contract.
type StakingPool struct {
	StakingPoolCaller     // Read-only binding to the contract
	StakingPoolTransactor // Write-only binding to the contract
	StakingPoolFilterer   // Log filterer for contract events
}

// StakingPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingPoolSession struct {
	Contract     *StakingPool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingPoolCallerSession struct {
	Contract *StakingPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StakingPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingPoolTransactorSession struct {
	Contract     *StakingPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StakingPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingPoolRaw struct {
	Contract *StakingPool // Generic contract binding to access the raw methods on
}

// StakingPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingPoolCallerRaw struct {
	Contract *StakingPoolCaller // Generic read-only contract binding to access the raw methods on
}

// StakingPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingPoolTransactorRaw struct {
	Contract *StakingPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingPool creates a new instance of StakingPool, bound to a specific deployed contract.
func NewStakingPool(address common.Address, backend bind.ContractBackend) (*StakingPool, error) {
	contract, err := bindStakingPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingPool{StakingPoolCaller: StakingPoolCaller{contract: contract}, StakingPoolTransactor: StakingPoolTransactor{contract: contract}, StakingPoolFilterer: StakingPoolFilterer{contract: contract}}, nil
}

// NewStakingPoolCaller creates a new read-only instance of StakingPool, bound to a specific deployed contract.
func NewStakingPoolCaller(address common.Address, caller bind.ContractCaller) (*StakingPoolCaller, error) {
	contract, err := bindStakingPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolCaller{contract: contract}, nil
}

// NewStakingPoolTransactor creates a new write-only instance of StakingPool, bound to a specific deployed contract.
func NewStakingPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingPoolTransactor, error) {
	contract, err := bindStakingPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolTransactor{contract: contract}, nil
}

// NewStakingPoolFilterer creates a new log filterer instance of StakingPool, bound to a specific deployed contract.
func NewStakingPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingPoolFilterer, error) {
	contract, err := bindStakingPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingPoolFilterer{contract: contract}, nil
}

// bindStakingPool binds a generic wrapper to an already deployed contract.
func bindStakingPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPool *StakingPoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StakingPool.Contract.StakingPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPool *StakingPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPool.Contract.StakingPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPool *StakingPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPool.Contract.StakingPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPool *StakingPoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StakingPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPool *StakingPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPool *StakingPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPool.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_StakingPool *StakingPoolCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingPool.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_StakingPool *StakingPoolSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _StakingPool.Contract.Balances(&_StakingPool.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_StakingPool *StakingPoolCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _StakingPool.Contract.Balances(&_StakingPool.CallOpts, arg0)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_StakingPool *StakingPoolCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StakingPool.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_StakingPool *StakingPoolSession) Operator() (common.Address, error) {
	return _StakingPool.Contract.Operator(&_StakingPool.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_StakingPool *StakingPoolCallerSession) Operator() (common.Address, error) {
	return _StakingPool.Contract.Operator(&_StakingPool.CallOpts)
}

// StakersCount is a free data retrieval call binding the contract method 0xa2e80c5e.
//
// Solidity: function stakersCount() view returns(uint256)
func (_StakingPool *StakingPoolCaller) StakersCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingPool.contract.Call(opts, out, "stakersCount")
	return *ret0, err
}

// StakersCount is a free data retrieval call binding the contract method 0xa2e80c5e.
//
// Solidity: function stakersCount() view returns(uint256)
func (_StakingPool *StakingPoolSession) StakersCount() (*big.Int, error) {
	return _StakingPool.Contract.StakersCount(&_StakingPool.CallOpts)
}

// StakersCount is a free data retrieval call binding the contract method 0xa2e80c5e.
//
// Solidity: function stakersCount() view returns(uint256)
func (_StakingPool *StakingPoolCallerSession) StakersCount() (*big.Int, error) {
	return _StakingPool.Contract.StakersCount(&_StakingPool.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_StakingPool *StakingPoolCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingPool.contract.Call(opts, out, "threshold")
	return *ret0, err
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_StakingPool *StakingPoolSession) Threshold() (*big.Int, error) {
	return _StakingPool.Contract.Threshold(&_StakingPool.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_StakingPool *StakingPoolCallerSession) Threshold() (*big.Int, error) {
	return _StakingPool.Contract.Threshold(&_StakingPool.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StakingPool *StakingPoolCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StakingPool.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StakingPool *StakingPoolSession) Token() (common.Address, error) {
	return _StakingPool.Contract.Token(&_StakingPool.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StakingPool *StakingPoolCallerSession) Token() (common.Address, error) {
	return _StakingPool.Contract.Token(&_StakingPool.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_StakingPool *StakingPoolCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StakingPool.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_StakingPool *StakingPoolSession) TokenAddress() (common.Address, error) {
	return _StakingPool.Contract.TokenAddress(&_StakingPool.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_StakingPool *StakingPoolCallerSession) TokenAddress() (common.Address, error) {
	return _StakingPool.Contract.TokenAddress(&_StakingPool.CallOpts)
}

// TotalBalance is a free data retrieval call binding the contract method 0xad7a672f.
//
// Solidity: function totalBalance() view returns(uint256)
func (_StakingPool *StakingPoolCaller) TotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingPool.contract.Call(opts, out, "totalBalance")
	return *ret0, err
}

// TotalBalance is a free data retrieval call binding the contract method 0xad7a672f.
//
// Solidity: function totalBalance() view returns(uint256)
func (_StakingPool *StakingPoolSession) TotalBalance() (*big.Int, error) {
	return _StakingPool.Contract.TotalBalance(&_StakingPool.CallOpts)
}

// TotalBalance is a free data retrieval call binding the contract method 0xad7a672f.
//
// Solidity: function totalBalance() view returns(uint256)
func (_StakingPool *StakingPoolCallerSession) TotalBalance() (*big.Int, error) {
	return _StakingPool.Contract.TotalBalance(&_StakingPool.CallOpts)
}

// WeightOf is a free data retrieval call binding the contract method 0xdd4bc101.
//
// Solidity: function weightOf(address staker) view returns(uint256)
func (_StakingPool *StakingPoolCaller) WeightOf(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingPool.contract.Call(opts, out, "weightOf", staker)
	return *ret0, err
}

// WeightOf is a free data retrieval call binding the contract method 0xdd4bc101.
//
// Solidity: function weightOf(address staker) view returns(uint256)
func (_StakingPool *StakingPoolSession) WeightOf(staker common.Address) (*big.Int, error) {
	return _StakingPool.Contract.WeightOf(&_StakingPool.CallOpts, staker)
}

// WeightOf is a free data retrieval call binding the contract method 0xdd4bc101.
//
// Solidity: function weightOf(address staker) view returns(uint256)
func (_StakingPool *StakingPoolCallerSession) WeightOf(staker common.Address) (*big.Int, error) {
	return _StakingPool.Contract.WeightOf(&_StakingPool.CallOpts, staker)
}

// Revoke is a paid mutator transaction binding the contract method 0x20c5429b.
//
// Solidity: function revoke(uint256 amount) returns()
func (_StakingPool *StakingPoolTransactor) Revoke(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakingPool.contract.Transact(opts, "revoke", amount)
}

// Revoke is a paid mutator transaction binding the contract method 0x20c5429b.
//
// Solidity: function revoke(uint256 amount) returns()
func (_StakingPool *StakingPoolSession) Revoke(amount *big.Int) (*types.Transaction, error) {
	return _StakingPool.Contract.Revoke(&_StakingPool.TransactOpts, amount)
}

// Revoke is a paid mutator transaction binding the contract method 0x20c5429b.
//
// Solidity: function revoke(uint256 amount) returns()
func (_StakingPool *StakingPoolTransactorSession) Revoke(amount *big.Int) (*types.Transaction, error) {
	return _StakingPool.Contract.Revoke(&_StakingPool.TransactOpts, amount)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _token) returns()
func (_StakingPool *StakingPoolTransactor) SetTokenAddress(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _StakingPool.contract.Transact(opts, "setTokenAddress", _token)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _token) returns()
func (_StakingPool *StakingPoolSession) SetTokenAddress(_token common.Address) (*types.Transaction, error) {
	return _StakingPool.Contract.SetTokenAddress(&_StakingPool.TransactOpts, _token)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _token) returns()
func (_StakingPool *StakingPoolTransactorSession) SetTokenAddress(_token common.Address) (*types.Transaction, error) {
	return _StakingPool.Contract.SetTokenAddress(&_StakingPool.TransactOpts, _token)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_StakingPool *StakingPoolTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakingPool.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_StakingPool *StakingPoolSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _StakingPool.Contract.Stake(&_StakingPool.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_StakingPool *StakingPoolTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _StakingPool.Contract.Stake(&_StakingPool.TransactOpts, amount)
}

// StakingPoolLogRevokeIterator is returned from FilterLogRevoke and is used to iterate over the raw logs and unpacked data for LogRevoke events raised by the StakingPool contract.
type StakingPoolLogRevokeIterator struct {
	Event *StakingPoolLogRevoke // Event containing the contract specifics and raw log

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
func (it *StakingPoolLogRevokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolLogRevoke)
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
		it.Event = new(StakingPoolLogRevoke)
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
func (it *StakingPoolLogRevokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolLogRevokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolLogRevoke represents a LogRevoke event raised by the StakingPool contract.
type StakingPoolLogRevoke struct {
	Staker      common.Address
	Value       *big.Int
	TotalAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogRevoke is a free log retrieval operation binding the contract event 0xdec5ce19f447be2b1838432cb938cf95c3c6675bfbc389ec94f1c2f79956a457.
//
// Solidity: event LogRevoke(address _staker, uint256 _value, uint256 _totalAmount)
func (_StakingPool *StakingPoolFilterer) FilterLogRevoke(opts *bind.FilterOpts) (*StakingPoolLogRevokeIterator, error) {

	logs, sub, err := _StakingPool.contract.FilterLogs(opts, "LogRevoke")
	if err != nil {
		return nil, err
	}
	return &StakingPoolLogRevokeIterator{contract: _StakingPool.contract, event: "LogRevoke", logs: logs, sub: sub}, nil
}

// WatchLogRevoke is a free log subscription operation binding the contract event 0xdec5ce19f447be2b1838432cb938cf95c3c6675bfbc389ec94f1c2f79956a457.
//
// Solidity: event LogRevoke(address _staker, uint256 _value, uint256 _totalAmount)
func (_StakingPool *StakingPoolFilterer) WatchLogRevoke(opts *bind.WatchOpts, sink chan<- *StakingPoolLogRevoke) (event.Subscription, error) {

	logs, sub, err := _StakingPool.contract.WatchLogs(opts, "LogRevoke")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolLogRevoke)
				if err := _StakingPool.contract.UnpackLog(event, "LogRevoke", log); err != nil {
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

// ParseLogRevoke is a log parse operation binding the contract event 0xdec5ce19f447be2b1838432cb938cf95c3c6675bfbc389ec94f1c2f79956a457.
//
// Solidity: event LogRevoke(address _staker, uint256 _value, uint256 _totalAmount)
func (_StakingPool *StakingPoolFilterer) ParseLogRevoke(log types.Log) (*StakingPoolLogRevoke, error) {
	event := new(StakingPoolLogRevoke)
	if err := _StakingPool.contract.UnpackLog(event, "LogRevoke", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingPoolLogSetTokenAddressIterator is returned from FilterLogSetTokenAddress and is used to iterate over the raw logs and unpacked data for LogSetTokenAddress events raised by the StakingPool contract.
type StakingPoolLogSetTokenAddressIterator struct {
	Event *StakingPoolLogSetTokenAddress // Event containing the contract specifics and raw log

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
func (it *StakingPoolLogSetTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolLogSetTokenAddress)
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
		it.Event = new(StakingPoolLogSetTokenAddress)
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
func (it *StakingPoolLogSetTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolLogSetTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolLogSetTokenAddress represents a LogSetTokenAddress event raised by the StakingPool contract.
type StakingPoolLogSetTokenAddress struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetTokenAddress is a free log retrieval operation binding the contract event 0xa664206ac1d008004105250bc16103eeae43c3550ab4dca3c187db27b62bd3ed.
//
// Solidity: event LogSetTokenAddress(address _token)
func (_StakingPool *StakingPoolFilterer) FilterLogSetTokenAddress(opts *bind.FilterOpts) (*StakingPoolLogSetTokenAddressIterator, error) {

	logs, sub, err := _StakingPool.contract.FilterLogs(opts, "LogSetTokenAddress")
	if err != nil {
		return nil, err
	}
	return &StakingPoolLogSetTokenAddressIterator{contract: _StakingPool.contract, event: "LogSetTokenAddress", logs: logs, sub: sub}, nil
}

// WatchLogSetTokenAddress is a free log subscription operation binding the contract event 0xa664206ac1d008004105250bc16103eeae43c3550ab4dca3c187db27b62bd3ed.
//
// Solidity: event LogSetTokenAddress(address _token)
func (_StakingPool *StakingPoolFilterer) WatchLogSetTokenAddress(opts *bind.WatchOpts, sink chan<- *StakingPoolLogSetTokenAddress) (event.Subscription, error) {

	logs, sub, err := _StakingPool.contract.WatchLogs(opts, "LogSetTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolLogSetTokenAddress)
				if err := _StakingPool.contract.UnpackLog(event, "LogSetTokenAddress", log); err != nil {
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

// ParseLogSetTokenAddress is a log parse operation binding the contract event 0xa664206ac1d008004105250bc16103eeae43c3550ab4dca3c187db27b62bd3ed.
//
// Solidity: event LogSetTokenAddress(address _token)
func (_StakingPool *StakingPoolFilterer) ParseLogSetTokenAddress(log types.Log) (*StakingPoolLogSetTokenAddress, error) {
	event := new(StakingPoolLogSetTokenAddress)
	if err := _StakingPool.contract.UnpackLog(event, "LogSetTokenAddress", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingPoolLogStakeIterator is returned from FilterLogStake and is used to iterate over the raw logs and unpacked data for LogStake events raised by the StakingPool contract.
type StakingPoolLogStakeIterator struct {
	Event *StakingPoolLogStake // Event containing the contract specifics and raw log

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
func (it *StakingPoolLogStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolLogStake)
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
		it.Event = new(StakingPoolLogStake)
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
func (it *StakingPoolLogStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolLogStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolLogStake represents a LogStake event raised by the StakingPool contract.
type StakingPoolLogStake struct {
	Staker      common.Address
	Value       *big.Int
	TotalAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogStake is a free log retrieval operation binding the contract event 0xa0d1fa51224fb08c7c0b48d481448ad6a31c945d1a68cac8758cac3a92d5001c.
//
// Solidity: event LogStake(address _staker, uint256 _value, uint256 _totalAmount)
func (_StakingPool *StakingPoolFilterer) FilterLogStake(opts *bind.FilterOpts) (*StakingPoolLogStakeIterator, error) {

	logs, sub, err := _StakingPool.contract.FilterLogs(opts, "LogStake")
	if err != nil {
		return nil, err
	}
	return &StakingPoolLogStakeIterator{contract: _StakingPool.contract, event: "LogStake", logs: logs, sub: sub}, nil
}

// WatchLogStake is a free log subscription operation binding the contract event 0xa0d1fa51224fb08c7c0b48d481448ad6a31c945d1a68cac8758cac3a92d5001c.
//
// Solidity: event LogStake(address _staker, uint256 _value, uint256 _totalAmount)
func (_StakingPool *StakingPoolFilterer) WatchLogStake(opts *bind.WatchOpts, sink chan<- *StakingPoolLogStake) (event.Subscription, error) {

	logs, sub, err := _StakingPool.contract.WatchLogs(opts, "LogStake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolLogStake)
				if err := _StakingPool.contract.UnpackLog(event, "LogStake", log); err != nil {
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

// ParseLogStake is a log parse operation binding the contract event 0xa0d1fa51224fb08c7c0b48d481448ad6a31c945d1a68cac8758cac3a92d5001c.
//
// Solidity: event LogStake(address _staker, uint256 _value, uint256 _totalAmount)
func (_StakingPool *StakingPoolFilterer) ParseLogStake(log types.Log) (*StakingPoolLogStake, error) {
	event := new(StakingPoolLogStake)
	if err := _StakingPool.contract.UnpackLog(event, "LogStake", log); err != nil {
		return nil, err
	}
	return event, nil
}
