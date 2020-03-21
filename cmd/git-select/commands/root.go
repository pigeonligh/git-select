package commands

import (
	"fmt"
	"os"

	"github.com/pigeonligh/git-select/models/config"

	"github.com/pigeonligh/git-select/models/opts"

	"github.com/spf13/cobra"
)

var (
	data *config.Map
)

// NewCommand returns a new instance of a cobra command
func NewCommand(CLIName string) *cobra.Command {
	var err error
	data, err = config.LoadConfig()
	if err != nil {
		fmt.Println("Errror @ load config : " + err.Error())
		os.Exit(1)
	}

	var opt opts.SelectOpts
	var command = &cobra.Command{
		Use:   CLIName,
		Short: "git-select is the command line to select user for git",
		Run: func(cmd *cobra.Command, args []string) {
			flags := cmd.Flags()
			tag, err := flags.GetString("tag")
			if err != nil {
				fmt.Println("Error" + err.Error())
				os.Exit(1)
			}
			for tag == "" {
				cmd.Help()
				return
			}
			global, err := flags.GetBool("global")
			if err != nil {
				fmt.Println("Error" + err.Error())
				os.Exit(1)
			}

			if err = data.Select(tag, global); err != nil {
				fmt.Println("Error: " + err.Error())
				os.Exit(1)
			}
		},
	}

	command.AddCommand(AddConfigCommand())
	command.AddCommand(ListConfigCommand())
	command.AddCommand(RemoveConfigCommand())

	command.Flags().StringVarP(&opt.Tag, "tag", "t", "", "select a user tag for git (necessary)")
	command.Flags().BoolVarP(&opt.Global, "global", "g", false, "set global or not")

	return command
}
