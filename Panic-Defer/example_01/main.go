package main

import "fmt"

// main.go
func main() {
	// Create a new store
	store := &StorageItemsSlice{
		Items: []Item{
			{Name: "Item 1", Price: 9.99},
			{Name: "Item 2", Price: 19.99},
		},
	}

	// Get an item from the store
	item := store.GetItem(2)
	//generate a panic because the index 2 does not exist

	// Print the item
	fmt.Printf("%+v\n", item)
}

// ./items/storage.go
type Item struct {
	Name  string
	Price float64
}

// StorageItemsSlice is a slice of items
type StorageItemsSlice struct {
	Items []Item
}

func (s *StorageItemsSlice) GetItem(id int) Item {
	return s.Items[id]
}
