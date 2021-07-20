package main

import (
	"github.com/Quantiex-Hub/cmd/bscrelayer/bsccontract"
	"github.com/Quantiex-Hub/cmd/bscrelayer/rpc/service"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Quantiex-Hub/flags"
	bsccommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	tmLog "github.com/tendermint/tendermint/libs/log"

	"github.com/Quantiex-Hub/cmd/bscrelayer/relayer"
	"github.com/Quantiex-Hub/cmd/bscrelayer/txs"
)

const (
	// FlagRPCURL defines the URL for the tendermint RPC connection
	FlagRPCURL = "rpc-url"
	// EnvPrefix defines the environment prefix for the root cmd
	EnvPrefix = "BSCRELAYER"
)

func init() {

	// Add --chain-id to persistent flags and mark it required
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}

	// Construct Root Command
	rootCmd.AddCommand(
		initRelayerCmd(),
		generateBindingsCmd(),
	)

	DefaultCLIHome := os.ExpandEnv("$HOME/.bscsub")
	executor := cli.PrepareMainCmd(rootCmd, EnvPrefix, os.ExpandEnv(DefaultCLIHome))
	err := executor.Execute()
	if err != nil {
		log.Fatal("failed executing CLI command", err)
	}
}

var rootCmd = &cobra.Command{
	Use:          "bscrelayer",
	Short:        "Streams live events from Ethereum and Binance and relays event information to the opposite chain",
	SilenceUsage: true,
}

//	initRelayerCmd
func initRelayerCmd() *cobra.Command {
	//nolint:lll
	initRelayerCmd := &cobra.Command{
		Use:     "init [web3Provider] [bridgeRegistryContractAddress]",
		Short:   "Validate credentials and initialize subscriptions to both chains",
		Args:    cobra.ExactArgs(2),
		Example: "bscrelayer init ws://localhost:8545/ 0x30753E4A8aad7F8597332E813735Def5dD395028",
		RunE:    RunInitRelayerCmd,
	}

	return initRelayerCmd
}

//	generateBindingsCmd : Generates ABIs and bindings for Bridge smart contracts which facilitate bsccontract interaction
func generateBindingsCmd() *cobra.Command {
	generateBindingsCmd := &cobra.Command{
		Use:     "gen",
		Short:   "Generates Bridge smart contracts ABIs and bindings",
		Args:    cobra.ExactArgs(0),
		Example: "gen",
		RunE:    RunGenerateBindingsCmd,
	}

	return generateBindingsCmd
}

// RunInitRelayerCmd executes initRelayerCmd
func RunInitRelayerCmd(cmd *cobra.Command, args []string) error {
	// Load the validator's Ethereum private key from environment variables
	privateKey, err := txs.LoadBscPrivateKey()
	if err != nil {
		return errors.Errorf("invalid [BINANCE_PRIVATE_KEY] environment variable")
	}

	// Validate and parse arguments
	if !relayer.IsWebsocketURL(args[0]) {
		return errors.Errorf("invalid [web3-provider]: %s", args[0])
	}
	web3Provider := args[0]

	//parameter  for bsccontract,maybe not used and finally hardcode in binance chain
	if !bsccommon.IsHexAddress(args[1]) {
		return errors.Errorf("invalid [bridge-registry-bsccontract-address]: %s", args[1])
	}
	registryContractAddress := bsccommon.HexToAddress(args[1])

	// Universal logger
	logger := tmLog.NewTMLogger(tmLog.NewSyncWriter(os.Stdout))

	// Initialize new Binance event listener
	binanceSub := relayer.NewBinanceSub(web3Provider, registryContractAddress, privateKey, logger)

	// 启动http服务
	service.StartHttpServer(&binanceSub)

	//go binanceSub.Start()
	binanceSub.Start()

	// Exit signal enables graceful shutdown
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal

	return nil
}

// RunGenerateBindingsCmd : executes the generateBindingsCmd
func RunGenerateBindingsCmd(cmd *cobra.Command, args []string) error {
	contracts := bsccontract.LoadBridgeContracts()

	// Compile contracts, generating bsccontract bins and abis
	err := bsccontract.CompileContracts(contracts)
	if err != nil {
		return err
	}

	// Generate bsccontract bindings from bins and abis
	return bsccontract.GenerateBindings(contracts)
}

func initConfig(cmd *cobra.Command) error {
	return viper.BindPFlag(flags.FlagChainID, cmd.PersistentFlags().Lookup(flags.FlagChainID))
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
