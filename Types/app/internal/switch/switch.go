package internal

import (
	"fmt"
	"time"
)

func SwitchFunction(num int) {
	switch num {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}
}
func SwitchFunctionMultipleCases(num int) {
	switch {
	case num > 0, num == 0:
		fmt.Println("Is positive or 0")
	case num < 0:
		fmt.Println("Is negative")
	default:
		fmt.Println("Invalid")
	}
}

func SwitchVariable() {
	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("Good Morning!", time.Hour())
	case time.Hour() < 17:
		fmt.Println("Good afternoon!", time.Hour())
	default:
		fmt.Println("Good evening!", time.Hour())

	}
}

func SwitchFallthrough(letter string) {
	switch {
	case letter == "a":
		fmt.Println("a")
		fallthrough
	case letter == "e":
		fmt.Println("e")
		fallthrough
	case letter == "i":
		fmt.Println("i")
		fallthrough
	case letter == "o":
		fmt.Println("o")
		fallthrough
	case letter == "u":
		fmt.Println("u")
	default:
		fmt.Println("Is not a vocal")
	}

}

func SwitchClassExample() {
	//switch without condition
	var age uint8 = 18
	switch {
	case age >= 150:
		fmt.Println("Are you inmortal?")
	case age >= 18:
		fmt.Println("You are an adult")
	default:
		fmt.Println("You're not an adult yet")
	}
}

func SwitchClassExampleTwo() {
	day := "sunday"
	switch day {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Printf("%s is a workday\n", day)
	default:
		fmt.Printf("%s is a weekend day\n", day)
	}
}
