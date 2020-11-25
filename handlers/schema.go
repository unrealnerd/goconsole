package handlers

import (
	"goconsole/repo"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

//SchemaHandle ... Gets a function that retreives the list of name spaces
func SchemaHandle() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		ddlRepo := repo.Ddl{}
		fmt.Println("schema called")

		if result, err := ddlRepo.GetAllNamespace(); err != nil {
			log.Println("Error when retreiving namespaces")
		} else {
			fmt.Println("---------List of Namespaces---------", strings.Join(result,"\n"))
		}

	}
}
