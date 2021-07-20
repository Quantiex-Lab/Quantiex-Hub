package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/Quantiex-Hub/cmd/ethrelayer/ethcontract"
	"github.com/Quantiex-Hub/cmd/ethrelayer/rpc/client"
	"math/big"
	"os"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ctypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	tmLog "github.com/tendermint/tendermint/libs/log"

	"github.com/Quantiex-Hub/cmd/ethrelayer/txs"
	"github.com/Quantiex-Hub/cmd/ethrelayer/types"
	xcommon "github.com/Quantiex-Hub/x/common/types"
)

// TODO: Move relay functionality out of EthereumSub into a new Relayer parent struct

// EthereumSub is an Ethereum listener that can relay txs to Binance and Ethereum
type EthereumSub struct {
	EthProvider             string
	RegistryContractAddress common.Address
	PrivateKey              *ecdsa.PrivateKey
	Logger                  tmLog.Logger
}

// NewEthereumSub initializes a new EthereumSub
func NewEthereumSub(ethProvider string, registryContractAddress common.Address,
	privateKey *ecdsa.PrivateKey, logger tmLog.Logger) (EthereumSub, error) {
	return EthereumSub{
		EthProvider:             ethProvider,
		RegistryContractAddress: registryContractAddress,
		PrivateKey:              privateKey,
		Logger:                  logger,
	}, nil
}

// Start an Ethereum chain subscription
func (sub EthereumSub) Start() {
	client, err := SetupWebsocketEthClient(sub.EthProvider)
	if err != nil {
		sub.Logger.Error(err.Error())
		os.Exit(1)
	}
	sub.Logger.Info("Started Ethereum websocket with provider:", sub.EthProvider)

	clientChainID, err := client.NetworkID(context.Background())
	if err != nil {
		sub.Logger.Error(err.Error())
		os.Exit(1)
	}

	// We will check logs for new events
	logs := make(chan ctypes.Log)

	// Start BridgeERC20Bank subscription, prepare ethcontract ABI and LockLog event signature
	bridgeERC20BankAddress, subBridgeERC20Bank := sub.startContractEventSub(logs, client, txs.BridgeERC20Bank)
	bridgeERC20BankContractABI := ethcontract.LoadABI(txs.BridgeERC20Bank)
	eventLogERC20LockSignature := bridgeERC20BankContractABI.Events[types.LogLock.String()].ID.Hex()
	eventLogERC20BurnSignature := bridgeERC20BankContractABI.Events[types.LogBurn.String()].ID.Hex()

	// Start QuantiexERC20Bridge subscription, prepare ethcontract ABI and LogNewProphecyClaim event signature
	_, subQuantiexERC20Bridge := sub.startContractEventSub(logs, client, txs.QuantiexERC20Bridge)
	quantiexERC20BridgeContractABI := ethcontract.LoadABI(txs.QuantiexERC20Bridge)
	eventLogERC20NewProphecyClaimSignature := quantiexERC20BridgeContractABI.Events[types.LogNewProphecyClaim.String()].ID.Hex()

	// Start BridgeERC721Bank subscription, prepare ethcontract ABI and LockLog event signature
	bridgeERC721BankAddress, subBridgeERC721Bank := sub.startContractEventSub(logs, client, txs.BridgeERC721Bank)
	bridgeERC721BankContractABI := ethcontract.LoadABI(txs.BridgeERC721Bank)
	eventLogERC721LockSignature := bridgeERC721BankContractABI.Events[types.LogLock.String()].ID.Hex()
	eventLogERC721BurnSignature := bridgeERC721BankContractABI.Events[types.LogBurn.String()].ID.Hex()

	// Start QuantiexERC721Bridge subscription, prepare ethcontract ABI and LogNewProphecyClaim event signature
	_, subQuantiexERC721Bridge := sub.startContractEventSub(logs, client, txs.QuantiexERC721Bridge)
	quantiexERC721BridgeContractABI := ethcontract.LoadABI(txs.QuantiexERC721Bridge)
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
				err = sub.handleERC20EthereumEvent(clientChainID, bridgeERC20BankAddress, bridgeERC20BankContractABI,
					types.LogBurn.String(), vLog)
			case eventLogERC20LockSignature:
				err = sub.handleERC20EthereumEvent(clientChainID, bridgeERC20BankAddress, bridgeERC20BankContractABI,
					types.LogLock.String(), vLog)
			case eventLogERC20NewProphecyClaimSignature:
				err = sub.handleLogERC20NewProphecyClaim(quantiexERC20BridgeContractABI, types.LogNewProphecyClaim.String(), vLog)
			case eventLogERC721BurnSignature:
				err = sub.handleERC721EthereumEvent(clientChainID, bridgeERC721BankAddress, bridgeERC721BankContractABI,
					types.LogBurn.String(), vLog)
			case eventLogERC721LockSignature:
				err = sub.handleERC721EthereumEvent(clientChainID, bridgeERC721BankAddress, bridgeERC721BankContractABI,
					types.LogLock.String(), vLog)
			case eventLogERC721NewProphecyClaimSignature:
				err = sub.handleLogERC721NewProphecyClaim(quantiexERC721BridgeContractABI, types.LogNewProphecyClaim.String(), vLog)
			}
			// TODO: Check local events store for status, if retryable, attempt relay again
			if err != nil {
				sub.Logger.Error(err.Error())
			}
		}
	}
}

// startContractEventSub : starts an event subscription on the specified Peggy ethcontract
func (sub EthereumSub) startContractEventSub(logs chan ctypes.Log, client *ethclient.Client,
	contractName txs.ContractRegistry) (common.Address, ethereum.Subscription) {
	// Get the ethcontract address for this subscription
	subContractAddress, err := txs.GetAddressFromBridgeRegistry(client, sub.RegistryContractAddress, contractName)
	if err != nil {
		sub.Logger.Error(err.Error())
	}

	// We need the address in []bytes for the query
	subQuery := ethereum.FilterQuery{
		Addresses: []common.Address{subContractAddress},
	}

	// Start the ethcontract subscription
	contractSub, err := client.SubscribeFilterLogs(context.Background(), subQuery, logs)
	if err != nil {
		sub.Logger.Error(err.Error())
	}
	sub.Logger.Info(fmt.Sprintf("Subscribed to %v ethcontract at address: %s", contractName, subContractAddress.Hex()))
	return subContractAddress, contractSub
}

// handleEthereumERC20Event unpacks an Ethereum event, converts it to a ProphecyClaim, and relays a tx to Binance
func (sub EthereumSub) handleERC20EthereumEvent(clientChainID *big.Int, bankAddress common.Address,
	bankContractABI abi.ABI, eventName string, cLog ctypes.Log) error {
	// Parse the event's attributes via ethcontract ABI
	event := types.EthereumERC20Event{}
	err := bankContractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}

	if event.ChainName != os.Getenv("CONNECT_TO_CHAIN") {
		return err
	}

	event.BankAddress = bankAddress
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

	prophecyClaim := xcommon.EthERC20ProphecyClaim{
		ClaimType:        event.ClaimType,
		EthereumSender:   xcommon.NewEthereumAddress(event.From.Hex()),
		BinanceReceiver:   xcommon.NewBinanceAddress(event.To.Hex()),
		Symbol:           event.Symbol,
		Amount:           event.Value.String(),
		TxHash:           txHash,
	}

	//send claim to EthRelayer
	_, err = client.SendERC20ProphecyClaimToBinance(prophecyClaim)
	if err != nil {
		return err
	}

	return nil
}

// Unpacks a handleLogERC20NewProphecyClaim event, builds a new OracleClaim, and relays it to Ethereum
func (sub EthereumSub) handleLogERC20NewProphecyClaim(contractABI abi.ABI, eventName string,
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
	return txs.RelayOracleClaimToEthereum(sub.EthProvider, sub.RegistryContractAddress, types.LogNewProphecyClaim,
		oracleClaim, sub.PrivateKey)
}

// handleEthereumERC721Event unpacks an Ethereum event, converts it to a ProphecyClaim, and relays a tx to Binance
func (sub EthereumSub) handleERC721EthereumEvent(clientChainID *big.Int, bankAddress common.Address,
	bankContractABI abi.ABI, eventName string, cLog ctypes.Log) error {
	// Parse the event's attributes via ethcontract ABI
	event := types.EthereumERC721Event{}
	err := bankContractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}

	if event.ChainName != os.Getenv("CONNECT_TO_CHAIN") {
		return err
	}

	event.BankAddress = bankAddress
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

	prophecyClaim := xcommon.EthERC721ProphecyClaim{
		ClaimType: event.ClaimType,
		ChainName: event.ChainName,
		EthereumSender: xcommon.NewEthereumAddress(event.From.Hex()),
		BinanceReceiver: xcommon.NewBinanceAddress(event.To.Hex()),
		Symbol: event.Symbol,
		TokenId: event.TokenId.String(),
		BaseURI: event.BaseURI,
		TokenURI: tokenURI,
		TxHash: txHash,
	}

	//send claim to EthRelayer
	_, err = client.SendERC721ProphecyClaimToBinance(prophecyClaim)
	if err != nil {
		return err
	}

	return nil
}

// Unpacks a handleLogERC721NewProphecyClaim event, builds a new OracleClaim, and relays it to Ethereum
func (sub EthereumSub) handleLogERC721NewProphecyClaim(contractABI abi.ABI, eventName string,
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
	return txs.RelayOracleClaimToEthereum(sub.EthProvider, sub.RegistryContractAddress, types.LogNewProphecyClaim,
		oracleClaim, sub.PrivateKey)
}