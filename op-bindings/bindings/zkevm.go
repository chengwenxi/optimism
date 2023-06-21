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
	BlockNumbers uint64
	Transactions []byte
	BlockWitness []byte
	PreStateRoot [32]byte
	WithdrawRoot [32]byte
	Signature    ZKEVMBatchSignature
}

// ZKEVMBatchSignature is an auto generated low-level Go binding around an user-defined struct.
type ZKEVMBatchSignature struct {
	Signers   [][]byte
	Signature []byte
}

// ZKEVMMetaData contains all meta data concerning the ZKEVM contract.
var ZKEVMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_submitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"res\",\"type\":\"string\"}],\"name\":\"ChallengeRes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengeDeposit\",\"type\":\"uint256\"}],\"name\":\"ChallengeState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SubmitBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FINALIZATION_PERIOD_SECONDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_DEPOSIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PORTAL\",\"outputs\":[{\"internalType\":\"contractOptimismPortal\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROOF_WINDOW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"}],\"name\":\"challengeState\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"challenges\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"challengeDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"}],\"name\":\"confirmBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"confirmBatchIndex\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_submitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"}],\"name\":\"isBatchInChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isUserInChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptimismPortal\",\"name\":\"_portal\",\"type\":\"address\"}],\"name\":\"setPortal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"storageBatchs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"originTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNumbers\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blockWitness\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"preStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"signers\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structZKEVM.BatchSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structZKEVM.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"submitBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200299b3803806200299b83398101604081905262000034916200031e565b62000040828262000048565b50506200037d565b600054610100900460ff1615808015620000695750600054600160ff909116105b806200009957506200008630620001d760201b62001bf71760201c565b15801562000099575060005460ff166001145b620001025760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff19166001179055801562000126576000805461ff0019166101001790555b62000130620001e6565b60678054606680546001600160a01b0319166001600160a01b038781169182179092556001600160e01b031990921690851617909155600090815260686020526040812080543492906200018690849062000356565b90915550508015620001d2576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6001600160a01b03163b151590565b600054610100900460ff16620002425760405162461bcd60e51b815260206004820152602b60248201526000805160206200297b83398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000f9565b6200024c6200024e565b565b600054610100900460ff16620002aa5760405162461bcd60e51b815260206004820152602b60248201526000805160206200297b83398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000f9565b6200024c33603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b80516001600160a01b03811681146200031957600080fd5b919050565b600080604083850312156200033257600080fd5b6200033d8362000301565b91506200034d6020840162000301565b90509250929050565b600082198211156200037857634e487b7160e01b600052601160045260246000fd5b500190565b6125ee806200038d6000396000f3fe6080604052600436106101965760003560e01c8063715018a6116100e1578063e1e158a51161008a578063f2fde38b11610064578063f2fde38b146105a0578063f4daa291146105c0578063f71b51f3146105d7578063fc7e286d1461060757600080fd5b8063e1e158a5146104ff578063e291c2f514610529578063eb1ec18f1461058b57600080fd5b80638dc45d9a116100bb5780638dc45d9a146103f6578063acc1245a14610423578063c7ab20391461044357600080fd5b8063715018a6146103a35780638d644bb7146103b85780638da5cb5b146103cb57600080fd5b8063485cc95511610143578063534db0e21161011d578063534db0e21461031157806359ef11201461033e5780636177fd181461035e57600080fd5b8063485cc955146102be5780634b21f578146102d15780634ff56192146102f157600080fd5b80632e1a7d4d116101745780632e1a7d4d146102445780633a4b66f114610264578063423fa8561461026c57600080fd5b8063032ecb381461019b5780630ff754ea146101bd57806323ec851814610214575b600080fd5b3480156101a757600080fd5b506101bb6101b63660046121ed565b610634565b005b3480156101c957600080fd5b506065546101ea9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561022057600080fd5b5061023461022f36600461222a565b61074e565b604051901515815260200161020b565b34801561025057600080fd5b506101bb61025f366004612247565b6108c9565b6101bb610ad0565b34801561027857600080fd5b506067546102a59074010000000000000000000000000000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff909116815260200161020b565b6101bb6102cc366004612260565b610cb3565b3480156102dd57600080fd5b506101bb6102ec366004612299565b610edc565b3480156102fd57600080fd5b506101bb61030c36600461222a565b61135c565b34801561031d57600080fd5b506067546101ea9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561034a57600080fd5b506101bb61035936600461230e565b6113ab565b34801561036a57600080fd5b5061023461037936600461222a565b73ffffffffffffffffffffffffffffffffffffffff16600090815260686020526040902054151590565b3480156103af57600080fd5b506101bb61166d565b6101bb6103c63660046121ed565b611681565b3480156103d757600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff166101ea565b34801561040257600080fd5b506066546101ea9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561042f57600080fd5b5061023461043e3660046121ed565b611ad3565b34801561044f57600080fd5b506104b561045e3660046121ed565b606b60205260009081526040902080546001820154600283015460039093015467ffffffffffffffff8316936801000000000000000090930473ffffffffffffffffffffffffffffffffffffffff16929060ff1685565b6040805167ffffffffffffffff909616865273ffffffffffffffffffffffffffffffffffffffff90941660208601529284019190915260608301521515608082015260a00161020b565b34801561050b57600080fd5b5061051b670de0b6b3a764000081565b60405190815260200161020b565b34801561053557600080fd5b5061056b6105443660046121ed565b60696020526000908152604090208054600182015460028301546003909301549192909184565b60408051948552602085019390935291830152606082015260800161020b565b34801561059757600080fd5b5061051b606481565b3480156105ac57600080fd5b506101bb6105bb36600461222a565b611b40565b3480156105cc57600080fd5b5061051b620186a081565b3480156105e357600080fd5b506102346105f23660046121ed565b606a6020526000908152604090205460ff1681565b34801561061357600080fd5b5061051b61062236600461222a565b60686020526000908152604090205481565b61063d81611ad3565b1561064757600080fd5b67ffffffffffffffff8116600090815260696020526040812060030154429061067490620186a0906123c0565b119050801561070a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43616e6e6f7420636f6e6669726d20626174636820696e73696465207468652060448201527f6368616c6c656e67652077696e646f770000000000000000000000000000000060648201526084015b60405180910390fd5b5067ffffffffffffffff166000908152606a6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b60675460009074010000000000000000000000000000000000000000900467ffffffffffffffff16810361078457506000919050565b60665473ffffffffffffffffffffffffffffffffffffffff9081169083160361080357606754606a906000906107de9060019074010000000000000000000000000000000000000000900467ffffffffffffffff166123d8565b67ffffffffffffffff16815260208101919091526040016000205460ff161592915050565b60005b60675467ffffffffffffffff74010000000000000000000000000000000000000000909104811690821610156108c05767ffffffffffffffff81166000908152606b602052604090205473ffffffffffffffffffffffffffffffffffffffff84811668010000000000000000909204161480156108a0575067ffffffffffffffff81166000908152606b602052604090206003015460ff16155b156108ae5750600192915050565b806108b881612401565b915050610806565b50506000919050565b6108d23361074e565b15610939576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5573657220697320696e206368616c6c656e67650000000000000000000000006044820152606401610701565b336000908152606860205260409020548161095357600080fd5b8181610967670de0b6b3a7640000836123c0565b1115610a915760665473ffffffffffffffffffffffffffffffffffffffff16331480156109b7575060675474010000000000000000000000000000000000000000900467ffffffffffffffff1615155b15610a915760675474010000000000000000000000000000000000000000900467ffffffffffffffff166000908152606960205260409020600301544290610a0390620186a0906123c0565b1115610a91576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f7375626d69747465722073686f756c64207761697420626174636820746f206260448201527f6520636f6e6669726d00000000000000000000000000000000000000000000006064820152608401610701565b81831115610a9c5750805b3360009081526068602052604081208054859290610abb908490612428565b90915550610acb90503382611c13565b505050565b60665473ffffffffffffffffffffffffffffffffffffffff16610b4f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f5375626d69747465722063616e6e6f74206265206164647265737328302900006044820152606401610701565b60665473ffffffffffffffffffffffffffffffffffffffff163314610bd0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f43616c6c6572206e6f74207375626d69747465720000000000000000000000006044820152606401610701565b60665473ffffffffffffffffffffffffffffffffffffffff16600090815260686020526040902054670de0b6b3a764000090610c0d9034906123c0565b1015610c75576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f446f206e6f74206861766520656e6f756768206465706f7369740000000000006044820152606401610701565b60665473ffffffffffffffffffffffffffffffffffffffff1660009081526068602052604081208054349290610cac9084906123c0565b9091555050565b600054610100900460ff1615808015610cd35750600054600160ff909116105b80610ced5750303b158015610ced575060005460ff166001145b610d79576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610701565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610dd757600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610ddf611cbf565b60678054606680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8781169182179092557fffffffff000000000000000000000000000000000000000000000000000000009092169085161790915560009081526068602052604081208054349290610e709084906123c0565b90915550508015610acb57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b60665473ffffffffffffffffffffffffffffffffffffffff16610f5b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f5375626d69747465722063616e6e6f74206265206164647265737328302900006044820152606401610701565b60665473ffffffffffffffffffffffffffffffffffffffff163314610fdc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f43616c6c6572206e6f74207375626d69747465720000000000000000000000006044820152606401610701565b60665473ffffffffffffffffffffffffffffffffffffffff16600090815260686020526040902054670de0b6b3a76400001115611075576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f496e73756666696369656e74207374616b696e6720616d6f756e7400000000006044820152606401610701565b606754819074010000000000000000000000000000000000000000900467ffffffffffffffff1660005b828110156112d657606361f6186000806110b883611d5e565b67ffffffffffffffff881660009081526069602052604090206002015491935091508989878181106110ec576110ec61243f565b90506020028101906110fe919061246e565b6060013514611169576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f5072657669657720737461746520726f6f74206e6f7420657175616c000000006044820152606401610701565b60006111ae8383878d8d8b8181106111835761118361243f565b9050602002810190611195919061246e565b6111a39060408101906124ac565b506060949350505050565b9050600061121e8b8b898181106111c7576111c761243f565b90506020028101906111d9919061246e565b6111e79060408101906124ac565b8d8d8b8181106111f9576111f961243f565b905060200281019061120b919061246e565b6112199060208101906121ed565b611df1565b9050600060208351026020840120905060405180608001604052808d8d8b81811061124b5761124b61243f565b905060200281019061125d919061246e565b608001358152602080820184905260408083018690524260609384015267ffffffffffffffff8d16600090815260698352819020845181559184015160018301558301516002820155910151600390910155886112b981612401565b9950505050505050505080806112ce90612518565b91505061109f565b50606780547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000067ffffffffffffffff84811682029290921792839055604051920416907f85d226f7b6c307388b0488cb75509c840f0154cbed044f4b3c4e52af7056911990600090a250505050565b611364611dfb565b606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b67ffffffffffffffff83166000908152606b602052604090205468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff166113f057600080fd5b67ffffffffffffffff83166000908152606b602052604090206003015460ff1615611477576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f4368616c6c656e676520616c72656164792066696e69736865640000000000006044820152606401610701565b67ffffffffffffffff83166000908152606b602052604081206002015442906114a2906064906123c0565b119050806114ee576114e9846040518060400160405280600781526020017f74696d656f757400000000000000000000000000000000000000000000000000815250611e7c565b611624565b81611555576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f496e76616c69642070726f6f66000000000000000000000000000000000000006044820152606401610701565b67ffffffffffffffff841660009081526069602052604090206001015461157f9084908490611f79565b6115e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f50726f7665206661696c656400000000000000000000000000000000000000006044820152606401610701565b611624846040518060400160405280600d81526020017f70726f7665207375636365737300000000000000000000000000000000000000815250611f94565b50505067ffffffffffffffff166000908152606b6020526040902060030180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b611675611dfb565b61167f600061209f565b565b60675473ffffffffffffffffffffffffffffffffffffffff16611700576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4368616c6c656e6765722063616e6e6f742062652061646472657373283029006044820152606401610701565b60675473ffffffffffffffffffffffffffffffffffffffff163314611781576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f43616c6c6572206e6f74206368616c6c656e67657200000000000000000000006044820152606401610701565b67ffffffffffffffff81166000908152606960205260408120600301549003611806576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4261746368206e6f7420657869737400000000000000000000000000000000006044820152606401610701565b67ffffffffffffffff81166000908152606b602052604090205468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff16156118a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f416c726561647920686173206368616c6c656e676500000000000000000000006044820152606401610701565b67ffffffffffffffff811660009081526069602052604081206003015442906118d690620186a0906123c0565b11905080611966576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f43616e6e6f74206368616c6c656e6765206261746368206f757473696465207460448201527f6865206368616c6c656e67652077696e646f77000000000000000000000000006064820152608401610701565b670de0b6b3a764000034101561197b57600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff16600090815260686020526040812080543492906119b29084906123c0565b90915550506040805160a08101825267ffffffffffffffff848116808352336020808501828152348688018181524260608901908152600060808a01818152888252606b8752908b902099518a54955173ffffffffffffffffffffffffffffffffffffffff1668010000000000000000027fffffffff00000000000000000000000000000000000000000000000000000000909616991698909817939093178855516001880155905160028701559351600390950180549515157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0090961695909517909455845190815292830191909152917fff4bd2fd14b0670e131544fdc1e7c53f06371ccac4a3a8e0445a93d02d68fb05910160405180910390a25050565b67ffffffffffffffff81166000908152606b602052604081205468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1615801590611b3a575067ffffffffffffffff82166000908152606b602052604090206003015460ff16155b92915050565b611b48611dfb565b73ffffffffffffffffffffffffffffffffffffffff8116611beb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610701565b611bf48161209f565b50565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6000611c30835a8460405180602001604052806000815250612116565b905080610acb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c656400000000000000000000000000000000000000000000000000000000006064820152608401610701565b600054610100900460ff16611d56576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610701565b61167f612130565b60008061f6188311611d77575060039261290492509050565b620493e08311611d8f5750600e926201107692509050565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f434952435549545f434f4e4649470000000000000000000000000000000000006044820152606401610701565b60005b9392505050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461167f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610701565b67ffffffffffffffff82166000908152606b602090815260408083205460665473ffffffffffffffffffffffffffffffffffffffff9081168552606890935290832080546801000000000000000090920490921692670de0b6b3a76400009291611ee7908490612428565b909155505073ffffffffffffffffffffffffffffffffffffffff811660009081526068602052604081208054670de0b6b3a76400009290611f299084906123c0565b925050819055508267ffffffffffffffff167fe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a8284604051611f6c929190612550565b60405180910390a2505050565b6000828103611f8a57506000611df4565b5060019392505050565b67ffffffffffffffff82166000908152606b6020908152604080832080546001909101546801000000000000000090910473ffffffffffffffffffffffffffffffffffffffff168085526068909352908320805492939192839290611ffa908490612428565b909155505060665473ffffffffffffffffffffffffffffffffffffffff16600090815260686020526040812080548392906120369084906123c0565b909155505060665460405167ffffffffffffffff8616917fe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a916120919173ffffffffffffffffffffffffffffffffffffffff16908790612550565b60405180910390a250505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600080600080845160208601878a8af19695505050505050565b600054610100900460ff166121c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610701565b61167f3361209f565b803567ffffffffffffffff811681146121e857600080fd5b919050565b6000602082840312156121ff57600080fd5b611df4826121d0565b73ffffffffffffffffffffffffffffffffffffffff81168114611bf457600080fd5b60006020828403121561223c57600080fd5b8135611df481612208565b60006020828403121561225957600080fd5b5035919050565b6000806040838503121561227357600080fd5b823561227e81612208565b9150602083013561228e81612208565b809150509250929050565b600080602083850312156122ac57600080fd5b823567ffffffffffffffff808211156122c457600080fd5b818501915085601f8301126122d857600080fd5b8135818111156122e757600080fd5b8660208260051b85010111156122fc57600080fd5b60209290920196919550909350505050565b60008060006040848603121561232357600080fd5b61232c846121d0565b9250602084013567ffffffffffffffff8082111561234957600080fd5b818601915086601f83011261235d57600080fd5b81358181111561236c57600080fd5b87602082850101111561237e57600080fd5b6020830194508093505050509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156123d3576123d3612391565b500190565b600067ffffffffffffffff838116908316818110156123f9576123f9612391565b039392505050565b600067ffffffffffffffff80831681810361241e5761241e612391565b6001019392505050565b60008282101561243a5761243a612391565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff418336030181126124a257600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126124e157600080fd5b83018035915067ffffffffffffffff8211156124fc57600080fd5b60200191503681900382131561251157600080fd5b9250929050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361254957612549612391565b5060010190565b73ffffffffffffffffffffffffffffffffffffffff8316815260006020604081840152835180604085015260005b8181101561259a5785810183015185820160600152820161257e565b818111156125ac576000606083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160600194935050505056fea164736f6c634300080f000a496e697469616c697a61626c653a20636f6e7472616374206973206e6f742069",
}

// ZKEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKEVMMetaData.ABI instead.
var ZKEVMABI = ZKEVMMetaData.ABI

// ZKEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZKEVMMetaData.Bin instead.
var ZKEVMBin = ZKEVMMetaData.Bin

// DeployZKEVM deploys a new Ethereum contract, binding an instance of ZKEVM to it.
func DeployZKEVM(auth *bind.TransactOpts, backend bind.ContractBackend, _submitter common.Address, _challenger common.Address) (common.Address, *types.Transaction, *ZKEVM, error) {
	parsed, err := ZKEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZKEVMBin), backend, _submitter, _challenger)
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

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_ZKEVM *ZKEVMCaller) FINALIZATIONPERIODSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "FINALIZATION_PERIOD_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_ZKEVM *ZKEVMSession) FINALIZATIONPERIODSECONDS() (*big.Int, error) {
	return _ZKEVM.Contract.FINALIZATIONPERIODSECONDS(&_ZKEVM.CallOpts)
}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_ZKEVM *ZKEVMCallerSession) FINALIZATIONPERIODSECONDS() (*big.Int, error) {
	return _ZKEVM.Contract.FINALIZATIONPERIODSECONDS(&_ZKEVM.CallOpts)
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

// PROOFWINDOW is a free data retrieval call binding the contract method 0xeb1ec18f.
//
// Solidity: function PROOF_WINDOW() view returns(uint256)
func (_ZKEVM *ZKEVMCaller) PROOFWINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "PROOF_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROOFWINDOW is a free data retrieval call binding the contract method 0xeb1ec18f.
//
// Solidity: function PROOF_WINDOW() view returns(uint256)
func (_ZKEVM *ZKEVMSession) PROOFWINDOW() (*big.Int, error) {
	return _ZKEVM.Contract.PROOFWINDOW(&_ZKEVM.CallOpts)
}

// PROOFWINDOW is a free data retrieval call binding the contract method 0xeb1ec18f.
//
// Solidity: function PROOF_WINDOW() view returns(uint256)
func (_ZKEVM *ZKEVMCallerSession) PROOFWINDOW() (*big.Int, error) {
	return _ZKEVM.Contract.PROOFWINDOW(&_ZKEVM.CallOpts)
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
// Solidity: function challenges(uint64 ) view returns(uint64 batchIndex, address challenger, uint256 challengeDeposit, uint256 startTime, bool finished)
func (_ZKEVM *ZKEVMCaller) Challenges(opts *bind.CallOpts, arg0 uint64) (struct {
	BatchIndex       uint64
	Challenger       common.Address
	ChallengeDeposit *big.Int
	StartTime        *big.Int
	Finished         bool
}, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "challenges", arg0)

	outstruct := new(struct {
		BatchIndex       uint64
		Challenger       common.Address
		ChallengeDeposit *big.Int
		StartTime        *big.Int
		Finished         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchIndex = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Challenger = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ChallengeDeposit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Finished = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Challenges is a free data retrieval call binding the contract method 0xc7ab2039.
//
// Solidity: function challenges(uint64 ) view returns(uint64 batchIndex, address challenger, uint256 challengeDeposit, uint256 startTime, bool finished)
func (_ZKEVM *ZKEVMSession) Challenges(arg0 uint64) (struct {
	BatchIndex       uint64
	Challenger       common.Address
	ChallengeDeposit *big.Int
	StartTime        *big.Int
	Finished         bool
}, error) {
	return _ZKEVM.Contract.Challenges(&_ZKEVM.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0xc7ab2039.
//
// Solidity: function challenges(uint64 ) view returns(uint64 batchIndex, address challenger, uint256 challengeDeposit, uint256 startTime, bool finished)
func (_ZKEVM *ZKEVMCallerSession) Challenges(arg0 uint64) (struct {
	BatchIndex       uint64
	Challenger       common.Address
	ChallengeDeposit *big.Int
	StartTime        *big.Int
	Finished         bool
}, error) {
	return _ZKEVM.Contract.Challenges(&_ZKEVM.CallOpts, arg0)
}

// ConfirmBatchIndex is a free data retrieval call binding the contract method 0xf71b51f3.
//
// Solidity: function confirmBatchIndex(uint64 ) view returns(bool)
func (_ZKEVM *ZKEVMCaller) ConfirmBatchIndex(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "confirmBatchIndex", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ConfirmBatchIndex is a free data retrieval call binding the contract method 0xf71b51f3.
//
// Solidity: function confirmBatchIndex(uint64 ) view returns(bool)
func (_ZKEVM *ZKEVMSession) ConfirmBatchIndex(arg0 uint64) (bool, error) {
	return _ZKEVM.Contract.ConfirmBatchIndex(&_ZKEVM.CallOpts, arg0)
}

// ConfirmBatchIndex is a free data retrieval call binding the contract method 0xf71b51f3.
//
// Solidity: function confirmBatchIndex(uint64 ) view returns(bool)
func (_ZKEVM *ZKEVMCallerSession) ConfirmBatchIndex(arg0 uint64) (bool, error) {
	return _ZKEVM.Contract.ConfirmBatchIndex(&_ZKEVM.CallOpts, arg0)
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

// IsBatchInChallenge is a free data retrieval call binding the contract method 0xacc1245a.
//
// Solidity: function isBatchInChallenge(uint64 batchIndex) view returns(bool)
func (_ZKEVM *ZKEVMCaller) IsBatchInChallenge(opts *bind.CallOpts, batchIndex uint64) (bool, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "isBatchInChallenge", batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBatchInChallenge is a free data retrieval call binding the contract method 0xacc1245a.
//
// Solidity: function isBatchInChallenge(uint64 batchIndex) view returns(bool)
func (_ZKEVM *ZKEVMSession) IsBatchInChallenge(batchIndex uint64) (bool, error) {
	return _ZKEVM.Contract.IsBatchInChallenge(&_ZKEVM.CallOpts, batchIndex)
}

// IsBatchInChallenge is a free data retrieval call binding the contract method 0xacc1245a.
//
// Solidity: function isBatchInChallenge(uint64 batchIndex) view returns(bool)
func (_ZKEVM *ZKEVMCallerSession) IsBatchInChallenge(batchIndex uint64) (bool, error) {
	return _ZKEVM.Contract.IsBatchInChallenge(&_ZKEVM.CallOpts, batchIndex)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address addr) view returns(bool)
func (_ZKEVM *ZKEVMCaller) IsStaked(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "isStaked", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address addr) view returns(bool)
func (_ZKEVM *ZKEVMSession) IsStaked(addr common.Address) (bool, error) {
	return _ZKEVM.Contract.IsStaked(&_ZKEVM.CallOpts, addr)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address addr) view returns(bool)
func (_ZKEVM *ZKEVMCallerSession) IsStaked(addr common.Address) (bool, error) {
	return _ZKEVM.Contract.IsStaked(&_ZKEVM.CallOpts, addr)
}

// IsUserInChallenge is a free data retrieval call binding the contract method 0x23ec8518.
//
// Solidity: function isUserInChallenge(address user) view returns(bool)
func (_ZKEVM *ZKEVMCaller) IsUserInChallenge(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "isUserInChallenge", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserInChallenge is a free data retrieval call binding the contract method 0x23ec8518.
//
// Solidity: function isUserInChallenge(address user) view returns(bool)
func (_ZKEVM *ZKEVMSession) IsUserInChallenge(user common.Address) (bool, error) {
	return _ZKEVM.Contract.IsUserInChallenge(&_ZKEVM.CallOpts, user)
}

// IsUserInChallenge is a free data retrieval call binding the contract method 0x23ec8518.
//
// Solidity: function isUserInChallenge(address user) view returns(bool)
func (_ZKEVM *ZKEVMCallerSession) IsUserInChallenge(user common.Address) (bool, error) {
	return _ZKEVM.Contract.IsUserInChallenge(&_ZKEVM.CallOpts, user)
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

// StorageBatchs is a free data retrieval call binding the contract method 0xe291c2f5.
//
// Solidity: function storageBatchs(uint64 ) view returns(bytes32 withdrawRoot, bytes32 commitment, bytes32 stateRoot, uint256 originTimestamp)
func (_ZKEVM *ZKEVMCaller) StorageBatchs(opts *bind.CallOpts, arg0 uint64) (struct {
	WithdrawRoot    [32]byte
	Commitment      [32]byte
	StateRoot       [32]byte
	OriginTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _ZKEVM.contract.Call(opts, &out, "storageBatchs", arg0)

	outstruct := new(struct {
		WithdrawRoot    [32]byte
		Commitment      [32]byte
		StateRoot       [32]byte
		OriginTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WithdrawRoot = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Commitment = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.StateRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.OriginTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StorageBatchs is a free data retrieval call binding the contract method 0xe291c2f5.
//
// Solidity: function storageBatchs(uint64 ) view returns(bytes32 withdrawRoot, bytes32 commitment, bytes32 stateRoot, uint256 originTimestamp)
func (_ZKEVM *ZKEVMSession) StorageBatchs(arg0 uint64) (struct {
	WithdrawRoot    [32]byte
	Commitment      [32]byte
	StateRoot       [32]byte
	OriginTimestamp *big.Int
}, error) {
	return _ZKEVM.Contract.StorageBatchs(&_ZKEVM.CallOpts, arg0)
}

// StorageBatchs is a free data retrieval call binding the contract method 0xe291c2f5.
//
// Solidity: function storageBatchs(uint64 ) view returns(bytes32 withdrawRoot, bytes32 commitment, bytes32 stateRoot, uint256 originTimestamp)
func (_ZKEVM *ZKEVMCallerSession) StorageBatchs(arg0 uint64) (struct {
	WithdrawRoot    [32]byte
	Commitment      [32]byte
	StateRoot       [32]byte
	OriginTimestamp *big.Int
}, error) {
	return _ZKEVM.Contract.StorageBatchs(&_ZKEVM.CallOpts, arg0)
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
// Solidity: function challengeState(uint64 batchIndex) payable returns()
func (_ZKEVM *ZKEVMTransactor) ChallengeState(opts *bind.TransactOpts, batchIndex uint64) (*types.Transaction, error) {
	return _ZKEVM.contract.Transact(opts, "challengeState", batchIndex)
}

// ChallengeState is a paid mutator transaction binding the contract method 0x8d644bb7.
//
// Solidity: function challengeState(uint64 batchIndex) payable returns()
func (_ZKEVM *ZKEVMSession) ChallengeState(batchIndex uint64) (*types.Transaction, error) {
	return _ZKEVM.Contract.ChallengeState(&_ZKEVM.TransactOpts, batchIndex)
}

// ChallengeState is a paid mutator transaction binding the contract method 0x8d644bb7.
//
// Solidity: function challengeState(uint64 batchIndex) payable returns()
func (_ZKEVM *ZKEVMTransactorSession) ChallengeState(batchIndex uint64) (*types.Transaction, error) {
	return _ZKEVM.Contract.ChallengeState(&_ZKEVM.TransactOpts, batchIndex)
}

// ConfirmBatch is a paid mutator transaction binding the contract method 0x032ecb38.
//
// Solidity: function confirmBatch(uint64 batchIndex) returns()
func (_ZKEVM *ZKEVMTransactor) ConfirmBatch(opts *bind.TransactOpts, batchIndex uint64) (*types.Transaction, error) {
	return _ZKEVM.contract.Transact(opts, "confirmBatch", batchIndex)
}

// ConfirmBatch is a paid mutator transaction binding the contract method 0x032ecb38.
//
// Solidity: function confirmBatch(uint64 batchIndex) returns()
func (_ZKEVM *ZKEVMSession) ConfirmBatch(batchIndex uint64) (*types.Transaction, error) {
	return _ZKEVM.Contract.ConfirmBatch(&_ZKEVM.TransactOpts, batchIndex)
}

// ConfirmBatch is a paid mutator transaction binding the contract method 0x032ecb38.
//
// Solidity: function confirmBatch(uint64 batchIndex) returns()
func (_ZKEVM *ZKEVMTransactorSession) ConfirmBatch(batchIndex uint64) (*types.Transaction, error) {
	return _ZKEVM.Contract.ConfirmBatch(&_ZKEVM.TransactOpts, batchIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _submitter, address _challenger) payable returns()
func (_ZKEVM *ZKEVMTransactor) Initialize(opts *bind.TransactOpts, _submitter common.Address, _challenger common.Address) (*types.Transaction, error) {
	return _ZKEVM.contract.Transact(opts, "initialize", _submitter, _challenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _submitter, address _challenger) payable returns()
func (_ZKEVM *ZKEVMSession) Initialize(_submitter common.Address, _challenger common.Address) (*types.Transaction, error) {
	return _ZKEVM.Contract.Initialize(&_ZKEVM.TransactOpts, _submitter, _challenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _submitter, address _challenger) payable returns()
func (_ZKEVM *ZKEVMTransactorSession) Initialize(_submitter common.Address, _challenger common.Address) (*types.Transaction, error) {
	return _ZKEVM.Contract.Initialize(&_ZKEVM.TransactOpts, _submitter, _challenger)
}

// ProveState is a paid mutator transaction binding the contract method 0x59ef1120.
//
// Solidity: function proveState(uint64 batchIndex, bytes proof) returns()
func (_ZKEVM *ZKEVMTransactor) ProveState(opts *bind.TransactOpts, batchIndex uint64, proof []byte) (*types.Transaction, error) {
	return _ZKEVM.contract.Transact(opts, "proveState", batchIndex, proof)
}

// ProveState is a paid mutator transaction binding the contract method 0x59ef1120.
//
// Solidity: function proveState(uint64 batchIndex, bytes proof) returns()
func (_ZKEVM *ZKEVMSession) ProveState(batchIndex uint64, proof []byte) (*types.Transaction, error) {
	return _ZKEVM.Contract.ProveState(&_ZKEVM.TransactOpts, batchIndex, proof)
}

// ProveState is a paid mutator transaction binding the contract method 0x59ef1120.
//
// Solidity: function proveState(uint64 batchIndex, bytes proof) returns()
func (_ZKEVM *ZKEVMTransactorSession) ProveState(batchIndex uint64, proof []byte) (*types.Transaction, error) {
	return _ZKEVM.Contract.ProveState(&_ZKEVM.TransactOpts, batchIndex, proof)
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

// SetPortal is a paid mutator transaction binding the contract method 0x4ff56192.
//
// Solidity: function setPortal(address _portal) returns()
func (_ZKEVM *ZKEVMTransactor) SetPortal(opts *bind.TransactOpts, _portal common.Address) (*types.Transaction, error) {
	return _ZKEVM.contract.Transact(opts, "setPortal", _portal)
}

// SetPortal is a paid mutator transaction binding the contract method 0x4ff56192.
//
// Solidity: function setPortal(address _portal) returns()
func (_ZKEVM *ZKEVMSession) SetPortal(_portal common.Address) (*types.Transaction, error) {
	return _ZKEVM.Contract.SetPortal(&_ZKEVM.TransactOpts, _portal)
}

// SetPortal is a paid mutator transaction binding the contract method 0x4ff56192.
//
// Solidity: function setPortal(address _portal) returns()
func (_ZKEVM *ZKEVMTransactorSession) SetPortal(_portal common.Address) (*types.Transaction, error) {
	return _ZKEVM.Contract.SetPortal(&_ZKEVM.TransactOpts, _portal)
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

// SubmitBatches is a paid mutator transaction binding the contract method 0x4b21f578.
//
// Solidity: function submitBatches((uint64,bytes,bytes,bytes32,bytes32,(bytes[],bytes))[] batches) returns()
func (_ZKEVM *ZKEVMTransactor) SubmitBatches(opts *bind.TransactOpts, batches []ZKEVMBatchData) (*types.Transaction, error) {
	return _ZKEVM.contract.Transact(opts, "submitBatches", batches)
}

// SubmitBatches is a paid mutator transaction binding the contract method 0x4b21f578.
//
// Solidity: function submitBatches((uint64,bytes,bytes,bytes32,bytes32,(bytes[],bytes))[] batches) returns()
func (_ZKEVM *ZKEVMSession) SubmitBatches(batches []ZKEVMBatchData) (*types.Transaction, error) {
	return _ZKEVM.Contract.SubmitBatches(&_ZKEVM.TransactOpts, batches)
}

// SubmitBatches is a paid mutator transaction binding the contract method 0x4b21f578.
//
// Solidity: function submitBatches((uint64,bytes,bytes,bytes32,bytes32,(bytes[],bytes))[] batches) returns()
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

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_ZKEVM *ZKEVMTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZKEVM.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_ZKEVM *ZKEVMSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ZKEVM.Contract.Withdraw(&_ZKEVM.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_ZKEVM *ZKEVMTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ZKEVM.Contract.Withdraw(&_ZKEVM.TransactOpts, amount)
}

// ZKEVMChallengeResIterator is returned from FilterChallengeRes and is used to iterate over the raw logs and unpacked data for ChallengeRes events raised by the ZKEVM contract.
type ZKEVMChallengeResIterator struct {
	Event *ZKEVMChallengeRes // Event containing the contract specifics and raw log

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
func (it *ZKEVMChallengeResIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMChallengeRes)
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
		it.Event = new(ZKEVMChallengeRes)
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
func (it *ZKEVMChallengeResIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMChallengeResIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMChallengeRes represents a ChallengeRes event raised by the ZKEVM contract.
type ZKEVMChallengeRes struct {
	BatchIndex *big.Int
	Winner     common.Address
	Res        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChallengeRes is a free log retrieval operation binding the contract event 0xe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a.
//
// Solidity: event ChallengeRes(uint256 indexed batchIndex, address winner, string res)
func (_ZKEVM *ZKEVMFilterer) FilterChallengeRes(opts *bind.FilterOpts, batchIndex []*big.Int) (*ZKEVMChallengeResIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _ZKEVM.contract.FilterLogs(opts, "ChallengeRes", batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMChallengeResIterator{contract: _ZKEVM.contract, event: "ChallengeRes", logs: logs, sub: sub}, nil
}

// WatchChallengeRes is a free log subscription operation binding the contract event 0xe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a.
//
// Solidity: event ChallengeRes(uint256 indexed batchIndex, address winner, string res)
func (_ZKEVM *ZKEVMFilterer) WatchChallengeRes(opts *bind.WatchOpts, sink chan<- *ZKEVMChallengeRes, batchIndex []*big.Int) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _ZKEVM.contract.WatchLogs(opts, "ChallengeRes", batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMChallengeRes)
				if err := _ZKEVM.contract.UnpackLog(event, "ChallengeRes", log); err != nil {
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

// ParseChallengeRes is a log parse operation binding the contract event 0xe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a.
//
// Solidity: event ChallengeRes(uint256 indexed batchIndex, address winner, string res)
func (_ZKEVM *ZKEVMFilterer) ParseChallengeRes(log types.Log) (*ZKEVMChallengeRes, error) {
	event := new(ZKEVMChallengeRes)
	if err := _ZKEVM.contract.UnpackLog(event, "ChallengeRes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMChallengeStateIterator is returned from FilterChallengeState and is used to iterate over the raw logs and unpacked data for ChallengeState events raised by the ZKEVM contract.
type ZKEVMChallengeStateIterator struct {
	Event *ZKEVMChallengeState // Event containing the contract specifics and raw log

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
func (it *ZKEVMChallengeStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMChallengeState)
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
		it.Event = new(ZKEVMChallengeState)
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
func (it *ZKEVMChallengeStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMChallengeStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMChallengeState represents a ChallengeState event raised by the ZKEVM contract.
type ZKEVMChallengeState struct {
	BatchIndex       *big.Int
	Challenger       common.Address
	ChallengeDeposit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterChallengeState is a free log retrieval operation binding the contract event 0xff4bd2fd14b0670e131544fdc1e7c53f06371ccac4a3a8e0445a93d02d68fb05.
//
// Solidity: event ChallengeState(uint256 indexed batchIndex, address challenger, uint256 challengeDeposit)
func (_ZKEVM *ZKEVMFilterer) FilterChallengeState(opts *bind.FilterOpts, batchIndex []*big.Int) (*ZKEVMChallengeStateIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _ZKEVM.contract.FilterLogs(opts, "ChallengeState", batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMChallengeStateIterator{contract: _ZKEVM.contract, event: "ChallengeState", logs: logs, sub: sub}, nil
}

// WatchChallengeState is a free log subscription operation binding the contract event 0xff4bd2fd14b0670e131544fdc1e7c53f06371ccac4a3a8e0445a93d02d68fb05.
//
// Solidity: event ChallengeState(uint256 indexed batchIndex, address challenger, uint256 challengeDeposit)
func (_ZKEVM *ZKEVMFilterer) WatchChallengeState(opts *bind.WatchOpts, sink chan<- *ZKEVMChallengeState, batchIndex []*big.Int) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _ZKEVM.contract.WatchLogs(opts, "ChallengeState", batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMChallengeState)
				if err := _ZKEVM.contract.UnpackLog(event, "ChallengeState", log); err != nil {
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

// ParseChallengeState is a log parse operation binding the contract event 0xff4bd2fd14b0670e131544fdc1e7c53f06371ccac4a3a8e0445a93d02d68fb05.
//
// Solidity: event ChallengeState(uint256 indexed batchIndex, address challenger, uint256 challengeDeposit)
func (_ZKEVM *ZKEVMFilterer) ParseChallengeState(log types.Log) (*ZKEVMChallengeState, error) {
	event := new(ZKEVMChallengeState)
	if err := _ZKEVM.contract.UnpackLog(event, "ChallengeState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
