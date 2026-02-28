package cmd

import (
	"log"

	"github.com/quangdtptit/go-cli/cmd/migrate"
	"github.com/quangdtptit/go-cli/cmd/serve"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "My application",
	Long:  "This is a sample CLI application",
}

func init() {
	rootCmd.AddCommand(serve.Cmd, migrate.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("[Execute] Error: ", err)
	}
}
