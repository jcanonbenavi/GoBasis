package internal

import (
	"errors"
	"fmt"
)

var (
	// ErrorMinimumTaxable is returned when the salary entered does not reach the taxable minimum
	ErrorMinimumTaxable = errors.New("Error: the salary entered does not reach the taxable minimum")
)

func Exercises() (err error) {
	var salary int = 1000
	if salary < 150000 {
		err = errors.New("Error: the salary entered does not reach the taxable minimum")
		return
	}
	return
}

func Exercises2() (salary int, err error) {
	salary = 1000
	if salary < 10000 {
		err = errors.New("Error: salary is less than 10000")
		return
	}
	return
}

func Exercises3() (salary int, err error) {
	salary = 1000
	if salary < 10000 {
		err = ErrorMinimumTaxable
		return
	}
	return
}

func Exercises4(salary int) (err error) {
	if salary < 10000 {
		err = fmt.Errorf("%w. salary: %d", ErrorMinimumTaxable, salary)
		return
	}
	return
}

func Exercises5(hourPerMonth int, valueHour int) (salary float64, err error) {
	if salary > 150.000 {
		salary = float64(hourPerMonth * valueHour)
		salary = salary - (salary * 0.10)
	}
	if hourPerMonth < 80 {
		err = errors.New("Error: the worker cannot have worked less than 80 hours per month")
		return
	}
	salary = float64(hourPerMonth * valueHour)
	return
}
