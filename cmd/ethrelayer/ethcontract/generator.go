package ethcontract

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	SolcCmdText   = "[SOLC_CMD]"
	DirectoryText = "[DiRECTORY]"
	ContractText  = "[CONTRACT]"
)

var (
	// BaseABIBINGEnCmd is the base command for ethcontract compilation to ABI and BIN
	BaseABIBINGenCmd = strings.Join([]string{"solc ",
		fmt.Sprintf("--%s ./eth_contracts/contracts/%s%s.sol ", SolcCmdText, DirectoryText, ContractText),
		fmt.Sprintf("-o ./cmd/ethrelayer/ethcontract/generated/%s/%s ", SolcCmdText, ContractText),
		"--overwrite ",
		"--allow-paths *,"},
		"")
	// BaseBindingGenCmd is the base command for ethcontract binding generation
	BaseBindingGenCmd = strings.Join([]string{"abigen ",
		fmt.Sprintf("--bin ./cmd/ethrelayer/ethcontract/generated/bin/%s/%s.bin ", ContractText, ContractText),
		fmt.Sprintf("--abi ./cmd/ethrelayer/ethcontract/generated/abi/%s/%s.abi ", ContractText, ContractText),
		fmt.Sprintf("--pkg %s ", ContractText),
		fmt.Sprintf("--type %s ", ContractText),
		fmt.Sprintf("--out ./cmd/ethrelayer/ethcontract/generated/bindings/%s/%s.go", ContractText, ContractText)},
		"")
)

// CompileContracts compiles contracts to BIN and ABI files
func CompileContracts(contracts BridgeContracts) error {
	for _, ethcontract := range contracts {
		// Construct generic BIN/ABI generation cmd with ethcontract's directory path and name
		baseDirectory := ""
		if ethcontract.String() == BridgeBank.String() {
			baseDirectory = ethcontract.String() + "/"
		}
		dirABIBINGenCmd := strings.Replace(BaseABIBINGenCmd, DirectoryText, baseDirectory, -1)
		contractABIBINGenCmd := strings.Replace(dirABIBINGenCmd, ContractText, ethcontract.String(), -1)

		// Segment BIN and ABI generation cmds
		contractBINGenCmd := strings.Replace(contractABIBINGenCmd, SolcCmdText, "bin", -1)
		err := execCmd(contractBINGenCmd)
		if err != nil {
			return err
		}

		contractABIGenCmd := strings.Replace(contractABIBINGenCmd, SolcCmdText, "abi", -1)
		err = execCmd(contractABIGenCmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// GenerateBindings generates bindings for each ethcontract
func GenerateBindings(contracts BridgeContracts) error {
	for _, ethcontract := range contracts {
		genBindingCmd := strings.Replace(BaseBindingGenCmd, ContractText, ethcontract.String(), -1)
		err := execCmd(genBindingCmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// execCmd executes a bash cmd
func execCmd(cmd string) error {
	_, err := exec.Command("sh", "-c", cmd).Output()
	fmt.Println("cmd: "+cmd)
	return err
}
