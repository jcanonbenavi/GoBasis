package main

import (
	internal "github.com/jcanonbenavi/app/internal/maps"
)

//internal "github.com/jcanonbenavi/app/internal/array"
//internal "github.com/jcanonbenavi/app/internal/for"
//internal "github.com/jcanonbenavi/app/internal/if"
//internal "github.com/jcanonbenavi/app/internal/switch"

func main() {
	/* if exercises
	internal.ClasicIfEvenorOdd(4)
	internal.ElseIfNegativeorPositive(-4)
	fmt.Println(internal.PowFunction(3, 2, 10))
	fmt.Printf("Your value for the purchase is $%.2f with a discount applied.\n", internal.IfShortVersionDiscountCalculator(100))
	*/

	// internal.SwitchFunction(3)
	// internal.SwitchVariable()
	// internal.SwitchFallthrough("i")
	// internal.SwitchClassExample()
	// internal.SwitchClassExampleTwo()

	//internal.StandarFor()
	// internal.ForLikeWhile()
	// internal.ForWithRange()

	// array
	/*apples := [2]string{"red", "green"}

	// read operation
	println("First apple is", apples[0])
	println("Second apple is", apples[1])

	//update second apple
	apples[1] = "yellow"
	println("Now, the second apple is", apples[1])*/

	//slice: It's similar to an array, but we don't have to specify the size

	// var s = []int{1, 2, 4, 5, 6, 7, 8, 9, 7, 6, 4, 2, 3}
	// fmt.Println(s)
	// //range
	// fmt.Println("Range [1:4]: ", s[1:4])

	// //slice with make
	// a := make([]int, 5) //print: 0 0 0 0 0 - Zero values
	// fmt.Println(a)
	// internal.SliceExample()

	//var MyMap map[int]string
	//println(MyMap) //print a nil and an error
	// var MyMap = map[int]string{}
	// println(MyMap) //print a pointer
	// MyMap[1] = "Value 1"
	// MyMap[2] = "Value 2"
	// fmt.Println(MyMap)
	// fmt.Println(MyMap[1])
	// fmt.Println(len(MyMap))
	// m := make(map[string]int)
	// fmt.Println(m) //print type map

	// //update second value
	// MyMap[2] = "Value 2.0"
	// fmt.Println(MyMap)
	// delete(MyMap, 2)
	// fmt.Println(MyMap)
	// internal.SlicesExampleFor()
	internal.MapAny()

	// internal.ExerciseOne()
	// internal.ExerciseTwo()
	// internal.ExerciseThree()
	// internal.ExerciseFour()

}
