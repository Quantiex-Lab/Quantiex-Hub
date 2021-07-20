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

// BinanceERC20Event struct is used by LogLock and LogBurn
type BinanceERC20Event struct {
	ChainID				*big.Int
	BridgeBankAddress	common.Address
	ID					[32]byte
	ChainName			string
	From				common.Address
	To					common.Address
	Token				common.Address
	Symbol				string
	Value				*big.Int
	Nonce				*big.Int
	ClaimType			xcommon.ClaimType
}

// String implements fmt.Stringer
func (e BinanceERC20Event) String() string {
	return fmt.Sprintf("\nChainID: %v\nBridgeBankAddress: %v\nChainName: %v\nSymbol: %v\nFrom %v"+
		"To: %v\nToken: %v\nValue: %v\nNonce: %v\nClaim type: %v",
		e.ChainID,
		e.BridgeBankAddress.Hex(),
		e.ChainName,
		e.Symbol,
		e.From.Hex(),
		e.To.Hex(),
		e.Token.Hex(),
		e.Value,
		e.Nonce,
		e.ClaimType.String())
}

// ProphecyClaimERC20Event struct which represents a LogNewProphecyClaim event
type ProphecyClaimERC20Event struct {
	ProphecyID       *big.Int
	ClaimType        uint8
	EthereumSender   common.Address
	BinanceReceiver  common.Address
	ValidatorAddress common.Address
	TokenAddress     common.Address
	Symbol           string
	Amount           *big.Int
	TxHash           string
}

// String implements fmt.Stringer
func (p ProphecyClaimERC20Event) String() string {
	return fmt.Sprintf("\nProphecyID: %v\nClaimType: %v\nEthereumSender: %v\n"+
		"BinanceReceiver: %v\nValidatorAddress: %v\nTokenAddress: %v\nSymbol: %v\nAmount: %v\nTxHash: %v\n\n",
		p.ProphecyID,
		p.ClaimType,
		p.EthereumSender.Hex(),
		p.BinanceReceiver.Hex(),
		p.ValidatorAddress.Hex(),
		p.TokenAddress.Hex(),
		p.Symbol,
		p.Amount,
		p.TxHash)
}



// BinanceERC721Event struct is used by LogLock and LogBurn
type BinanceERC721Event struct {
	ChainID				*big.Int
	BridgeBankAddress	common.Address
	ID					[32]byte
	ChainName			string
	From				common.Address
	To					common.Address
	Token				common.Address
	Symbol				string
	TokenId				*big.Int
	BaseURI				string
	TokenURI			string
	Nonce				*big.Int
	ClaimType			xcommon.ClaimType
}

// String implements fmt.Stringer
func (e BinanceERC721Event) String() string {
	return fmt.Sprintf("\nChainID: %v\nBridgeBankAddress: %v\nChainName: %v\nFrom: %v\nTo %v"+
		"Token: %v\nSymbol: %v\nTokenId: %v\nBaseURI: %v\nTokenURI: %v\nNonce %v\nClaimType: %v",
		e.ChainID,
		e.BridgeBankAddress.Hex(),
		e.ChainName,
		e.From.Hex(),
		e.To.Hex(),
		e.Token.Hex(),
		e.Symbol,
		e.TokenId,
		e.BaseURI,
		e.TokenURI,
		e.Nonce,
		e.ClaimType.String())
}

// ProphecyClaimERC721Event struct which represents a LogNewProphecyClaim event
type ProphecyClaimERC721Event struct {
	ProphecyID			*big.Int
	ClaimType			uint8
	EthereumSender		common.Address
	BinanceReceiver		common.Address
	ValidatorAddress	common.Address
	TokenAddress		common.Address
	Symbol				string
	TokenId				*big.Int
	TokenURI			string
	TxHash				string
}

// String implements fmt.Stringer
func (p ProphecyClaimERC721Event) String() string {
	return fmt.Sprintf("\nProphecyID: %v\nClaimType: %v\nEthereumSender: %v\n"+
		"BinanceReceiver: %v\nValidatorAddress: %v\nTokenAddress: %v\nSymbol: %v\nTokenId: %v\nTokenURI: %v\nTxHash: %v\n\n",
		p.ProphecyID,
		p.ClaimType,
		p.EthereumSender.Hex(),
		p.BinanceReceiver.Hex(),
		p.ValidatorAddress.Hex(),
		p.TokenAddress.Hex(),
		p.Symbol,
		p.TokenId,
		p.TokenURI,
		p.TxHash)
}
