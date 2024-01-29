package main

import "github.com/jcanonbenavi/app/internal"

func main() {

	student_one := internal.NewStudent("Juan", "2021-01-01")
	course_one := internal.NewBootcamp("MELI", "Go")
	course_one.RegisterStudent(student_one)
	course_one.PrintStudents()

	student_two := internal.NewStudent("Pedro", "2021-01-01")
	course_one.RegisterStudent(student_two)
	course_one.PrintStudents()

}
