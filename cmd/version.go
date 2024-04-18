package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(version)
}

var version = &cobra.Command{
  Use: "version",
  Short: "Print version",
  Long: "This function prints the package version",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("DB v0")
  },
}
