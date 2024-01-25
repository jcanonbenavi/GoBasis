package main

import (
	"fmt"

	"github.com/jcanonbenavi/app/internal"
	"github.com/jcanonbenavi/app/internal/repository"
)

func main() {
	// bytes, err := internal.Readfile("../customer.txt")
	// fmt.Println(string(bytes))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Ejecución finalizada")

	//initialize map

	defer func() {
		fmt.Println("End of execution")
	}()

	mapCliente := repository.NewClientMap()
	//create a new client
	client := internal.Client{
		ID:    1,
		Name:  "Juan",
		File:  "customer.txt",
		Phone: 123456789,
		Home:  "Calle 123",
	}

	err := mapCliente.Save(&client)
	if err != nil {
		fmt.Println(err)
		return
	}

	clienteTwo := internal.Client{
		Name:  "Juana",
		File:  "customers.txt",
		Phone: 12345678,
		Home:  "Calle 12",
	}
	err = mapCliente.Save(&clienteTwo)

	// clienteThree := internal.Client{
	// 	Name:  "Juana",
	// 	File:  "customers.txt",
	// 	Phone: 12345678,
	// 	Home:  "Calle 12",
	// }
	// err = mapCliente.Save(&clienteThree)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(mapCliente.GetAll())

}
