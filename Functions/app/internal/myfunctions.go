package internal

import (
	"fmt"
)

func Load(salary float64) float64 {
	if salary > 5000 {
		return salary * 0.17
	} else if salary > 15000 {
		return salary * 0.27
	} else {
		return 0
	}
}

func Average(notes ...float64) (float64, error) {
	var total float64
	for _, note := range notes {
		if note < 0 {
			return note, fmt.Errorf("Not negative numbers, please")
		}
		total = total + note/float64(len(notes))
	}
	return total, nil
}

const (
	CategoryC = "C"
	CategoryB = "B"
	CategoryA = "A"
)

func Salary(minutes int, category string) (float64, error) {
	hours := float64(minutes / 60)
	switch category {
	case CategoryA:
		salary := hours * 3.000
		salary = salary + (salary * 0.50)
		return salary, nil
	case CategoryB:
		salary := hours * 1.500
		salary = salary + (salary * 0.20)
		return salary, nil
	case CategoryC:
		salary := hours * 1.000
		return salary, nil
	default:
		fmt.Println("Invalid category")
		return 0.0, nil
	}
}
