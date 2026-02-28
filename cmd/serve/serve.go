package serve

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	_defaultPort = 8080
)

var Cmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long:  "Use this cli to start server",
	Run: func(cmd *cobra.Command, args []string) {
		port := cmd.Flag("port").Value.String()
		fmt.Println("Server started with port", port)
	},
}

func init() {
	Cmd.Flags().IntP("port", "p", _defaultPort, "Port to serve on")
}
