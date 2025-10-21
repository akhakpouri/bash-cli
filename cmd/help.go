package cmd

import "github.com/spf13/cobra"

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

}

func init() {
	root.AddCommand(helpCommand)
}
