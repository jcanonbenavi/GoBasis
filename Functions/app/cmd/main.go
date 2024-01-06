package main

import (
	"fmt"

	"github.com/jcanonbenavi/app/internal"
)

func main() {
	fmt.Println(internal.Calculate(50, 80, internal.Multiplication))
	fmt.Println(internal.Square(10))
	//fmt.Println(internal.MyFunctionWithElipsis(2, 4, 5, 3, 3, 5, 6))
	fmt.Println(internal.CalculateOp(internal.AdditionOp, 2, 3, 2, 1, 2, 3, 4, 5, 6))
	result, err := internal.DivisionErrors(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Division result: ", result)
}
