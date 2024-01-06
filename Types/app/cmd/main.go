package main

import (
	"fmt"

	internal "github.com/jcanonbenavi/app/internal/if"
)

func main() {
	internal.ClasicIfEvenorOdd(4)
	internal.ElseIfNegativeorPositive(-4)
	fmt.Println(internal.PowFunction(3, 2, 10))
	fmt.Printf("Your value for the purchase is $%.2f with a discount applied.\n", internal.IfShortVersionDiscountCalculator(100))
}
