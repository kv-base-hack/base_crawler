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

// SwapTopic1MetaData contains all meta data concerning the SwapTopic1 contract.
var SwapTopic1MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"}]",
}

// SwapTopic1ABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapTopic1MetaData.ABI instead.
var SwapTopic1ABI = SwapTopic1MetaData.ABI

// SwapTopic1 is an auto generated Go binding around an Ethereum contract.
type SwapTopic1 struct {
	SwapTopic1Caller     // Read-only binding to the contract
	SwapTopic1Transactor // Write-only binding to the contract
	SwapTopic1Filterer   // Log filterer for contract events
}

// SwapTopic1Caller is an auto generated read-only Go binding around an Ethereum contract.
type SwapTopic1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTopic1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapTopic1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTopic1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapTopic1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTopic1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapTopic1Session struct {
	Contract     *SwapTopic1       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapTopic1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapTopic1CallerSession struct {
	Contract *SwapTopic1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SwapTopic1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapTopic1TransactorSession struct {
	Contract     *SwapTopic1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SwapTopic1Raw is an auto generated low-level Go binding around an Ethereum contract.
type SwapTopic1Raw struct {
	Contract *SwapTopic1 // Generic contract binding to access the raw methods on
}

// SwapTopic1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapTopic1CallerRaw struct {
	Contract *SwapTopic1Caller // Generic read-only contract binding to access the raw methods on
}

// SwapTopic1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapTopic1TransactorRaw struct {
	Contract *SwapTopic1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapTopic1 creates a new instance of SwapTopic1, bound to a specific deployed contract.
func NewSwapTopic1(address common.Address, backend bind.ContractBackend) (*SwapTopic1, error) {
	contract, err := bindSwapTopic1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapTopic1{SwapTopic1Caller: SwapTopic1Caller{contract: contract}, SwapTopic1Transactor: SwapTopic1Transactor{contract: contract}, SwapTopic1Filterer: SwapTopic1Filterer{contract: contract}}, nil
}

// NewSwapTopic1Caller creates a new read-only instance of SwapTopic1, bound to a specific deployed contract.
func NewSwapTopic1Caller(address common.Address, caller bind.ContractCaller) (*SwapTopic1Caller, error) {
	contract, err := bindSwapTopic1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTopic1Caller{contract: contract}, nil
}

// NewSwapTopic1Transactor creates a new write-only instance of SwapTopic1, bound to a specific deployed contract.
func NewSwapTopic1Transactor(address common.Address, transactor bind.ContractTransactor) (*SwapTopic1Transactor, error) {
	contract, err := bindSwapTopic1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTopic1Transactor{contract: contract}, nil
}

// NewSwapTopic1Filterer creates a new log filterer instance of SwapTopic1, bound to a specific deployed contract.
func NewSwapTopic1Filterer(address common.Address, filterer bind.ContractFilterer) (*SwapTopic1Filterer, error) {
	contract, err := bindSwapTopic1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapTopic1Filterer{contract: contract}, nil
}

// bindSwapTopic1 binds a generic wrapper to an already deployed contract.
func bindSwapTopic1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapTopic1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapTopic1 *SwapTopic1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapTopic1.Contract.SwapTopic1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapTopic1 *SwapTopic1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapTopic1.Contract.SwapTopic1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapTopic1 *SwapTopic1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapTopic1.Contract.SwapTopic1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapTopic1 *SwapTopic1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapTopic1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapTopic1 *SwapTopic1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapTopic1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapTopic1 *SwapTopic1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapTopic1.Contract.contract.Transact(opts, method, params...)
}

// SwapTopic1SwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the SwapTopic1 contract.
type SwapTopic1SwapIterator struct {
	Event *SwapTopic1Swap // Event containing the contract specifics and raw log

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
func (it *SwapTopic1SwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapTopic1Swap)
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
		it.Event = new(SwapTopic1Swap)
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
func (it *SwapTopic1SwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapTopic1SwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapTopic1Swap represents a Swap event raised by the SwapTopic1 contract.
type SwapTopic1Swap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_SwapTopic1 *SwapTopic1Filterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*SwapTopic1SwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SwapTopic1.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SwapTopic1SwapIterator{contract: _SwapTopic1.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_SwapTopic1 *SwapTopic1Filterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *SwapTopic1Swap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SwapTopic1.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapTopic1Swap)
				if err := _SwapTopic1.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_SwapTopic1 *SwapTopic1Filterer) ParseSwap(log types.Log) (*SwapTopic1Swap, error) {
	event := new(SwapTopic1Swap)
	if err := _SwapTopic1.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
