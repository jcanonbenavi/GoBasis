package internal

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e *Employee) PrintEmployee() {
	println(e.ID, e.Position, e.Person.ID, e.Person.Name, e.Person.DateOfBirth)
}
