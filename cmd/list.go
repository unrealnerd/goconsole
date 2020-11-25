package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list items",
	Long:  `list <entity> - will list the specified entity
			--limit - option will let you limit the number of items in the result
			--offset - used as a start index from nth item to retreive`,
	// Run: handlers.ListHandle(),
}
var (
	limit  int
	offset int
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().IntVarP(&limit, "limit", "l", 100, "limits the number of items to display. By default shows upto 100")
	listCmd.PersistentFlags().IntVarP(&offset, "offset", "o", 0, "offset is the index of the item from where to retrive. default is 0")
}
