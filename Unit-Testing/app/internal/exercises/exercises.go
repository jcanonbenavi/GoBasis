package internal

import (
	"errors"
	"fmt"
)

const (
	CategoryC = "C"
	CategoryB = "B"
	CategoryA = "A"
)

func Salary(minutes int, category string) (salary float64, err error) {
	hours := float64(minutes / 60)
	switch category {
	case CategoryA:
		salary = hours * 3.000
		salary = salary + (salary * 0.50)
		return
	case CategoryB:
		salary = hours * 1.500
		salary = salary + (salary * 0.20)
		return
	case CategoryC:
		salary = hours * 1.000
		return
	default:
		err = fmt.Errorf("Invalid category")
		return
	}
}

func Load(salary float64) (total float64) {
	if salary > 5.000 && salary < 15.000 {
		total = salary * 0.17
		return
	} else if salary > 15.000 {
		total = salary * 0.27
		return
	} else {
		total = 0
		return
	}
}

func Average(notes ...float64) (calcute float64, err error) {
	var total float64
	for _, note := range notes {
		if note < 0 {
			err = errors.New("Not negative numbers, please")
			return
		}
		total = total + note/float64(len(notes))
	}
	return total, nil
}
