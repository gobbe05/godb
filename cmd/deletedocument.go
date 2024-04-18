package cmd

import (
	"db/dbreadwrite"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(deletedocument)
}

var deletedocument = &cobra.Command{
  Use: "deletedocument [collection] [id]",
  Short: "Deletes a document from a collection",
  Long: "Deletes a document from a collection using a collection name and an id",
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) < 2 {
      panic("Please pass all the required arguments")
    }
    collectionName  := args[0]
    id, err := strconv.Atoi(args[1])
    if err != nil {
      panic(err)
    }
    dbreadwrite.DeleteDocument(collectionName, id)
    dbreadwrite.GetCollection(collectionName)
  },
}
