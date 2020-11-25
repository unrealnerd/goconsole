package cmd

import (
	"goconsole/config"
	"goconsole/handlers"
	"fmt"

	"github.com/spf13/cobra"
)

var namespace string

// tableCmd represents the table command
var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "list tables in schema",
	Long: `	list tables <namespace> - will list the tables in the specified schema
			--limit - option will let you limit the number of items in the result
			--offset - used as a start index from nth item to retreive`,
	Run: handlers.TableHandle(),
}

func init() {
	listCmd.AddCommand(tableCmd)
	fmt.Println("Default Namespace: ", config.Data.DefaultNameSpace)
	tableCmd.PersistentFlags().StringVarP(&namespace, "schema", "s", config.Data.DefaultNameSpace, "Schema name")

}
