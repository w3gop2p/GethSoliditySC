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

// TokenERC20MetaData contains all meta data concerning the TokenERC20 contract.
var TokenERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526040518060400160405280600c81526020017f4d79546f6b656e455243323000000000000000000000000000000000000000008152505f908161004791906103ef565b506040518060400160405280600381526020017f4d544b00000000000000000000000000000000000000000000000000000000008152506001908161008c91906103ef565b50601260025f6101000a81548160ff021916908360ff1602179055503480156100b3575f80fd5b5060405161152a38038061152a83398181016040528101906100d591906104ec565b60025f9054906101000a900460ff1660ff16600a6100f39190610673565b816100fe91906106bd565b60038190555060035460045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055503373ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6003546040516101a7919061070d565b60405180910390a350610726565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061023057607f821691505b602082108103610243576102426101ec565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026102a57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261026a565b6102af868361026a565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6102f36102ee6102e9846102c7565b6102d0565b6102c7565b9050919050565b5f819050919050565b61030c836102d9565b610320610318826102fa565b848454610276565b825550505050565b5f90565b610334610328565b61033f818484610303565b505050565b5b81811015610362576103575f8261032c565b600181019050610345565b5050565b601f8211156103a75761037881610249565b6103818461025b565b81016020851015610390578190505b6103a461039c8561025b565b830182610344565b50505b505050565b5f82821c905092915050565b5f6103c75f19846008026103ac565b1980831691505092915050565b5f6103df83836103b8565b9150826002028217905092915050565b6103f8826101b5565b67ffffffffffffffff811115610411576104106101bf565b5b61041b8254610219565b610426828285610366565b5f60209050601f831160018114610457575f8415610445578287015190505b61044f85826103d4565b8655506104b6565b601f19841661046586610249565b5f5b8281101561048c57848901518255600182019150602085019450602081019050610467565b868310156104a957848901516104a5601f8916826103b8565b8355505b6001600288020188555050505b505050505050565b5f80fd5b6104cb816102c7565b81146104d5575f80fd5b50565b5f815190506104e6816104c2565b92915050565b5f60208284031215610501576105006104be565b5b5f61050e848285016104d8565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f8160011c9050919050565b5f808291508390505b60018511156105995780860481111561057557610574610517565b5b60018516156105845780820291505b808102905061059285610544565b9450610559565b94509492505050565b5f826105b1576001905061066c565b816105be575f905061066c565b81600181146105d457600281146105de5761060d565b600191505061066c565b60ff8411156105f0576105ef610517565b5b8360020a91508482111561060757610606610517565b5b5061066c565b5060208310610133831016604e8410600b84101617156106425782820a90508381111561063d5761063c610517565b5b61066c565b61064f8484846001610550565b9250905081840481111561066657610665610517565b5b81810290505b9392505050565b5f61067d826102c7565b9150610688836102c7565b92506106b57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846105a2565b905092915050565b5f6106c7826102c7565b91506106d2836102c7565b92508282026106e0816102c7565b915082820484148315176106f7576106f6610517565b5b5092915050565b610707816102c7565b82525050565b5f6020820190506107205f8301846106fe565b92915050565b610df7806107335f395ff3fe608060405234801561000f575f80fd5b5060043610610091575f3560e01c8063313ce56711610064578063313ce5671461013157806370a082311461014f57806395d89b411461017f578063a9059cbb1461019d578063dd62ed3e146101cd57610091565b806306fdde0314610095578063095ea7b3146100b357806318160ddd146100e357806323b872dd14610101575b5f80fd5b61009d6101fd565b6040516100aa91906109ca565b60405180910390f35b6100cd60048036038101906100c89190610a7b565b610288565b6040516100da9190610ad3565b60405180910390f35b6100eb610375565b6040516100f89190610afb565b60405180910390f35b61011b60048036038101906101169190610b14565b61037e565b6040516101289190610ad3565b60405180910390f35b61013961065e565b6040516101469190610b7f565b60405180910390f35b61016960048036038101906101649190610b98565b610670565b6040516101769190610afb565b60405180910390f35b6101876106b6565b60405161019491906109ca565b60405180910390f35b6101b760048036038101906101b29190610a7b565b610742565b6040516101c49190610ad3565b60405180910390f35b6101e760048036038101906101e29190610bc3565b6108d8565b6040516101f49190610afb565b60405180910390f35b5f805461020990610c2e565b80601f016020809104026020016040519081016040528092919081815260200182805461023590610c2e565b80156102805780601f1061025757610100808354040283529160200191610280565b820191905f5260205f20905b81548152906001019060200180831161026357829003601f168201915b505050505081565b5f8160055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516103639190610afb565b60405180910390a36001905092915050565b5f600354905090565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156103ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f690610ca8565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156104ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b190610d10565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546105069190610d5b565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546105599190610d8e565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546105e79190610d5b565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161064b9190610afb565b60405180910390a3600190509392505050565b60025f9054906101000a900460ff1681565b5f60045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b600180546106c390610c2e565b80601f01602080910402602001604051908101604052809291908181526020018280546106ef90610c2e565b801561073a5780601f106107115761010080835404028352916020019161073a565b820191905f5260205f20905b81548152906001019060200180831161071d57829003601f168201915b505050505081565b5f8160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156107c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ba90610ca8565b60405180910390fd5b8160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461080f9190610d5b565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546108629190610d8e565b925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516108c69190610afb565b60405180910390a36001905092915050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61099c8261095a565b6109a68185610964565b93506109b6818560208601610974565b6109bf81610982565b840191505092915050565b5f6020820190508181035f8301526109e28184610992565b905092915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610a17826109ee565b9050919050565b610a2781610a0d565b8114610a31575f80fd5b50565b5f81359050610a4281610a1e565b92915050565b5f819050919050565b610a5a81610a48565b8114610a64575f80fd5b50565b5f81359050610a7581610a51565b92915050565b5f8060408385031215610a9157610a906109ea565b5b5f610a9e85828601610a34565b9250506020610aaf85828601610a67565b9150509250929050565b5f8115159050919050565b610acd81610ab9565b82525050565b5f602082019050610ae65f830184610ac4565b92915050565b610af581610a48565b82525050565b5f602082019050610b0e5f830184610aec565b92915050565b5f805f60608486031215610b2b57610b2a6109ea565b5b5f610b3886828701610a34565b9350506020610b4986828701610a34565b9250506040610b5a86828701610a67565b9150509250925092565b5f60ff82169050919050565b610b7981610b64565b82525050565b5f602082019050610b925f830184610b70565b92915050565b5f60208284031215610bad57610bac6109ea565b5b5f610bba84828501610a34565b91505092915050565b5f8060408385031215610bd957610bd86109ea565b5b5f610be685828601610a34565b9250506020610bf785828601610a34565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610c4557607f821691505b602082108103610c5857610c57610c01565b5b50919050565b7f496e73756666696369656e742062616c616e63650000000000000000000000005f82015250565b5f610c92601483610964565b9150610c9d82610c5e565b602082019050919050565b5f6020820190508181035f830152610cbf81610c86565b9050919050565b7f416c6c6f77616e636520657863656564656400000000000000000000000000005f82015250565b5f610cfa601283610964565b9150610d0582610cc6565b602082019050919050565b5f6020820190508181035f830152610d2781610cee565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610d6582610a48565b9150610d7083610a48565b9250828203905081811115610d8857610d87610d2e565b5b92915050565b5f610d9882610a48565b9150610da383610a48565b9250828201905080821115610dbb57610dba610d2e565b5b9291505056fea2646970667358221220ca02b6a0383b69b8d063e35aefb7b345c4a893e1ba7b9a21db0fb9913ae2837164736f6c634300081a0033",
}

// TokenERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenERC20MetaData.ABI instead.
var TokenERC20ABI = TokenERC20MetaData.ABI

// TokenERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenERC20MetaData.Bin instead.
var TokenERC20Bin = TokenERC20MetaData.Bin

// DeployTokenERC20 deploys a new Ethereum contract, binding an instance of TokenERC20 to it.
func DeployTokenERC20(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *TokenERC20, error) {
	parsed, err := TokenERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenERC20Bin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenERC20{TokenERC20Caller: TokenERC20Caller{contract: contract}, TokenERC20Transactor: TokenERC20Transactor{contract: contract}, TokenERC20Filterer: TokenERC20Filterer{contract: contract}}, nil
}

// TokenERC20 is an auto generated Go binding around an Ethereum contract.
type TokenERC20 struct {
	TokenERC20Caller     // Read-only binding to the contract
	TokenERC20Transactor // Write-only binding to the contract
	TokenERC20Filterer   // Log filterer for contract events
}

// TokenERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type TokenERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenERC20Session struct {
	Contract     *TokenERC20       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenERC20CallerSession struct {
	Contract *TokenERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenERC20TransactorSession struct {
	Contract     *TokenERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type TokenERC20Raw struct {
	Contract *TokenERC20 // Generic contract binding to access the raw methods on
}

// TokenERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenERC20CallerRaw struct {
	Contract *TokenERC20Caller // Generic read-only contract binding to access the raw methods on
}

// TokenERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenERC20TransactorRaw struct {
	Contract *TokenERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenERC20 creates a new instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20(address common.Address, backend bind.ContractBackend) (*TokenERC20, error) {
	contract, err := bindTokenERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenERC20{TokenERC20Caller: TokenERC20Caller{contract: contract}, TokenERC20Transactor: TokenERC20Transactor{contract: contract}, TokenERC20Filterer: TokenERC20Filterer{contract: contract}}, nil
}

// NewTokenERC20Caller creates a new read-only instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20Caller(address common.Address, caller bind.ContractCaller) (*TokenERC20Caller, error) {
	contract, err := bindTokenERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenERC20Caller{contract: contract}, nil
}

// NewTokenERC20Transactor creates a new write-only instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*TokenERC20Transactor, error) {
	contract, err := bindTokenERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenERC20Transactor{contract: contract}, nil
}

// NewTokenERC20Filterer creates a new log filterer instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*TokenERC20Filterer, error) {
	contract, err := bindTokenERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenERC20Filterer{contract: contract}, nil
}

// bindTokenERC20 binds a generic wrapper to an already deployed contract.
func bindTokenERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenERC20 *TokenERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenERC20.Contract.TokenERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenERC20 *TokenERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenERC20.Contract.TokenERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenERC20 *TokenERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenERC20.Contract.TokenERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenERC20 *TokenERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenERC20 *TokenERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenERC20 *TokenERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenERC20 *TokenERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenERC20 *TokenERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TokenERC20.Contract.Allowance(&_TokenERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenERC20 *TokenERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TokenERC20.Contract.Allowance(&_TokenERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenERC20 *TokenERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenERC20 *TokenERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _TokenERC20.Contract.BalanceOf(&_TokenERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenERC20 *TokenERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TokenERC20.Contract.BalanceOf(&_TokenERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenERC20 *TokenERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TokenERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenERC20 *TokenERC20Session) Decimals() (uint8, error) {
	return _TokenERC20.Contract.Decimals(&_TokenERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenERC20 *TokenERC20CallerSession) Decimals() (uint8, error) {
	return _TokenERC20.Contract.Decimals(&_TokenERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenERC20 *TokenERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenERC20 *TokenERC20Session) Name() (string, error) {
	return _TokenERC20.Contract.Name(&_TokenERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenERC20 *TokenERC20CallerSession) Name() (string, error) {
	return _TokenERC20.Contract.Name(&_TokenERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenERC20 *TokenERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenERC20 *TokenERC20Session) Symbol() (string, error) {
	return _TokenERC20.Contract.Symbol(&_TokenERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenERC20 *TokenERC20CallerSession) Symbol() (string, error) {
	return _TokenERC20.Contract.Symbol(&_TokenERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenERC20 *TokenERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenERC20 *TokenERC20Session) TotalSupply() (*big.Int, error) {
	return _TokenERC20.Contract.TotalSupply(&_TokenERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenERC20 *TokenERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _TokenERC20.Contract.TotalSupply(&_TokenERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Approve(&_TokenERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Approve(&_TokenERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Transfer(&_TokenERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Transfer(&_TokenERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.TransferFrom(&_TokenERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.TransferFrom(&_TokenERC20.TransactOpts, sender, recipient, amount)
}

// TokenERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TokenERC20 contract.
type TokenERC20ApprovalIterator struct {
	Event *TokenERC20Approval // Event containing the contract specifics and raw log

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
func (it *TokenERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC20Approval)
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
		it.Event = new(TokenERC20Approval)
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
func (it *TokenERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC20Approval represents a Approval event raised by the TokenERC20 contract.
type TokenERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TokenERC20 *TokenERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC20ApprovalIterator{contract: _TokenERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TokenERC20 *TokenERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC20Approval)
				if err := _TokenERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TokenERC20 *TokenERC20Filterer) ParseApproval(log types.Log) (*TokenERC20Approval, error) {
	event := new(TokenERC20Approval)
	if err := _TokenERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TokenERC20 contract.
type TokenERC20TransferIterator struct {
	Event *TokenERC20Transfer // Event containing the contract specifics and raw log

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
func (it *TokenERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC20Transfer)
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
		it.Event = new(TokenERC20Transfer)
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
func (it *TokenERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC20Transfer represents a Transfer event raised by the TokenERC20 contract.
type TokenERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TokenERC20 *TokenERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC20TransferIterator{contract: _TokenERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TokenERC20 *TokenERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC20Transfer)
				if err := _TokenERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_TokenERC20 *TokenERC20Filterer) ParseTransfer(log types.Log) (*TokenERC20Transfer, error) {
	event := new(TokenERC20Transfer)
	if err := _TokenERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
