package internal

import "fmt"

type PersonJsonStruct struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address,omitempty"`
	Password  string `json:"-"` //hide
}

func (person PersonJsonStruct) FullName() {
	fmt.Println(person.FirstName, person.LastName)
}
