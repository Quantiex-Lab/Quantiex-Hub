package txs

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	oracle "github.com/Quantiex-Hub/cmd/bscrelayer/bsccontract/generated/bindings/oracle"
	quantiexbridge "github.com/Quantiex-Hub/cmd/bscrelayer/bsccontract/generated/bindings/quantiexbridge"
	"github.com/Quantiex-Hub/cmd/bscrelayer/types"

	xcommon "github.com/Quantiex-Hub/x/common/types"
)

const (
	// GasLimit the gas limit in Gwei used for transactions sent with TransactOpts
	GasLimit = uint64(3000000)
	DefaultPrefix = "PEGGY"
)

// RelayProphecyClaimToBinance relays the provided ProphecyClaim to QuantiexBridge bsccontract on the Binance network
func RelayProphecyClaimToBinance(provider string, contractAddress common.Address, event types.Event,
	claim xcommon.EthProphecyClaim, key *ecdsa.PrivateKey) error {
	// Initialize client service, validator's tx auth, and target bsccontract address
	client, auth, target := initRelayConfig(provider, contractAddress, event, key)

	// Initialize QuantiexBridge instance
	fmt.Println("\nFetching QuantiexBridge bsccontract...")
	quantiexBridgeInstance, err := quantiexbridge.NewQuantiexBridge(target, client)
	if err != nil {
		return err
	}

	// Send transaction
	fmt.Println("Sending new ProphecyClaim to QuantiexBridge...")
	if event == types.MsgBurn {
		if !strings.Contains(claim.Symbol, DefaultPrefix) {
			log.Fatal("Can only relay burns of 'PEGGY' prefixed coins")
		}
		strs := strings.SplitAfter(claim.Symbol, DefaultPrefix)
		claim.Symbol = strings.ToUpper(strings.Join(strs[1:], ""))
	} else {
		claim.Symbol = strings.ToUpper(claim.Symbol)
	}

	sender := common.HexToAddress(claim.EthereumSender.String())
	receiver := common.HexToAddress(claim.BinanceReceiver.String())
	amount := big.NewInt(0)
	amount.SetString(claim.Amount, 10)

	tx, err := quantiexBridgeInstance.NewProphecyClaim(auth, uint8(event),
		sender, receiver, claim.Symbol, amount, claim.TxHash)
	if err != nil {
		return err
	}
	fmt.Println("NewProphecyClaim tx hash:", tx.Hash().Hex())

	// Get the transaction receipt
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return err
	}

	switch receipt.Status {
	case 0:
		fmt.Println("Tx Status: 0 - Failed")
	case 1:
		fmt.Println("Tx Status: 1 - Successful")
	}
	return nil
}

// RelayOracleClaimToBinance relays the provided OracleClaim to Oracle bsccontract on the Binance network
func RelayOracleClaimToBinance(provider string, contractAddress common.Address, event types.Event,
	claim OracleClaim, key *ecdsa.PrivateKey) error {
	// Initialize client service, validator's tx auth, and target bsccontract address
	client, auth, target := initRelayConfig(provider, contractAddress, event, key)

	// Initialize Oracle instance
	fmt.Println("\nFetching Oracle bsccontract...")
	oracleInstance, err := oracle.NewOracle(target, client)
	if err != nil {
		return err
	}

	// Send transaction
	fmt.Println("Sending new OracleClaim to Oracle...")
	tx, err := oracleInstance.NewOracleClaim(auth, claim.ProphecyID, claim.Message, claim.Signature)
	if err != nil {
		return err
	}
	fmt.Println("NewOracleClaim tx hash:", tx.Hash().Hex())

	// Get the transaction receipt
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return err
	}

	switch receipt.Status {
	case 0:
		fmt.Println("Tx Status: 0 - Failed")
	case 1:
		fmt.Println("Tx Status: 1 - Successful")
	}

	return nil
}

// initRelayConfig set up Binance client, validator's transaction auth, and the target bsccontract's address
func initRelayConfig(provider string, registry common.Address, event types.Event, key *ecdsa.PrivateKey,
) (*ethclient.Client, *bind.TransactOpts, common.Address) {
	// Start Binance client
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatal(err)
	}

	// Load the validator's address
	sender, err := LoadBscSender()
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), sender)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Set up TransactOpts auth's tx signature authorization
	transactOptsAuth := bind.NewKeyedTransactor(key)
	transactOptsAuth.Nonce = big.NewInt(int64(nonce))
	transactOptsAuth.Value = big.NewInt(0) // in wei
	transactOptsAuth.GasLimit = GasLimit
	transactOptsAuth.GasPrice = gasPrice

	var targetContract ContractRegistry
	switch event {
	// ProphecyClaims are sent to the QuantiexBridge bsccontract
	case types.MsgBurn, types.MsgLock:
		targetContract = QuantiexBridge
	// OracleClaims are sent to the Oracle bsccontract
	case types.LogNewProphecyClaim:
		targetContract = Oracle
	default:
		panic("invalid target bsccontract address")
	}

	// Get the specific bsccontract's address
	target, err := GetAddressFromBridgeRegistry(client, registry, targetContract)
	if err != nil {
		log.Fatal(err)
	}
	return client, transactOptsAuth, target
}
