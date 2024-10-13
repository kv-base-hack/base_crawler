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

// SwapTopic6MetaData contains all meta data concerning the SwapTopic6 contract.
var SwapTopic6MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"protocolFeesToken0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"protocolFeesToken1\",\"type\":\"uint128\"}],\"name\":\"Swap\",\"type\":\"event\"}]",
}

// SwapTopic6ABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapTopic6MetaData.ABI instead.
var SwapTopic6ABI = SwapTopic6MetaData.ABI

// SwapTopic6 is an auto generated Go binding around an Ethereum contract.
type SwapTopic6 struct {
	SwapTopic6Caller     // Read-only binding to the contract
	SwapTopic6Transactor // Write-only binding to the contract
	SwapTopic6Filterer   // Log filterer for contract events
}

// SwapTopic6Caller is an auto generated read-only Go binding around an Ethereum contract.
type SwapTopic6Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTopic6Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapTopic6Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTopic6Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapTopic6Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTopic6Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapTopic6Session struct {
	Contract     *SwapTopic6       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapTopic6CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapTopic6CallerSession struct {
	Contract *SwapTopic6Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SwapTopic6TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapTopic6TransactorSession struct {
	Contract     *SwapTopic6Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SwapTopic6Raw is an auto generated low-level Go binding around an Ethereum contract.
type SwapTopic6Raw struct {
	Contract *SwapTopic6 // Generic contract binding to access the raw methods on
}

// SwapTopic6CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapTopic6CallerRaw struct {
	Contract *SwapTopic6Caller // Generic read-only contract binding to access the raw methods on
}

// SwapTopic6TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapTopic6TransactorRaw struct {
	Contract *SwapTopic6Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapTopic6 creates a new instance of SwapTopic6, bound to a specific deployed contract.
func NewSwapTopic6(address common.Address, backend bind.ContractBackend) (*SwapTopic6, error) {
	contract, err := bindSwapTopic6(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapTopic6{SwapTopic6Caller: SwapTopic6Caller{contract: contract}, SwapTopic6Transactor: SwapTopic6Transactor{contract: contract}, SwapTopic6Filterer: SwapTopic6Filterer{contract: contract}}, nil
}

// NewSwapTopic6Caller creates a new read-only instance of SwapTopic6, bound to a specific deployed contract.
func NewSwapTopic6Caller(address common.Address, caller bind.ContractCaller) (*SwapTopic6Caller, error) {
	contract, err := bindSwapTopic6(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTopic6Caller{contract: contract}, nil
}

// NewSwapTopic6Transactor creates a new write-only instance of SwapTopic6, bound to a specific deployed contract.
func NewSwapTopic6Transactor(address common.Address, transactor bind.ContractTransactor) (*SwapTopic6Transactor, error) {
	contract, err := bindSwapTopic6(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTopic6Transactor{contract: contract}, nil
}

// NewSwapTopic6Filterer creates a new log filterer instance of SwapTopic6, bound to a specific deployed contract.
func NewSwapTopic6Filterer(address common.Address, filterer bind.ContractFilterer) (*SwapTopic6Filterer, error) {
	contract, err := bindSwapTopic6(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapTopic6Filterer{contract: contract}, nil
}

// bindSwapTopic6 binds a generic wrapper to an already deployed contract.
func bindSwapTopic6(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapTopic6MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapTopic6 *SwapTopic6Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapTopic6.Contract.SwapTopic6Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapTopic6 *SwapTopic6Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapTopic6.Contract.SwapTopic6Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapTopic6 *SwapTopic6Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapTopic6.Contract.SwapTopic6Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapTopic6 *SwapTopic6CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapTopic6.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapTopic6 *SwapTopic6TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapTopic6.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapTopic6 *SwapTopic6TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapTopic6.Contract.contract.Transact(opts, method, params...)
}

// SwapTopic6SwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the SwapTopic6 contract.
type SwapTopic6SwapIterator struct {
	Event *SwapTopic6Swap // Event containing the contract specifics and raw log

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
func (it *SwapTopic6SwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapTopic6Swap)
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
		it.Event = new(SwapTopic6Swap)
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
func (it *SwapTopic6SwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapTopic6SwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapTopic6Swap represents a Swap event raised by the SwapTopic6 contract.
type SwapTopic6Swap struct {
	Sender             common.Address
	Recipient          common.Address
	Amount0            *big.Int
	Amount1            *big.Int
	SqrtPriceX96       *big.Int
	Liquidity          *big.Int
	Tick               *big.Int
	ProtocolFeesToken0 *big.Int
	ProtocolFeesToken1 *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick, uint128 protocolFeesToken0, uint128 protocolFeesToken1)
func (_SwapTopic6 *SwapTopic6Filterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*SwapTopic6SwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SwapTopic6.contract.FilterLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &SwapTopic6SwapIterator{contract: _SwapTopic6.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick, uint128 protocolFeesToken0, uint128 protocolFeesToken1)
func (_SwapTopic6 *SwapTopic6Filterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *SwapTopic6Swap, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SwapTopic6.contract.WatchLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapTopic6Swap)
				if err := _SwapTopic6.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick, uint128 protocolFeesToken0, uint128 protocolFeesToken1)
func (_SwapTopic6 *SwapTopic6Filterer) ParseSwap(log types.Log) (*SwapTopic6Swap, error) {
	event := new(SwapTopic6Swap)
	if err := _SwapTopic6.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
