package internal

import (
	"errors"
)

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
