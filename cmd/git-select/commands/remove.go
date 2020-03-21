package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RemoveConfigCommand returns a new instance of a cobra command for removing config
func RemoveConfigCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:                   "rm tag",
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			tag := args[0]
			if err := data.Remove(tag); err != nil {
				fmt.Println("Error: " + err.Error())
			} else {
				fmt.Println(tag + " is removed.")
			}
		},
	}

	return command
}
