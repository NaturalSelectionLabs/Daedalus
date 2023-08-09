package convert

import (
	"fmt"
	"github.com/naturalSelectionLabs/daedalus/internal/converter"
	"github.com/naturalSelectionLabs/daedalus/pkg/types"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

var oldFile string

var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "action functions",
	Run: func(cmd *cobra.Command, args []string) {
		if oldFile == "" {
			fmt.Println("Please provide a configuration file using -f flag.")
			return
		}

		workflow := handleSingleWorkflow(oldFile)

		fmt.Println(workflow)
	},
}

func init() {
	actionCmd.Flags().StringVarP(&oldFile, "file", "f", "", "from old file")
}

func handleSingleWorkflow(filePath string, options ...converter.FlowOption) *types.Workflow {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading the configuration file: %s\n", err)
		return nil
	}

	var workflow types.Workflow

	_ = yaml.Unmarshal(file, &workflow)

	newWorkflow := converter.MigrateWorkflow(workflow, options...)

	return &newWorkflow
}
