package internal

import "fmt"

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

func GetAll(products *ProductsSlice) {
	for _, i := range *products {
		fmt.Printf("ID: %d, Nombre: %s, Precio: %.2f, Descripción: %s, Categoría: %s\n", i.ID,
			i.Name, i.Price, i.Description, i.Category)
	}
}
