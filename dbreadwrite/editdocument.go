package dbreadwrite

import (
	"db/types"
	"encoding/json"
	"io"
	"os"
)

func EditDocument(filePath string, id int, newData map[string]interface{}) {
  j, err := os.Open(filePath)
  if err != nil {
    panic(err)
  }
  defer j.Close()

  b, err := io.ReadAll(j) 
  if err != nil {
    panic(err)
  }
  //Create collection
  var collection types.Collection
  // Parse json and insert into collection
  err = json.Unmarshal([]byte(b), &collection)
  if err != nil {
    panic(err)
  }
  // Get the document
  document, err := collection.GetById(id)
  if err != nil {
    panic(err)
  }

  for k, v := range newData {
    if k != "id" {
      document[k] = v
    }
  }

  _, err = collection.UpdateById(id, document)
  colBytes, err := json.Marshal(collection)
  if err != nil {
    panic(err)
  }
  err = os.WriteFile(filePath, colBytes, 6440)
  if err != nil {
    panic(err)
  }
}
