package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "zap",
	Short: "zap is cli tool for performing basic operations such as deep copy, move, etc",
	Long:  "zap is cli tool for performing basic operations such as deep copy, move, etc. It will dig deep into your file system to recusrive operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while executing zap. '%s'", err)
		os.Exit(1)
	}
}
