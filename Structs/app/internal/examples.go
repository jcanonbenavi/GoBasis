package internal

type FinalSalary func(int, int) int

type Author struct {
	Name      string
	Language  string
	Marticles int
	Pay       int
	Salary    FinalSalary
}

type Nameoperation struct {
	Name string
	Do   func()
}
