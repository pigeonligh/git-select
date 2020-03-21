package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pigeonligh/git-select/models/opts"

	"github.com/spf13/cobra"
)

// AddConfigCommand returns a new instance of a cobra command for adding config
func AddConfigCommand() *cobra.Command {
	var opt opts.ConfigOpts

	var command = &cobra.Command{
		Use:   "add",
		Short: "Add an new user config",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Add an user config.")

			reader := bufio.NewReader(os.Stdin)

			flags := cmd.Flags()
			name, err := flags.GetString("name")
			for name == "" && err == nil {
				fmt.Print("Enter your name: ")
				name, err = reader.ReadString('\n')
				name = strings.TrimSpace(name)
			}
			if err != nil {
				fmt.Println("Error: " + err.Error())
				os.Exit(1)
			}
			email, err := flags.GetString("email")
			for email == "" && err == nil {
				fmt.Print("Enter your email: ")
				email, err = reader.ReadString('\n')
				email = strings.TrimSpace(email)
			}
			if err != nil {
				fmt.Println("Error: " + err.Error())
				os.Exit(1)
			}
			key, err := flags.GetString("key")
			for key == "" && err == nil {
				fmt.Print("Enter file in which the ssh-key is (such as ~/.ssh/id_rsa): ")
				key, err = reader.ReadString('\n')
				key = strings.TrimSpace(key)
			}
			if err != nil {
				fmt.Println("Error: " + err.Error())
				os.Exit(1)
			}

			tag, err := flags.GetString("tag")
			for data.CheckTag(tag) && err == nil {
				if tag != "" {
					fmt.Println("The tag is used.")
				}
				fmt.Print("Enter a tag for the new config: ")
				tag, err = reader.ReadString('\n')
				tag = strings.TrimSpace(tag)
			}
			if err != nil {
				fmt.Println("Error: " + err.Error())
				os.Exit(1)
			}

			err = data.Add(opts.ConfigOpts{
				Name:    name,
				Email:   email,
				KeyPath: key,
				Tag:     tag,
			})
			if err != nil {
				fmt.Println("Error: " + err.Error())
				os.Exit(1)
			}
		},
	}

	command.Flags().StringVarP(&opt.Name, "name", "n", "", "override user.name")
	command.Flags().StringVarP(&opt.Email, "email", "e", "", "override user.email")
	command.Flags().StringVarP(&opt.KeyPath, "key", "k", "", "override ssh-key")
	command.Flags().StringVarP(&opt.Tag, "tag", "t", "", "tag for the new config")

	return command
}
