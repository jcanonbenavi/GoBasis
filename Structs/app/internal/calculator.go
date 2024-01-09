package internal

type Operation struct {
	Name   string
	Input  [2]float64
	Output float64
}

type Calculator struct {
	Count     int
	Limit     int
	Operation []Operation
}

func (c *Calculator) Add(a, b float64) (result float64, err string) {
	if c.Count >= c.Limit {
		err = "Limit reached"
		return
	}
	result = a + b

	op := Operation{
		Name:   "Add",
		Input:  [2]float64{a, b},
		Output: result,
	}
	c.Operation = append(c.Operation, op)
	return
}
