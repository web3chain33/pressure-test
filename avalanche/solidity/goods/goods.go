// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package goods

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

// GoodsMetaData contains all meta data concerning the Goods contract.
var GoodsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"batchMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"batchTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSuccessNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GoodsABI is the input ABI used to generate the binding from.
// Deprecated: Use GoodsMetaData.ABI instead.
var GoodsABI = GoodsMetaData.ABI

// Goods is an auto generated Go binding around an Ethereum contract.
type Goods struct {
	GoodsCaller     // Read-only binding to the contract
	GoodsTransactor // Write-only binding to the contract
	GoodsFilterer   // Log filterer for contract events
}

// GoodsCaller is an auto generated read-only Go binding around an Ethereum contract.
type GoodsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoodsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GoodsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoodsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GoodsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoodsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GoodsSession struct {
	Contract     *Goods            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GoodsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GoodsCallerSession struct {
	Contract *GoodsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GoodsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GoodsTransactorSession struct {
	Contract     *GoodsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GoodsRaw is an auto generated low-level Go binding around an Ethereum contract.
type GoodsRaw struct {
	Contract *Goods // Generic contract binding to access the raw methods on
}

// GoodsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GoodsCallerRaw struct {
	Contract *GoodsCaller // Generic read-only contract binding to access the raw methods on
}

// GoodsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GoodsTransactorRaw struct {
	Contract *GoodsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGoods creates a new instance of Goods, bound to a specific deployed contract.
func NewGoods(address common.Address, backend bind.ContractBackend) (*Goods, error) {
	contract, err := bindGoods(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Goods{GoodsCaller: GoodsCaller{contract: contract}, GoodsTransactor: GoodsTransactor{contract: contract}, GoodsFilterer: GoodsFilterer{contract: contract}}, nil
}

// NewGoodsCaller creates a new read-only instance of Goods, bound to a specific deployed contract.
func NewGoodsCaller(address common.Address, caller bind.ContractCaller) (*GoodsCaller, error) {
	contract, err := bindGoods(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GoodsCaller{contract: contract}, nil
}

// NewGoodsTransactor creates a new write-only instance of Goods, bound to a specific deployed contract.
func NewGoodsTransactor(address common.Address, transactor bind.ContractTransactor) (*GoodsTransactor, error) {
	contract, err := bindGoods(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GoodsTransactor{contract: contract}, nil
}

// NewGoodsFilterer creates a new log filterer instance of Goods, bound to a specific deployed contract.
func NewGoodsFilterer(address common.Address, filterer bind.ContractFilterer) (*GoodsFilterer, error) {
	contract, err := bindGoods(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GoodsFilterer{contract: contract}, nil
}

// bindGoods binds a generic wrapper to an already deployed contract.
func bindGoods(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GoodsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Goods *GoodsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Goods.Contract.GoodsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Goods *GoodsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Goods.Contract.GoodsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Goods *GoodsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Goods.Contract.GoodsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Goods *GoodsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Goods.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Goods *GoodsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Goods.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Goods *GoodsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Goods.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Goods *GoodsCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Goods.contract.Call(opts, &out, "balanceOf", account, id)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Goods *GoodsSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Goods.Contract.BalanceOf(&_Goods.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Goods *GoodsCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Goods.Contract.BalanceOf(&_Goods.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Goods *GoodsCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Goods.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)
	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Goods *GoodsSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Goods.Contract.BalanceOfBatch(&_Goods.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Goods *GoodsCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Goods.Contract.BalanceOfBatch(&_Goods.CallOpts, accounts, ids)
}

// GetSuccessNum is a free data retrieval call binding the contract method 0x77aa72c0.
//
// Solidity: function getSuccessNum() view returns(uint256)
func (_Goods *GoodsCaller) GetSuccessNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goods.contract.Call(opts, &out, "getSuccessNum")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// GetSuccessNum is a free data retrieval call binding the contract method 0x77aa72c0.
//
// Solidity: function getSuccessNum() view returns(uint256)
func (_Goods *GoodsSession) GetSuccessNum() (*big.Int, error) {
	return _Goods.Contract.GetSuccessNum(&_Goods.CallOpts)
}

// GetSuccessNum is a free data retrieval call binding the contract method 0x77aa72c0.
//
// Solidity: function getSuccessNum() view returns(uint256)
func (_Goods *GoodsCallerSession) GetSuccessNum() (*big.Int, error) {
	return _Goods.Contract.GetSuccessNum(&_Goods.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Goods *GoodsCaller) IsApprovedForAll(opts *bind.CallOpts, account, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Goods.contract.Call(opts, &out, "isApprovedForAll", account, operator)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Goods *GoodsSession) IsApprovedForAll(account, operator common.Address) (bool, error) {
	return _Goods.Contract.IsApprovedForAll(&_Goods.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Goods *GoodsCallerSession) IsApprovedForAll(account, operator common.Address) (bool, error) {
	return _Goods.Contract.IsApprovedForAll(&_Goods.CallOpts, account, operator)
}

// SuccessNum is a free data retrieval call binding the contract method 0x4cd7e538.
//
// Solidity: function successNum() view returns(uint256)
func (_Goods *GoodsCaller) SuccessNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goods.contract.Call(opts, &out, "successNum")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// SuccessNum is a free data retrieval call binding the contract method 0x4cd7e538.
//
// Solidity: function successNum() view returns(uint256)
func (_Goods *GoodsSession) SuccessNum() (*big.Int, error) {
	return _Goods.Contract.SuccessNum(&_Goods.CallOpts)
}

// SuccessNum is a free data retrieval call binding the contract method 0x4cd7e538.
//
// Solidity: function successNum() view returns(uint256)
func (_Goods *GoodsCallerSession) SuccessNum() (*big.Int, error) {
	return _Goods.Contract.SuccessNum(&_Goods.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Goods *GoodsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Goods.contract.Call(opts, &out, "supportsInterface", interfaceId)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Goods *GoodsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Goods.Contract.SupportsInterface(&_Goods.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Goods *GoodsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Goods.Contract.SupportsInterface(&_Goods.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Goods *GoodsCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Goods.contract.Call(opts, &out, "uri", arg0)
	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Goods *GoodsSession) Uri(arg0 *big.Int) (string, error) {
	return _Goods.Contract.Uri(&_Goods.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Goods *GoodsCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _Goods.Contract.Uri(&_Goods.CallOpts, arg0)
}

// BatchMint is a paid mutator transaction binding the contract method 0x4684d7e9.
//
// Solidity: function batchMint(address owner, uint256[] ids) returns()
func (_Goods *GoodsTransactor) BatchMint(opts *bind.TransactOpts, owner common.Address, ids []*big.Int) (*types.Transaction, error) {
	return _Goods.contract.Transact(opts, "batchMint", owner, ids)
}

func (_Goods *GoodsTransactor) Contract() *bind.BoundContract {
	return _Goods.contract
}

// BatchMint is a paid mutator transaction binding the contract method 0x4684d7e9.
//
// Solidity: function batchMint(address owner, uint256[] ids) returns()
func (_Goods *GoodsSession) BatchMint(owner common.Address, ids []*big.Int) (*types.Transaction, error) {
	return _Goods.Contract.BatchMint(&_Goods.TransactOpts, owner, ids)
}

// BatchMint is a paid mutator transaction binding the contract method 0x4684d7e9.
//
// Solidity: function batchMint(address owner, uint256[] ids) returns()
func (_Goods *GoodsTransactorSession) BatchMint(owner common.Address, ids []*big.Int) (*types.Transaction, error) {
	return _Goods.Contract.BatchMint(&_Goods.TransactOpts, owner, ids)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x3593cebc.
//
// Solidity: function batchTransfer(address from, address to, uint256[] ids) returns()
func (_Goods *GoodsTransactor) BatchTransfer(opts *bind.TransactOpts, from, to common.Address, ids []*big.Int) (*types.Transaction, error) {
	return _Goods.contract.Transact(opts, "batchTransfer", from, to, ids)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x3593cebc.
//
// Solidity: function batchTransfer(address from, address to, uint256[] ids) returns()
func (_Goods *GoodsSession) BatchTransfer(from, to common.Address, ids []*big.Int) (*types.Transaction, error) {
	return _Goods.Contract.BatchTransfer(&_Goods.TransactOpts, from, to, ids)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x3593cebc.
//
// Solidity: function batchTransfer(address from, address to, uint256[] ids) returns()
func (_Goods *GoodsTransactorSession) BatchTransfer(from, to common.Address, ids []*big.Int) (*types.Transaction, error) {
	return _Goods.Contract.BatchTransfer(&_Goods.TransactOpts, from, to, ids)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 id) returns()
func (_Goods *GoodsTransactor) Mint(opts *bind.TransactOpts, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Goods.contract.Transact(opts, "mint", to, id)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 id) returns()
func (_Goods *GoodsSession) Mint(to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Goods.Contract.Mint(&_Goods.TransactOpts, to, id)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 id) returns()
func (_Goods *GoodsTransactorSession) Mint(to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Goods.Contract.Mint(&_Goods.TransactOpts, to, id)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Goods *GoodsTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from, to common.Address, ids, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Goods.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Goods *GoodsSession) SafeBatchTransferFrom(from, to common.Address, ids, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Goods.Contract.SafeBatchTransferFrom(&_Goods.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Goods *GoodsTransactorSession) SafeBatchTransferFrom(from, to common.Address, ids, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Goods.Contract.SafeBatchTransferFrom(&_Goods.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Goods *GoodsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from, to common.Address, id, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Goods.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Goods *GoodsSession) SafeTransferFrom(from, to common.Address, id, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Goods.Contract.SafeTransferFrom(&_Goods.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Goods *GoodsTransactorSession) SafeTransferFrom(from, to common.Address, id, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Goods.Contract.SafeTransferFrom(&_Goods.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Goods *GoodsTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Goods.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Goods *GoodsSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Goods.Contract.SetApprovalForAll(&_Goods.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Goods *GoodsTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Goods.Contract.SetApprovalForAll(&_Goods.TransactOpts, operator, approved)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 id) returns()
func (_Goods *GoodsTransactor) Transfer(opts *bind.TransactOpts, from, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Goods.contract.Transact(opts, "transfer", from, to, id)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 id) returns()
func (_Goods *GoodsSession) Transfer(from, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Goods.Contract.Transfer(&_Goods.TransactOpts, from, to, id)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 id) returns()
func (_Goods *GoodsTransactorSession) Transfer(from, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Goods.Contract.Transfer(&_Goods.TransactOpts, from, to, id)
}

// GoodsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Goods contract.
type GoodsApprovalForAllIterator struct {
	Event *GoodsApprovalForAll // Event containing the contract specifics and raw log

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
func (it *GoodsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoodsApprovalForAll)
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
		it.Event = new(GoodsApprovalForAll)
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
func (it *GoodsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoodsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoodsApprovalForAll represents a ApprovalForAll event raised by the Goods contract.
type GoodsApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Goods *GoodsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account, operator []common.Address) (*GoodsApprovalForAllIterator, error) {
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Goods.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &GoodsApprovalForAllIterator{contract: _Goods.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Goods *GoodsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *GoodsApprovalForAll, account, operator []common.Address) (event.Subscription, error) {
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Goods.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoodsApprovalForAll)
				if err := _Goods.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Goods *GoodsFilterer) ParseApprovalForAll(log types.Log) (*GoodsApprovalForAll, error) {
	event := new(GoodsApprovalForAll)
	if err := _Goods.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GoodsTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Goods contract.
type GoodsTransferBatchIterator struct {
	Event *GoodsTransferBatch // Event containing the contract specifics and raw log

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
func (it *GoodsTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoodsTransferBatch)
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
		it.Event = new(GoodsTransferBatch)
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
func (it *GoodsTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoodsTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoodsTransferBatch represents a TransferBatch event raised by the Goods contract.
type GoodsTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Goods *GoodsFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator, from, to []common.Address) (*GoodsTransferBatchIterator, error) {
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Goods.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GoodsTransferBatchIterator{contract: _Goods.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Goods *GoodsFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *GoodsTransferBatch, operator, from, to []common.Address) (event.Subscription, error) {
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Goods.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoodsTransferBatch)
				if err := _Goods.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Goods *GoodsFilterer) ParseTransferBatch(log types.Log) (*GoodsTransferBatch, error) {
	event := new(GoodsTransferBatch)
	if err := _Goods.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GoodsTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Goods contract.
type GoodsTransferSingleIterator struct {
	Event *GoodsTransferSingle // Event containing the contract specifics and raw log

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
func (it *GoodsTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoodsTransferSingle)
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
		it.Event = new(GoodsTransferSingle)
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
func (it *GoodsTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoodsTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoodsTransferSingle represents a TransferSingle event raised by the Goods contract.
type GoodsTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Goods *GoodsFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator, from, to []common.Address) (*GoodsTransferSingleIterator, error) {
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Goods.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GoodsTransferSingleIterator{contract: _Goods.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Goods *GoodsFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *GoodsTransferSingle, operator, from, to []common.Address) (event.Subscription, error) {
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Goods.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoodsTransferSingle)
				if err := _Goods.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Goods *GoodsFilterer) ParseTransferSingle(log types.Log) (*GoodsTransferSingle, error) {
	event := new(GoodsTransferSingle)
	if err := _Goods.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GoodsURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Goods contract.
type GoodsURIIterator struct {
	Event *GoodsURI // Event containing the contract specifics and raw log

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
func (it *GoodsURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoodsURI)
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
		it.Event = new(GoodsURI)
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
func (it *GoodsURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoodsURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoodsURI represents a URI event raised by the Goods contract.
type GoodsURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Goods *GoodsFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*GoodsURIIterator, error) {
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Goods.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &GoodsURIIterator{contract: _Goods.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Goods *GoodsFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *GoodsURI, id []*big.Int) (event.Subscription, error) {
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Goods.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoodsURI)
				if err := _Goods.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Goods *GoodsFilterer) ParseURI(log types.Log) (*GoodsURI, error) {
	event := new(GoodsURI)
	if err := _Goods.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
