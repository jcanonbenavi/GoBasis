package main

import (
	"errors"
	"fmt"

	"github.com/jcanonbenavi/app/internal"
	"github.com/jcanonbenavi/app/internal/repository"
)

func main() {
	movieRepo := repository.NewMovieMap(10)

	mv := internal.Movie{
		Title:  "Batman",
		Year:   1989,
		Rating: 8.5,
	}

	err := movieRepo.Save(&mv)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrLimitReached):
			fmt.Println("code 01: limit reached")
		case errors.Is(err, repository.ErrInvalidMovie):
			fmt.Println("code 02: invalid movie")
		case errors.Is(err, repository.ErrConstraintViolation):
			fmt.Println("code 03: constraint violation")
		default:
			fmt.Println("code 04: unknown error")
		}
		return
	}
	_, salaryerr := internal.Exercises2()
	errWrap := fmt.Errorf("error wrap: %w", salaryerr)
	fmt.Println(errors.Is(errWrap, salaryerr))
	_, salaryerr2 := internal.Exercises3()
	fmt.Println(errors.Is(salaryerr2, internal.ErrorMinimumTaxable))
	err = internal.Exercises4(10000)
	if err != nil {
		println(err.Error())
	}

	salary, err := internal.Exercises5(10, 19)
	if err != nil {
		println(err.Error())
		return
	}
	println(salary)
}
