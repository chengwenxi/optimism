// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
)

// L2ToL1MessagePasserMetaData contains all meta data concerning the L2ToL1MessagePasser contract.
var L2ToL1MessagePasserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"MessagePassed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawerBalanceBurnt\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSAGE_VERSION\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"branches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initiateWithdrawal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sentMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e060405234801561001057600080fd5b5060016080819052600060a081905260c081905280610bdc61004a823960006103d8015260006103af015260006103860152610bdc6000f3fe60806040526004361061007f5760003560e01c806383cc76601161004e57806383cc76601461014c578063c2b3e5ac1461017a578063d4b9f4fa1461018d578063ecc70428146101a357600080fd5b80633f827a5a146100a857806344df8e70146100d557806354fd4d50146100ea57806382e3702d1461010c57600080fd5b366100a3576100a133620186a0604051806020016040528060008152506101fa565b005b600080fd5b3480156100b457600080fd5b506100bd600181565b60405161ffff90911681526020015b60405180910390f35b3480156100e157600080fd5b506100a1610347565b3480156100f657600080fd5b506100ff61037f565b6040516100cc9190610808565b34801561011857600080fd5b5061013c610127366004610822565b60006020819052908152604090205460ff1681565b60405190151581526020016100cc565b34801561015857600080fd5b5061016c610167366004610822565b610422565b6040519081526020016100cc565b6100a161018836600461086a565b6101fa565b34801561019957600080fd5b5061016c60025481565b3480156101af57600080fd5b5061016c6001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b60006102906040518060c001604052806102546001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b815233602082015273ffffffffffffffffffffffffffffffffffffffff871660408201523460608201526080810186905260a001849052610439565b9050600061029d82610486565b905073ffffffffffffffffffffffffffffffffffffffff8516336103016001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b7f3034a173cd04c8b47937b2f95d00794df516ce57df60210b71de7ecf6d2fe4f3348888888860405161033895949392919061096e565b60405180910390a45050505050565b4761035181610616565b60405181907f7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f90600090a250565b60606103aa7f0000000000000000000000000000000000000000000000000000000000000000610645565b6103d37f0000000000000000000000000000000000000000000000000000000000000000610645565b6103fc7f0000000000000000000000000000000000000000000000000000000000000000610645565b60405160200161040e939291906109a1565b604051602081830303815290604052905090565b602b816028811061043257600080fd5b0154905081565b80516020808301516040808501516060860151608087015160a08801519351600097610469979096959101610a17565b604051602081830303815290604052805190602001209050919050565b6001546000907dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1682825b7dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8316156105a3576104e0600284610a9d565b7dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660000361054c5781602b826028811061051b5761051b610ada565b0155610545826003836028811061053457610534610ada565b015460009182526020526040902090565b9150610577565b610574602b826028811061056257610562610ada565b01548360009182526020526040902090565b91505b7d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600193841c1692016104af565b81602b82602881106105b7576105b7610ada565b0155506002819055600180547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8082168301167fffff0000000000000000000000000000000000000000000000000000000000009091161790559392505050565b8060405161062390610782565b6040518091039082f0905080158015610640573d6000803e3d6000fd5b505050565b60608160000361068857505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156106b2578061069c81610b38565b91506106ab9050600a83610b70565b915061068c565b60008167ffffffffffffffff8111156106cd576106cd61083b565b6040519080825280601f01601f1916602001820160405280156106f7576020820181803683370190505b5090505b841561077a5761070c600183610b84565b9150610719600a86610b9b565b610724906030610baf565b60f81b81838151811061073957610739610ada565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610773600a86610b70565b94506106fb565b949350505050565b600880610bc883390190565b60005b838110156107a9578181015183820152602001610791565b838111156107b8576000848401525b50505050565b600081518084526107d681602086016020860161078e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061081b60208301846107be565b9392505050565b60006020828403121561083457600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060006060848603121561087f57600080fd5b833573ffffffffffffffffffffffffffffffffffffffff811681146108a357600080fd5b925060208401359150604084013567ffffffffffffffff808211156108c757600080fd5b818601915086601f8301126108db57600080fd5b8135818111156108ed576108ed61083b565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156109335761093361083b565b8160405282815289602084870101111561094c57600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b85815284602082015260a06040820152600061098d60a08301866107be565b606083019490945250608001529392505050565b600084516109b381846020890161078e565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516109ef816001850160208a0161078e565b60019201918201528351610a0a81600284016020880161078e565b0160020195945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a0830152610a6260c08301846107be565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60007dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80841680610ace57610ace610a6e565b92169190910692915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610b6957610b69610b09565b5060010190565b600082610b7f57610b7f610a6e565b500490565b600082821015610b9657610b96610b09565b500390565b600082610baa57610baa610a6e565b500690565b60008219821115610bc257610bc2610b09565b50019056fe608060405230fffea164736f6c634300080f000a",
}

// L2ToL1MessagePasserABI is the input ABI used to generate the binding from.
// Deprecated: Use L2ToL1MessagePasserMetaData.ABI instead.
var L2ToL1MessagePasserABI = L2ToL1MessagePasserMetaData.ABI

// L2ToL1MessagePasserBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2ToL1MessagePasserMetaData.Bin instead.
var L2ToL1MessagePasserBin = L2ToL1MessagePasserMetaData.Bin

// DeployL2ToL1MessagePasser deploys a new Ethereum contract, binding an instance of L2ToL1MessagePasser to it.
func DeployL2ToL1MessagePasser(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2ToL1MessagePasser, error) {
	parsed, err := L2ToL1MessagePasserMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2ToL1MessagePasserBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2ToL1MessagePasser{L2ToL1MessagePasserCaller: L2ToL1MessagePasserCaller{contract: contract}, L2ToL1MessagePasserTransactor: L2ToL1MessagePasserTransactor{contract: contract}, L2ToL1MessagePasserFilterer: L2ToL1MessagePasserFilterer{contract: contract}}, nil
}

// L2ToL1MessagePasser is an auto generated Go binding around an Ethereum contract.
type L2ToL1MessagePasser struct {
	L2ToL1MessagePasserCaller     // Read-only binding to the contract
	L2ToL1MessagePasserTransactor // Write-only binding to the contract
	L2ToL1MessagePasserFilterer   // Log filterer for contract events
}

// L2ToL1MessagePasserCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ToL1MessagePasserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ToL1MessagePasserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ToL1MessagePasserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ToL1MessagePasserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ToL1MessagePasserSession struct {
	Contract     *L2ToL1MessagePasser // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// L2ToL1MessagePasserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ToL1MessagePasserCallerSession struct {
	Contract *L2ToL1MessagePasserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// L2ToL1MessagePasserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ToL1MessagePasserTransactorSession struct {
	Contract     *L2ToL1MessagePasserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// L2ToL1MessagePasserRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ToL1MessagePasserRaw struct {
	Contract *L2ToL1MessagePasser // Generic contract binding to access the raw methods on
}

// L2ToL1MessagePasserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserCallerRaw struct {
	Contract *L2ToL1MessagePasserCaller // Generic read-only contract binding to access the raw methods on
}

// L2ToL1MessagePasserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserTransactorRaw struct {
	Contract *L2ToL1MessagePasserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ToL1MessagePasser creates a new instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasser(address common.Address, backend bind.ContractBackend) (*L2ToL1MessagePasser, error) {
	contract, err := bindL2ToL1MessagePasser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasser{L2ToL1MessagePasserCaller: L2ToL1MessagePasserCaller{contract: contract}, L2ToL1MessagePasserTransactor: L2ToL1MessagePasserTransactor{contract: contract}, L2ToL1MessagePasserFilterer: L2ToL1MessagePasserFilterer{contract: contract}}, nil
}

// NewL2ToL1MessagePasserCaller creates a new read-only instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasserCaller(address common.Address, caller bind.ContractCaller) (*L2ToL1MessagePasserCaller, error) {
	contract, err := bindL2ToL1MessagePasser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserCaller{contract: contract}, nil
}

// NewL2ToL1MessagePasserTransactor creates a new write-only instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasserTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ToL1MessagePasserTransactor, error) {
	contract, err := bindL2ToL1MessagePasser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserTransactor{contract: contract}, nil
}

// NewL2ToL1MessagePasserFilterer creates a new log filterer instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasserFilterer(address common.Address, filterer bind.ContractFilterer) (*L2ToL1MessagePasserFilterer, error) {
	contract, err := bindL2ToL1MessagePasser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserFilterer{contract: contract}, nil
}

// bindL2ToL1MessagePasser binds a generic wrapper to an already deployed contract.
func bindL2ToL1MessagePasser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ToL1MessagePasserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ToL1MessagePasser.Contract.L2ToL1MessagePasserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.L2ToL1MessagePasserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.L2ToL1MessagePasserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ToL1MessagePasser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.contract.Transact(opts, method, params...)
}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) MESSAGEVERSION(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "MESSAGE_VERSION")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) MESSAGEVERSION() (uint16, error) {
	return _L2ToL1MessagePasser.Contract.MESSAGEVERSION(&_L2ToL1MessagePasser.CallOpts)
}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) MESSAGEVERSION() (uint16, error) {
	return _L2ToL1MessagePasser.Contract.MESSAGEVERSION(&_L2ToL1MessagePasser.CallOpts)
}

// Branches is a free data retrieval call binding the contract method 0x83cc7660.
//
// Solidity: function branches(uint256 ) view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) Branches(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "branches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Branches is a free data retrieval call binding the contract method 0x83cc7660.
//
// Solidity: function branches(uint256 ) view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) Branches(arg0 *big.Int) ([32]byte, error) {
	return _L2ToL1MessagePasser.Contract.Branches(&_L2ToL1MessagePasser.CallOpts, arg0)
}

// Branches is a free data retrieval call binding the contract method 0x83cc7660.
//
// Solidity: function branches(uint256 ) view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) Branches(arg0 *big.Int) ([32]byte, error) {
	return _L2ToL1MessagePasser.Contract.Branches(&_L2ToL1MessagePasser.CallOpts, arg0)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) MessageNonce() (*big.Int, error) {
	return _L2ToL1MessagePasser.Contract.MessageNonce(&_L2ToL1MessagePasser.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) MessageNonce() (*big.Int, error) {
	return _L2ToL1MessagePasser.Contract.MessageNonce(&_L2ToL1MessagePasser.CallOpts)
}

// MessageRoot is a free data retrieval call binding the contract method 0xd4b9f4fa.
//
// Solidity: function messageRoot() view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) MessageRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "messageRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageRoot is a free data retrieval call binding the contract method 0xd4b9f4fa.
//
// Solidity: function messageRoot() view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) MessageRoot() ([32]byte, error) {
	return _L2ToL1MessagePasser.Contract.MessageRoot(&_L2ToL1MessagePasser.CallOpts)
}

// MessageRoot is a free data retrieval call binding the contract method 0xd4b9f4fa.
//
// Solidity: function messageRoot() view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) MessageRoot() ([32]byte, error) {
	return _L2ToL1MessagePasser.Contract.MessageRoot(&_L2ToL1MessagePasser.CallOpts)
}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) SentMessages(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "sentMessages", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) SentMessages(arg0 [32]byte) (bool, error) {
	return _L2ToL1MessagePasser.Contract.SentMessages(&_L2ToL1MessagePasser.CallOpts, arg0)
}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) SentMessages(arg0 [32]byte) (bool, error) {
	return _L2ToL1MessagePasser.Contract.SentMessages(&_L2ToL1MessagePasser.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) Version() (string, error) {
	return _L2ToL1MessagePasser.Contract.Version(&_L2ToL1MessagePasser.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) Version() (string, error) {
	return _L2ToL1MessagePasser.Contract.Version(&_L2ToL1MessagePasser.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) Burn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.Transact(opts, "burn")
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) Burn() (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Burn(&_L2ToL1MessagePasser.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) Burn() (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Burn(&_L2ToL1MessagePasser.TransactOpts)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xc2b3e5ac.
//
// Solidity: function initiateWithdrawal(address _target, uint256 _gasLimit, bytes _data) payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) InitiateWithdrawal(opts *bind.TransactOpts, _target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.Transact(opts, "initiateWithdrawal", _target, _gasLimit, _data)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xc2b3e5ac.
//
// Solidity: function initiateWithdrawal(address _target, uint256 _gasLimit, bytes _data) payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) InitiateWithdrawal(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.InitiateWithdrawal(&_L2ToL1MessagePasser.TransactOpts, _target, _gasLimit, _data)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xc2b3e5ac.
//
// Solidity: function initiateWithdrawal(address _target, uint256 _gasLimit, bytes _data) payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) InitiateWithdrawal(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.InitiateWithdrawal(&_L2ToL1MessagePasser.TransactOpts, _target, _gasLimit, _data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) Receive() (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Receive(&_L2ToL1MessagePasser.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) Receive() (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Receive(&_L2ToL1MessagePasser.TransactOpts)
}

// L2ToL1MessagePasserMessagePassedIterator is returned from FilterMessagePassed and is used to iterate over the raw logs and unpacked data for MessagePassed events raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserMessagePassedIterator struct {
	Event *L2ToL1MessagePasserMessagePassed // Event containing the contract specifics and raw log

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
func (it *L2ToL1MessagePasserMessagePassedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ToL1MessagePasserMessagePassed)
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
		it.Event = new(L2ToL1MessagePasserMessagePassed)
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
func (it *L2ToL1MessagePasserMessagePassedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ToL1MessagePasserMessagePassedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ToL1MessagePasserMessagePassed represents a MessagePassed event raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserMessagePassed struct {
	Nonce          *big.Int
	Sender         common.Address
	Target         common.Address
	Value          *big.Int
	GasLimit       *big.Int
	Data           []byte
	WithdrawalHash [32]byte
	RootHash       [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMessagePassed is a free log retrieval operation binding the contract event 0x3034a173cd04c8b47937b2f95d00794df516ce57df60210b71de7ecf6d2fe4f3.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash, bytes32 rootHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) FilterMessagePassed(opts *bind.FilterOpts, nonce []*big.Int, sender []common.Address, target []common.Address) (*L2ToL1MessagePasserMessagePassedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.FilterLogs(opts, "MessagePassed", nonceRule, senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserMessagePassedIterator{contract: _L2ToL1MessagePasser.contract, event: "MessagePassed", logs: logs, sub: sub}, nil
}

// WatchMessagePassed is a free log subscription operation binding the contract event 0x3034a173cd04c8b47937b2f95d00794df516ce57df60210b71de7ecf6d2fe4f3.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash, bytes32 rootHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) WatchMessagePassed(opts *bind.WatchOpts, sink chan<- *L2ToL1MessagePasserMessagePassed, nonce []*big.Int, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.WatchLogs(opts, "MessagePassed", nonceRule, senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ToL1MessagePasserMessagePassed)
				if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "MessagePassed", log); err != nil {
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

// ParseMessagePassed is a log parse operation binding the contract event 0x3034a173cd04c8b47937b2f95d00794df516ce57df60210b71de7ecf6d2fe4f3.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash, bytes32 rootHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) ParseMessagePassed(log types.Log) (*L2ToL1MessagePasserMessagePassed, error) {
	event := new(L2ToL1MessagePasserMessagePassed)
	if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "MessagePassed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ToL1MessagePasserWithdrawerBalanceBurntIterator is returned from FilterWithdrawerBalanceBurnt and is used to iterate over the raw logs and unpacked data for WithdrawerBalanceBurnt events raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserWithdrawerBalanceBurntIterator struct {
	Event *L2ToL1MessagePasserWithdrawerBalanceBurnt // Event containing the contract specifics and raw log

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
func (it *L2ToL1MessagePasserWithdrawerBalanceBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
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
		it.Event = new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
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
func (it *L2ToL1MessagePasserWithdrawerBalanceBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ToL1MessagePasserWithdrawerBalanceBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ToL1MessagePasserWithdrawerBalanceBurnt represents a WithdrawerBalanceBurnt event raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserWithdrawerBalanceBurnt struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawerBalanceBurnt is a free log retrieval operation binding the contract event 0x7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f.
//
// Solidity: event WithdrawerBalanceBurnt(uint256 indexed amount)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) FilterWithdrawerBalanceBurnt(opts *bind.FilterOpts, amount []*big.Int) (*L2ToL1MessagePasserWithdrawerBalanceBurntIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.FilterLogs(opts, "WithdrawerBalanceBurnt", amountRule)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserWithdrawerBalanceBurntIterator{contract: _L2ToL1MessagePasser.contract, event: "WithdrawerBalanceBurnt", logs: logs, sub: sub}, nil
}

// WatchWithdrawerBalanceBurnt is a free log subscription operation binding the contract event 0x7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f.
//
// Solidity: event WithdrawerBalanceBurnt(uint256 indexed amount)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) WatchWithdrawerBalanceBurnt(opts *bind.WatchOpts, sink chan<- *L2ToL1MessagePasserWithdrawerBalanceBurnt, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.WatchLogs(opts, "WithdrawerBalanceBurnt", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
				if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "WithdrawerBalanceBurnt", log); err != nil {
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

// ParseWithdrawerBalanceBurnt is a log parse operation binding the contract event 0x7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f.
//
// Solidity: event WithdrawerBalanceBurnt(uint256 indexed amount)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) ParseWithdrawerBalanceBurnt(log types.Log) (*L2ToL1MessagePasserWithdrawerBalanceBurnt, error) {
	event := new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
	if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "WithdrawerBalanceBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
