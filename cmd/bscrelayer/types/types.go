package types

import (
	"fmt"
	"math/big"

	xcommon "github.com/Quantiex-Hub/x/common/types"
	"github.com/ethereum/go-ethereum/common"
)

// Event enum containing supported chain events
type Event byte

const (
	// Unsupported is an invalid Binance or Ethereum event
	Unsupported Event = iota
	// MsgBurn is a Ethereum msg of type MsgBurn
	MsgBurn
	// MsgLock is a Ethereum msg of type MsgLock
	MsgLock
	// LogLock is for Binance event LogLock
	LogLock
	// LogBurn is for Binance event LogBurn
	LogBurn
	// LogNewProphecyClaim is an Ethereum event named 'LogNewProphecyClaim'
	LogNewProphecyClaim
)

// String returns the event type as a string
func (d Event) String() string {
	return [...]string{"unsupported", "burn", "lock", "LogLock", "LogBurn", "LogNewProphecyClaim"}[d]
}

// BinanceEvent struct is used by LogLock and LogBurn
type BinanceEvent struct {
	BinanceChainID         *big.Int
	BridgeContractAddress common.Address
	ID                    [32]byte
	ChainName             string
	From                  common.Address
	To                    common.Address
	Token                 common.Address
	Symbol                string
	Value                 *big.Int
	Nonce                 *big.Int
	ClaimType             xcommon.ClaimType
}

// String implements fmt.Stringer
func (e BinanceEvent) String() string {
	return fmt.Sprintf("\nChain ID: %v\nChain Name: %v\nBridge bsccontract address: %v\nToken symbol: %v\nToken "+
		"bsccontract address: %v\nSender: %v\nRecipient: %v\nValue: %v\nNonce: %v\nClaim type: %v",
		e.BinanceChainID, e.ChainName, e.BridgeContractAddress.Hex(), e.Symbol, e.Token.Hex(), e.From.Hex(),
		e.To.Hex(), e.Value, e.Nonce, e.ClaimType.String())
}

// ProphecyClaimEvent struct which represents a LogNewProphecyClaim event
type ProphecyClaimEvent struct {
	EthereumSender   common.Address
	Symbol           string
	ProphecyID       *big.Int
	Amount           *big.Int
	BinanceReceiver   common.Address
	ValidatorAddress common.Address
	TokenAddress     common.Address
	ClaimType        uint8
	TxHash           string
}

// String implements fmt.Stringer
func (p ProphecyClaimEvent) String() string {
	return fmt.Sprintf("\nProphecy ID: %v\nClaim Type: %v\nSender: %v\n"+
		"Recipient: %v\nSymbol: %v\nToken: %v\nAmount: %v\nValidator: %v\nTxHash: %v\n\n",
		p.ProphecyID, p.ClaimType, p.EthereumSender.Hex(), p.BinanceReceiver.Hex(),
		p.Symbol, p.TokenAddress.Hex(), p.Amount, p.ValidatorAddress.Hex(), p.TxHash)
}
