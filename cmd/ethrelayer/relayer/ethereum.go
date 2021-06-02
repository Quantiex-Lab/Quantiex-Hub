package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/Quantiex-Hub/cmd/ethrelayer/ethcontract"
	"github.com/Quantiex-Hub/cmd/ethrelayer/rpc/client"
	"math/big"
	"os"

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

// LoadValidatorCredentials : loads validator's credentials (address, moniker, and passphrase)
func LoadValidatorCredentials(validatorFrom string) ([]byte, string, error) {
	// Get the validator's name and account address using their moniker
	//validatorAccAddress, validatorName, err := sdkContext.GetFromFields(inBuf, validatorFrom, false)
	//if err != nil {
	//	return []byte{}, "", err
	//}
	//validatorAddress := sdk.ValAddress(validatorAccAddress)

	// Confirm that the key is valid
	var err error
	return []byte{}, "", err
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

	// Start BridgeBank subscription, prepare ethcontract ABI and LockLog event signature
	bridgeBankAddress, subBridgeBank := sub.startContractEventSub(logs, client, txs.BridgeBank)
	bridgeBankContractABI := ethcontract.LoadABI(txs.BridgeBank)
	eventLogLockSignature := bridgeBankContractABI.Events[types.LogLock.String()].ID.Hex()
	eventLogBurnSignature := bridgeBankContractABI.Events[types.LogBurn.String()].ID.Hex()

	// Start QuantiexBridge subscription, prepare ethcontract ABI and LogNewProphecyClaim event signature
	quantiexBridgeAddress, subQuantiexBridge := sub.startContractEventSub(logs, client, txs.QuantiexBridge)
	quantiexBridgeContractABI := ethcontract.LoadABI(txs.QuantiexBridge)
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
				err = sub.handleEthereumEvent(clientChainID, bridgeBankAddress, bridgeBankContractABI,
					types.LogBurn.String(), vLog)
			case eventLogLockSignature:
				err = sub.handleEthereumEvent(clientChainID, bridgeBankAddress, bridgeBankContractABI,
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

// handleEthereumEvent unpacks an Ethereum event, converts it to a ProphecyClaim, and relays a tx to Binance
func (sub EthereumSub) handleEthereumEvent(clientChainID *big.Int, contractAddress common.Address,
	contractABI abi.ABI, eventName string, cLog ctypes.Log) error {
	// Parse the event's attributes via ethcontract ABI
	event := types.EthereumEvent{}
	err := contractABI.Unpack(&event, eventName, cLog.Data)
	if err != nil {
		sub.Logger.Error("error unpacking: %v", err)
	}

	if event.ChainName != os.Getenv("CONNECT_TO_CHAIN") {
		return err
	}

	event.BridgeContractAddress = contractAddress
	event.EthereumChainID = clientChainID
	if eventName == types.LogBurn.String() {
		event.ClaimType = xcommon.BurnText
	} else {
		event.ClaimType = xcommon.LockText
	}
	sub.Logger.Info(event.String())

	// Add the event to the record
	txHash := cLog.TxHash.Hex()
	types.NewEventWrite(txHash, event)

	prophecyClaim := xcommon.EthProphecyClaim{
		ClaimType:        event.ClaimType,
		EthereumSender:   xcommon.NewEthereumAddress(event.From.Hex()),
		BinanceReceiver:   xcommon.NewBinanceAddress(event.To.Hex()),
		Symbol:           event.Symbol,
		Amount:           event.Value.String(),
		TxHash:           txHash,
	}

	//send claim to EthRelayer
	_, err = client.SendProphecyClaimToBinance(prophecyClaim)
	if err != nil {
		return err
	}

	return nil
}

// Unpacks a handleLogNewProphecyClaim event, builds a new OracleClaim, and relays it to Ethereum
func (sub EthereumSub) handleLogNewProphecyClaim(contractAddress common.Address, contractABI abi.ABI,
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
	return txs.RelayOracleClaimToEthereum(sub.EthProvider, contractAddress, types.LogNewProphecyClaim,
		oracleClaim, sub.PrivateKey)
}
