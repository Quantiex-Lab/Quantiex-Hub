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
	// QuantiexBridge enables validators to make ProphecyClaims
	QuantiexBridge
	// BridgeBank manages protocol assets on both Ethereum and Binance
	BridgeBank
	// StakingPool manages staking assets
	StakingPool
)

// BridgeContractToString returns the string associated with a BridgeContract
var BridgeContractToString = [...]string{"BridgeRegistry", "Valset", "Oracle", "QuantiexBridge", "BridgeBank", "StakingPool"}

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
		QuantiexBridge,
		BridgeBank,
		StakingPool,
	}
}
