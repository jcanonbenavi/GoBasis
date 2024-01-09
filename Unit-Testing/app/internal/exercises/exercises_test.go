package internal_test

import (
	"testing"

	internal "github.com/jcanonbenavi/app/internal/exercises"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Run("success - case 01: salary > 5000 - 6000 * 0.17 = 1020", func(t *testing.T) {
		salary := 6.000
		result := internal.Load(salary)
		expectedResult := 1.020
		require.Equal(t, expectedResult, result)
	})

	t.Run("success - case 02: salary > 15000 - 20000 * 0.27 = 5400", func(t *testing.T) {
		salary := 20.000
		result := internal.Load(salary)
		expectedResult := 5.400
		require.Equal(t, expectedResult, result)
	})

	t.Run("success - case 03: salary < 5000 - 4000 * 0 = 0", func(t *testing.T) {
		salary := 4.000
		result := internal.Load(salary)
		expectedResult := 0.0
		require.Equal(t, expectedResult, result)
	})
}

func TestAverage(t *testing.T) {
	t.Run("success - 01: average of 5.0, 4.5, 3.5, 3.0, 2.0 ", func(t *testing.T) {
		result, err := internal.Average(5.0, 4.5, 3.5, 3.0, 2.0)
		expectedResult := 3.5999999999999996
		require.Equal(t, result, expectedResult)
		require.Nil(t, err)
	})

	t.Run("failure - 01: average of 5.0, 4.5, 3.5, 3.0, -2.0 ", func(t *testing.T) {
		result, err := internal.Average(5.0, 4.5, 3.5, -3.0, 2.0)
		expectedResult := 0.0
		expectedError := "Not negative numbers, please"
		require.Equal(t, result, expectedResult)
		require.Equal(t, err, expectedError)
	})
}

func TestSalary(t *testing.T) {
	t.Run("success - 01: 120 minutes, category A", func(t *testing.T) {
		result, err := internal.Salary(120, internal.CategoryA)
		expectedResult := 9.000
		require.Equal(t, result, expectedResult)
		require.Nil(t, err)
	})

	t.Run("success - 01: 120 minutes, category B", func(t *testing.T) {
		result, err := internal.Salary(120, internal.CategoryB)
		expectedResult := 3.600
		require.Equal(t, result, expectedResult)
		require.Nil(t, err)
	})

	t.Run("success - 01: 120 minutes, category C", func(t *testing.T) {
		result, err := internal.Salary(120, internal.CategoryC)
		expectedResult := 2.000
		require.Equal(t, result, expectedResult)
		require.Nil(t, err)
	})

	t.Run("failure - 01: 120 minutes, category D", func(t *testing.T) {
		result, err := internal.Salary(120, "D")
		expectedResult := 0.0
		expectedError := "Invalid category"
		require.Equal(t, result, expectedResult)
		require.Equal(t, err, expectedError)
	})

}
