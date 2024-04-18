package cmd

import (
	"db/dbreadwrite"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(createdocument)
}

var createdocument =  &cobra.Command{
  Use: "createdocument [collection] [newdata]",
  Short: "Creates a new document",
  Long: `Create a new document. Please surround json in '' such as '{"name": "example"}'`,
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) == 0 {
      panic("Please pass a collection name!")
    }
    
    // Create the document map
    d := make(map[string]interface{})
    if len(args) > 1 {
      fmt.Println(args[1])
      err := json.Unmarshal([]byte(args[1]), &d)
      if err != nil {
        panic(err)
      }
    }
    collectionName := args[0]
    id := dbreadwrite.CreateDocument("./data/" + collectionName + ".json", d)
    fmt.Println("Created a new document with id : " + string(id))
    dbreadwrite.GetDocument(collectionName, id)
  },
}
