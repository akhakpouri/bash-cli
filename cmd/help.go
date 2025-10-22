package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helpCommand = &cobra.Command{
	Use:     "help",
	Aliases: []string{"help"},
	Short:   "Help about any command",
	Long:    "Help about any command",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		help()
	},
}

func help() {
	message := `zap helps with basic operation commands.
	
Commands:
	cp	copy, and rename, a file from one location to another recursively.

Usage: zap [flags] [options]

Use "zap options" for a list of global command-line options (applies to all commands).`
	fmt.Println(message)
}

func init() {
	root.AddCommand(helpCommand)
}
