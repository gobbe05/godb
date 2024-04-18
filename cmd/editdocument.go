package cmd

import (
	"db/dbreadwrite"
	"encoding/json"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(editdocument)
}

var editdocument = &cobra.Command{
  Use: "editdocument [collection] [id] [json]",
  Short: "Updates a document",
  Long: "Updates a document using a collection name, an id and new values",
  Run: func(cmd *cobra.Command, args []string) {
    // Check so that all fields are passed
    if len(args) < 3 {
      panic("Please enter all the requried fields!")
    }
    
    collectionName := args[0] // Gets the name of the collection from args
    id, err := strconv.Atoi(args[1]) // Gets the id from args and returns an error if it can't be converted
    if err != nil {
      panic(err)
    }
    j := args[2] // Gets the json from args

    d := make(map[string]interface{}) // Creates a new document
    err = json.Unmarshal([]byte(j), &d) // Casts the json from args to a byte array, parses it and saves it in the document map
    if err != nil {
      panic(err)
    }

    dbreadwrite.EditDocument("./data/" + collectionName + ".json", id, d)
    dbreadwrite.GetDocument(collectionName, id) 
  },
}
