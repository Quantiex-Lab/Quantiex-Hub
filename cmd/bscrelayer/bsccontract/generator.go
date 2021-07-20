package bsccontract

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
	// BaseABIBINGEnCmd is the base command for bsccontract compilation to ABI and BIN
	BaseABIBINGenCmd = strings.Join([]string{"solc ",
		fmt.Sprintf("--%s ./bsc_contracts/contracts/%s%s.sol ", SolcCmdText, DirectoryText, ContractText),
		fmt.Sprintf("-o ./cmd/bscrelayer/bsccontract/generated/%s/%s ", SolcCmdText, ContractText),
		"--overwrite ",
		"--allow-paths *,"},
		"")
	// BaseBindingGenCmd is the base command for bsccontract binding generation
	BaseBindingGenCmd = strings.Join([]string{"abigen ",
		fmt.Sprintf("--bin ./cmd/bscrelayer/bsccontract/generated/bin/%s/%s.bin ", ContractText, ContractText),
		fmt.Sprintf("--abi ./cmd/bscrelayer/bsccontract/generated/abi/%s/%s.abi ", ContractText, ContractText),
		fmt.Sprintf("--pkg %s ", ContractText),
		fmt.Sprintf("--type %s ", ContractText),
		fmt.Sprintf("--out ./cmd/bscrelayer/bsccontract/generated/bindings/%s/%s.go", ContractText, ContractText)},
		"")
)

// CompileContracts compiles contracts to BIN and ABI files
func CompileContracts(contracts BridgeContracts) error {
	for _, bsccontract := range contracts {
		// Construct generic BIN/ABI generation cmd with bsccontract's directory path and name
		baseDirectory := ""
		if bsccontract.String() == BridgeERC20Bank.String() {
			baseDirectory = bsccontract.String() + "/"
		} else if bsccontract.String() == BridgeERC721Bank.String() {
			baseDirectory = bsccontract.String() + "/"
		}
		dirABIBINGenCmd := strings.Replace(BaseABIBINGenCmd, DirectoryText, baseDirectory, -1)
		contractABIBINGenCmd := strings.Replace(dirABIBINGenCmd, ContractText, bsccontract.String(), -1)

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

// GenerateBindings generates bindings for each bsccontract
func GenerateBindings(contracts BridgeContracts) error {
	for _, bsccontract := range contracts {
		genBindingCmd := strings.Replace(BaseBindingGenCmd, ContractText, bsccontract.String(), -1)
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
	fmt.Println("CMD: "+cmd)
	return err
}
