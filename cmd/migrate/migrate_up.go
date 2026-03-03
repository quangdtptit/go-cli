package migrate

import (
	"fmt"
	"log"

	"github.com/quangdtptit/go-cli/pkg/migration"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply all up migration files",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		err := migration.Up(path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("migrate up successfully")
	},
}
