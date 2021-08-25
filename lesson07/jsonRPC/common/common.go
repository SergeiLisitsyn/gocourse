package common

import (
	"errors"
	"fmt"
)

// List struct represents content of item's List as a map of items
type List struct {
	Database map[int]Item
}

// Struct represents item for storing in registr
type Item struct {
	ID   int
	Name string
}

//Struct to register RPCserver
type API int

var id int = 1
var l = NewList()

// Function input from STDIN new items

//Method return List of items
func (api *API) GetList(payrol string, reply *List) error {

	if len(l.Database) == 0 {
		return errors.New("List is empty")
	}
	*reply = *l
	//	fmt.Println("Income List request")
	return nil
}

// Register method add item to the list
func (api *API) Register(name []byte, reply *Item) error {
	getItem := Item{
		ID:   0,
		Name: "existedName",
	}
	// check if items with payloaded name exists in the database
	var ok bool = true
	for _, n := range l.Database {
		if n.Name == string(name) {
			ok = false
			break
		}
	}
	if !ok {
		return fmt.Errorf("Item with name %v already exists", string(name))
	}
	//set reply value
	getItem.ID = len(l.Database) + 1
	getItem.Name = string(name)
	l.Database[getItem.ID] = getItem
	*reply = getItem
	fmt.Println(getItem.Name)
	fmt.Println(l.Database)
	return nil
}

// NewList function return a new instance of List
func NewList() *List {
	return &List{
		Database: make(map[int]Item),
	}
}
