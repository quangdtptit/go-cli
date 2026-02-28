package migrate

import (
	"log"

	"github.com/quangdtptit/go-cli/pkg/migration"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Apply all down migration files",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		err := migration.Down(path)
		if err != nil {
			log.Fatal(err)
		}
	},
}
