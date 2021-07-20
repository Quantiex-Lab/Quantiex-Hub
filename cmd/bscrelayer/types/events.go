package types

import "log"

// TODO: This should be moved to new 'events' directory and expanded so that it can
// serve as a local store of witnessed events and allow for re-trying failed relays.

// ERC20EventRecords map of transaction hashes to EthereumEvent structs
var ERC20EventRecords = make(map[string]BinanceERC20Event)

// ERC721EventRecords map of transaction hashes to EthereumEvent structs
var ERC721EventRecords = make(map[string]BinanceERC721Event)

// IsEventRecorded checks the sessions stored events for this transaction hash
func IsERC20EventRecorded(txHash string) bool {
	return ERC20EventRecords[txHash].Nonce != nil
}

// NewEventWrite add a validator's address to the official claims list
func NewERC20EventWrite(txHash string, event BinanceERC20Event) {
	ERC20EventRecords[txHash] = event
}

// PrintERC20EventByTx prints any witnessed events associated with a given transaction hash
func PrintERC20EventByTx(txHash string) {
	if IsERC20EventRecorded(txHash) {
		log.Println(ERC20EventRecords[txHash].String())
	} else {
		log.Printf("\nNo records from this session for tx: %v\n", txHash)
	}
}

// PrintERC20Events prints all the claims made on this event
func PrintERC20Events() {
	// For each claim, print the validator which submitted the claim
	for txHash, event := range ERC20EventRecords {
		log.Printf("\nTransaction: %v\n", txHash)
		log.Println(event.String())
	}
}



// NewERC721EventWrite add a validator's address to the official claims list
func NewERC721EventWrite(txHash string, event BinanceERC721Event) {
	ERC721EventRecords[txHash] = event
}

// IsERC721EventRecorded checks the sessions stored events for this transaction hash
func IsERC721EventRecorded(txHash string) bool {
	return ERC721EventRecords[txHash].Nonce != nil
}

// PrintERC721EventByTx prints any witnessed events associated with a given transaction hash
func PrintERC721EventByTx(txHash string) {
	if IsERC721EventRecorded(txHash) {
		log.Println(ERC721EventRecords[txHash].String())
	} else {
		log.Printf("\nNo records from this session for tx: %v\n", txHash)
	}
}

// PrintERC721Events prints all the claims made on this event
func PrintERC721Events() {
	// For each claim, print the validator which submitted the claim
	for txHash, event := range ERC721EventRecords {
		log.Printf("\nTransaction: %v\n", txHash)
		log.Println(event.String())
	}
}
