package command

import (
	"github.com/naturalSelectionLabs/daedalus/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	port  int
	token string
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "serve web ui",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run(port)
	},
}

func init() {
	viper.SetEnvPrefix("")
	viper.BindEnv("GITHUB_TOKEN")
	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port number for the web server")
	serverCmd.Flags().StringVarP(&token, "github-token", "", viper.GetString("GITHUB_TOKEN"), "Github API access token")
	server.InitGithubClient(token)
}
