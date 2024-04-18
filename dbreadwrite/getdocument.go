package dbreadwrite

import (
	"db/types"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func GetDocument(collectionName string, id int) {
  j, err := os.Open("./data/" + collectionName + ".json")
  if err != nil {
    panic(err)
  }

  b, _ := io.ReadAll(j)

  // Create collection
  var collection types.Collection
  err = json.Unmarshal([]byte(b), &collection)
  if err != nil {
    panic(err)
  }

  d, err := collection.GetById(id)
  if err != nil {
    panic(err)
  }

  jFormatted, err := json.MarshalIndent(d, "", "    ")
  if err != nil {
    panic(err)
  }

  fmt.Println(string(jFormatted))

}
