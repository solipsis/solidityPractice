// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SpawnABI is the input ABI used to generate the binding from.
const SpawnABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"createTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"subAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"payable\":true,\"type\":\"fallback\"}]"

// SpawnBin is the compiled bytecode used for deploying new contracts.
const SpawnBin = `0x6060604052341561000c57fe5b5b6103258061001c6000396000f300606060405236156100495763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ac552d8114610052578063e58ad4531461008b575b6100505b5b565b005b341561005a57fe5b6100626100c4565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b341561009357fe5b61006261012a565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60006113886100d1610146565b90815260405190819003602001906000f08015156100eb57fe5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff92831617908190551690505b90565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6040516101a3806101578339019056006060604052341561000c57fe5b6040516020806101a383398101604052515b60018054600160a060020a03191633600160a060020a031617905560008190555b505b610153806100506000396000f300606060405236156100495763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166342cde4e881146100aa5780638da5cb5b146100cc575b6100a85b60005473ffffffffffffffffffffffffffffffffffffffff301631106100a55760015460405173ffffffffffffffffffffffffffffffffffffffff9182169130163180156108fc02916000818181858888f150505050505b5b565b005b34156100b257fe5b6100ba610105565b60408051918252519081900360200190f35b34156100d457fe5b6100dc61010b565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005481565b60015473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a723058205c62777a2bfc6020600c2359e10af906dd024ba520979257d5be70350a8a05720029a165627a7a72305820dd24a301572c06800c2cfa0dabee3b83aea6ab03fa331653e8db31e94f8649b40029`

// DeploySpawn deploys a new Ethereum contract, binding an instance of Spawn to it.
func DeploySpawn(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Spawn, error) {
	parsed, err := abi.JSON(strings.NewReader(SpawnABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SpawnBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Spawn{SpawnCaller: SpawnCaller{contract: contract}, SpawnTransactor: SpawnTransactor{contract: contract}}, nil
}

// Spawn is an auto generated Go binding around an Ethereum contract.
type Spawn struct {
	SpawnCaller     // Read-only binding to the contract
	SpawnTransactor // Write-only binding to the contract
}

// SpawnCaller is an auto generated read-only Go binding around an Ethereum contract.
type SpawnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpawnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SpawnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpawnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SpawnSession struct {
	Contract     *Spawn            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpawnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SpawnCallerSession struct {
	Contract *SpawnCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SpawnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SpawnTransactorSession struct {
	Contract     *SpawnTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpawnRaw is an auto generated low-level Go binding around an Ethereum contract.
type SpawnRaw struct {
	Contract *Spawn // Generic contract binding to access the raw methods on
}

// SpawnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SpawnCallerRaw struct {
	Contract *SpawnCaller // Generic read-only contract binding to access the raw methods on
}

// SpawnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SpawnTransactorRaw struct {
	Contract *SpawnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpawn creates a new instance of Spawn, bound to a specific deployed contract.
func NewSpawn(address common.Address, backend bind.ContractBackend) (*Spawn, error) {
	contract, err := bindSpawn(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Spawn{SpawnCaller: SpawnCaller{contract: contract}, SpawnTransactor: SpawnTransactor{contract: contract}}, nil
}

// NewSpawnCaller creates a new read-only instance of Spawn, bound to a specific deployed contract.
func NewSpawnCaller(address common.Address, caller bind.ContractCaller) (*SpawnCaller, error) {
	contract, err := bindSpawn(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SpawnCaller{contract: contract}, nil
}

// NewSpawnTransactor creates a new write-only instance of Spawn, bound to a specific deployed contract.
func NewSpawnTransactor(address common.Address, transactor bind.ContractTransactor) (*SpawnTransactor, error) {
	contract, err := bindSpawn(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SpawnTransactor{contract: contract}, nil
}

// bindSpawn binds a generic wrapper to an already deployed contract.
func bindSpawn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SpawnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spawn *SpawnRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Spawn.Contract.SpawnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spawn *SpawnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spawn.Contract.SpawnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spawn *SpawnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spawn.Contract.SpawnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spawn *SpawnCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Spawn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spawn *SpawnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spawn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spawn *SpawnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spawn.Contract.contract.Transact(opts, method, params...)
}

// SubAddress is a free data retrieval call binding the contract method 0xe58ad453.
//
// Solidity: function subAddress() constant returns(address)
func (_Spawn *SpawnCaller) SubAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Spawn.contract.Call(opts, out, "subAddress")
	return *ret0, err
}

// SubAddress is a free data retrieval call binding the contract method 0xe58ad453.
//
// Solidity: function subAddress() constant returns(address)
func (_Spawn *SpawnSession) SubAddress() (common.Address, error) {
	return _Spawn.Contract.SubAddress(&_Spawn.CallOpts)
}

// SubAddress is a free data retrieval call binding the contract method 0xe58ad453.
//
// Solidity: function subAddress() constant returns(address)
func (_Spawn *SpawnCallerSession) SubAddress() (common.Address, error) {
	return _Spawn.Contract.SubAddress(&_Spawn.CallOpts)
}

// CreateTransfer is a paid mutator transaction binding the contract method 0x01ac552d.
//
// Solidity: function createTransfer() returns(address)
func (_Spawn *SpawnTransactor) CreateTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spawn.contract.Transact(opts, "createTransfer")
}

// CreateTransfer is a paid mutator transaction binding the contract method 0x01ac552d.
//
// Solidity: function createTransfer() returns(address)
func (_Spawn *SpawnSession) CreateTransfer() (*types.Transaction, error) {
	return _Spawn.Contract.CreateTransfer(&_Spawn.TransactOpts)
}

// CreateTransfer is a paid mutator transaction binding the contract method 0x01ac552d.
//
// Solidity: function createTransfer() returns(address)
func (_Spawn *SpawnTransactorSession) CreateTransfer() (*types.Transaction, error) {
	return _Spawn.Contract.CreateTransfer(&_Spawn.TransactOpts)
}

// TransferABI is the input ABI used to generate the binding from.
const TransferABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"}]"

// TransferBin is the compiled bytecode used for deploying new contracts.
const TransferBin = `0x6060604052341561000c57fe5b6040516020806101a383398101604052515b60018054600160a060020a03191633600160a060020a031617905560008190555b505b610153806100506000396000f300606060405236156100495763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166342cde4e881146100aa5780638da5cb5b146100cc575b6100a85b60005473ffffffffffffffffffffffffffffffffffffffff301631106100a55760015460405173ffffffffffffffffffffffffffffffffffffffff9182169130163180156108fc02916000818181858888f150505050505b5b565b005b34156100b257fe5b6100ba610105565b60408051918252519081900360200190f35b34156100d457fe5b6100dc61010b565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005481565b60015473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a723058205c62777a2bfc6020600c2359e10af906dd024ba520979257d5be70350a8a05720029`

// DeployTransfer deploys a new Ethereum contract, binding an instance of Transfer to it.
func DeployTransfer(auth *bind.TransactOpts, backend bind.ContractBackend, _threshold *big.Int) (common.Address, *types.Transaction, *Transfer, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TransferBin), backend, _threshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Transfer{TransferCaller: TransferCaller{contract: contract}, TransferTransactor: TransferTransactor{contract: contract}}, nil
}

// Transfer is an auto generated Go binding around an Ethereum contract.
type Transfer struct {
	TransferCaller     // Read-only binding to the contract
	TransferTransactor // Write-only binding to the contract
}

// TransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferSession struct {
	Contract     *Transfer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferCallerSession struct {
	Contract *TransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferTransactorSession struct {
	Contract     *TransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferRaw struct {
	Contract *Transfer // Generic contract binding to access the raw methods on
}

// TransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferCallerRaw struct {
	Contract *TransferCaller // Generic read-only contract binding to access the raw methods on
}

// TransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferTransactorRaw struct {
	Contract *TransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransfer creates a new instance of Transfer, bound to a specific deployed contract.
func NewTransfer(address common.Address, backend bind.ContractBackend) (*Transfer, error) {
	contract, err := bindTransfer(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transfer{TransferCaller: TransferCaller{contract: contract}, TransferTransactor: TransferTransactor{contract: contract}}, nil
}

// NewTransferCaller creates a new read-only instance of Transfer, bound to a specific deployed contract.
func NewTransferCaller(address common.Address, caller bind.ContractCaller) (*TransferCaller, error) {
	contract, err := bindTransfer(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TransferCaller{contract: contract}, nil
}

// NewTransferTransactor creates a new write-only instance of Transfer, bound to a specific deployed contract.
func NewTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferTransactor, error) {
	contract, err := bindTransfer(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TransferTransactor{contract: contract}, nil
}

// bindTransfer binds a generic wrapper to an already deployed contract.
func bindTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.TransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.Contract.TransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer.Contract.TransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Transfer *TransferCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Transfer.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Transfer *TransferSession) Owner() (common.Address, error) {
	return _Transfer.Contract.Owner(&_Transfer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Transfer *TransferCallerSession) Owner() (common.Address, error) {
	return _Transfer.Contract.Owner(&_Transfer.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() constant returns(uint256)
func (_Transfer *TransferCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Transfer.contract.Call(opts, out, "threshold")
	return *ret0, err
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() constant returns(uint256)
func (_Transfer *TransferSession) Threshold() (*big.Int, error) {
	return _Transfer.Contract.Threshold(&_Transfer.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() constant returns(uint256)
func (_Transfer *TransferCallerSession) Threshold() (*big.Int, error) {
	return _Transfer.Contract.Threshold(&_Transfer.CallOpts)
}
