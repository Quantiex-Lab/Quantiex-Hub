package txs

import (
	"context"
	bridgeregistry "github.com/Quantiex-Hub/cmd/bscrelayer/bsccontract/generated/bindings/bridgeregistry"
	bscBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// TODO: Update BridgeRegistry bsccontract so that all bridge bsccontract addresses can be queried
//		in one transaction. Then refactor ContractRegistry to a map and store it under new
//		Relayer struct.

// ContractRegistry is an enum for the bridge bsccontract types
type ContractRegistry byte

const (
	// Valset valset bsccontract
	Valset ContractRegistry = iota + 1
	// Oracle bsccontract
	Oracle
	// BridgeBank bridgeBank bsccontract
	BridgeBank
	// QuantiexBridge quantiexBridge bsccontract
	QuantiexBridge
)

// String returns the event type as a string
func (d ContractRegistry) String() string {
	return [...]string{"valset", "oracle", "bridgebank", "binancebridge"}[d-1]
}

// GetAddressFromBridgeRegistry queries the requested bsccontract address from the BridgeRegistry bsccontract
func GetAddressFromBridgeRegistry(client *ethclient.Client, registry common.Address, target ContractRegistry,
) (common.Address, error) {
	sender, err := LoadBscSender()
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set up CallOpts auth
	//var xxx *big.Int
	//var sender .Address
	auth := bscBind.CallOpts{
		Pending:     true,
		From:        sender,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	// Initialize BridgeRegistry instance
	registryInstance, err := bridgeregistry.NewBridgeRegistry(registry, client)
	if err != nil {
		log.Fatal(err)
	}

	var address common.Address
	switch target {
	case Valset:
		address, err = registryInstance.Valset(&auth)
	case Oracle:
		address, err = registryInstance.Oracle(&auth)
	case BridgeBank:
		address, err = registryInstance.BridgeBank(&auth)
	case QuantiexBridge:
		address, err = registryInstance.QuantiexBridge(&auth)
	default:
		panic("invalid target bsccontract address")
	}

	if err != nil {
		log.Fatal(err)
	}

	return address, nil
}


