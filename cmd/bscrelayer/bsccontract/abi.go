package bsccontract

// -------------------------------------------------------
//    Contract Contains functionality for loading the
//				 smart bsccontract
// -------------------------------------------------------

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Quantiex-Hub/cmd/bscrelayer/txs"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// File paths to Peggy smart bsccontract ABIs
const (
	BridgeERC20BankABI   = "/generated/abi/BridgeERC20Bank/BridgeERC20Bank.abi"
	BridgeERC721BankABI   = "/generated/abi/BridgeERC721Bank/BridgeERC721Bank.abi"
	QuantiexERC20BridgeABI = "/generated/abi/QuantiexERC20Bridge/QuantiexERC20Bridge.abi"
	QuantiexERC721BridgeABI = "/generated/abi/QuantiexERC721Bridge/QuantiexERC721Bridge.abi"
)

// LoadABI loads a smart bsccontract as an abi.ABI
func LoadABI(contractType txs.ContractRegistry) abi.ABI {
	var (
		_, b, _, _ = runtime.Caller(0)
		dir        = filepath.Dir(b)
	)

	var filePath string
	switch contractType {
	case txs.QuantiexERC20Bridge:
		filePath = QuantiexERC20BridgeABI
	case txs.QuantiexERC721Bridge:
		filePath = QuantiexERC721BridgeABI
	case txs.BridgeERC20Bank:
		filePath = BridgeERC20BankABI
	case txs.BridgeERC721Bank:
		filePath = BridgeERC721BankABI
	}

	// Read the file containing the bsccontract's ABI
	contractRaw, err := ioutil.ReadFile(dir + filePath)
	if err != nil {
		panic(err)
	}

	// Convert the raw abi into a usable format
	contractABI, err := abi.JSON(strings.NewReader(string(contractRaw)))
	if err != nil {
		panic(err)
	}
	return contractABI
}
