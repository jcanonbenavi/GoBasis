package internal

import "fmt"

func PointerExample() {
	var myVar string = "Original Message"
	var ptr *string = &myVar
	fmt.Printf("My var value: %v\n", myVar)
	fmt.Printf("My var direction: %v\n", &myVar)
	fmt.Printf("Ptr value: %v\n", ptr)
	fmt.Printf("Ptr direction: %v\n", &ptr)

	fmt.Printf("Value pointed to by the pointer (*ptr) %v\n", *ptr)

	*ptr = "New Message"
	fmt.Printf("New value of my var: %v\n", myVar)

}

func Increase(v *int) {
	*v++
	fmt.Printf("Value of v in Increase() %v\n", *v)
}
