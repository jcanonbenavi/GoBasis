package internal

import "errors"

const (
	Addition       = "+"
	Subtraction    = "-"
	Multiplication = "*"
	Division       = "/"
)

// both are float
func Calculate(a, b float64, operation string) float64 {
	switch operation {
	case Addition:
		return a + b
	case Subtraction:
		return a - b
	case Multiplication:
		return a * b
	case Division:
		if b != 0 {
			return a / b
		}
	}
	return 0
}

func Square(n int) (result int) {
	result = n * n
	return
}

func MyFunctionWithElipsis(values ...float64) float64 {
	var total float64
	for _, value := range values {
		total += value
	}
	return total
}

func DivisionErrors(value1, value2 float64) (float64, error) {
	if value2 == 0 {
		return 0, errors.New("Can not divide by zero")
	}
	return value1 / value2, nil
}

func ClosureFunction() func(firstName, lastName string) (fullname string) {
	// separator
	separator := "-"

	return func(firstName, lastName string) (fullname string) {
		// default values
		defaultFirstName := "John"
		defaultLastName := "Doe"
		if firstName != "" {
			defaultFirstName = firstName
		}
		if lastName != "" {
			defaultLastName = lastName
		}

		fullname = defaultFirstName + separator + defaultLastName
		return
	}
}
