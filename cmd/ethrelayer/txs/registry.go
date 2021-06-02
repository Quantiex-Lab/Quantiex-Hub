package txs

import (
	"context"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"log"
	bridgeregistry "github.com/Quantiex-Hub/cmd/ethrelayer/ethcontract/generated/bindings/bridgeregistry"
)

// TODO: Update BridgeRegistry ethcontract so that all bridge ethcontract addresses can be queried
//		in one transaction. Then refactor ContractRegistry to a map and store it under new
//		Relayer struct.

// ContractRegistry is an enum for the bridge ethcontract types
type ContractRegistry byte

const (
	// Valset valset ethcontract
	Valset ContractRegistry = iota + 1
	// Oracle ethcontract
	Oracle
	// BridgeBank  ethcontract
	BridgeBank
	// QuantiexBridge ethcontract
	QuantiexBridge
)

// String returns the event type as a string
func (d ContractRegistry) String() string {
	return [...]string{"valset", "oracle", "bridgebank", "binancebridge"}[d-1]
}

// GetAddressFromBridgeRegistry queries the requested ethcontract address from the BridgeRegistry ethcontract
func GetAddressFromBridgeRegistry(client *ethclient.Client, registry common.Address, target ContractRegistry,
) (common.Address, error) {
	sender, err := LoadSender()
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set up CallOpts auth
	auth := ethBind.CallOpts{
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
		panic("invalid target ethcontract address")
	}

	if err != nil {
		log.Fatal(err)
	}

	return address, nil
}

