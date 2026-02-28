package migrate

import (
	"github.com/spf13/cobra"
)

const (
	_defaultPath = "./migrations"
)

var Cmd = &cobra.Command{
	Use:   "migrate",
	Short: "Database migrate cli",
	Long:  "Use this cli to manage your database migrations.",
}

func init() {
	// add migrate command
	Cmd.AddCommand(createCmd, upCmd, downCmd)
	// option
	Cmd.PersistentFlags().StringP("path", "p", _defaultPath, "migrations path")
}
