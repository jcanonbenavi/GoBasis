package main

import (
	"fmt"

	"github.com/jcanonbenavi/app/internal"
	exercise1 "github.com/jcanonbenavi/app/internal/exercise_1"
	exercise2 "github.com/jcanonbenavi/app/internal/exercise_2"
)

func main() {
	book := internal.Book{
		Title:  "Pride and Prejudice",
		Author: "Jane Austen",
		Pages:  200,
		Publisher: internal.Publisher{
			Name: "Somebody",
			Founders: []string{
				"Somebody",
				"John",
			},
		},
	}

	fmt.Println("Number of founders", book.TotalFounders())
	fmt.Println(book.ShowInfo())

	movie := internal.Movie{
		Title:    "Pride and Prejudice",
		Director: "Joe Wright",
		IMB:      8.8,
		Publisher: internal.Publisher{
			Name: "Somebody",
			Founders: []string{
				"Somebody",
				"John",
				"Pepe",
			},
		},
	}

	fmt.Println("Number of founders", movie.TotalFounders())
	fmt.Println(movie.ShowInfo())

	internal.PointerExample()

	v := 10

	internal.Increase(&v)

	// Create a new person
	person := exercise1.Person{
		FirstName: "John",
		LastName:  "Doe",
	}

	// update the first name
	person.SetFirstName("Jane")

	// print the full name
	println(person.FullName())

	employee := exercise2.Employee{
		Name:     "John Doe",
		Position: "Software Engineer",
		Company: exercise2.Company{
			Name: "Labs L.L.C.",
			Location: exercise2.Location{
				City:    "Buenos Aires",
				Country: "Argentina",
			},
		},
	}

	// print the information of the employee
	employee.Information()

	// create a new location
	newLocation := exercise2.Location{
		City:    "New York",
		Country: "United States",
	}
	employee.Company.MigrateLocation(newLocation)

	// print the information of the employee
	employee.Information()
}
