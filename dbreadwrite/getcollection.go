package dbreadwrite

import (
	"db/types"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func GetCollection(name string) {
  j, err := os.Open("./data/" + name + ".json")
  if err != nil {
    panic(err)
  }
  b,_ := io.ReadAll(j)
  
  // Create collection
  var collection types.Collection
  err = json.Unmarshal([]byte(b), &collection)
  if err != nil {
    panic(err)
  }
  jFormatted, err := json.MarshalIndent(collection, "", "    ")
  if err != nil {
    panic(err)
  }
  fmt.Println(string(jFormatted))
}
