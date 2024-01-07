package internal

import (
	"fmt"
	"sort"
)

const (
	Minimum    = "minimum"
	AverageTwo = "average"
	Maximum    = "maximum"
)

func Operation(operation string, values ...float64) (float64, error) {
	switch operation {
	case Minimum:
		return minimumfunc(values)
	case AverageTwo:
		return averageFunc(values)
	case Maximum:
		return maxFunc(values)
	default:
		return 0, fmt.Errorf("Undefined operation")
	}
}

func minimumfunc(values []float64) (float64, error) {
	sort.Float64s(values)
	return values[0], nil
}

func maxFunc(values []float64) (float64, error) {
	sort.Float64s(values)
	return values[len(values)-1], nil
}

func averageFunc(values []float64) (float64, error) {
	var total float64
	for _, note := range values {
		if note < 0 {
			return note, fmt.Errorf("Not negative numbers, please")
		}
		total = total + note/float64(len(values))
	}
	return total, nil
}
