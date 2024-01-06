package internal

import (
	"fmt"
	"math"
)

func ClasicIfEvenorOdd(number int) {
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}
}

func ElseIfNegativeorPositive(number int) {
	if number < 0 {
		fmt.Printf("%d is negative \n", number)
	} else if number > 0 {
		fmt.Printf("%d is positive \n", number)
	} else {
		fmt.Printf("Is %d \n", number)
	}
}

func PowFunction(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func IfShortVersionDiscountCalculator(amount float64) float64 {
	if amount >= 100 {
		return amount - (amount * 0.10)
	}
	return amount - (amount * 0.5)
}

func ClassExercise(salary float64) {
	if salary <= 3000 {
		fmt.Println("This person must pay taxes")
	} else if salary <= 4000 {
		fmt.Printf("They must pay %4.2f of their salary \n", salary*0.10)
	} else {
		fmt.Printf("They must pay %4.2f of their salary \n", salary*0.15)
	}
}
