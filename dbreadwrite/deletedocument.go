package dbreadwrite

import (
	"db/types"
	"encoding/json"
	"io"
	"os"
)

func DeleteDocument(collectionName string, id int) {
  j, err := os.Open("./data/" + collectionName + ".json") // Open the file
  if err != nil {
    panic(err)
  }

  b, err := io.ReadAll(j) // Read the file to a byte array
  if err != nil {
    panic(err)
  }

  var collection types.Collection // Create collection
  
  err = json.Unmarshal([]byte(b), &collection) // Parse json
  if err != nil {
    panic(err)
  }

  err = collection.DeleteById(id)
  if err != nil {
    panic(err)
  }
  
  jBytes, err := json.Marshal(collection) // Turn the collection struct to json and save it as a byte array
  if err != nil {
    panic(err)
  }

  err = os.WriteFile("./data/" + collectionName + ".json", jBytes, 6440) // Write the json byte array to the file
  if err != nil {
    panic(err)
  }
}
