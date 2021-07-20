package ethcontract

// BridgeContract is an enum containing supported ethcontract names
type BridgeContract int

const (
	// BridgeRegistry registers deployed addresses of the other contracts
	BridgeRegistry BridgeContract = iota + 1
	// Valset manages the validator set and signature verification
	Valset
	// Oracle enables validators to make OracleClaims and processes ProphecyClaims
	Oracle
	// QuantiexERC20Bridge enables validators to make ProphecyClaims
	QuantiexERC20Bridge
	// QuantiexERC721Bridge enables validators to make ProphecyClaims
	QuantiexERC721Bridge
	// BridgeERC20Bank manages protocol assets on both Ethereum and Binance
	BridgeERC20Bank
	// BridgeERC721Bank manages protocol assets on both Ethereum and Binance
	BridgeERC721Bank
	// StakingPool manages staking assets
	StakingPool
)

// BridgeContractToString returns the string associated with a BridgeContract
var BridgeContractToString = [...]string{"BridgeRegistry", "Valset", "Oracle", "QuantiexERC20Bridge",
	"QuantiexERC721Bridge", "BridgeERC20Bank", "BridgeERC721Bank", "StakingPool"}

// String returns the BridgeContract as a string
func (d BridgeContract) String() string {
	return BridgeContractToString[d-1]
}

// BridgeContracts is a slice of BridgeContract
type BridgeContracts []BridgeContract

// LoadBridgeContracts loads populated slice of BridgeContract
func LoadBridgeContracts() BridgeContracts {
	return BridgeContracts{
		BridgeRegistry,
		Valset,
		Oracle,
		QuantiexERC20Bridge,
		QuantiexERC721Bridge,
		BridgeERC20Bank,
		BridgeERC721Bank,
		StakingPool,
	}
}
