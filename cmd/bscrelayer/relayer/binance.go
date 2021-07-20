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
	"strings"

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
	bridgeERC20BankAddress, subBridgeERC20Bank := sub.startContractEventSub(logs, client, txs.BridgeERC20Bank)
	bridgeERC20BankContractABI := bsccontract.LoadABI(txs.BridgeERC20Bank)
	eventLogERC20LockSignature := bridgeERC20BankContractABI.Events[types.LogLock.String()].ID.Hex()
	eventLogERC20BurnSignature := bridgeERC20BankContractABI.Events[types.LogBurn.String()].ID.Hex()

	// Start QuantiexBridge subscription, prepare ethcontract ABI and LogNewProphecyClaim event signature
	_, subQuantiexERC20Bridge := sub.startContractEventSub(logs, client, txs.QuantiexERC20Bridge)
	quantiexERC20BridgeContractABI := bsccontract.LoadABI(txs.QuantiexERC20Bridge)
	eventLogERC20NewProphecyClaimSignature := quantiexERC20BridgeContractABI.Events[types.LogNewProphecyClaim.String()].ID.Hex()

	// Start BridgeERC721Bank subscription, prepare ethcontract ABI and LockLog event signature
	bridgeERC721BankAddress, subBridgeERC721Bank := sub.startContractEventSub(logs, client, txs.BridgeERC721Bank)
	bridgeERC721BankContractABI := bsccontract.LoadABI(txs.BridgeERC721Bank)
	eventLogERC721LockSignature := bridgeERC721BankContractABI.Events[types.LogLock.String()].ID.Hex()
	eventLogERC721BurnSignature := bridgeERC721BankContractABI.Events[types.LogBurn.String()].ID.Hex()

	// Start QuantiexERC721Bridge subscription, prepare ethcontract ABI and LogNewProphecyClaim event signature
	_, subQuantiexERC721Bridge := sub.startContractEventSub(logs, client, txs.QuantiexERC721Bridge)
	quantiexERC721BridgeContractABI := bsccontract.LoadABI(txs.QuantiexERC721Bridge)
	eventLogERC721NewProphecyClaimSignature := quantiexERC721BridgeContractABI.Events[types.LogNewProphecyClaim.String()].ID.Hex()

	for {
		select {
		// Handle any errors
		case err := <-subBridgeERC20Bank.Err():
			sub.Logger.Error(err.Error())
		case err := <-subQuantiexERC20Bridge.Err():
			sub.Logger.Error(err.Error())
		case err := <-subBridgeERC721Bank.Err():
			sub.Logger.Error(err.Error())
		case err := <-subQuantiexERC721Bridge.Err():
			sub.Logger.Error(err.Error())
		//vLog is raw event data
		case vLog := <-logs:
			sub.Logger.Info(fmt.Sprintf("Witnessed tx %s on block %d\n", vLog.TxHash.Hex(), vLog.BlockNumber))
			var err error
			switch vLog.Topics[0].Hex() {
			case eventLogERC20BurnSignature:
				err = sub.handleERC20BinanceEvent(clientChainID, bridgeERC20BankAddress, bridgeERC20BankContractABI,
					types.LogBurn.String(), vLog)
			case eventLogERC20LockSignature:
				err = sub.handleERC20BinanceEvent(clientChainID, bridgeERC20BankAddress, bridgeERC20BankContractABI,
					types.LogLock.String(), vLog)
			case eventLogERC20NewProphecyClaimSignature:
				err = sub.handleLogERC20NewProphecyClaim(quantiexERC20BridgeContractABI, types.LogNewProphecyClaim.String(), vLog)
			case eventLogERC721BurnSignature:
				err = sub.handleERC721BinanceEvent(clientChainID, bridgeERC721BankAddress, bridgeERC721BankContractABI,
					types.LogBurn.String(), vLog)
			case eventLogERC721LockSignature:
				err = sub.handleERC721BinanceEvent(clientChainID, bridgeERC721BankAddress, bridgeERC721BankContractABI,
					types.LogLock.String(), vLog)
			case eventLogERC721NewProphecyClaimSignature:
				err = sub.handleLogERC721NewProphecyClaim(quantiexERC721BridgeContractABI,types.LogNewProphecyClaim.String(), vLog)
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
func (sub BinanceSub) handleERC20BinanceEvent(clientChainID *big.Int, bridgeBankAddress common.Address,
	contractABI abi.ABI, eventName string, cLog ctypes.Log) error {
	// Parse the event's attributes via bsccontract ABI
	event := types.BinanceERC20Event{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	event.BridgeBankAddress = bridgeBankAddress
	event.ChainID = clientChainID

	if eventName == types.LogBurn.String() {
		event.ClaimType = xcommon.BurnText
	} else {
		event.ClaimType = xcommon.LockText
	}
	sub.Logger.Info(event.String())

	// Add the event to the record
	txHash := cLog.TxHash.Hex()
	types.NewERC20EventWrite(txHash, event)

	prophecyClaim := xcommon.BscERC20ProphecyClaim{
		ChainName:        event.ChainName,
		ClaimType:        event.ClaimType,
		BinanceSender:     xcommon.NewBinanceAddress(event.From.Hex()),
		EthereumReceiver: xcommon.NewEthereumAddress(event.To.Hex()),
		Symbol:           event.Symbol,
		Amount:           event.Value.String(),
		TxHash:           txHash,
	}

	//send claim to EthRelayer
	_, err = client.SendERC20ProphecyClaimToEthereum(prophecyClaim)
	if err != nil {
		return err
	}

	return nil
}

// Unpacks a handleLogNewProphecyClaim event, builds a new OracleClaim, and relays it to Ethereum
func (sub BinanceSub) handleLogERC20NewProphecyClaim(contractABI abi.ABI, eventName string,
	cLog ctypes.Log) error {
	// Parse the event's attributes via ethcontract ABI
	event := types.ProphecyClaimERC20Event{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	sub.Logger.Info(event.String())

	oracleClaim, err := txs.ERC20ProphecyClaimToSignedOracleClaim(event, sub.PrivateKey)
	if err != nil {
		return err
	}
	//return txs.RelayOracleClaimToBinance(sub.BscProvider, contractAddress, types.LogNewProphecyClaim,
	//	oracleClaim, sub.PrivateKey)
	return txs.RelayOracleClaimToBinance(sub.BscProvider, sub.RegistryContractAddress, types.LogNewProphecyClaim,
		oracleClaim, sub.PrivateKey)
}

// handleEthereumEvent unpacks an Ethereum event, converts it to a ProphecyClaim, and relays a tx to Binance
func (sub BinanceSub) handleERC721BinanceEvent(clientChainID *big.Int, bridgeBankAddress common.Address,
	contractABI abi.ABI, eventName string, cLog ctypes.Log) error {
	// Parse the event's attributes via bsccontract ABI
	event := types.BinanceERC721Event{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	event.BridgeBankAddress = bridgeBankAddress
	event.ChainID = clientChainID

	if eventName == types.LogBurn.String() {
		event.ClaimType = xcommon.BurnText
	} else {
		event.ClaimType = xcommon.LockText
	}
	sub.Logger.Info(event.String())

	// Add the event to the record
	txHash := cLog.TxHash.Hex()
	types.NewERC721EventWrite(txHash, event)

	tokenURI := event.TokenURI
	if strings.HasPrefix(tokenURI, event.BaseURI) {
		tokenURI = tokenURI[len(event.BaseURI):]
	}

	prophecyClaim := xcommon.BscERC721ProphecyClaim{
		ClaimType: event.ClaimType,
		ChainName: event.ChainName,
		BinanceSender: xcommon.NewBinanceAddress(event.From.Hex()),
		EthereumReceiver: xcommon.NewEthereumAddress(event.To.Hex()),
		Symbol: event.Symbol,
		TokenId: event.TokenId.String(),
		BaseURI: event.BaseURI,
		TokenURI: tokenURI,
		TxHash: txHash,
	}

	//send claim to EthRelayer
	_, err = client.SendERC721ProphecyClaimToEthereum(prophecyClaim)
	if err != nil {
		return err
	}

	return nil
}

// Unpacks a handleLogNewProphecyClaim event, builds a new OracleClaim, and relays it to Ethereum
func (sub BinanceSub) handleLogERC721NewProphecyClaim(contractABI abi.ABI, eventName string,
	cLog ctypes.Log) error {
	// Parse the event's attributes via ethcontract ABI
	event := types.ProphecyClaimERC721Event{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}
	sub.Logger.Info(event.String())

	oracleClaim, err := txs.ERC721ProphecyClaimToSignedOracleClaim(event, sub.PrivateKey)
	if err != nil {
		return err
	}
	return txs.RelayOracleClaimToBinance(sub.BscProvider, sub.RegistryContractAddress, types.LogNewProphecyClaim,
		oracleClaim, sub.PrivateKey)
}