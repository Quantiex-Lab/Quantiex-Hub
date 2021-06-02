package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/Quantiex-Hub/cmd/bscrelayer/bsccontract"
	"github.com/Quantiex-Hub/cmd/bscrelayer/rpc/client"
	"github.com/Quantiex-Hub/cmd/bscrelayer/txs"
	"github.com/Quantiex-Hub/cmd/bscrelayer/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ctypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"

	xcommon "github.com/Quantiex-Hub/x/common/types"
	bsccommon "github.com/ethereum/go-ethereum/common"
	tmLog "github.com/tendermint/tendermint/libs/log"
)

// TODO: Move relay functionality out of BinanceSub into a new Relayer parent struct

// BinanceSub defines a Binance listener that relays events to Ethereum and Binance
type BinanceSub struct {
	BscProvider             string
	RegistryContractAddress bsccommon.Address
	PrivateKey              *ecdsa.PrivateKey
	Logger                  tmLog.Logger
}

// NewBinanceSub initializes a new BinanceSub
func NewBinanceSub(bscProvider string, registryContractAddress bsccommon.Address,
	privateKey *ecdsa.PrivateKey, logger tmLog.Logger) BinanceSub {
	return BinanceSub{
		BscProvider:             bscProvider,
		RegistryContractAddress: registryContractAddress,
		PrivateKey:              privateKey,
		Logger:                  logger,
	}
}

// Start a Binance chain subscription
func (sub BinanceSub) Start() {
	client, err := SetupWebsocketBscClient(sub.BscProvider)
	if err != nil {
		sub.Logger.Error(err.Error())
		os.Exit(1)
	}
	sub.Logger.Info("Started Binance websocket with provider:", sub.BscProvider)

	clientChainID, err := client.NetworkID(context.Background())
	if err != nil {
		println(clientChainID)
		sub.Logger.Error(err.Error())
		os.Exit(1)
	}
	//We will check logs for new events
	logs := make(chan ctypes.Log)
	//Start BridgeBank subscription, prepare bsccontract ABI and LockLog event signature
	bridgeBankAddress, subBridgeBank := sub.startContractEventSub(logs, client, txs.BridgeBank)
	bridgeBankContractABI := bsccontract.LoadABI(txs.BridgeBank)
	eventLogLockSignature := bridgeBankContractABI.Events[types.LogLock.String()].ID.Hex()
	eventLogBurnSignature := bridgeBankContractABI.Events[types.LogBurn.String()].ID.Hex()

	// Start QuantiexBridge subscription, prepare ethcontract ABI and LogNewProphecyClaim event signature
	quantiexBridgeAddress, subQuantiexBridge := sub.startContractEventSub(logs, client, txs.QuantiexBridge)
	quantiexBridgeContractABI := bsccontract.LoadABI(txs.QuantiexBridge)
	eventLogNewProphecyClaimSignature := quantiexBridgeContractABI.Events[types.LogNewProphecyClaim.String()].ID.Hex()

	for {
		select {
		// Handle any errors
		case err := <-subBridgeBank.Err():
			sub.Logger.Error(err.Error())
		case err := <-subQuantiexBridge.Err():
			sub.Logger.Error(err.Error())
		//vLog is raw event data
		case vLog := <-logs:
			sub.Logger.Info(fmt.Sprintf("Witnessed tx %s on block %d\n", vLog.TxHash.Hex(), vLog.BlockNumber))
			var err error
			switch vLog.Topics[0].Hex() {
			case eventLogBurnSignature:
				err = sub.handleBinanceEvent(clientChainID, bridgeBankAddress, bridgeBankContractABI,
					types.LogBurn.String(), vLog)
			case eventLogLockSignature:
				err = sub.handleBinanceEvent(clientChainID, bridgeBankAddress, bridgeBankContractABI,
					types.LogLock.String(), vLog)
			case eventLogNewProphecyClaimSignature:
				err = sub.handleLogNewProphecyClaim(quantiexBridgeAddress, quantiexBridgeContractABI,
					types.LogNewProphecyClaim.String(), vLog)
			}
			// TODO: Check local events store for status, if retryable, attempt relay again
			if err != nil {
				sub.Logger.Error(err.Error())
			}
		}
	}
}

// startContractEventSub : starts an event subscription on the specified Peggy bsccontract
func (sub BinanceSub) startContractEventSub(logs chan ctypes.Log, client *ethclient.Client,
	contractName txs.ContractRegistry) (common.Address, ethereum.Subscription) {
	// Get the bsccontract address for this subscription
	subContractAddress, err := txs.GetAddressFromBridgeRegistry(client, sub.RegistryContractAddress, contractName)
	if err != nil {
		sub.Logger.Error(err.Error())
	}
	//var subContractAddress common.Address
	// We need the address in []bytes for the query
	subQuery := ethereum.FilterQuery{
		Addresses: []common.Address{subContractAddress},
	}
	println(&subQuery)

	// Start the bsccontract subscription
	contractSub, err := client.SubscribeFilterLogs(context.Background(), subQuery, logs)
	if err != nil {
		sub.Logger.Error(err.Error())
	}
	//var contractSub ethereum.Subscription
	sub.Logger.Info(fmt.Sprintf("Subscribed to %v bsccontract at address: %s", contractName, subContractAddress.Hex()))
	return subContractAddress, contractSub
}

// handleEthereumEvent unpacks an Ethereum event, converts it to a ProphecyClaim, and relays a tx to Binance
func (sub BinanceSub) handleBinanceEvent(clientChainID *big.Int, contractAddress common.Address,
	contractABI abi.ABI, eventName string, cLog ctypes.Log) error {
	// Parse the event's attributes via bsccontract ABI
	event := types.BinanceEvent{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	event.BridgeContractAddress = contractAddress
	event.BinanceChainID = clientChainID

	if eventName == types.LogBurn.String() {
		event.ClaimType = xcommon.BurnText
	} else {
		event.ClaimType = xcommon.LockText
	}
	sub.Logger.Info(event.String())

	// Add the event to the record
	txHash := cLog.TxHash.Hex()
	types.NewEventWrite(txHash, event)

	prophecyClaim := xcommon.BscProphecyClaim{
		ChainName:        event.ChainName,
		ClaimType:        event.ClaimType,
		BinanceSender:     xcommon.NewBinanceAddress(event.From.Hex()),
		EthereumReceiver: xcommon.NewEthereumAddress(event.To.Hex()),
		Symbol:           event.Symbol,
		Amount:           event.Value.String(),
		TxHash:           txHash,
	}

	//send claim to EthRelayer
	_, err = client.SendProphecyClaimToEthereum(prophecyClaim)
	if err != nil {
		return err
	}

	return nil
}

// Unpacks a handleLogNewProphecyClaim event, builds a new OracleClaim, and relays it to Ethereum
func (sub BinanceSub) handleLogNewProphecyClaim(contractAddress common.Address, contractABI abi.ABI,
	eventName string, cLog ctypes.Log) error {
	// Parse the event's attributes via ethcontract ABI
	event := types.ProphecyClaimEvent{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	sub.Logger.Info(event.String())

	oracleClaim, err := txs.ProphecyClaimToSignedOracleClaim(event, sub.PrivateKey)
	if err != nil {
		return err
	}
	return txs.RelayOracleClaimToBinance(sub.BscProvider, contractAddress, types.LogNewProphecyClaim,
		oracleClaim, sub.PrivateKey)
}
