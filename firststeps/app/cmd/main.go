package main

import "fmt"

func main() {
	Name := "Alexandra"
	Age := 29

	fmt.Printf("My name is %s, and I'm %d years old \n", Name, Age)

	Blue := "blue"
	Pink := "pink"
	Black := "black"

	fmt.Printf("My favorite colors are %s, %s, and %s \n", Blue, Pink, Black)

	BookName := "Harry Potter"
	Release := 1997
	Readed := false

	fmt.Printf("The book title is %s and it was published on %d, but I haven't (%t) read the book \n", BookName, Release, Readed)

	State := "Querétaro"
	Latitude := -100.3880600
	Length := 20.5880600

	fmt.Printf("The state of %s has a latitude of %e and a length of %e \n", State, Latitude, Length)

	/*
			var student1 string = 23
		  var grade float64 = "A"
		  var isEnrolled = "yes"
		  var number_of_subjects int
		  number_of_subjects := 5
	*/

	//My corrections

	var student1 int = 23 //23 isn't a string
	fmt.Println(student1)

	var grade string = "A" //"A" isn't a float64
	fmt.Println(grade)

	var isEnrolled = true //Maybe with "yes" it can work but it's easier with a boolean
	fmt.Println(isEnrolled)

	var NumberOfSubjects int //I modified the variable's name because, in Go, as a convention,
	// we don't use underscores. We use camelcase
	NumberOfSubjects = 5 // I changed the := because we didn't declare a new variable; we assigned a new value.
	fmt.Println(NumberOfSubjects)

}
