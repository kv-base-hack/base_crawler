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

// Transfer1MetaData contains all meta data concerning the Transfer1 contract.
var Transfer1MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]",
}

// Transfer1ABI is the input ABI used to generate the binding from.
// Deprecated: Use Transfer1MetaData.ABI instead.
var Transfer1ABI = Transfer1MetaData.ABI

// Transfer1 is an auto generated Go binding around an Ethereum contract.
type Transfer1 struct {
	Transfer1Caller     // Read-only binding to the contract
	Transfer1Transactor // Write-only binding to the contract
	Transfer1Filterer   // Log filterer for contract events
}

// Transfer1Caller is an auto generated read-only Go binding around an Ethereum contract.
type Transfer1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Transfer1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Transfer1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Transfer1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Transfer1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Transfer1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Transfer1Session struct {
	Contract     *Transfer1        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Transfer1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Transfer1CallerSession struct {
	Contract *Transfer1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Transfer1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Transfer1TransactorSession struct {
	Contract     *Transfer1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Transfer1Raw is an auto generated low-level Go binding around an Ethereum contract.
type Transfer1Raw struct {
	Contract *Transfer1 // Generic contract binding to access the raw methods on
}

// Transfer1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Transfer1CallerRaw struct {
	Contract *Transfer1Caller // Generic read-only contract binding to access the raw methods on
}

// Transfer1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Transfer1TransactorRaw struct {
	Contract *Transfer1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTransfer1 creates a new instance of Transfer1, bound to a specific deployed contract.
func NewTransfer1(address common.Address, backend bind.ContractBackend) (*Transfer1, error) {
	contract, err := bindTransfer1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transfer1{Transfer1Caller: Transfer1Caller{contract: contract}, Transfer1Transactor: Transfer1Transactor{contract: contract}, Transfer1Filterer: Transfer1Filterer{contract: contract}}, nil
}

// NewTransfer1Caller creates a new read-only instance of Transfer1, bound to a specific deployed contract.
func NewTransfer1Caller(address common.Address, caller bind.ContractCaller) (*Transfer1Caller, error) {
	contract, err := bindTransfer1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Transfer1Caller{contract: contract}, nil
}

// NewTransfer1Transactor creates a new write-only instance of Transfer1, bound to a specific deployed contract.
func NewTransfer1Transactor(address common.Address, transactor bind.ContractTransactor) (*Transfer1Transactor, error) {
	contract, err := bindTransfer1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Transfer1Transactor{contract: contract}, nil
}

// NewTransfer1Filterer creates a new log filterer instance of Transfer1, bound to a specific deployed contract.
func NewTransfer1Filterer(address common.Address, filterer bind.ContractFilterer) (*Transfer1Filterer, error) {
	contract, err := bindTransfer1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Transfer1Filterer{contract: contract}, nil
}

// bindTransfer1 binds a generic wrapper to an already deployed contract.
func bindTransfer1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Transfer1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer1 *Transfer1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transfer1.Contract.Transfer1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer1 *Transfer1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer1.Contract.Transfer1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer1 *Transfer1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer1.Contract.Transfer1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer1 *Transfer1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transfer1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer1 *Transfer1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer1 *Transfer1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer1.Contract.contract.Transact(opts, method, params...)
}

// Transfer1TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Transfer1 contract.
type Transfer1TransferIterator struct {
	Event *Transfer1Transfer // Event containing the contract specifics and raw log

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
func (it *Transfer1TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Transfer1Transfer)
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
		it.Event = new(Transfer1Transfer)
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
func (it *Transfer1TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Transfer1TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Transfer1Transfer represents a Transfer event raised by the Transfer1 contract.
type Transfer1Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Transfer1 *Transfer1Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Transfer1TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Transfer1.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Transfer1TransferIterator{contract: _Transfer1.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Transfer1 *Transfer1Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Transfer1Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Transfer1.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Transfer1Transfer)
				if err := _Transfer1.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Transfer1 *Transfer1Filterer) ParseTransfer(log types.Log) (*Transfer1Transfer, error) {
	event := new(Transfer1Transfer)
	if err := _Transfer1.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
