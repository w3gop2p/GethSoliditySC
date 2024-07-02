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

// MyNFTMetaData contains all meta data concerning the MyNFT contract.
var MyNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526040518060400160405280600581526020017f4d794e46540000000000000000000000000000000000000000000000000000008152505f9081610047919061030a565b506040518060400160405280600481526020017f4d4e4654000000000000000000000000000000000000000000000000000000008152506001908161008c919061030a565b50348015610098575f80fd5b50604051611bd7380380611bd783398181016040528101906100ba91906104f9565b80600690816100c9919061030a565b5050610540565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061014b57607f821691505b60208210810361015e5761015d610107565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610185565b6101ca8683610185565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61020e610209610204846101e2565b6101eb565b6101e2565b9050919050565b5f819050919050565b610227836101f4565b61023b61023382610215565b848454610191565b825550505050565b5f90565b61024f610243565b61025a81848461021e565b505050565b5b8181101561027d576102725f82610247565b600181019050610260565b5050565b601f8211156102c25761029381610164565b61029c84610176565b810160208510156102ab578190505b6102bf6102b785610176565b83018261025f565b50505b505050565b5f82821c905092915050565b5f6102e25f19846008026102c7565b1980831691505092915050565b5f6102fa83836102d3565b9150826002028217905092915050565b610313826100d0565b67ffffffffffffffff81111561032c5761032b6100da565b5b6103368254610134565b610341828285610281565b5f60209050601f831160018114610372575f8415610360578287015190505b61036a85826102ef565b8655506103d1565b601f19841661038086610164565b5f5b828110156103a757848901518255600182019150602085019450602081019050610382565b868310156103c457848901516103c0601f8916826102d3565b8355505b6001600288020188555050505b505050505050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b61040b826103f2565b810181811067ffffffffffffffff8211171561042a576104296100da565b5b80604052505050565b5f61043c6103d9565b90506104488282610402565b919050565b5f67ffffffffffffffff821115610467576104666100da565b5b610470826103f2565b9050602081019050919050565b8281835e5f83830152505050565b5f61049d6104988461044d565b610433565b9050828152602081018484840111156104b9576104b86103ee565b5b6104c484828561047d565b509392505050565b5f82601f8301126104e0576104df6103ea565b5b81516104f084826020860161048b565b91505092915050565b5f6020828403121561050e5761050d6103e2565b5b5f82015167ffffffffffffffff81111561052b5761052a6103e6565b5b610537848285016104cc565b91505092915050565b61168a8061054d5f395ff3fe608060405234801561000f575f80fd5b506004361061009c575f3560e01c80636a627842116100645780636a627842146101565780636c0360eb1461018657806370a08231146101a457806395d89b41146101d4578063c87b56dd146101f25761009c565b806306fdde03146100a0578063081812fc146100be578063095ea7b3146100ee57806323b872dd1461010a5780636352211e14610126575b5f80fd5b6100a8610222565b6040516100b59190610e1f565b60405180910390f35b6100d860048036038101906100d39190610e76565b6102ad565b6040516100e59190610ee0565b60405180910390f35b61010860048036038101906101039190610f23565b610384565b005b610124600480360381019061011f9190610f61565b61051b565b005b610140600480360381019061013b9190610e76565b610752565b60405161014d9190610ee0565b60405180910390f35b610170600480360381019061016b9190610fb1565b6107fe565b60405161017d9190610feb565b60405180910390f35b61018e61098f565b60405161019b9190610e1f565b60405180910390f35b6101be60048036038101906101b99190610fb1565b610a1f565b6040516101cb9190610feb565b60405180910390f35b6101dc610ad3565b6040516101e99190610e1f565b60405180910390f35b61020c60048036038101906102079190610e76565b610b5f565b6040516102199190610e1f565b60405180910390f35b5f805461022e90611031565b80601f016020809104026020016040519081016040528092919081815260200182805461025a90611031565b80156102a55780601f1061027c576101008083540402835291602001916102a5565b820191905f5260205f20905b81548152906001019060200180831161028857829003601f168201915b505050505081565b5f8073ffffffffffffffffffffffffffffffffffffffff1660025f8481526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361034c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610343906110ab565b60405180910390fd5b60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f61038e82610752565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036103fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f590611113565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461046c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104639061117b565b60405180910390fd5b8260045f8481526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550818373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a4505050565b61052481610752565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610591576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610588906111e3565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036105ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f69061124b565b60405180910390fd5b60035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f81548092919061064c90611296565b919050555060035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f81548092919061069e906112bd565b91905055508160025f8381526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b5f8060025f8481526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036107f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ec906110ab565b60405180910390fd5b80915050919050565b5f8073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361086d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108649061134e565b60405180910390fd5b60055f81548092919061087f906112bd565b91905055505f600554905060035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8154809291906108d7906112bd565b91905055508260025f8381526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808373ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a480915050919050565b60606006805461099e90611031565b80601f01602080910402602001604051908101604052809291908181526020018280546109ca90611031565b8015610a155780601f106109ec57610100808354040283529160200191610a15565b820191905f5260205f20905b8154815290600101906020018083116109f857829003601f168201915b5050505050905090565b5f8073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a8e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a85906113dc565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60018054610ae090611031565b80601f0160208091040260200160405190810160405280929190818152602001828054610b0c90611031565b8015610b575780601f10610b2e57610100808354040283529160200191610b57565b820191905f5260205f20905b815481529060010190602001808311610b3a57829003601f168201915b505050505081565b60605f73ffffffffffffffffffffffffffffffffffffffff1660025f8481526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610bff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bf6906110ab565b60405180910390fd5b6006610c0a83610c31565b604051602001610c1b9291906114c6565b6040516020818303038152906040529050919050565b60605f8203610c77576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610daa565b5f8290505f5b5f8214610ca6578080610c8f906112bd565b915050600a82610c9f9190611516565b9150610c7d565b5f8167ffffffffffffffff811115610cc157610cc0611546565b5b6040519080825280601f01601f191660200182016040528015610cf35781602001600182028036833780820191505090505b5090505f8290505b5f8614610da257600181610d0f9190611573565b90505f600a8088610d209190611516565b610d2a91906115a6565b87610d359190611573565b6030610d4191906115f3565b90505f8160f81b905080848481518110610d5e57610d5d611627565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a88610d999190611516565b97505050610cfb565b819450505050505b919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610df182610daf565b610dfb8185610db9565b9350610e0b818560208601610dc9565b610e1481610dd7565b840191505092915050565b5f6020820190508181035f830152610e378184610de7565b905092915050565b5f80fd5b5f819050919050565b610e5581610e43565b8114610e5f575f80fd5b50565b5f81359050610e7081610e4c565b92915050565b5f60208284031215610e8b57610e8a610e3f565b5b5f610e9884828501610e62565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610eca82610ea1565b9050919050565b610eda81610ec0565b82525050565b5f602082019050610ef35f830184610ed1565b92915050565b610f0281610ec0565b8114610f0c575f80fd5b50565b5f81359050610f1d81610ef9565b92915050565b5f8060408385031215610f3957610f38610e3f565b5b5f610f4685828601610f0f565b9250506020610f5785828601610e62565b9150509250929050565b5f805f60608486031215610f7857610f77610e3f565b5b5f610f8586828701610f0f565b9350506020610f9686828701610f0f565b9250506040610fa786828701610e62565b9150509250925092565b5f60208284031215610fc657610fc5610e3f565b5b5f610fd384828501610f0f565b91505092915050565b610fe581610e43565b82525050565b5f602082019050610ffe5f830184610fdc565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061104857607f821691505b60208210810361105b5761105a611004565b5b50919050565b7f546f6b656e20494420646f6573206e6f742065786973740000000000000000005f82015250565b5f611095601783610db9565b91506110a082611061565b602082019050919050565b5f6020820190508181035f8301526110c281611089565b9050919050565b7f417070726f76616c20746f2063757272656e74206f776e6572000000000000005f82015250565b5f6110fd601983610db9565b9150611108826110c9565b602082019050919050565b5f6020820190508181035f83015261112a816110f1565b9050919050565b7f417070726f76652063616c6c6572206973206e6f74206f776e657200000000005f82015250565b5f611165601b83610db9565b915061117082611131565b602082019050919050565b5f6020820190508181035f83015261119281611159565b9050919050565b7f5472616e736665722066726f6d20696e636f7272656374206f776e65720000005f82015250565b5f6111cd601d83610db9565b91506111d882611199565b602082019050919050565b5f6020820190508181035f8301526111fa816111c1565b9050919050565b7f5472616e7366657220746f20746865207a65726f2061646472657373000000005f82015250565b5f611235601c83610db9565b915061124082611201565b602082019050919050565b5f6020820190508181035f83015261126281611229565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6112a082610e43565b91505f82036112b2576112b1611269565b5b600182039050919050565b5f6112c782610e43565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036112f9576112f8611269565b5b600182019050919050565b7f4d696e7420746f20746865207a65726f206164647265737300000000000000005f82015250565b5f611338601883610db9565b915061134382611304565b602082019050919050565b5f6020820190508181035f8301526113658161132c565b9050919050565b7f41646472657373207a65726f206973206e6f7420612076616c6964206f776e655f8201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b5f6113c6602183610db9565b91506113d18261136c565b604082019050919050565b5f6020820190508181035f8301526113f3816113ba565b9050919050565b5f81905092915050565b5f819050815f5260205f209050919050565b5f815461142281611031565b61142c81866113fa565b9450600182165f8114611446576001811461145b5761148d565b60ff198316865281151582028601935061148d565b61146485611404565b5f5b8381101561148557815481890152600182019150602081019050611466565b838801955050505b50505092915050565b5f6114a082610daf565b6114aa81856113fa565b93506114ba818560208601610dc9565b80840191505092915050565b5f6114d18285611416565b91506114dd8284611496565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61152082610e43565b915061152b83610e43565b92508261153b5761153a6114e9565b5b828204905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f61157d82610e43565b915061158883610e43565b92508282039050818111156115a05761159f611269565b5b92915050565b5f6115b082610e43565b91506115bb83610e43565b92508282026115c981610e43565b915082820484148315176115e0576115df611269565b5b5092915050565b5f60ff82169050919050565b5f6115fd826115e7565b9150611608836115e7565b9250828201905060ff81111561162157611620611269565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea2646970667358221220ba2d2a1592b17df405d5a3a215a85771564c0d5b16fb403d26d11dd7f0056f8364736f6c634300081a0033",
}

// MyNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use MyNFTMetaData.ABI instead.
var MyNFTABI = MyNFTMetaData.ABI

// MyNFTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyNFTMetaData.Bin instead.
var MyNFTBin = MyNFTMetaData.Bin

// DeployMyNFT deploys a new Ethereum contract, binding an instance of MyNFT to it.
func DeployMyNFT(auth *bind.TransactOpts, backend bind.ContractBackend, baseURI_ string) (common.Address, *types.Transaction, *MyNFT, error) {
	parsed, err := MyNFTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyNFTBin), backend, baseURI_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyNFT{MyNFTCaller: MyNFTCaller{contract: contract}, MyNFTTransactor: MyNFTTransactor{contract: contract}, MyNFTFilterer: MyNFTFilterer{contract: contract}}, nil
}

// MyNFT is an auto generated Go binding around an Ethereum contract.
type MyNFT struct {
	MyNFTCaller     // Read-only binding to the contract
	MyNFTTransactor // Write-only binding to the contract
	MyNFTFilterer   // Log filterer for contract events
}

// MyNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyNFTSession struct {
	Contract     *MyNFT            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyNFTCallerSession struct {
	Contract *MyNFTCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MyNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyNFTTransactorSession struct {
	Contract     *MyNFTTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyNFTRaw struct {
	Contract *MyNFT // Generic contract binding to access the raw methods on
}

// MyNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyNFTCallerRaw struct {
	Contract *MyNFTCaller // Generic read-only contract binding to access the raw methods on
}

// MyNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyNFTTransactorRaw struct {
	Contract *MyNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyNFT creates a new instance of MyNFT, bound to a specific deployed contract.
func NewMyNFT(address common.Address, backend bind.ContractBackend) (*MyNFT, error) {
	contract, err := bindMyNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyNFT{MyNFTCaller: MyNFTCaller{contract: contract}, MyNFTTransactor: MyNFTTransactor{contract: contract}, MyNFTFilterer: MyNFTFilterer{contract: contract}}, nil
}

// NewMyNFTCaller creates a new read-only instance of MyNFT, bound to a specific deployed contract.
func NewMyNFTCaller(address common.Address, caller bind.ContractCaller) (*MyNFTCaller, error) {
	contract, err := bindMyNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyNFTCaller{contract: contract}, nil
}

// NewMyNFTTransactor creates a new write-only instance of MyNFT, bound to a specific deployed contract.
func NewMyNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*MyNFTTransactor, error) {
	contract, err := bindMyNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyNFTTransactor{contract: contract}, nil
}

// NewMyNFTFilterer creates a new log filterer instance of MyNFT, bound to a specific deployed contract.
func NewMyNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*MyNFTFilterer, error) {
	contract, err := bindMyNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyNFTFilterer{contract: contract}, nil
}

// bindMyNFT binds a generic wrapper to an already deployed contract.
func bindMyNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MyNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyNFT *MyNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyNFT.Contract.MyNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyNFT *MyNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyNFT.Contract.MyNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyNFT *MyNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyNFT.Contract.MyNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyNFT *MyNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyNFT *MyNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyNFT *MyNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MyNFT *MyNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MyNFT *MyNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyNFT.Contract.BalanceOf(&_MyNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MyNFT *MyNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyNFT.Contract.BalanceOf(&_MyNFT.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_MyNFT *MyNFTCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_MyNFT *MyNFTSession) BaseURI() (string, error) {
	return _MyNFT.Contract.BaseURI(&_MyNFT.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_MyNFT *MyNFTCallerSession) BaseURI() (string, error) {
	return _MyNFT.Contract.BaseURI(&_MyNFT.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MyNFT *MyNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MyNFT *MyNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MyNFT.Contract.GetApproved(&_MyNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MyNFT *MyNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MyNFT.Contract.GetApproved(&_MyNFT.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyNFT *MyNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyNFT *MyNFTSession) Name() (string, error) {
	return _MyNFT.Contract.Name(&_MyNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyNFT *MyNFTCallerSession) Name() (string, error) {
	return _MyNFT.Contract.Name(&_MyNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MyNFT *MyNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MyNFT *MyNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MyNFT.Contract.OwnerOf(&_MyNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MyNFT *MyNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MyNFT.Contract.OwnerOf(&_MyNFT.CallOpts, tokenId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyNFT *MyNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyNFT *MyNFTSession) Symbol() (string, error) {
	return _MyNFT.Contract.Symbol(&_MyNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyNFT *MyNFTCallerSession) Symbol() (string, error) {
	return _MyNFT.Contract.Symbol(&_MyNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MyNFT *MyNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MyNFT *MyNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MyNFT.Contract.TokenURI(&_MyNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MyNFT *MyNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MyNFT.Contract.TokenURI(&_MyNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MyNFT *MyNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MyNFT *MyNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT.Contract.Approve(&_MyNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MyNFT *MyNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT.Contract.Approve(&_MyNFT.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_MyNFT *MyNFTTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MyNFT.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_MyNFT *MyNFTSession) Mint(to common.Address) (*types.Transaction, error) {
	return _MyNFT.Contract.Mint(&_MyNFT.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_MyNFT *MyNFTTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _MyNFT.Contract.Mint(&_MyNFT.TransactOpts, to)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MyNFT *MyNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MyNFT *MyNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT.Contract.TransferFrom(&_MyNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MyNFT *MyNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT.Contract.TransferFrom(&_MyNFT.TransactOpts, from, to, tokenId)
}

// MyNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MyNFT contract.
type MyNFTApprovalIterator struct {
	Event *MyNFTApproval // Event containing the contract specifics and raw log

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
func (it *MyNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyNFTApproval)
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
		it.Event = new(MyNFTApproval)
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
func (it *MyNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyNFTApproval represents a Approval event raised by the MyNFT contract.
type MyNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MyNFT *MyNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MyNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MyNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MyNFTApprovalIterator{contract: _MyNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MyNFT *MyNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MyNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MyNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyNFTApproval)
				if err := _MyNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MyNFT *MyNFTFilterer) ParseApproval(log types.Log) (*MyNFTApproval, error) {
	event := new(MyNFTApproval)
	if err := _MyNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyNFT contract.
type MyNFTTransferIterator struct {
	Event *MyNFTTransfer // Event containing the contract specifics and raw log

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
func (it *MyNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyNFTTransfer)
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
		it.Event = new(MyNFTTransfer)
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
func (it *MyNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyNFTTransfer represents a Transfer event raised by the MyNFT contract.
type MyNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MyNFT *MyNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MyNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MyNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MyNFTTransferIterator{contract: _MyNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MyNFT *MyNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MyNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyNFTTransfer)
				if err := _MyNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MyNFT *MyNFTFilterer) ParseTransfer(log types.Log) (*MyNFTTransfer, error) {
	event := new(MyNFTTransfer)
	if err := _MyNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
