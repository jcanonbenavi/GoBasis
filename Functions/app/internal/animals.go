package internal

import "fmt"

type Animal string

const (
	Dog Animal = "Dog"
	Cat Animal = "Cat"
)

type Animals func(int) (int, error)

func AnimalsOrchestrator(operator Animal) (a Animals, err error) {
	switch operator {
	case Dog:
		a = dogFunc
	case Cat:
		a = catFunc
	default:
		return nil, fmt.Errorf("No such type of animal was found in the list")
	}
	return
}

func dogFunc(amount int) (food int, err error) {
	food = amount * 10
	return
}

func catFunc(amount int) (food int, err error) {
	food = amount * 5
	return
}
