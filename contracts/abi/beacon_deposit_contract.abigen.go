// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// BeaconDepositContractMetaData contains all meta data concerning the BeaconDepositContract contract.
var BeaconDepositContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"NATIVE_ASSET\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"credentials\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"redirect\",\"inputs\":[{\"name\":\"fromPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"withdrawalCredentials\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"credentials\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"signature\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"index\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Redirect\",\"inputs\":[{\"name\":\"fromPubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"toPubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"credentials\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"index\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"fromPubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"credentials\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"withdrawalCredentials\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"index\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DepositNotMultipleOfGwei\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositValueTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientRedirectAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientWithdrawalAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCredentialsLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPubKeyLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignatureLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]}]",
	Bin: "0x60806040525f80546001600160a01b03191673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee179055348015610034575f80fd5b50610c50806100425f395ff3fe60806040526004361061003e575f3560e01c806304a3267f146100425780635b70fa2914610063578063bf53253b14610076578063bf9b6a55146100c6575b5f80fd5b34801561004d575f80fd5b5061006161005c366004610866565b6100e5565b005b6100616100713660046108e1565b6102b0565b348015610081575f80fd5b5061009d73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee81565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100d1575f80fd5b506100616100e0366004610866565b61047c565b68929eee149b4bd212685c156101025763ab143c065f526004601cfd5b3068929eee149b4bd212685d60308414610148576040517f9f10647200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208214610182576040517fb39bca1600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610192600a6407735940006109e1565b67ffffffffffffffff168167ffffffffffffffff1610156101df576040517f1203093700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160208082527f01000000000000000000000000000000000000000000000000000000000000003317908201528082019091527fe8d1abf8fc2a88152f5f2e1e5e3634f882fc7db420a88ca0b64eb78400f5851a90869086906001805488918891889167ffffffffffffffff909116905f61025c83610a07565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040516102959796959493929190610ad5565b60405180910390a15f68929eee149b4bd212685d5050505050565b68929eee149b4bd212685c156102cd5763ab143c065f526004601cfd5b3068929eee149b4bd212685d60308614610313576040517f9f10647200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020841461034d576040517fb39bca1600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608114610387576040517f4be6321b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5473ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff1111111111111111111111111111111111111112016103d4576103cd610605565b92506103dd565b6103dd836106ec565b7f68af751683498a9f9be59fe8b0d52a64dd155255d85cdb29fea30b1e3f891d46878787878787875f601481819054906101000a900467ffffffffffffffff1661042690610a07565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905560405161045f989796959493929190610b38565b60405180910390a15f68929eee149b4bd212685d50505050505050565b68929eee149b4bd212685c156104995763ab143c065f526004601cfd5b3068929eee149b4bd212685d6030841415806104b6575060308214155b156104ed576040517f9f10647200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104fd600a6407735940006109e1565b67ffffffffffffffff168167ffffffffffffffff16101561054a576040517f0494a69c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160208082527f01000000000000000000000000000000000000000000000000000000000000003317908201528082019091527f488862f49032f13a357c663883290b0e784d0f814979e39f33f5d40ce70493f9908690869086908690866001600881819054906101000a900467ffffffffffffffff166105cd90610a07565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790556040516102959796959493929190610b9d565b5f610614633b9aca0034610bd7565b1561064b576040517f40567b3800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61065a633b9aca0034610bea565b905067ffffffffffffffff81111561069e576040517f2aa6673400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6407735940008110156106dd576040517f0e1eddda00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106e75f346107ed565b919050565b5f5473ffffffffffffffffffffffffffffffffffffffff16639dc29fac3361072267ffffffffffffffff8516633b9aca00610bfd565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044015f604051808303815f87803b15801561078a575f80fd5b505af115801561079c573d5f803e3d5ffd5b50505064077359400067ffffffffffffffff8316101590506107ea576040517f0e1eddda00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b5f385f3884865af16108065763b12d13eb5f526004601cfd5b5050565b5f8083601f84011261081a575f80fd5b50813567ffffffffffffffff811115610831575f80fd5b602083019150836020828501011115610848575f80fd5b9250929050565b803567ffffffffffffffff811681146106e7575f80fd5b5f805f805f6060868803121561087a575f80fd5b853567ffffffffffffffff80821115610891575f80fd5b61089d89838a0161080a565b909750955060208801359150808211156108b5575f80fd5b506108c28882890161080a565b90945092506108d590506040870161084f565b90509295509295909350565b5f805f805f805f6080888a0312156108f7575f80fd5b873567ffffffffffffffff8082111561090e575f80fd5b61091a8b838c0161080a565b909950975060208a0135915080821115610932575f80fd5b61093e8b838c0161080a565b909750955085915061095260408b0161084f565b945060608a0135915080821115610967575f80fd5b506109748a828b0161080a565b989b979a50959850939692959293505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f67ffffffffffffffff808416806109fb576109fb610987565b92169190910492915050565b5f67ffffffffffffffff808316818103610a2357610a236109b4565b6001019392505050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b5f81518084525f5b81811015610a9857602081850181015186830182015201610a7c565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60a081525f610ae860a08301898b610a2d565b8281036020840152610afa8189610a74565b90508281036040840152610b0f818789610a2d565b91505067ffffffffffffffff808516606084015280841660808401525098975050505050505050565b60a081525f610b4b60a083018a8c610a2d565b8281036020840152610b5e81898b610a2d565b905067ffffffffffffffff80881660408501528382036060850152610b84828789610a2d565b9250808516608085015250509998505050505050505050565b60a081525f610bb060a08301898b610a2d565b8281036020840152610bc381888a610a2d565b90508281036040840152610b0f8187610a74565b5f82610be557610be5610987565b500690565b5f82610bf857610bf8610987565b500490565b8082028115828204841417610c1457610c146109b4565b9291505056fea26469706673582212201ec37cbc5af93eedfe109978e9b98288adfefc690e7d3ea560c11faec993503764736f6c63430008180033",
}

// BeaconDepositContractABI is the input ABI used to generate the binding from.
// Deprecated: Use BeaconDepositContractMetaData.ABI instead.
var BeaconDepositContractABI = BeaconDepositContractMetaData.ABI

// BeaconDepositContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BeaconDepositContractMetaData.Bin instead.
var BeaconDepositContractBin = BeaconDepositContractMetaData.Bin

// DeployBeaconDepositContract deploys a new Ethereum contract, binding an instance of BeaconDepositContract to it.
func DeployBeaconDepositContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BeaconDepositContract, error) {
	parsed, err := BeaconDepositContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BeaconDepositContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BeaconDepositContract{BeaconDepositContractCaller: BeaconDepositContractCaller{contract: contract}, BeaconDepositContractTransactor: BeaconDepositContractTransactor{contract: contract}, BeaconDepositContractFilterer: BeaconDepositContractFilterer{contract: contract}}, nil
}

// BeaconDepositContract is an auto generated Go binding around an Ethereum contract.
type BeaconDepositContract struct {
	BeaconDepositContractCaller     // Read-only binding to the contract
	BeaconDepositContractTransactor // Write-only binding to the contract
	BeaconDepositContractFilterer   // Log filterer for contract events
}

// BeaconDepositContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeaconDepositContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconDepositContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeaconDepositContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconDepositContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeaconDepositContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconDepositContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeaconDepositContractSession struct {
	Contract     *BeaconDepositContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BeaconDepositContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeaconDepositContractCallerSession struct {
	Contract *BeaconDepositContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// BeaconDepositContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeaconDepositContractTransactorSession struct {
	Contract     *BeaconDepositContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// BeaconDepositContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeaconDepositContractRaw struct {
	Contract *BeaconDepositContract // Generic contract binding to access the raw methods on
}

// BeaconDepositContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeaconDepositContractCallerRaw struct {
	Contract *BeaconDepositContractCaller // Generic read-only contract binding to access the raw methods on
}

// BeaconDepositContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeaconDepositContractTransactorRaw struct {
	Contract *BeaconDepositContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeaconDepositContract creates a new instance of BeaconDepositContract, bound to a specific deployed contract.
func NewBeaconDepositContract(address common.Address, backend bind.ContractBackend) (*BeaconDepositContract, error) {
	contract, err := bindBeaconDepositContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeaconDepositContract{BeaconDepositContractCaller: BeaconDepositContractCaller{contract: contract}, BeaconDepositContractTransactor: BeaconDepositContractTransactor{contract: contract}, BeaconDepositContractFilterer: BeaconDepositContractFilterer{contract: contract}}, nil
}

// NewBeaconDepositContractCaller creates a new read-only instance of BeaconDepositContract, bound to a specific deployed contract.
func NewBeaconDepositContractCaller(address common.Address, caller bind.ContractCaller) (*BeaconDepositContractCaller, error) {
	contract, err := bindBeaconDepositContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconDepositContractCaller{contract: contract}, nil
}

// NewBeaconDepositContractTransactor creates a new write-only instance of BeaconDepositContract, bound to a specific deployed contract.
func NewBeaconDepositContractTransactor(address common.Address, transactor bind.ContractTransactor) (*BeaconDepositContractTransactor, error) {
	contract, err := bindBeaconDepositContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconDepositContractTransactor{contract: contract}, nil
}

// NewBeaconDepositContractFilterer creates a new log filterer instance of BeaconDepositContract, bound to a specific deployed contract.
func NewBeaconDepositContractFilterer(address common.Address, filterer bind.ContractFilterer) (*BeaconDepositContractFilterer, error) {
	contract, err := bindBeaconDepositContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeaconDepositContractFilterer{contract: contract}, nil
}

// bindBeaconDepositContract binds a generic wrapper to an already deployed contract.
func bindBeaconDepositContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeaconDepositContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeaconDepositContract *BeaconDepositContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconDepositContract.Contract.BeaconDepositContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeaconDepositContract *BeaconDepositContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeaconDepositContract.Contract.BeaconDepositContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeaconDepositContract *BeaconDepositContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeaconDepositContract.Contract.BeaconDepositContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeaconDepositContract *BeaconDepositContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconDepositContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeaconDepositContract *BeaconDepositContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeaconDepositContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeaconDepositContract *BeaconDepositContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeaconDepositContract.Contract.contract.Transact(opts, method, params...)
}

// NATIVEASSET is a free data retrieval call binding the contract method 0xbf53253b.
//
// Solidity: function NATIVE_ASSET() view returns(address)
func (_BeaconDepositContract *BeaconDepositContractCaller) NATIVEASSET(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeaconDepositContract.contract.Call(opts, &out, "NATIVE_ASSET")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEASSET is a free data retrieval call binding the contract method 0xbf53253b.
//
// Solidity: function NATIVE_ASSET() view returns(address)
func (_BeaconDepositContract *BeaconDepositContractSession) NATIVEASSET() (common.Address, error) {
	return _BeaconDepositContract.Contract.NATIVEASSET(&_BeaconDepositContract.CallOpts)
}

// NATIVEASSET is a free data retrieval call binding the contract method 0xbf53253b.
//
// Solidity: function NATIVE_ASSET() view returns(address)
func (_BeaconDepositContract *BeaconDepositContractCallerSession) NATIVEASSET() (common.Address, error) {
	return _BeaconDepositContract.Contract.NATIVEASSET(&_BeaconDepositContract.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x5b70fa29.
//
// Solidity: function deposit(bytes pubkey, bytes credentials, uint64 amount, bytes signature) payable returns()
func (_BeaconDepositContract *BeaconDepositContractTransactor) Deposit(opts *bind.TransactOpts, pubkey []byte, credentials []byte, amount uint64, signature []byte) (*types.Transaction, error) {
	return _BeaconDepositContract.contract.Transact(opts, "deposit", pubkey, credentials, amount, signature)
}

// Deposit is a paid mutator transaction binding the contract method 0x5b70fa29.
//
// Solidity: function deposit(bytes pubkey, bytes credentials, uint64 amount, bytes signature) payable returns()
func (_BeaconDepositContract *BeaconDepositContractSession) Deposit(pubkey []byte, credentials []byte, amount uint64, signature []byte) (*types.Transaction, error) {
	return _BeaconDepositContract.Contract.Deposit(&_BeaconDepositContract.TransactOpts, pubkey, credentials, amount, signature)
}

// Deposit is a paid mutator transaction binding the contract method 0x5b70fa29.
//
// Solidity: function deposit(bytes pubkey, bytes credentials, uint64 amount, bytes signature) payable returns()
func (_BeaconDepositContract *BeaconDepositContractTransactorSession) Deposit(pubkey []byte, credentials []byte, amount uint64, signature []byte) (*types.Transaction, error) {
	return _BeaconDepositContract.Contract.Deposit(&_BeaconDepositContract.TransactOpts, pubkey, credentials, amount, signature)
}

// Redirect is a paid mutator transaction binding the contract method 0xbf9b6a55.
//
// Solidity: function redirect(bytes fromPubkey, bytes toPubkey, uint64 amount) returns()
func (_BeaconDepositContract *BeaconDepositContractTransactor) Redirect(opts *bind.TransactOpts, fromPubkey []byte, toPubkey []byte, amount uint64) (*types.Transaction, error) {
	return _BeaconDepositContract.contract.Transact(opts, "redirect", fromPubkey, toPubkey, amount)
}

// Redirect is a paid mutator transaction binding the contract method 0xbf9b6a55.
//
// Solidity: function redirect(bytes fromPubkey, bytes toPubkey, uint64 amount) returns()
func (_BeaconDepositContract *BeaconDepositContractSession) Redirect(fromPubkey []byte, toPubkey []byte, amount uint64) (*types.Transaction, error) {
	return _BeaconDepositContract.Contract.Redirect(&_BeaconDepositContract.TransactOpts, fromPubkey, toPubkey, amount)
}

// Redirect is a paid mutator transaction binding the contract method 0xbf9b6a55.
//
// Solidity: function redirect(bytes fromPubkey, bytes toPubkey, uint64 amount) returns()
func (_BeaconDepositContract *BeaconDepositContractTransactorSession) Redirect(fromPubkey []byte, toPubkey []byte, amount uint64) (*types.Transaction, error) {
	return _BeaconDepositContract.Contract.Redirect(&_BeaconDepositContract.TransactOpts, fromPubkey, toPubkey, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x04a3267f.
//
// Solidity: function withdraw(bytes pubkey, bytes withdrawalCredentials, uint64 amount) returns()
func (_BeaconDepositContract *BeaconDepositContractTransactor) Withdraw(opts *bind.TransactOpts, pubkey []byte, withdrawalCredentials []byte, amount uint64) (*types.Transaction, error) {
	return _BeaconDepositContract.contract.Transact(opts, "withdraw", pubkey, withdrawalCredentials, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x04a3267f.
//
// Solidity: function withdraw(bytes pubkey, bytes withdrawalCredentials, uint64 amount) returns()
func (_BeaconDepositContract *BeaconDepositContractSession) Withdraw(pubkey []byte, withdrawalCredentials []byte, amount uint64) (*types.Transaction, error) {
	return _BeaconDepositContract.Contract.Withdraw(&_BeaconDepositContract.TransactOpts, pubkey, withdrawalCredentials, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x04a3267f.
//
// Solidity: function withdraw(bytes pubkey, bytes withdrawalCredentials, uint64 amount) returns()
func (_BeaconDepositContract *BeaconDepositContractTransactorSession) Withdraw(pubkey []byte, withdrawalCredentials []byte, amount uint64) (*types.Transaction, error) {
	return _BeaconDepositContract.Contract.Withdraw(&_BeaconDepositContract.TransactOpts, pubkey, withdrawalCredentials, amount)
}

// BeaconDepositContractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BeaconDepositContract contract.
type BeaconDepositContractDepositIterator struct {
	Event *BeaconDepositContractDeposit // Event containing the contract specifics and raw log

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
func (it *BeaconDepositContractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconDepositContractDeposit)
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
		it.Event = new(BeaconDepositContractDeposit)
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
func (it *BeaconDepositContractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconDepositContractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconDepositContractDeposit represents a Deposit event raised by the BeaconDepositContract contract.
type BeaconDepositContractDeposit struct {
	Pubkey      []byte
	Credentials []byte
	Amount      uint64
	Signature   []byte
	Index       uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x68af751683498a9f9be59fe8b0d52a64dd155255d85cdb29fea30b1e3f891d46.
//
// Solidity: event Deposit(bytes pubkey, bytes credentials, uint64 amount, bytes signature, uint64 index)
func (_BeaconDepositContract *BeaconDepositContractFilterer) FilterDeposit(opts *bind.FilterOpts) (*BeaconDepositContractDepositIterator, error) {

	logs, sub, err := _BeaconDepositContract.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &BeaconDepositContractDepositIterator{contract: _BeaconDepositContract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x68af751683498a9f9be59fe8b0d52a64dd155255d85cdb29fea30b1e3f891d46.
//
// Solidity: event Deposit(bytes pubkey, bytes credentials, uint64 amount, bytes signature, uint64 index)
func (_BeaconDepositContract *BeaconDepositContractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BeaconDepositContractDeposit) (event.Subscription, error) {

	logs, sub, err := _BeaconDepositContract.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconDepositContractDeposit)
				if err := _BeaconDepositContract.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x68af751683498a9f9be59fe8b0d52a64dd155255d85cdb29fea30b1e3f891d46.
//
// Solidity: event Deposit(bytes pubkey, bytes credentials, uint64 amount, bytes signature, uint64 index)
func (_BeaconDepositContract *BeaconDepositContractFilterer) ParseDeposit(log types.Log) (*BeaconDepositContractDeposit, error) {
	event := new(BeaconDepositContractDeposit)
	if err := _BeaconDepositContract.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeaconDepositContractRedirectIterator is returned from FilterRedirect and is used to iterate over the raw logs and unpacked data for Redirect events raised by the BeaconDepositContract contract.
type BeaconDepositContractRedirectIterator struct {
	Event *BeaconDepositContractRedirect // Event containing the contract specifics and raw log

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
func (it *BeaconDepositContractRedirectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconDepositContractRedirect)
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
		it.Event = new(BeaconDepositContractRedirect)
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
func (it *BeaconDepositContractRedirectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconDepositContractRedirectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconDepositContractRedirect represents a Redirect event raised by the BeaconDepositContract contract.
type BeaconDepositContractRedirect struct {
	FromPubkey  []byte
	ToPubkey    []byte
	Credentials []byte
	Amount      uint64
	Index       uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRedirect is a free log retrieval operation binding the contract event 0x488862f49032f13a357c663883290b0e784d0f814979e39f33f5d40ce70493f9.
//
// Solidity: event Redirect(bytes fromPubkey, bytes toPubkey, bytes credentials, uint64 amount, uint64 index)
func (_BeaconDepositContract *BeaconDepositContractFilterer) FilterRedirect(opts *bind.FilterOpts) (*BeaconDepositContractRedirectIterator, error) {

	logs, sub, err := _BeaconDepositContract.contract.FilterLogs(opts, "Redirect")
	if err != nil {
		return nil, err
	}
	return &BeaconDepositContractRedirectIterator{contract: _BeaconDepositContract.contract, event: "Redirect", logs: logs, sub: sub}, nil
}

// WatchRedirect is a free log subscription operation binding the contract event 0x488862f49032f13a357c663883290b0e784d0f814979e39f33f5d40ce70493f9.
//
// Solidity: event Redirect(bytes fromPubkey, bytes toPubkey, bytes credentials, uint64 amount, uint64 index)
func (_BeaconDepositContract *BeaconDepositContractFilterer) WatchRedirect(opts *bind.WatchOpts, sink chan<- *BeaconDepositContractRedirect) (event.Subscription, error) {

	logs, sub, err := _BeaconDepositContract.contract.WatchLogs(opts, "Redirect")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconDepositContractRedirect)
				if err := _BeaconDepositContract.contract.UnpackLog(event, "Redirect", log); err != nil {
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

// ParseRedirect is a log parse operation binding the contract event 0x488862f49032f13a357c663883290b0e784d0f814979e39f33f5d40ce70493f9.
//
// Solidity: event Redirect(bytes fromPubkey, bytes toPubkey, bytes credentials, uint64 amount, uint64 index)
func (_BeaconDepositContract *BeaconDepositContractFilterer) ParseRedirect(log types.Log) (*BeaconDepositContractRedirect, error) {
	event := new(BeaconDepositContractRedirect)
	if err := _BeaconDepositContract.contract.UnpackLog(event, "Redirect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeaconDepositContractWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the BeaconDepositContract contract.
type BeaconDepositContractWithdrawalIterator struct {
	Event *BeaconDepositContractWithdrawal // Event containing the contract specifics and raw log

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
func (it *BeaconDepositContractWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconDepositContractWithdrawal)
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
		it.Event = new(BeaconDepositContractWithdrawal)
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
func (it *BeaconDepositContractWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconDepositContractWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconDepositContractWithdrawal represents a Withdrawal event raised by the BeaconDepositContract contract.
type BeaconDepositContractWithdrawal struct {
	FromPubkey            []byte
	Credentials           []byte
	WithdrawalCredentials []byte
	Amount                uint64
	Index                 uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xe8d1abf8fc2a88152f5f2e1e5e3634f882fc7db420a88ca0b64eb78400f5851a.
//
// Solidity: event Withdrawal(bytes fromPubkey, bytes credentials, bytes withdrawalCredentials, uint64 amount, uint64 index)
func (_BeaconDepositContract *BeaconDepositContractFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*BeaconDepositContractWithdrawalIterator, error) {

	logs, sub, err := _BeaconDepositContract.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &BeaconDepositContractWithdrawalIterator{contract: _BeaconDepositContract.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xe8d1abf8fc2a88152f5f2e1e5e3634f882fc7db420a88ca0b64eb78400f5851a.
//
// Solidity: event Withdrawal(bytes fromPubkey, bytes credentials, bytes withdrawalCredentials, uint64 amount, uint64 index)
func (_BeaconDepositContract *BeaconDepositContractFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *BeaconDepositContractWithdrawal) (event.Subscription, error) {

	logs, sub, err := _BeaconDepositContract.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconDepositContractWithdrawal)
				if err := _BeaconDepositContract.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xe8d1abf8fc2a88152f5f2e1e5e3634f882fc7db420a88ca0b64eb78400f5851a.
//
// Solidity: event Withdrawal(bytes fromPubkey, bytes credentials, bytes withdrawalCredentials, uint64 amount, uint64 index)
func (_BeaconDepositContract *BeaconDepositContractFilterer) ParseWithdrawal(log types.Log) (*BeaconDepositContractWithdrawal, error) {
	event := new(BeaconDepositContractWithdrawal)
	if err := _BeaconDepositContract.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
