package main

import (
	internal "github.com/jcanonbenavi/app/internal/Exercises"
)

func main() {
	// result := internal.Author{
	// 	Name:      "Jorge",
	// 	Language:  "Go",
	// 	Marticles: 10,
	// 	Pay:       100,
	// 	Salary: func(Ma int, Pay int) int {
	// 		return Ma * Pay
	// 	},
	// }
	// ShowTime := internal.Nameoperation{
	// 	Name: "ShowTime",
	// 	Do: func() {
	// 		println("ShowTime")
	// 	},
	// }
	// ShowTime.Do()
	// ShowDate := internal.Nameoperation{
	// 	Name: "ShowDate",
	// 	Do: func() {
	// 		print("ShowDate\n")
	// 	},
	// }
	// ShowDate.Do()
	// fmt.Println(result)

	// //var p internal.Person
	// //fmt.Printf("%#v\n", p) //0 values

	// /*________________________*/
	// person_one := internal.Person{
	// 	Name:       "Juan",
	// 	Gender:     "Men",
	// 	Age:        32,
	// 	Profession: "IT",
	// 	Likes: internal.Preferences{
	// 		Foods:  "Chicken",
	// 		Movies: "Harry Potter",
	// 		Series: "Doctor House",
	// 		Animes: "Dragon Ball",
	// 		Sports: "Futbol",
	// 	},
	// 	Address: internal.Address{
	// 		Street:       "Monte Calvario",
	// 		Neighborhood: "Universidades",
	// 		City:         "Querétaro",
	// 	},
	// }

	// person_two := internal.Person{
	// 	Name:       "Alexandra",
	// 	Gender:     "Women",
	// 	Age:        29,
	// 	Profession: "Software developer",
	// 	Likes: internal.Preferences{
	// 		Foods:  "Pizza",
	// 		Movies: "Harry Potter",
	// 		Series: "Doctor House",
	// 		Animes: "Digimon",
	// 		Sports: "Futbol",
	// 	},
	// 	Address: internal.Address{
	// 		Street:       "Monte Calvario",
	// 		Neighborhood: "Universidades",
	// 		City:         "Querétaro",
	// 	},
	// }

	// fmt.Println(person_one)
	// fmt.Println(person_two)

	// //fmt.Println(person_one.Age) print the age of the person one

	// /* Updating values
	// person_two.Likes.Movies = "N/A"
	// fmt.Println(person_two)
	// person_two.Likes = internal.Preferences{"Pizza", "N/A", "N/A", "N/A", "N/A"}
	// fmt.Println(person_two) */

	// myPerson := internal.PersonJsonStruct{
	// 	FirstName: "Maya",
	// 	LastName:  "Canon",
	// 	Address:   "",         //not value
	// 	Password:  "password", //hide
	// }
	// personAsJson, err := json.Marshal(myPerson)
	// fmt.Println(string(personAsJson))
	// fmt.Println(err) //nill
	// //myPerson.FullName()

	// /* ____________Composition exercise____________	*/

	// circle := internal.Circle{
	// 	Radius: 5.0,
	// }
	// fmt.Println("Value of the area", circle.Area())
	// ptr := &circle.Radius
	// circle.SetRadius(ptr, 9.0)
	// //fmt.Println(*ptr)
	// //*ptr = 7.0
	// fmt.Println("New value of the radius", circle.Radius)
	// fmt.Println("Value of the area", circle.Area())

	// /*____________________*/

	// myCar := internal.Car{
	// 	Engine:        internal.Engine{Horsepower: 250, Type: "V4"},
	// 	Chasis:        internal.Chasis{Material: "Steel"},
	// 	Bodywork:      internal.Bodywork{Color: "Red"},
	// 	NumberOfDoors: 4,
	// }

	// myMotorcycle := internal.Motorcycle{
	// 	Engine:   internal.Engine{Horsepower: 150, Type: "V2"},
	// 	Chasis:   internal.Chasis{Material: "Aluminum"},
	// 	Bodywork: internal.Bodywork{Color: "Black"},
	// }

	// myCar.Start()
	// myCar.DisplayInfo()

	// myMotorcycle.Start()
	// myMotorcycle.DisplayInfo()

	// /*fmt.Println(internal.NewHuman(
	// 	"", 0,
	// ))

	// fmt.Println(internal.NewHuman(
	// 	"Alexa", 29,
	// )) */

	// ProductOne := internal.Product{
	// 	ID:          1,
	// 	Name:        "Product One",
	// 	Price:       100.0,
	// 	Description: "Product One Description",
	// 	Category:    "Category One",
	// }
	// ProductTwo := internal.Product{
	// 	ID:          2,
	// 	Name:        "Product Two",
	// 	Price:       200.0,
	// 	Description: "Product Two Description",
	// 	Category:    "Category Two",
	// }
	// productArray := internal.ProductsSlice{}

	// ProductOne.Save(&productArray)
	// ProductTwo.Save(&productArray)
	// internal.GetAll(productArray)
	// id, err := internal.GetById(6, productArray)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(id)

	employee := internal.Employee{
		ID:       1,
		Position: "Position One",
		Person: internal.Person{
			ID:          1,
			Name:        "Person One",
			DateOfBirth: "01/01/2000",
		},
	}
	employee.PrintEmployee()
}
