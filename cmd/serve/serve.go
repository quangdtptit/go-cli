package serve

import (
	"github.com/quangdtptit/go-cli/config"
	"github.com/quangdtptit/go-cli/internal/app"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long:  "Use this cli to start server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := config.LoadConfig()
		app.Run(cfg)
	},
}
