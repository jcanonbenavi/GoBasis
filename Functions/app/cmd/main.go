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
	result, err := internal.DivisionErrors(5.0, 2.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Division result: ", result)
	resultAverage, err := internal.Average(5.0, 4.5, 3.5, 3.0, 2.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The average of the student's grades is: %f \n", resultAverage)

	resultSalary, err := internal.Salary(800, internal.CategoryA)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The worker's salary is: %f \n", resultSalary)
	minFunc, err := internal.Operation(internal.Minimum, 2, 3, 3, 4, 10, 2, 4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The minimum of the student's grades is:", minFunc)
	averageFunc, err := internal.Operation(internal.AverageTwo, 2, 3, 3, 4, 1, 2, 4, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The average of the student's grades is:", averageFunc)
	maxFunc, err := internal.Operation(internal.Maximum, 2, 3, 3, 4, 1, 2, 4, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The maximum of the student's grades is:", maxFunc)

	animalDog, err := internal.AnimalsOrchestrator(internal.Dog)
	if err != nil {
		fmt.Println(err)
		return
	}
	resultAnimals, err := animalDog(2)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The amount of necessary food is: %d\n", resultAnimals)

}
