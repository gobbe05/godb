package types

import (
	"errors"
)

type Collection struct {
  Name string `json:"name"`
  Documents  []map[string]interface{}
}

func (collection *Collection) GetLastId() int {
  if len(collection.Documents) == 0 {
    return 0
  }
  lastDoc := collection.Documents[len(collection.Documents)-1]
  return int(lastDoc["id"].(float64))
}

func (collection *Collection) GetById(id int) (map[string]interface{}, error){
  for _, document := range collection.Documents {
    idInt := int(document["id"].(float64))
    if id == idInt {
      return document, nil
    }
  }
  return make(map[string]interface{}),errors.New("GetById : Document with the inserted id could not be found!")
}

func (collection *Collection) DeleteById(id int) error {
  for i, document := range collection.Documents {
    idInt := int(document["id"].(float64))
    if idInt == id {
      collection.Documents = append(collection.Documents[:i], collection.Documents[i+1:]...)
      return nil
    }
  }
  return errors.New("From DeleteById : Could not find document with id!")
}

func (collection *Collection) UpdateById(id int, d map[string]interface{}) (map[string]interface{}, error) {
  for i, document := range collection.Documents {
    idInt := int(document["id"].(float64))
    if  idInt == id {
      collection.Documents[i] = d
      return collection.Documents[i], nil
    } 
  }
  return make(map[string]interface{}), errors.New("Document with the inserted id could not be found")
}
