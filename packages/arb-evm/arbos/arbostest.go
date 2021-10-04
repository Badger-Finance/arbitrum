package arbos

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var (
	setNonceABI   abi.Method
	setBalanceABI abi.Method
	setCodeABI    abi.Method
	setStorageABI abi.Method
	storeABI      abi.Method
)

func init() {
	arbostest, err := abi.JSON(strings.NewReader(arboscontracts.ArbosTestABI))
	if err != nil {
		panic(err)
	}

	setNonceABI = arbostest.Methods["setNonce"]
	setBalanceABI = arbostest.Methods["setBalance"]
	setCodeABI = arbostest.Methods["setCode"]
	setStorageABI = arbostest.Methods["setStorage"]
	storeABI = arbostest.Methods["store"]
}

func SetNonceData(address common.Address, nonce uint64) []byte {
	args, err := setNonceABI.Inputs.Pack(address, new(big.Int).SetUint64(nonce))
	if err != nil {
		panic(err)
	}
	return append(setNonceABI.ID, args...)
}

func SetBalanceData(address common.Address, balance *big.Int) []byte {
	args, err := setBalanceABI.Inputs.Pack(address, balance)
	if err != nil {
		panic(err)
	}
	return append(setBalanceABI.ID, args...)
}

func SetCodeData(address common.Address, code []byte) []byte {
	args, err := setCodeABI.Inputs.Pack(address, code)
	if err != nil {
		panic(err)
	}
	return append(setCodeABI.ID, args...)
}

func SetStorageData(address common.Address, storage map[common.Hash]common.Hash) []byte {
	args, err := setStorageABI.Inputs.Pack(address, storage)
	if err != nil {
		panic(err)
	}
	return append(setStorageABI.ID, args...)
}

func StoreData(address common.Address, key, val common.Hash) []byte {
	args, err := storeABI.Inputs.Pack(address, key, val)
	if err != nil {
		panic(err)
	}
	return append(storeABI.ID, args...)
}
