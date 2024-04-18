package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use: "db",
  Short: "Very short",
  Long: "Little longer",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hello this is a db app. Please use --help to get more information.")
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
