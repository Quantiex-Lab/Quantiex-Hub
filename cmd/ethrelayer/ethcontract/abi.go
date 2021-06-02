package ethcontract

// -------------------------------------------------------
//    Contract Contains functionality for loading the
//				 smart ethcontract
// -------------------------------------------------------

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Quantiex-Hub/cmd/ethrelayer/txs"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// File paths to Peggy smart ethcontract ABIs
const (
	BridgeBankABI   = "/generated/abi/BridgeBank/BridgeBank.abi"
	QuantiexBridgeABI = "/generated/abi/QuantiexBridge/QuantiexBridge.abi"
)

// LoadABI loads a smart ethcontract as an abi.ABI
func LoadABI(contractType txs.ContractRegistry) abi.ABI {
	var (
		_, b, _, _ = runtime.Caller(0)
		dir        = filepath.Dir(b)
	)

	var filePath string
	switch contractType {
	case txs.QuantiexBridge:
		filePath = QuantiexBridgeABI
	case txs.BridgeBank:
		filePath = BridgeBankABI
	}

	// Read the file containing the ethcontract's ABI
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
