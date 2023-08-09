package command

import (
	"fmt"
	"github.com/naturalSelectionLabs/daedalus/pkg/command/convert"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "daedalus",
	Short: "Build your project CI/CD tool",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(
		uiCmd,
		convert.Cmd,
	)

}
