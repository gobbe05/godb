package dbreadwrite

import (
	"db/types"
	"encoding/json"
	"io"
	"os"
)

func CreateDocument(filePath string, d map[string]interface{}) int {
  j, err := os.Open(filePath) // Open the json file
  if err != nil {
    panic(err)
  }
  defer j.Close() // Wait to close until function returns

  b,_ := io.ReadAll(j) // Read the file to a byte array

  var collection types.Collection // Create the collection variable
  err = json.Unmarshal([]byte(b), &collection) // Read the json data and assign it to the collection
  if err != nil {
    panic(err)
  }

  id := collection.GetLastId() + 1
  
  d["id"] = id // Add id to the passed object
  collection.Documents = append(collection.Documents, d) // Add the object to the collection
  jBytes, err := json.Marshal(collection) // Turn the collection struct in to json
  if err != nil {
    panic(err)
  }
  err = os.WriteFile(filePath, jBytes, 0644) // Write to the file
  if err != nil {
    panic(err)
  }
  return id // Return id of the object
}
