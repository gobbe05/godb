package cmd

import (
	"db/dbreadwrite"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(getdocument)
}

var getdocument = &cobra.Command{
  Use: "getdocument [collection] [id]",
  Short: "Gets a document",
  Long: "Gets a document using a collection and an id",
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) < 2 {
      panic("Please pass all the needed arguments")
    }
    collectionName := args[0] // Gets the collection name
    id, err := strconv.Atoi(args[1]) // Gets the id from args and returns an error if it can't be converted
    if err != nil {
      panic(err)
    }

    dbreadwrite.GetDocument(collectionName, id)
  },
}
