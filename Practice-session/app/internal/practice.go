package internal

import "fmt"

func NewStudent(name string, entryDate string) Student {
	return Student{
		Name:      name,
		EntryDate: entryDate,
	}
}

func NewBootcamp(organization string, language string) Bootcamp {
	return Bootcamp{
		Organization: organization,
		Language:     language,
	}
}

type Student struct {
	Name      string
	EntryDate string
}

type Bootcamp struct {
	Organization string
	Language     string
	// TODO: Add Students field
	Students []Student
}

// Pointer to Bootcamp struct to modify the struct
func (b *Bootcamp) RegisterStudent(s Student) {
	// Bootcamp struct has a Students field that is a slice of Student structs
	// Append the new student to the Students slice
	(*b).Students = append((*b).Students, s)
}

func (b *Bootcamp) PrintStudents() {
	fmt.Printf(
		"Welcome to the %s bootcamp, this bootcamp is language %s and we have %d students.\n",
		b.Organization,
		b.Language,
		len(b.Students),
	)
}
