package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Quantiex-Hub/flags"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	tmLog "github.com/tendermint/tendermint/libs/log"

	"github.com/Quantiex-Hub/cmd/ethrelayer/ethcontract"
	"github.com/Quantiex-Hub/cmd/ethrelayer/relayer"
	"github.com/Quantiex-Hub/cmd/ethrelayer/rpc/service"
	"github.com/Quantiex-Hub/cmd/ethrelayer/txs"
)

const (
	// FlagRPCURL defines the URL for the tendermint RPC connection
	FlagRPCURL = "rpc-url"
	// EnvPrefix defines the environment prefix for the root cmd
	EnvPrefix = "ETHRELAYER"
)

func init() {

	// Add --chain-id to persistent flags and mark it required
	rootCmd.PersistentFlags().String(flags.FlagChainID, "", "Chain ID of ETH node")
	rootCmd.PersistentFlags().String(FlagRPCURL, "", "RPC URL of ETH node")
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}

	// Construct Root Command
	rootCmd.AddCommand(
		initRelayerCmd(),
		generateBindingsCmd(),
	)

	DefaultCLIHome := os.ExpandEnv("$HOME/.ethsub")
	executor := cli.PrepareMainCmd(rootCmd, EnvPrefix, os.ExpandEnv(DefaultCLIHome))
	err := executor.Execute()
	if err != nil {
		log.Fatal("failed executing CLI command", err)
	}
}

var rootCmd = &cobra.Command{
	Use:          "ethrelayer",
	Short:        "Streams live events from Ethereum and relays event information to the opposite chain BSC",
	SilenceUsage: true,
}

//	initRelayerCmd
func initRelayerCmd() *cobra.Command {
	//nolint:lll
	initRelayerCmd := &cobra.Command{
		Use:     "init [web3Provider] [bridgeRegistryContractAddress]",
		Short:   "Validate credentials and initialize subscriptions to both chains",
		Args:    cobra.ExactArgs(2),
		Example: "ethrelayer init ws://localhost:8545/ 0x30753E4A8aad7F8597332E813735Def5dD395028",
		RunE:    RunInitRelayerCmd,
	}

	return initRelayerCmd
}

//	generateBindingsCmd : Generates ABIs and bindings for Bridge smart contracts which facilitate ethcontract interaction
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
	privateKey, err := txs.LoadPrivateKey()
	if err != nil {
		return errors.Errorf("invalid [ETHEREUM_PRIVATE_KEY] environment variable")
	}

	// Validate and parse arguments
	if !relayer.IsWebsocketURL(args[0]) {
		return errors.Errorf("invalid [web3-provider]: %s", args[0])
	}
	web3Provider := args[0]

	if !common.IsHexAddress(args[1]) {
		return errors.Errorf("invalid [bridge-registry-ethcontract-address]: %s", args[1])
	}
	registryContractAddress := common.HexToAddress(args[1])

	// Universal logger
	logger := tmLog.NewTMLogger(tmLog.NewSyncWriter(os.Stdout))

	ethSub, err := relayer.NewEthereumSub(web3Provider, registryContractAddress, privateKey, logger)
	if err != nil {
		return err
	}

	// 启动http服务
	service.StartHttpServer(&ethSub)

	go ethSub.Start()

	// Exit signal enables graceful shutdown
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal

	return nil
}

// RunGenerateBindingsCmd : executes the generateBindingsCmd
func RunGenerateBindingsCmd(cmd *cobra.Command, args []string) error {
	contracts := ethcontract.LoadBridgeContracts()

	// Compile contracts, generating ethcontract bins and abis
	err := ethcontract.CompileContracts(contracts)
	if err != nil {
		return err
	}

	// Generate ethcontract bindings from bins and abis
	return ethcontract.GenerateBindings(contracts)
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
