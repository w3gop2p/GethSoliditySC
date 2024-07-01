// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigen

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

// EthSenderMetaData contains all meta data concerning the EthSender contract.
var EthSenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthSent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getTotalSent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sendEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalSent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506106e78061005b5f395ff3fe60806040526004361061004d575f3560e01c806306e99fef1461005857806308f81a931461006e5780632e1a7d4d146100aa578063806c0b2b146100d25780638da5cb5b1461010e57610054565b3661005457005b5f80fd5b348015610063575f80fd5b5061006c610138565b005b348015610079575f80fd5b50610094600480360381019061008f9190610498565b610286565b6040516100a191906104db565b60405180910390f35b3480156100b5575f80fd5b506100d060048036038101906100cb919061051e565b61029b565b005b3480156100dd575f80fd5b506100f860048036038101906100f39190610498565b6103d1565b60405161010591906104db565b60405180910390f35b348015610119575f80fd5b50610122610417565b60405161012f9190610558565b60405180910390f35b680246ddf97976680000471015610184576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017b906105cb565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc680246ddf9797668000090811502906040515f60405180830381858888f193505050501580156101d0573d5f803e3d5ffd5b50680246ddf9797668000060015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546102269190610616565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167f78f5cdad99320ec2ba57132d7dffb1d125775c823239e60ff5e9300fd4ac898c680246ddf9797668000060405161027c91906104db565b60405180910390a2565b6001602052805f5260405f205f915090505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610328576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031f90610693565b60405180910390fd5b8047101561036b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610362906105cb565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f193505050501580156103cd573d5f803e3d5ffd5b5050565b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6104678261043e565b9050919050565b6104778161045d565b8114610481575f80fd5b50565b5f813590506104928161046e565b92915050565b5f602082840312156104ad576104ac61043a565b5b5f6104ba84828501610484565b91505092915050565b5f819050919050565b6104d5816104c3565b82525050565b5f6020820190506104ee5f8301846104cc565b92915050565b6104fd816104c3565b8114610507575f80fd5b50565b5f81359050610518816104f4565b92915050565b5f602082840312156105335761053261043a565b5b5f6105408482850161050a565b91505092915050565b6105528161045d565b82525050565b5f60208201905061056b5f830184610549565b92915050565b5f82825260208201905092915050565b7f4e6f7420656e6f7567682062616c616e636520696e20636f6e747261637400005f82015250565b5f6105b5601e83610571565b91506105c082610581565b602082019050919050565b5f6020820190508181035f8301526105e2816105a9565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610620826104c3565b915061062b836104c3565b9250828201905080821115610643576106426105e9565b5b92915050565b7f43616c6c6572206973206e6f7420746865206f776e65720000000000000000005f82015250565b5f61067d601783610571565b915061068882610649565b602082019050919050565b5f6020820190508181035f8301526106aa81610671565b905091905056fea26469706673582212203466d0d51680d94ad2c74f0e1517ecaaebf306b489be5558034bf26ffbf4f41b64736f6c634300081a0033",
}

// EthSenderABI is the input ABI used to generate the binding from.
// Deprecated: Use EthSenderMetaData.ABI instead.
var EthSenderABI = EthSenderMetaData.ABI

// EthSenderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthSenderMetaData.Bin instead.
var EthSenderBin = EthSenderMetaData.Bin

// DeployEthSender deploys a new Ethereum contract, binding an instance of EthSender to it.
func DeployEthSender(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthSender, error) {
	parsed, err := EthSenderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthSenderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthSender{EthSenderCaller: EthSenderCaller{contract: contract}, EthSenderTransactor: EthSenderTransactor{contract: contract}, EthSenderFilterer: EthSenderFilterer{contract: contract}}, nil
}

// EthSender is an auto generated Go binding around an Ethereum contract.
type EthSender struct {
	EthSenderCaller     // Read-only binding to the contract
	EthSenderTransactor // Write-only binding to the contract
	EthSenderFilterer   // Log filterer for contract events
}

// EthSenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthSenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthSenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthSenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthSenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthSenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthSenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthSenderSession struct {
	Contract     *EthSender        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthSenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthSenderCallerSession struct {
	Contract *EthSenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EthSenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthSenderTransactorSession struct {
	Contract     *EthSenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthSenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthSenderRaw struct {
	Contract *EthSender // Generic contract binding to access the raw methods on
}

// EthSenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthSenderCallerRaw struct {
	Contract *EthSenderCaller // Generic read-only contract binding to access the raw methods on
}

// EthSenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthSenderTransactorRaw struct {
	Contract *EthSenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthSender creates a new instance of EthSender, bound to a specific deployed contract.
func NewEthSender(address common.Address, backend bind.ContractBackend) (*EthSender, error) {
	contract, err := bindEthSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthSender{EthSenderCaller: EthSenderCaller{contract: contract}, EthSenderTransactor: EthSenderTransactor{contract: contract}, EthSenderFilterer: EthSenderFilterer{contract: contract}}, nil
}

// NewEthSenderCaller creates a new read-only instance of EthSender, bound to a specific deployed contract.
func NewEthSenderCaller(address common.Address, caller bind.ContractCaller) (*EthSenderCaller, error) {
	contract, err := bindEthSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthSenderCaller{contract: contract}, nil
}

// NewEthSenderTransactor creates a new write-only instance of EthSender, bound to a specific deployed contract.
func NewEthSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*EthSenderTransactor, error) {
	contract, err := bindEthSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthSenderTransactor{contract: contract}, nil
}

// NewEthSenderFilterer creates a new log filterer instance of EthSender, bound to a specific deployed contract.
func NewEthSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*EthSenderFilterer, error) {
	contract, err := bindEthSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthSenderFilterer{contract: contract}, nil
}

// bindEthSender binds a generic wrapper to an already deployed contract.
func bindEthSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthSenderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthSender *EthSenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthSender.Contract.EthSenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthSender *EthSenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthSender.Contract.EthSenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthSender *EthSenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthSender.Contract.EthSenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthSender *EthSenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthSender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthSender *EthSenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthSender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthSender *EthSenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthSender.Contract.contract.Transact(opts, method, params...)
}

// GetTotalSent is a free data retrieval call binding the contract method 0x806c0b2b.
//
// Solidity: function getTotalSent(address _address) view returns(uint256)
func (_EthSender *EthSenderCaller) GetTotalSent(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EthSender.contract.Call(opts, &out, "getTotalSent", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSent is a free data retrieval call binding the contract method 0x806c0b2b.
//
// Solidity: function getTotalSent(address _address) view returns(uint256)
func (_EthSender *EthSenderSession) GetTotalSent(_address common.Address) (*big.Int, error) {
	return _EthSender.Contract.GetTotalSent(&_EthSender.CallOpts, _address)
}

// GetTotalSent is a free data retrieval call binding the contract method 0x806c0b2b.
//
// Solidity: function getTotalSent(address _address) view returns(uint256)
func (_EthSender *EthSenderCallerSession) GetTotalSent(_address common.Address) (*big.Int, error) {
	return _EthSender.Contract.GetTotalSent(&_EthSender.CallOpts, _address)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthSender *EthSenderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthSender.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthSender *EthSenderSession) Owner() (common.Address, error) {
	return _EthSender.Contract.Owner(&_EthSender.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthSender *EthSenderCallerSession) Owner() (common.Address, error) {
	return _EthSender.Contract.Owner(&_EthSender.CallOpts)
}

// TotalSent is a free data retrieval call binding the contract method 0x08f81a93.
//
// Solidity: function totalSent(address ) view returns(uint256)
func (_EthSender *EthSenderCaller) TotalSent(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EthSender.contract.Call(opts, &out, "totalSent", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSent is a free data retrieval call binding the contract method 0x08f81a93.
//
// Solidity: function totalSent(address ) view returns(uint256)
func (_EthSender *EthSenderSession) TotalSent(arg0 common.Address) (*big.Int, error) {
	return _EthSender.Contract.TotalSent(&_EthSender.CallOpts, arg0)
}

// TotalSent is a free data retrieval call binding the contract method 0x08f81a93.
//
// Solidity: function totalSent(address ) view returns(uint256)
func (_EthSender *EthSenderCallerSession) TotalSent(arg0 common.Address) (*big.Int, error) {
	return _EthSender.Contract.TotalSent(&_EthSender.CallOpts, arg0)
}

// SendEth is a paid mutator transaction binding the contract method 0x06e99fef.
//
// Solidity: function sendEth() returns()
func (_EthSender *EthSenderTransactor) SendEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthSender.contract.Transact(opts, "sendEth")
}

// SendEth is a paid mutator transaction binding the contract method 0x06e99fef.
//
// Solidity: function sendEth() returns()
func (_EthSender *EthSenderSession) SendEth() (*types.Transaction, error) {
	return _EthSender.Contract.SendEth(&_EthSender.TransactOpts)
}

// SendEth is a paid mutator transaction binding the contract method 0x06e99fef.
//
// Solidity: function sendEth() returns()
func (_EthSender *EthSenderTransactorSession) SendEth() (*types.Transaction, error) {
	return _EthSender.Contract.SendEth(&_EthSender.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_EthSender *EthSenderTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _EthSender.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_EthSender *EthSenderSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _EthSender.Contract.Withdraw(&_EthSender.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_EthSender *EthSenderTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _EthSender.Contract.Withdraw(&_EthSender.TransactOpts, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EthSender *EthSenderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthSender.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EthSender *EthSenderSession) Receive() (*types.Transaction, error) {
	return _EthSender.Contract.Receive(&_EthSender.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EthSender *EthSenderTransactorSession) Receive() (*types.Transaction, error) {
	return _EthSender.Contract.Receive(&_EthSender.TransactOpts)
}

// EthSenderEthSentIterator is returned from FilterEthSent and is used to iterate over the raw logs and unpacked data for EthSent events raised by the EthSender contract.
type EthSenderEthSentIterator struct {
	Event *EthSenderEthSent // Event containing the contract specifics and raw log

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
func (it *EthSenderEthSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthSenderEthSent)
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
		it.Event = new(EthSenderEthSent)
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
func (it *EthSenderEthSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthSenderEthSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthSenderEthSent represents a EthSent event raised by the EthSender contract.
type EthSenderEthSent struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthSent is a free log retrieval operation binding the contract event 0x78f5cdad99320ec2ba57132d7dffb1d125775c823239e60ff5e9300fd4ac898c.
//
// Solidity: event EthSent(address indexed recipient, uint256 amount)
func (_EthSender *EthSenderFilterer) FilterEthSent(opts *bind.FilterOpts, recipient []common.Address) (*EthSenderEthSentIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EthSender.contract.FilterLogs(opts, "EthSent", recipientRule)
	if err != nil {
		return nil, err
	}
	return &EthSenderEthSentIterator{contract: _EthSender.contract, event: "EthSent", logs: logs, sub: sub}, nil
}

// WatchEthSent is a free log subscription operation binding the contract event 0x78f5cdad99320ec2ba57132d7dffb1d125775c823239e60ff5e9300fd4ac898c.
//
// Solidity: event EthSent(address indexed recipient, uint256 amount)
func (_EthSender *EthSenderFilterer) WatchEthSent(opts *bind.WatchOpts, sink chan<- *EthSenderEthSent, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EthSender.contract.WatchLogs(opts, "EthSent", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthSenderEthSent)
				if err := _EthSender.contract.UnpackLog(event, "EthSent", log); err != nil {
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

// ParseEthSent is a log parse operation binding the contract event 0x78f5cdad99320ec2ba57132d7dffb1d125775c823239e60ff5e9300fd4ac898c.
//
// Solidity: event EthSent(address indexed recipient, uint256 amount)
func (_EthSender *EthSenderFilterer) ParseEthSent(log types.Log) (*EthSenderEthSent, error) {
	event := new(EthSenderEthSent)
	if err := _EthSender.contract.UnpackLog(event, "EthSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
