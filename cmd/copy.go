package cmd

import (
	"github.com/spf13/cobra"
)

var copyCommand = &cobra.Command{
	Use:     "cpr",
	Aliases: []string{"copy", "recursive"},
	Short:   "Recusrive copy",
	Long:    "Look for a file name recursively in a directory",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		RecusiveCopy(args[0],
			args[1],
			args[2])
	},
}

func init() {
	root.AddCommand(copyCommand)
}
