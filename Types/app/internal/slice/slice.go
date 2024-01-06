package internal

import "fmt"

func SliceExample() {
	// 3: len - 8: cap
	names := make([]string, 3, 4)
	names[0] = "John"
	names[1] = "Paul"
	names[2] = "George"
	fmt.Printf("names: %v, len: %d, cap: %d\n", names, len(names), cap(names))
	names = append(names, "Ringo", "Other") //The capacity doubles when we add a new element. Now is 8
	fmt.Printf("names: %v, len: %d, cap: %d\n", names, len(names), cap(names))
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[1:4]) //4 not included

	//another way - 02-01-2024

	c := [...]float64{1, 2, 3, 4, 5}
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))
}
