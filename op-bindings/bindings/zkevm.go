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

// ZKEVMBatchData is an auto generated low-level Go binding around an user-defined struct.
type ZKEVMBatchData struct {
	BlockWitness   []byte
	StateRoot      [32]byte
	Timestamp      *big.Int
	Transactions   []byte
	GlobalExitRoot [32]byte
}

// ZKEVMMetaData contains all meta data concerning the ZKEVM contract.
var ZKEVMMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SubmitBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_DEPOSIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PORTAL\",\"outputs\":[{\"internalType\":\"contractOptimismPortal\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_WINDOW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batch\",\"type\":\"uint64\"}],\"name\":\"challengeState\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"challenges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptimismPortal\",\"name\":\"_portal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_submitter\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"originTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batch\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"stateRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blockWitness\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structZKEVM.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"submitBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506114e5806100206000396000f3fe6080604052600436106101445760003560e01c8063715018a6116100c0578063956d719111610074578063e1e158a511610059578063e1e158a5146103d3578063f2fde38b146103ef578063fc7e286d1461040f57600080fd5b8063956d719114610379578063c7ab2039146103a657600080fd5b80638da5cb5b116100a55780638da5cb5b1461030c5780638dc45d9a14610337578063913b21501461036457600080fd5b8063715018a6146102e45780638d644bb7146102f957600080fd5b80633a4b66f111610117578063485cc955116100fc578063485cc95514610284578063534db0e21461029757806359ef1120146102c457600080fd5b80633a4b66f11461022a578063423fa8561461023257600080fd5b806301117700146101495780630ff754ea1461016b5780631b4082c9146101c25780632b259441146101fd575b600080fd5b34801561015557600080fd5b506101696101643660046111aa565b61043c565b005b34801561017757600080fd5b506065546101989073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101ce57600080fd5b506101ef6101dd36600461123c565b606b6020526000908152604090205481565b6040519081526020016101b9565b34801561020957600080fd5b506101ef61021836600461123c565b606a6020526000908152604090205481565b610169610781565b34801561023e57600080fd5b5060675461026b9074010000000000000000000000000000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016101b9565b610169610292366004611280565b61097e565b3480156102a357600080fd5b506067546101989073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102d057600080fd5b506101696102df3660046112b9565b610bb0565b3480156102f057600080fd5b50610169610c44565b61016961030736600461123c565b610c58565b34801561031857600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff16610198565b34801561034357600080fd5b506066546101989073ffffffffffffffffffffffffffffffffffffffff1681565b34801561037057600080fd5b506101ef606481565b34801561038557600080fd5b506101ef61039436600461123c565b60696020526000908152604090205481565b3480156103b257600080fd5b506101ef6103c136600461123c565b606c6020526000908152604090205481565b3480156103df57600080fd5b506101ef670de0b6b3a764000081565b3480156103fb57600080fd5b5061016961040a36600461133c565b610e29565b34801561041b57600080fd5b506101ef61042a36600461133c565b60686020526000908152604090205481565b60665473ffffffffffffffffffffffffffffffffffffffff1633146104c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f43616c6c6572206e6f74207375626d697474657200000000000000000000000060448201526064015b60405180910390fd5b60665473ffffffffffffffffffffffffffffffffffffffff16600090815260686020526040902054670de0b6b3a7640000111561055b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f446f206e6f74206861766520656e6f75676874206465706f736974000000000060448201526064016104b9565b60665473ffffffffffffffffffffffffffffffffffffffff16600090815260686020526040902054670de0b6b3a76400001061059657600080fd5b606754819074010000000000000000000000000000000000000000900467ffffffffffffffff1660005b828110156106fb57606361f6186000806105d983610ee0565b91509150600061061e8383878d8d8b8181106105f7576105f7611359565b90506020028101906106099190611388565b61061390806113c6565b506060949350505050565b80516020908102908201209091508761063681611461565b67ffffffffffffffff8116600090815260696020526040902083905598508b90508a8881811061066857610668611359565b905060200281019061067a9190611388565b67ffffffffffffffff89166000908152606a6020908152604090912091013590558a8a888181106106ad576106ad611359565b90506020028101906106bf9190611388565b67ffffffffffffffff89166000908152606b6020526040908190209101359055508594506106f39350849250611488915050565b9150506105c0565b50606780547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000067ffffffffffffffff84811682029290921792839055604051920416907f85d226f7b6c307388b0488cb75509c840f0154cbed044f4b3c4e52af7056911990600090a250505050565b60665473ffffffffffffffffffffffffffffffffffffffff163314610802576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f43616c6c6572206e6f74207375626d697474657200000000000000000000000060448201526064016104b9565b60665473ffffffffffffffffffffffffffffffffffffffff16600090815260686020526040902054670de0b6b3a7640000111561089b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f446f206e6f74206861766520656e6f75676874206465706f736974000000000060448201526064016104b9565b60665473ffffffffffffffffffffffffffffffffffffffff16600090815260686020526040902054670de0b6b3a7640000906108d89034906114c0565b1015610940576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f446f206e6f74206861766520656e6f756768206465706f73697400000000000060448201526064016104b9565b60665473ffffffffffffffffffffffffffffffffffffffff16600090815260686020526040812080543492906109779084906114c0565b9091555050565b600054610100900460ff161580801561099e5750600054600160ff909116105b806109b85750303b1580156109b8575060005460ff166001145b610a44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104b9565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610aa257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610aaa610f73565b606780547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1690556065805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560668054928516929091168217905560009081526068602052604081208054349290610b439084906114c0565b90915550508015610bab57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b67ffffffffffffffff83166000908152606b60205260408120544290610bd8906064906114c0565b67ffffffffffffffff86166000908152606c602052604081205492909111925003610c0257600080fd5b8015610c265781610c1257600080fd5b67ffffffffffffffff841660005260696020525b50505067ffffffffffffffff166000908152606c6020526040812055565b610c4c611012565b610c566000611093565b565b60675473ffffffffffffffffffffffffffffffffffffffff1615610cd8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f416c726561647920696e206368616c6c656e676500000000000000000000000060448201526064016104b9565b67ffffffffffffffff81166000908152606b60205260408120544290610d00906064906114c0565b11905080610d90576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f43616e6e6f74206368616c6c656e676520626174636865206f7574736964652060448201527f746865206368616c6c656e67652077696e646f7700000000000000000000000060648201526084016104b9565b606780547fffffffffffffffffffffffff00000000000000000000000000000000000000001633179055670de0b6b3a7640000341015610dcf57600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff1660009081526068602052604081208054349290610e069084906114c0565b90915550505067ffffffffffffffff166000908152606c60205260409020349055565b610e31611012565b73ffffffffffffffffffffffffffffffffffffffff8116610ed4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104b9565b610edd81611093565b50565b60008061f6188311610ef9575060039261290492509050565b620493e08311610f115750600e926201107692509050565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f434952435549545f434f4e46494700000000000000000000000000000000000060448201526064016104b9565b600054610100900460ff1661100a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104b9565b610c5661110a565b60335473ffffffffffffffffffffffffffffffffffffffff163314610c56576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104b9565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff166111a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104b9565b610c5633611093565b600080602083850312156111bd57600080fd5b823567ffffffffffffffff808211156111d557600080fd5b818501915085601f8301126111e957600080fd5b8135818111156111f857600080fd5b8660208260051b850101111561120d57600080fd5b60209290920196919550909350505050565b803567ffffffffffffffff8116811461123757600080fd5b919050565b60006020828403121561124e57600080fd5b6112578261121f565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610edd57600080fd5b6000806040838503121561129357600080fd5b823561129e8161125e565b915060208301356112ae8161125e565b809150509250929050565b6000806000604084860312156112ce57600080fd5b6112d78461121f565b9250602084013567ffffffffffffffff808211156112f457600080fd5b818601915086601f83011261130857600080fd5b81358181111561131757600080fd5b87602082850101111561132957600080fd5b6020830194508093505050509250925092565b60006020828403121561134e57600080fd5b81356112578161125e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618336030181126113bc57600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126113fb57600080fd5b83018035915067ffffffffffffffff82111561141657600080fd5b60200191503681900382131561142b57600080fd5b9250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681810361147e5761147e611432565b6001019392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036114b9576114b9611432565b5060010190565b600082198211156114d3576114d3611432565b50019056fea164736f6c634300080f000a",
}

// ZKEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKEVMMetaData.ABI instead.
var ZKEVMABI = ZKEVMMetaData.ABI

// ZKEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZKEVMMetaData.Bin instead.
var ZKEVMBin = ZKEVMMetaData.Bin

// DeployZKEVM deploys a new Ethereum contract, binding an instance of ZKEVM to it.
func DeployZKEVM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZKEVM, error) {
	parsed, err := ZKEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZKEVMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZKEVM{ZKEVMCaller: ZKEVMCaller{contract: contract}, ZKEVMTransactor: ZKEVMTransactor{contract: contract}, ZKEVMFilterer: ZKEVMFilterer{contract: contract}}, nil
}

// ZKEVM is an auto generated Go binding around an Ethereum contract.
type ZKEVM struct {
	ZKEVMCaller     // Read-only binding to the contract
	ZKEVMTransactor // Write-only binding to the contract
	ZKEVMFilterer   // Log filterer for contract events
}

// ZKEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZKEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKEVMSession struct {
	Contract     *ZKEVM            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZKEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKEVMCallerSession struct {
	Contract *ZKEVMCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZKEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKEVMTransactorSession struct {
	Contract     *ZKEVMTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZKEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZKEVMRaw struct {
	Contract *ZKEVM // Generic contract binding to access the raw methods on
}

// ZKEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKEVMCallerRaw struct {
	Contract *ZKEVMCaller // Generic read-only contract binding to access the raw methods on
}

// ZKEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKEVMTransactorRaw struct {
	Contract *ZKEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZKEVM creates a new instance of ZKEVM, bound to a specific deployed contract.
func NewZKEVM(address common.Address, backend bind.ContractBackend) (*ZKEVM, error) {
	contract, err := bindZKEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKEVM{ZKEVMCaller: ZKEVMCaller{contract: contract}, ZKEVMTransactor: ZKEVMTransactor{contract: contract}, ZKEVMFilterer: ZKEVMFilterer{contract: contract}}, nil
}

// NewZKEVMCaller creates a new read-only instance of ZKEVM, bound to a specific deployed contract.
func NewZKEVMCaller(address common.Address, caller bind.ContractCaller) (*ZKEVMCaller, error) {
	contract, err := bindZKEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKEVMCaller{contract: contract}, nil
}

// NewZKEVMTransactor creates a new write-only instance of ZKEVM, bound to a specific deployed contract.
func NewZKEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKEVMTransactor, error) {
	contract, err := bindZKEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKEVMTransactor{contract: contract}, nil
}

// NewZKEVMFilterer creates a new log filterer instance of ZKEVM, bound to a specific deployed contract.
func NewZKEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKEVMFilterer, error) {
	contract, err := bindZKEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKEVMFilterer{contract: contract}, nil
}

// bindZKEVM binds a generic wrapper to an already deployed contract.
func bindZKEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZKEVMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKEVM *ZKEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKEVM.Contract.ZKEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKEVM *ZKEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKEVM.Contract.ZKEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKEVM *ZKEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKEVM.Contract.ZKEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKEVM *ZKEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKEVM *ZKEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKEVM *ZKEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKEVM.Contract.contract.Transact(opts, method, params...)
}

// MINDEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() view returns(uint256)
func (_ZKEVM *ZKEVMCaller) MINDEPOSIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "MIN_DEPOSIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINDEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() view returns(uint256)
func (_ZKEVM *ZKEVMSession) MINDEPOSIT() (*big.Int, error) {
	return _ZKEVM.Contract.MINDEPOSIT(&_ZKEVM.CallOpts)
}

// MINDEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() view returns(uint256)
func (_ZKEVM *ZKEVMCallerSession) MINDEPOSIT() (*big.Int, error) {
	return _ZKEVM.Contract.MINDEPOSIT(&_ZKEVM.CallOpts)
}

// PORTAL is a free data retrieval call binding the contract method 0x0ff754ea.
//
// Solidity: function PORTAL() view returns(address)
func (_ZKEVM *ZKEVMCaller) PORTAL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "PORTAL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PORTAL is a free data retrieval call binding the contract method 0x0ff754ea.
//
// Solidity: function PORTAL() view returns(address)
func (_ZKEVM *ZKEVMSession) PORTAL() (common.Address, error) {
	return _ZKEVM.Contract.PORTAL(&_ZKEVM.CallOpts)
}

// PORTAL is a free data retrieval call binding the contract method 0x0ff754ea.
//
// Solidity: function PORTAL() view returns(address)
func (_ZKEVM *ZKEVMCallerSession) PORTAL() (common.Address, error) {
	return _ZKEVM.Contract.PORTAL(&_ZKEVM.CallOpts)
}

// WINDOW is a free data retrieval call binding the contract method 0x913b2150.
//
// Solidity: function _WINDOW() view returns(uint256)
func (_ZKEVM *ZKEVMCaller) WINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WINDOW is a free data retrieval call binding the contract method 0x913b2150.
//
// Solidity: function _WINDOW() view returns(uint256)
func (_ZKEVM *ZKEVMSession) WINDOW() (*big.Int, error) {
	return _ZKEVM.Contract.WINDOW(&_ZKEVM.CallOpts)
}

// WINDOW is a free data retrieval call binding the contract method 0x913b2150.
//
// Solidity: function _WINDOW() view returns(uint256)
func (_ZKEVM *ZKEVMCallerSession) WINDOW() (*big.Int, error) {
	return _ZKEVM.Contract.WINDOW(&_ZKEVM.CallOpts)
}

// Challenger is a free data retrieval call binding the contract method 0x534db0e2.
//
// Solidity: function challenger() view returns(address)
func (_ZKEVM *ZKEVMCaller) Challenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "challenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Challenger is a free data retrieval call binding the contract method 0x534db0e2.
//
// Solidity: function challenger() view returns(address)
func (_ZKEVM *ZKEVMSession) Challenger() (common.Address, error) {
	return _ZKEVM.Contract.Challenger(&_ZKEVM.CallOpts)
}

// Challenger is a free data retrieval call binding the contract method 0x534db0e2.
//
// Solidity: function challenger() view returns(address)
func (_ZKEVM *ZKEVMCallerSession) Challenger() (common.Address, error) {
	return _ZKEVM.Contract.Challenger(&_ZKEVM.CallOpts)
}

// Challenges is a free data retrieval call binding the contract method 0xc7ab2039.
//
// Solidity: function challenges(uint64 ) view returns(uint256)
func (_ZKEVM *ZKEVMCaller) Challenges(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "challenges", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Challenges is a free data retrieval call binding the contract method 0xc7ab2039.
//
// Solidity: function challenges(uint64 ) view returns(uint256)
func (_ZKEVM *ZKEVMSession) Challenges(arg0 uint64) (*big.Int, error) {
	return _ZKEVM.Contract.Challenges(&_ZKEVM.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0xc7ab2039.
//
// Solidity: function challenges(uint64 ) view returns(uint256)
func (_ZKEVM *ZKEVMCallerSession) Challenges(arg0 uint64) (*big.Int, error) {
	return _ZKEVM.Contract.Challenges(&_ZKEVM.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x956d7191.
//
// Solidity: function commitments(uint64 ) view returns(bytes32)
func (_ZKEVM *ZKEVMCaller) Commitments(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x956d7191.
//
// Solidity: function commitments(uint64 ) view returns(bytes32)
func (_ZKEVM *ZKEVMSession) Commitments(arg0 uint64) ([32]byte, error) {
	return _ZKEVM.Contract.Commitments(&_ZKEVM.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x956d7191.
//
// Solidity: function commitments(uint64 ) view returns(bytes32)
func (_ZKEVM *ZKEVMCallerSession) Commitments(arg0 uint64) ([32]byte, error) {
	return _ZKEVM.Contract.Commitments(&_ZKEVM.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256)
func (_ZKEVM *ZKEVMCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "deposits", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256)
func (_ZKEVM *ZKEVMSession) Deposits(arg0 common.Address) (*big.Int, error) {
	return _ZKEVM.Contract.Deposits(&_ZKEVM.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256)
func (_ZKEVM *ZKEVMCallerSession) Deposits(arg0 common.Address) (*big.Int, error) {
	return _ZKEVM.Contract.Deposits(&_ZKEVM.CallOpts, arg0)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_ZKEVM *ZKEVMCaller) LastBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "lastBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_ZKEVM *ZKEVMSession) LastBatchSequenced() (uint64, error) {
	return _ZKEVM.Contract.LastBatchSequenced(&_ZKEVM.CallOpts)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_ZKEVM *ZKEVMCallerSession) LastBatchSequenced() (uint64, error) {
	return _ZKEVM.Contract.LastBatchSequenced(&_ZKEVM.CallOpts)
}

// OriginTimestamps is a free data retrieval call binding the contract method 0x1b4082c9.
//
// Solidity: function originTimestamps(uint64 ) view returns(uint256)
func (_ZKEVM *ZKEVMCaller) OriginTimestamps(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "originTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OriginTimestamps is a free data retrieval call binding the contract method 0x1b4082c9.
//
// Solidity: function originTimestamps(uint64 ) view returns(uint256)
func (_ZKEVM *ZKEVMSession) OriginTimestamps(arg0 uint64) (*big.Int, error) {
	return _ZKEVM.Contract.OriginTimestamps(&_ZKEVM.CallOpts, arg0)
}

// OriginTimestamps is a free data retrieval call binding the contract method 0x1b4082c9.
//
// Solidity: function originTimestamps(uint64 ) view returns(uint256)
func (_ZKEVM *ZKEVMCallerSession) OriginTimestamps(arg0 uint64) (*big.Int, error) {
	return _ZKEVM.Contract.OriginTimestamps(&_ZKEVM.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZKEVM *ZKEVMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZKEVM *ZKEVMSession) Owner() (common.Address, error) {
	return _ZKEVM.Contract.Owner(&_ZKEVM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZKEVM *ZKEVMCallerSession) Owner() (common.Address, error) {
	return _ZKEVM.Contract.Owner(&_ZKEVM.CallOpts)
}

// StateRoots is a free data retrieval call binding the contract method 0x2b259441.
//
// Solidity: function stateRoots(uint64 ) view returns(bytes32)
func (_ZKEVM *ZKEVMCaller) StateRoots(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "stateRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateRoots is a free data retrieval call binding the contract method 0x2b259441.
//
// Solidity: function stateRoots(uint64 ) view returns(bytes32)
func (_ZKEVM *ZKEVMSession) StateRoots(arg0 uint64) ([32]byte, error) {
	return _ZKEVM.Contract.StateRoots(&_ZKEVM.CallOpts, arg0)
}

// StateRoots is a free data retrieval call binding the contract method 0x2b259441.
//
// Solidity: function stateRoots(uint64 ) view returns(bytes32)
func (_ZKEVM *ZKEVMCallerSession) StateRoots(arg0 uint64) ([32]byte, error) {
	return _ZKEVM.Contract.StateRoots(&_ZKEVM.CallOpts, arg0)
}

// Submitter is a free data retrieval call binding the contract method 0x8dc45d9a.
//
// Solidity: function submitter() view returns(address)
func (_ZKEVM *ZKEVMCaller) Submitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "submitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Submitter is a free data retrieval call binding the contract method 0x8dc45d9a.
//
// Solidity: function submitter() view returns(address)
func (_ZKEVM *ZKEVMSession) Submitter() (common.Address, error) {
	return _ZKEVM.Contract.Submitter(&_ZKEVM.CallOpts)
}

// Submitter is a free data retrieval call binding the contract method 0x8dc45d9a.
//
// Solidity: function submitter() view returns(address)
func (_ZKEVM *ZKEVMCallerSession) Submitter() (common.Address, error) {
	return _ZKEVM.Contract.Submitter(&_ZKEVM.CallOpts)
}

// ChallengeState is a paid mutator transaction binding the contract method 0x8d644bb7.
//
// Solidity: function challengeState(uint64 batch) payable returns()
func (_ZKEVM *ZKEVMTransactor) ChallengeState(opts *bind.TransactOpts, batch uint64) (*types.Transaction, error) {
	return _ZKEVM.contract.Transact(opts, "challengeState", batch)
}

// ChallengeState is a paid mutator transaction binding the contract method 0x8d644bb7.
//
// Solidity: function challengeState(uint64 batch) payable returns()
func (_ZKEVM *ZKEVMSession) ChallengeState(batch uint64) (*types.Transaction, error) {
	return _ZKEVM.Contract.ChallengeState(&_ZKEVM.TransactOpts, batch)
}

// ChallengeState is a paid mutator transaction binding the contract method 0x8d644bb7.
//
// Solidity: function challengeState(uint64 batch) payable returns()
func (_ZKEVM *ZKEVMTransactorSession) ChallengeState(batch uint64) (*types.Transaction, error) {
	return _ZKEVM.Contract.ChallengeState(&_ZKEVM.TransactOpts, batch)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _portal, address _submitter) payable returns()
func (_ZKEVM *ZKEVMTransactor) Initialize(opts *bind.TransactOpts, _portal common.Address, _submitter common.Address) (*types.Transaction, error) {
	return _ZKEVM.contract.Transact(opts, "initialize", _portal, _submitter)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _portal, address _submitter) payable returns()
func (_ZKEVM *ZKEVMSession) Initialize(_portal common.Address, _submitter common.Address) (*types.Transaction, error) {
	return _ZKEVM.Contract.Initialize(&_ZKEVM.TransactOpts, _portal, _submitter)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _portal, address _submitter) payable returns()
func (_ZKEVM *ZKEVMTransactorSession) Initialize(_portal common.Address, _submitter common.Address) (*types.Transaction, error) {
	return _ZKEVM.Contract.Initialize(&_ZKEVM.TransactOpts, _portal, _submitter)
}

// ProveState is a paid mutator transaction binding the contract method 0x59ef1120.
//
// Solidity: function proveState(uint64 batch, bytes proof) returns()
func (_ZKEVM *ZKEVMTransactor) ProveState(opts *bind.TransactOpts, batch uint64, proof []byte) (*types.Transaction, error) {
	return _ZKEVM.contract.Transact(opts, "proveState", batch, proof)
}

// ProveState is a paid mutator transaction binding the contract method 0x59ef1120.
//
// Solidity: function proveState(uint64 batch, bytes proof) returns()
func (_ZKEVM *ZKEVMSession) ProveState(batch uint64, proof []byte) (*types.Transaction, error) {
	return _ZKEVM.Contract.ProveState(&_ZKEVM.TransactOpts, batch, proof)
}

// ProveState is a paid mutator transaction binding the contract method 0x59ef1120.
//
// Solidity: function proveState(uint64 batch, bytes proof) returns()
func (_ZKEVM *ZKEVMTransactorSession) ProveState(batch uint64, proof []byte) (*types.Transaction, error) {
	return _ZKEVM.Contract.ProveState(&_ZKEVM.TransactOpts, batch, proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZKEVM *ZKEVMTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKEVM.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZKEVM *ZKEVMSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZKEVM.Contract.RenounceOwnership(&_ZKEVM.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZKEVM *ZKEVMTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZKEVM.Contract.RenounceOwnership(&_ZKEVM.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_ZKEVM *ZKEVMTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKEVM.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_ZKEVM *ZKEVMSession) Stake() (*types.Transaction, error) {
	return _ZKEVM.Contract.Stake(&_ZKEVM.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_ZKEVM *ZKEVMTransactorSession) Stake() (*types.Transaction, error) {
	return _ZKEVM.Contract.Stake(&_ZKEVM.TransactOpts)
}

// SubmitBatches is a paid mutator transaction binding the contract method 0x01117700.
//
// Solidity: function submitBatches((bytes,bytes32,uint256,bytes,bytes32)[] batches) returns()
func (_ZKEVM *ZKEVMTransactor) SubmitBatches(opts *bind.TransactOpts, batches []ZKEVMBatchData) (*types.Transaction, error) {
	return _ZKEVM.contract.Transact(opts, "submitBatches", batches)
}

// SubmitBatches is a paid mutator transaction binding the contract method 0x01117700.
//
// Solidity: function submitBatches((bytes,bytes32,uint256,bytes,bytes32)[] batches) returns()
func (_ZKEVM *ZKEVMSession) SubmitBatches(batches []ZKEVMBatchData) (*types.Transaction, error) {
	return _ZKEVM.Contract.SubmitBatches(&_ZKEVM.TransactOpts, batches)
}

// SubmitBatches is a paid mutator transaction binding the contract method 0x01117700.
//
// Solidity: function submitBatches((bytes,bytes32,uint256,bytes,bytes32)[] batches) returns()
func (_ZKEVM *ZKEVMTransactorSession) SubmitBatches(batches []ZKEVMBatchData) (*types.Transaction, error) {
	return _ZKEVM.Contract.SubmitBatches(&_ZKEVM.TransactOpts, batches)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZKEVM *ZKEVMTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZKEVM.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZKEVM *ZKEVMSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZKEVM.Contract.TransferOwnership(&_ZKEVM.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZKEVM *ZKEVMTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZKEVM.Contract.TransferOwnership(&_ZKEVM.TransactOpts, newOwner)
}

// ZKEVMInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZKEVM contract.
type ZKEVMInitializedIterator struct {
	Event *ZKEVMInitialized // Event containing the contract specifics and raw log

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
func (it *ZKEVMInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMInitialized)
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
		it.Event = new(ZKEVMInitialized)
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
func (it *ZKEVMInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMInitialized represents a Initialized event raised by the ZKEVM contract.
type ZKEVMInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZKEVM *ZKEVMFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZKEVMInitializedIterator, error) {

	logs, sub, err := _ZKEVM.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZKEVMInitializedIterator{contract: _ZKEVM.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZKEVM *ZKEVMFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZKEVMInitialized) (event.Subscription, error) {

	logs, sub, err := _ZKEVM.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMInitialized)
				if err := _ZKEVM.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZKEVM *ZKEVMFilterer) ParseInitialized(log types.Log) (*ZKEVMInitialized, error) {
	event := new(ZKEVMInitialized)
	if err := _ZKEVM.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZKEVM contract.
type ZKEVMOwnershipTransferredIterator struct {
	Event *ZKEVMOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZKEVMOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMOwnershipTransferred)
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
		it.Event = new(ZKEVMOwnershipTransferred)
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
func (it *ZKEVMOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMOwnershipTransferred represents a OwnershipTransferred event raised by the ZKEVM contract.
type ZKEVMOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZKEVM *ZKEVMFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZKEVMOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZKEVM.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMOwnershipTransferredIterator{contract: _ZKEVM.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZKEVM *ZKEVMFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZKEVMOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZKEVM.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMOwnershipTransferred)
				if err := _ZKEVM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZKEVM *ZKEVMFilterer) ParseOwnershipTransferred(log types.Log) (*ZKEVMOwnershipTransferred, error) {
	event := new(ZKEVMOwnershipTransferred)
	if err := _ZKEVM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMSubmitBatchesIterator is returned from FilterSubmitBatches and is used to iterate over the raw logs and unpacked data for SubmitBatches events raised by the ZKEVM contract.
type ZKEVMSubmitBatchesIterator struct {
	Event *ZKEVMSubmitBatches // Event containing the contract specifics and raw log

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
func (it *ZKEVMSubmitBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMSubmitBatches)
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
		it.Event = new(ZKEVMSubmitBatches)
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
func (it *ZKEVMSubmitBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMSubmitBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMSubmitBatches represents a SubmitBatches event raised by the ZKEVM contract.
type ZKEVMSubmitBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubmitBatches is a free log retrieval operation binding the contract event 0x85d226f7b6c307388b0488cb75509c840f0154cbed044f4b3c4e52af70569119.
//
// Solidity: event SubmitBatches(uint64 indexed numBatch)
func (_ZKEVM *ZKEVMFilterer) FilterSubmitBatches(opts *bind.FilterOpts, numBatch []uint64) (*ZKEVMSubmitBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _ZKEVM.contract.FilterLogs(opts, "SubmitBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMSubmitBatchesIterator{contract: _ZKEVM.contract, event: "SubmitBatches", logs: logs, sub: sub}, nil
}

// WatchSubmitBatches is a free log subscription operation binding the contract event 0x85d226f7b6c307388b0488cb75509c840f0154cbed044f4b3c4e52af70569119.
//
// Solidity: event SubmitBatches(uint64 indexed numBatch)
func (_ZKEVM *ZKEVMFilterer) WatchSubmitBatches(opts *bind.WatchOpts, sink chan<- *ZKEVMSubmitBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _ZKEVM.contract.WatchLogs(opts, "SubmitBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMSubmitBatches)
				if err := _ZKEVM.contract.UnpackLog(event, "SubmitBatches", log); err != nil {
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

// ParseSubmitBatches is a log parse operation binding the contract event 0x85d226f7b6c307388b0488cb75509c840f0154cbed044f4b3c4e52af70569119.
//
// Solidity: event SubmitBatches(uint64 indexed numBatch)
func (_ZKEVM *ZKEVMFilterer) ParseSubmitBatches(log types.Log) (*ZKEVMSubmitBatches, error) {
	event := new(ZKEVMSubmitBatches)
	if err := _ZKEVM.contract.UnpackLog(event, "SubmitBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
