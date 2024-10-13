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

// SwapTopic2MetaData contains all meta data concerning the SwapTopic2 contract.
var SwapTopic2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"}]",
}

// SwapTopic2ABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapTopic2MetaData.ABI instead.
var SwapTopic2ABI = SwapTopic2MetaData.ABI

// SwapTopic2 is an auto generated Go binding around an Ethereum contract.
type SwapTopic2 struct {
	SwapTopic2Caller     // Read-only binding to the contract
	SwapTopic2Transactor // Write-only binding to the contract
	SwapTopic2Filterer   // Log filterer for contract events
}

// SwapTopic2Caller is an auto generated read-only Go binding around an Ethereum contract.
type SwapTopic2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTopic2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapTopic2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTopic2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapTopic2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTopic2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapTopic2Session struct {
	Contract     *SwapTopic2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapTopic2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapTopic2CallerSession struct {
	Contract *SwapTopic2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SwapTopic2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapTopic2TransactorSession struct {
	Contract     *SwapTopic2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SwapTopic2Raw is an auto generated low-level Go binding around an Ethereum contract.
type SwapTopic2Raw struct {
	Contract *SwapTopic2 // Generic contract binding to access the raw methods on
}

// SwapTopic2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapTopic2CallerRaw struct {
	Contract *SwapTopic2Caller // Generic read-only contract binding to access the raw methods on
}

// SwapTopic2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapTopic2TransactorRaw struct {
	Contract *SwapTopic2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapTopic2 creates a new instance of SwapTopic2, bound to a specific deployed contract.
func NewSwapTopic2(address common.Address, backend bind.ContractBackend) (*SwapTopic2, error) {
	contract, err := bindSwapTopic2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapTopic2{SwapTopic2Caller: SwapTopic2Caller{contract: contract}, SwapTopic2Transactor: SwapTopic2Transactor{contract: contract}, SwapTopic2Filterer: SwapTopic2Filterer{contract: contract}}, nil
}

// NewSwapTopic2Caller creates a new read-only instance of SwapTopic2, bound to a specific deployed contract.
func NewSwapTopic2Caller(address common.Address, caller bind.ContractCaller) (*SwapTopic2Caller, error) {
	contract, err := bindSwapTopic2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTopic2Caller{contract: contract}, nil
}

// NewSwapTopic2Transactor creates a new write-only instance of SwapTopic2, bound to a specific deployed contract.
func NewSwapTopic2Transactor(address common.Address, transactor bind.ContractTransactor) (*SwapTopic2Transactor, error) {
	contract, err := bindSwapTopic2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTopic2Transactor{contract: contract}, nil
}

// NewSwapTopic2Filterer creates a new log filterer instance of SwapTopic2, bound to a specific deployed contract.
func NewSwapTopic2Filterer(address common.Address, filterer bind.ContractFilterer) (*SwapTopic2Filterer, error) {
	contract, err := bindSwapTopic2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapTopic2Filterer{contract: contract}, nil
}

// bindSwapTopic2 binds a generic wrapper to an already deployed contract.
func bindSwapTopic2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapTopic2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapTopic2 *SwapTopic2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapTopic2.Contract.SwapTopic2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapTopic2 *SwapTopic2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapTopic2.Contract.SwapTopic2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapTopic2 *SwapTopic2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapTopic2.Contract.SwapTopic2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapTopic2 *SwapTopic2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapTopic2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapTopic2 *SwapTopic2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapTopic2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapTopic2 *SwapTopic2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapTopic2.Contract.contract.Transact(opts, method, params...)
}

// SwapTopic2SwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the SwapTopic2 contract.
type SwapTopic2SwapIterator struct {
	Event *SwapTopic2Swap // Event containing the contract specifics and raw log

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
func (it *SwapTopic2SwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapTopic2Swap)
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
		it.Event = new(SwapTopic2Swap)
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
func (it *SwapTopic2SwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapTopic2SwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapTopic2Swap represents a Swap event raised by the SwapTopic2 contract.
type SwapTopic2Swap struct {
	Sender       common.Address
	Recipient    common.Address
	Amount0      *big.Int
	Amount1      *big.Int
	SqrtPriceX96 *big.Int
	Liquidity    *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_SwapTopic2 *SwapTopic2Filterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*SwapTopic2SwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SwapTopic2.contract.FilterLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &SwapTopic2SwapIterator{contract: _SwapTopic2.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_SwapTopic2 *SwapTopic2Filterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *SwapTopic2Swap, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SwapTopic2.contract.WatchLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapTopic2Swap)
				if err := _SwapTopic2.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_SwapTopic2 *SwapTopic2Filterer) ParseSwap(log types.Log) (*SwapTopic2Swap, error) {
	event := new(SwapTopic2Swap)
	if err := _SwapTopic2.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
