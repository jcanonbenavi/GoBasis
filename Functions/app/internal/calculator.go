package internal

const (
	AdditionOp       = "+"
	SubtractionOp    = "-"
	MultiplicationOp = "*"
	DivisionOp       = "/"
)

func opSummatory(value1, value2 float64) float64 {
	return value1 + value2
}

func opSubtraction(value1, value2 float64) float64 {
	return value1 - value2
}

func opMultiplication(value1, value2 float64) float64 {
	return value1 * value2
}

func opDivision(value1, value2 float64) float64 {
	if value2 == 0 {
		return 0
	}
	return value1 / value2
}

func CalculateOp(operator string, values ...float64) float64 {
	switch operator {
	case AdditionOp:
		return operationsOrchestrator(values, opSummatory)
	case SubtractionOp:
		return operationsOrchestrator(values, opSubtraction)
	case MultiplicationOp:
		return operationsOrchestrator(values, opMultiplication)
	case DivisionOp:
		return operationsOrchestrator(values, opDivision)
	}
	return 0
}

func operationsOrchestrator(values []float64, operation func(value1, value2 float64) float64) float64 {
	var result float64
	for _, value := range values {
		result = operation(result, value)
	}
	return result
}
