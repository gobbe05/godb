package cmd

import (
  "fmt"

  "db/dbreadwrite"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(getcollection)
}

var getcollection = &cobra.Command{
  Use: "getcollection",
  Short: "Gets collection",
  Long: "Gets a collection based on a name argument",
  Run: func(cmd *cobra.Command, args []string) {
    // Makes sure that a name argument is passed
    if len(args) == 0 {
      fmt.Println("Please pass a name as an argument")
    }
    dbreadwrite.GetCollection(args[0])
  },
}
