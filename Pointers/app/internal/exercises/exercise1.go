package exercises

import "fmt"

type Student struct {
	Name            string
	LastName        string
	DNI             int
	InscriptionDate string
}

func (s *Student) Detail() (values string) {
	values = fmt.Sprintf("Student\n-Name: %s\n-LastName: %s\n-DNI: %d\n-InscriptionDate: %s\n", s.Name, s.LastName, s.DNI, s.InscriptionDate)
	return
}
