package types

import (
	"fmt"
	xcommon "github.com/Quantiex-Hub/x/common/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Event enum containing supported chain events
type Event byte

const (
	// Unsupported is an invalid Binance or Ethereum event
	Unsupported Event = iota
	// MsgBurn is a Binance msg of type MsgBurn
	MsgBurn
	// MsgLock is a Binance msg of type MsgLock
	MsgLock
	// LogLock is for Ethereum event LogLock
	LogLock
	// LogBurn is for Ethereum event LogBurn
	LogBurn
	// LogNewProphecyClaim is an Ethereum event named 'LogNewProphecyClaim'
	LogNewProphecyClaim
)

// String returns the event type as a string
func (d Event) String() string {
	return [...]string{"unsupported", "burn", "lock", "LogLock",
		"LogBurn", "LogNewProphecyClaim"}[d]
}

// EthereumERC20Event struct is used by LogLock and LogBurn
type EthereumERC20Event struct {
	ChainID               *big.Int
	BankAddress           common.Address
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
func (e EthereumERC20Event) String() string {
	return fmt.Sprintf("\nChainID: %v\nERC20BankAddress: %v\nChainName: %v\nSymbol: %v\nFrom %v"+
		"To: %v\nToken: %v\nValue: %v\nNonce: %v\nClaim type: %v",
		e.ChainID,
		e.BankAddress.Hex(),
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
	ChainName        string
	ClaimType        uint8
	BinanceSender    common.Address
	EthereumReceiver common.Address
	ValidatorAddress common.Address
	TokenAddress     common.Address
	Symbol           string
	Amount           *big.Int
	TxHash           string
}

// String implements fmt.Stringer
func (p ProphecyClaimERC20Event) String() string {
	return fmt.Sprintf("\nProphecyID: %v\nChainName: %v\nClaimType: %v\nBinanceSender: %v\n" +
		"EthereumReceiver: %v\nValidatorAddress: %v\nSymbol: %v\nTokenAddress: %v\nAmount: %v\nTxHash: %v\n\n",
		p.ProphecyID,
		p.ChainName,
		p.ClaimType,
		p.BinanceSender.Hex(),
		p.EthereumReceiver.Hex(),
		p.ValidatorAddress.Hex(),
		p.TokenAddress.Hex(),
		p.Symbol,
		p.Amount,
		p.TxHash)
}




// EthereumERC721Event struct is used by LogLock and LogBurn
type EthereumERC721Event struct {
	ChainID			*big.Int
	BankAddress	    common.Address
	ClaimType		xcommon.ClaimType
	ChainName		string
	From			common.Address
	To				common.Address
	Token			common.Address
	Symbol			string
	TokenId			*big.Int
	BaseURI			string
	TokenURI		string
	Nonce			*big.Int
}

// String implements fmt.Stringer
func (e EthereumERC721Event) String() string {
	return fmt.Sprintf("\nChainID: %v\nERC721BankAddress: %v\nClaimType: %v\nChainName: %v\n" +
		"Sender: %v\nRecipient: %v\nToken: %v\nSymbol: %v\nTokenId: %v\nBaseURI: %v\nTokenURI: %v\nNonce: %v\n\n",
		e.ChainID,
		e.BankAddress.Hex(),
		e.ClaimType.String(),
		e.ChainName,
		e.From.Hex(),
		e.To.Hex(),
		e.Token.Hex(),
		e.Symbol,
		e.TokenId,
		e.BaseURI,
		e.TokenURI,
		e.Nonce,
	)
}

// ProphecyClaimERC721Event struct which represents a LogNewProphecyClaim event
type ProphecyClaimERC721Event struct {
	ProphecyID       *big.Int
	ChainName        string
	ClaimType        uint8
	BinanceSender    common.Address
	EthereumReceiver common.Address
	ValidatorAddress common.Address
	TokenAddress     common.Address
	Symbol           string
	TokenId          *big.Int
	TokenURI         string
	TxHash           string
}

// String implements fmt.Stringer
func (p ProphecyClaimERC721Event) String() string {
	return fmt.Sprintf("ProphecyID: %v\nChainName: %v\nClaimType: %v\nBinanceSender: %v\nEthereumReceiver: %v\n" +
		"ValidatorAddress: %v\nTokenAddress: %v\nSymbol: %v\nTokenId: %v\nTokenURI: %v\nTxHash: %v\n\n",
		p.ProphecyID,
		p.ChainName,
		p.ClaimType,
		p.BinanceSender.Hex(),
		p.EthereumReceiver.Hex(),
		p.ValidatorAddress.Hex(),
		p.TokenAddress.Hex(),
		p.Symbol,
		p.TokenId,
		p.TokenURI,
		p.TxHash)
}
