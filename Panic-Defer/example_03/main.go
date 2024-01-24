package main

import "fmt"

// main.go
func main() {
	// Create a new store
	store := &StorageItemsSlice{}
	// store := &StorageItemsSlice{
	// 	Items: map[int]Item{}, to initialize the map
	// }

	// Add an item to the store
	// generate a panic because the map not initialized
	store.AddItem(&Item{Name: "Item 1", Price: 9.99})
	fmt.Println(store)
}

// ./items/storage.go
type Item struct {
	Name  string
	Price float64
}

type StorageItemsSlice struct {
	Items  map[int]Item
	LastId int
}

func (s *StorageItemsSlice) GetItem(id int) Item {
	return s.Items[id]
}

func (s *StorageItemsSlice) AddItem(item *Item) {
	// validate item
	if item.Name == "" {
		return
	}

	// add item to map
	s.LastId++
	s.Items[s.LastId] = *item
}
