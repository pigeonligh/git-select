package main

import (
	"fmt"
	"os"

	"github.com/pigeonligh/git-select/cmd/git-select/commands"
)

func main() {
	if err := commands.NewCommand("git-select").Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
