package cmd

import (
	"db/dbreadwrite"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(createcollection)
}

var createcollection = &cobra.Command{
  Use: "createcollection",
  Short: "Creates collection",
  Long: "Creates collection",
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) == 0 {
      fmt.Println("Please pass a name as an argument")
    }
    collectionName := args[0]
    dbreadwrite.CreateCollection(collectionName)
    fmt.Println("Created collection with the name : " + args[0])
    dbreadwrite.GetCollection(collectionName)
  },
}
