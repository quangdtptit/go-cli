package migrate

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [migrate_name]",
	Short: "Create a new migration file",
	Long:  "Use this cli to create a new migration file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_ = createMigrationFile(args[0], cmd)
	},
}

func createMigrationFile(name string, cmd *cobra.Command) error {
	timestamp := time.Now().Format("20060101150405")
	path := cmd.Flag("path").Value.String()

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		fmt.Println("Error creating migrations directory:", err)
		return err
	}

	// up file
	upFile := fmt.Sprintf("%s__%s.up.sql", timestamp, name)
	if _, err := os.Create(filepath.Join(path, upFile)); err != nil {
		fmt.Println("Error creating up migration file:", err)
		return err
	} else {
		fmt.Println("Created up migration file successfully", upFile)
	}

	// down file
	downFile := fmt.Sprintf("%s__%s.down.sql", timestamp, name)
	if _, err := os.Create(filepath.Join(path, downFile)); err != nil {
		fmt.Println("Error creating down migration file:", err)
		return err
	} else {
		fmt.Println("Created down migration file successfully", downFile)
	}
	return nil
}
