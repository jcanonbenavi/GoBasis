package internal

import (
	"errors"
	"fmt"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

type ProductsSlice []Product

func (p *Product) Save(productSlice *ProductsSlice) {
	*productSlice = append(*productSlice, *p)
}

func GetAll(products ProductsSlice) {
	for _, i := range products {
		fmt.Printf("ID: %d, Nombre: %s, Precio: %.2f, Descripción: %s, Categoría: %s\n", i.ID,
			i.Name, i.Price, i.Description, i.Category)
	}
}

func GetById(id int, products ProductsSlice) (product Product, err error) {
	for _, i := range products {
		if i.ID == id {
			product = i
			return
		}
	}
	err = errors.New("Product not found")
	return
}
