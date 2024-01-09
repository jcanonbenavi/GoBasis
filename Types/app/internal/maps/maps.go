package internal

import "fmt"

//key - value

//there are two ways to do it

//1//
// MyMap := map[string]int{}
// myMap := make(map[string]string)
//The make function allocates and initializes a hash map data structure and returns a map value that points to it.

/*
map[KeyType]ValueType*/

func SlicesExampleFor() {
	var students = map[string]int{"Benjamin": 20, "Nahuel": 26}
	for key, element := range students {
		fmt.Println("Key: ", key, "=>", "Element: ", element)

	}
}

func MapAny() {
	var students = map[int]any{}
	students[1] = "Benjamin"
	students[2] = 26
	students[3] = true
	fmt.Println(students)
}
