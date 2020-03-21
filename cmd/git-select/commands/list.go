package commands

import (
	"fmt"

	"github.com/bndr/gotabulate"
	"github.com/spf13/cobra"
)

// ListConfigCommand returns a new instance of a cobra command for listing config
func ListConfigCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:                   "list",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			var tableData [][]interface{}
			var empty = true
			for tag, value := range data.Data {
				tableData = append(tableData, []interface{}{tag, value.Name, value.Email, value.KeyPath})
				empty = false
			}
			if empty {
				fmt.Println("No config.")
			} else {
				table := gotabulate.Create(tableData)
				table.SetHeaders([]string{"Name", "Email", "KeyPath"})
				fmt.Println(table.Render("grid"))
			}
		},
	}

	return command
}
