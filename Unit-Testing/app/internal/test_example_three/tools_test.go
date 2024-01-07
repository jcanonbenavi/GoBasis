package tools_test

import (
	"testing"

	internal "github.com/jcanonbenavi/app/internal/test_example_three"
	"github.com/stretchr/testify/require"
)

func TestElementAtIndex_ExistentIndex(t *testing.T) {
	// arrange
	slice := []int{1, 2, 3, 4, 5}
	index := 0

	expected := 1

	// act
	result, err := internal.ElementAtIndex(slice, index)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestElementAtIndex_NonExistingIndex(t *testing.T) {
	// arrange
	slice := []int{1, 2, 3, 4, 5}
	idx := len(slice)

	expected := 0

	// act
	result, err := internal.ElementAtIndex(slice, idx)

	// assert
	require.Error(t, err)
	require.ErrorIs(t, err, internal.ErrIndexOutOfRange)
	require.EqualError(t, err, internal.ErrIndexOutOfRange.Error())
	require.Equal(t, expected, result)
}
