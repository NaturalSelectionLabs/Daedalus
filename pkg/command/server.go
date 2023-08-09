package command

import (
	"github.com/naturalSelectionLabs/daedalus/pkg/server"
	"github.com/spf13/cobra"
)

var (
	port int
)

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "serve web ui",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run(port)
	},
}

func init() {
	uiCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port number for the web server")
}
