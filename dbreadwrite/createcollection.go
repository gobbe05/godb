package dbreadwrite

import (
  "fmt"
  "os"
)

func CreateCollection(name string) {
  // os.Stat gets file info and returns an error if the file doesnt exists
  // This means that if the file exists no err will be returned and a panic will occur
  if _, err := os.Stat("./data/" + name + ".json"); err == nil {
    panic("File already exists")
  }
  // Create json file
  file, err := os.Create("./data/" + name + ".json")
  if err != nil {
    panic(err)
  }
  preJson := fmt.Sprintf("{\"name\":\"%s\", \"documents\": []}", name) // Format json
  file.WriteString(preJson) // Write json to file
}
