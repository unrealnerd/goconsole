package cmd

import (
	"github.com/spf13/cobra"
	"goconsole/handlers"
)

// schemaCmd represents the schema command
var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "list schema",
	Long:  `list schema`,
	Run:   handlers.SchemaHandle(),
}

func init() {
	listCmd.AddCommand(schemaCmd)
}
