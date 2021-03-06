// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BridgeRegistry

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

// BridgeRegistryABI is the input ABI used to generate the binding from.
const BridgeRegistryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_quantiexERC20Bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quantiexERC721Bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeERC20Bank\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeERC721Bank\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_valset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingPool\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_quantiexERC20Bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_quantiexERC721Bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_bridgeERC20Bank\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_bridgeERC721Bank\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_valset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_stakingPool\",\"type\":\"address\"}],\"name\":\"LogContractsRegistered\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeERC20Bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeERC721Bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"quantiexERC20Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"quantiexERC721Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BridgeRegistryBin is the compiled bytecode used for deploying new contracts.
var BridgeRegistryBin = "0x608060405234801561001057600080fd5b506040516108aa3803806108aa833981810160405260e081101561003357600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190505050866000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f626ca7cf58dcae558e83ad85a3cbf2c646a3f55925624c86f40d3b0ab55bf5d96000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200197505050505050505060405180910390a1505050505050506103c6806104e46000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80637dc0d1d01161005b5780637dc0d1d0146101605780637f54af0c146101aa578063c572e051146101f4578063cd9b07c41461023e5761007d565b80630c56ae3b146100825780633bfdab0a146100cc57806378c0911514610116575b600080fd5b61008a610288565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100d46102ae565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61011e6102d3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101686102f9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101b261031f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101fc610345565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61024661036b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea265627a7a7231582023e8a93089115ddf3cf7433636c3e4f5523dc641c382697fe48109c181ec21fe64736f6c63430005110032"

// DeployBridgeRegistry deploys a new Ethereum contract, binding an instance of BridgeRegistry to it.
func DeployBridgeRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _quantiexERC20Bridge common.Address, _quantiexERC721Bridge common.Address, _bridgeERC20Bank common.Address, _bridgeERC721Bank common.Address, _oracle common.Address, _valset common.Address, _stakingPool common.Address) (common.Address, *types.Transaction, *BridgeRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeRegistryBin), backend, _quantiexERC20Bridge, _quantiexERC721Bridge, _bridgeERC20Bank, _bridgeERC721Bank, _oracle, _valset, _stakingPool)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeRegistry{BridgeRegistryCaller: BridgeRegistryCaller{contract: contract}, BridgeRegistryTransactor: BridgeRegistryTransactor{contract: contract}, BridgeRegistryFilterer: BridgeRegistryFilterer{contract: contract}}, nil
}

// BridgeRegistry is an auto generated Go binding around an Ethereum contract.
type BridgeRegistry struct {
	BridgeRegistryCaller     // Read-only binding to the contract
	BridgeRegistryTransactor // Write-only binding to the contract
	BridgeRegistryFilterer   // Log filterer for contract events
}

// BridgeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeRegistrySession struct {
	Contract     *BridgeRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeRegistryCallerSession struct {
	Contract *BridgeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BridgeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeRegistryTransactorSession struct {
	Contract     *BridgeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BridgeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRegistryRaw struct {
	Contract *BridgeRegistry // Generic contract binding to access the raw methods on
}

// BridgeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeRegistryCallerRaw struct {
	Contract *BridgeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeRegistryTransactorRaw struct {
	Contract *BridgeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeRegistry creates a new instance of BridgeRegistry, bound to a specific deployed contract.
func NewBridgeRegistry(address common.Address, backend bind.ContractBackend) (*BridgeRegistry, error) {
	contract, err := bindBridgeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeRegistry{BridgeRegistryCaller: BridgeRegistryCaller{contract: contract}, BridgeRegistryTransactor: BridgeRegistryTransactor{contract: contract}, BridgeRegistryFilterer: BridgeRegistryFilterer{contract: contract}}, nil
}

// NewBridgeRegistryCaller creates a new read-only instance of BridgeRegistry, bound to a specific deployed contract.
func NewBridgeRegistryCaller(address common.Address, caller bind.ContractCaller) (*BridgeRegistryCaller, error) {
	contract, err := bindBridgeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryCaller{contract: contract}, nil
}

// NewBridgeRegistryTransactor creates a new write-only instance of BridgeRegistry, bound to a specific deployed contract.
func NewBridgeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeRegistryTransactor, error) {
	contract, err := bindBridgeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryTransactor{contract: contract}, nil
}

// NewBridgeRegistryFilterer creates a new log filterer instance of BridgeRegistry, bound to a specific deployed contract.
func NewBridgeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeRegistryFilterer, error) {
	contract, err := bindBridgeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryFilterer{contract: contract}, nil
}

// bindBridgeRegistry binds a generic wrapper to an already deployed contract.
func bindBridgeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeRegistry *BridgeRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeRegistry.Contract.BridgeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeRegistry *BridgeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.BridgeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeRegistry *BridgeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.BridgeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeRegistry *BridgeRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeRegistry *BridgeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeRegistry *BridgeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeRegistry.Contract.contract.Transact(opts, method, params...)
}

// BridgeERC20Bank is a free data retrieval call binding the contract method 0x78c09115.
//
// Solidity: function bridgeERC20Bank() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) BridgeERC20Bank(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "bridgeERC20Bank")
	return *ret0, err
}

// BridgeERC20Bank is a free data retrieval call binding the contract method 0x78c09115.
//
// Solidity: function bridgeERC20Bank() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) BridgeERC20Bank() (common.Address, error) {
	return _BridgeRegistry.Contract.BridgeERC20Bank(&_BridgeRegistry.CallOpts)
}

// BridgeERC20Bank is a free data retrieval call binding the contract method 0x78c09115.
//
// Solidity: function bridgeERC20Bank() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) BridgeERC20Bank() (common.Address, error) {
	return _BridgeRegistry.Contract.BridgeERC20Bank(&_BridgeRegistry.CallOpts)
}

// BridgeERC721Bank is a free data retrieval call binding the contract method 0xc572e051.
//
// Solidity: function bridgeERC721Bank() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) BridgeERC721Bank(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "bridgeERC721Bank")
	return *ret0, err
}

// BridgeERC721Bank is a free data retrieval call binding the contract method 0xc572e051.
//
// Solidity: function bridgeERC721Bank() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) BridgeERC721Bank() (common.Address, error) {
	return _BridgeRegistry.Contract.BridgeERC721Bank(&_BridgeRegistry.CallOpts)
}

// BridgeERC721Bank is a free data retrieval call binding the contract method 0xc572e051.
//
// Solidity: function bridgeERC721Bank() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) BridgeERC721Bank() (common.Address, error) {
	return _BridgeRegistry.Contract.BridgeERC721Bank(&_BridgeRegistry.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) Oracle() (common.Address, error) {
	return _BridgeRegistry.Contract.Oracle(&_BridgeRegistry.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) Oracle() (common.Address, error) {
	return _BridgeRegistry.Contract.Oracle(&_BridgeRegistry.CallOpts)
}

// QuantiexERC20Bridge is a free data retrieval call binding the contract method 0x3bfdab0a.
//
// Solidity: function quantiexERC20Bridge() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) QuantiexERC20Bridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "quantiexERC20Bridge")
	return *ret0, err
}

// QuantiexERC20Bridge is a free data retrieval call binding the contract method 0x3bfdab0a.
//
// Solidity: function quantiexERC20Bridge() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) QuantiexERC20Bridge() (common.Address, error) {
	return _BridgeRegistry.Contract.QuantiexERC20Bridge(&_BridgeRegistry.CallOpts)
}

// QuantiexERC20Bridge is a free data retrieval call binding the contract method 0x3bfdab0a.
//
// Solidity: function quantiexERC20Bridge() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) QuantiexERC20Bridge() (common.Address, error) {
	return _BridgeRegistry.Contract.QuantiexERC20Bridge(&_BridgeRegistry.CallOpts)
}

// QuantiexERC721Bridge is a free data retrieval call binding the contract method 0xcd9b07c4.
//
// Solidity: function quantiexERC721Bridge() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) QuantiexERC721Bridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "quantiexERC721Bridge")
	return *ret0, err
}

// QuantiexERC721Bridge is a free data retrieval call binding the contract method 0xcd9b07c4.
//
// Solidity: function quantiexERC721Bridge() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) QuantiexERC721Bridge() (common.Address, error) {
	return _BridgeRegistry.Contract.QuantiexERC721Bridge(&_BridgeRegistry.CallOpts)
}

// QuantiexERC721Bridge is a free data retrieval call binding the contract method 0xcd9b07c4.
//
// Solidity: function quantiexERC721Bridge() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) QuantiexERC721Bridge() (common.Address, error) {
	return _BridgeRegistry.Contract.QuantiexERC721Bridge(&_BridgeRegistry.CallOpts)
}

// StakingPool is a free data retrieval call binding the contract method 0x0c56ae3b.
//
// Solidity: function stakingPool() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) StakingPool(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "stakingPool")
	return *ret0, err
}

// StakingPool is a free data retrieval call binding the contract method 0x0c56ae3b.
//
// Solidity: function stakingPool() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) StakingPool() (common.Address, error) {
	return _BridgeRegistry.Contract.StakingPool(&_BridgeRegistry.CallOpts)
}

// StakingPool is a free data retrieval call binding the contract method 0x0c56ae3b.
//
// Solidity: function stakingPool() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) StakingPool() (common.Address, error) {
	return _BridgeRegistry.Contract.StakingPool(&_BridgeRegistry.CallOpts)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_BridgeRegistry *BridgeRegistryCaller) Valset(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeRegistry.contract.Call(opts, out, "valset")
	return *ret0, err
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_BridgeRegistry *BridgeRegistrySession) Valset() (common.Address, error) {
	return _BridgeRegistry.Contract.Valset(&_BridgeRegistry.CallOpts)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_BridgeRegistry *BridgeRegistryCallerSession) Valset() (common.Address, error) {
	return _BridgeRegistry.Contract.Valset(&_BridgeRegistry.CallOpts)
}

// BridgeRegistryLogContractsRegisteredIterator is returned from FilterLogContractsRegistered and is used to iterate over the raw logs and unpacked data for LogContractsRegistered events raised by the BridgeRegistry contract.
type BridgeRegistryLogContractsRegisteredIterator struct {
	Event *BridgeRegistryLogContractsRegistered // Event containing the contract specifics and raw log

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
func (it *BridgeRegistryLogContractsRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRegistryLogContractsRegistered)
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
		it.Event = new(BridgeRegistryLogContractsRegistered)
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
func (it *BridgeRegistryLogContractsRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRegistryLogContractsRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRegistryLogContractsRegistered represents a LogContractsRegistered event raised by the BridgeRegistry contract.
type BridgeRegistryLogContractsRegistered struct {
	QuantiexERC20Bridge  common.Address
	QuantiexERC721Bridge common.Address
	BridgeERC20Bank      common.Address
	BridgeERC721Bank     common.Address
	Oracle               common.Address
	Valset               common.Address
	StakingPool          common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterLogContractsRegistered is a free log retrieval operation binding the contract event 0x626ca7cf58dcae558e83ad85a3cbf2c646a3f55925624c86f40d3b0ab55bf5d9.
//
// Solidity: event LogContractsRegistered(address _quantiexERC20Bridge, address _quantiexERC721Bridge, address _bridgeERC20Bank, address _bridgeERC721Bank, address _oracle, address _valset, address _stakingPool)
func (_BridgeRegistry *BridgeRegistryFilterer) FilterLogContractsRegistered(opts *bind.FilterOpts) (*BridgeRegistryLogContractsRegisteredIterator, error) {

	logs, sub, err := _BridgeRegistry.contract.FilterLogs(opts, "LogContractsRegistered")
	if err != nil {
		return nil, err
	}
	return &BridgeRegistryLogContractsRegisteredIterator{contract: _BridgeRegistry.contract, event: "LogContractsRegistered", logs: logs, sub: sub}, nil
}

// WatchLogContractsRegistered is a free log subscription operation binding the contract event 0x626ca7cf58dcae558e83ad85a3cbf2c646a3f55925624c86f40d3b0ab55bf5d9.
//
// Solidity: event LogContractsRegistered(address _quantiexERC20Bridge, address _quantiexERC721Bridge, address _bridgeERC20Bank, address _bridgeERC721Bank, address _oracle, address _valset, address _stakingPool)
func (_BridgeRegistry *BridgeRegistryFilterer) WatchLogContractsRegistered(opts *bind.WatchOpts, sink chan<- *BridgeRegistryLogContractsRegistered) (event.Subscription, error) {

	logs, sub, err := _BridgeRegistry.contract.WatchLogs(opts, "LogContractsRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRegistryLogContractsRegistered)
				if err := _BridgeRegistry.contract.UnpackLog(event, "LogContractsRegistered", log); err != nil {
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

// ParseLogContractsRegistered is a log parse operation binding the contract event 0x626ca7cf58dcae558e83ad85a3cbf2c646a3f55925624c86f40d3b0ab55bf5d9.
//
// Solidity: event LogContractsRegistered(address _quantiexERC20Bridge, address _quantiexERC721Bridge, address _bridgeERC20Bank, address _bridgeERC721Bank, address _oracle, address _valset, address _stakingPool)
func (_BridgeRegistry *BridgeRegistryFilterer) ParseLogContractsRegistered(log types.Log) (*BridgeRegistryLogContractsRegistered, error) {
	event := new(BridgeRegistryLogContractsRegistered)
	if err := _BridgeRegistry.contract.UnpackLog(event, "LogContractsRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}
