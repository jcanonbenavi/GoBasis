package internal

import "fmt"

type animal interface {
	MakeSound()
}

type Dog struct {
	Name string
}

func (d Dog) MakeSound() {
	fmt.Println("Gou")
}

type Cat struct {
	Name string
}

func (d Cat) MakeSound() {
	fmt.Println("Miau")
}
