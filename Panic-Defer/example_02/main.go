package main

// main.go
func main() {
	// Create a new store
	store := &StorageItemsSlice{
		Items: []Item{
			{Name: "Item 1", Price: 9.99},
			{Name: "Item 2", Price: 19.99},
		},
	}

	// Add an item to the store
	// generate a panic because the item is nil
	store.AddItem(nil)
}

// ./items/storage.go
type Item struct {
	Name  string
	Price float64
}

type StorageItemsSlice struct {
	Items []Item
}

func (s *StorageItemsSlice) GetItem(id int) Item {
	return s.Items[id]
}

// new function to add an item to the store
func (s *StorageItemsSlice) AddItem(item *Item) {
	// validate item
	if (*item).Name == "" {
		return
	}

	// add item to slice
	s.Items = append(s.Items, *item)
}
