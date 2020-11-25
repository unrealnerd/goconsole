package handlers

import (
	"goconsole/repo"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

//TableHandle ...
func TableHandle() func(cmd *cobra.Command, args []string) {

	return func(cmd *cobra.Command, args []string) {
		namespace, err := cmd.Flags().GetString("schema")

		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Namespace : ", namespace)
		
		limit, err := cmd.Flags().GetInt("limit")
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Limit: ", limit)

		offset, err := cmd.Flags().GetInt("offset")
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Offset: ", offset)

		ddlRepo := repo.Ddl{}
		if result, err := ddlRepo.GetAllTablesIn(namespace, limit, offset); err != nil {
			log.Println("Error when retreiving tables in namespace", namespace)
		} else {
			fmt.Println("---------List of Tables---------\n", strings.Join(result, "\n"))
		}

	}
}
