package exercise1

type Person struct {
	// FirstName is the first name of the person
	FirstName string
	// LastName is the last name of the person
	LastName string
}

// FullName returns the full name of the person
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// SetFirstName sets the first name of the person
func (p *Person) SetFirstName(firstName string) {
	(*p).FirstName = firstName
}
