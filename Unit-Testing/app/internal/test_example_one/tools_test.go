package internal_test

import (
	"testing"

	internal "github.com/jcanonbenavi/app/internal/test_example_one"
)

func TestElementAtIndex(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	index := 0

	expected := 1

	result := internal.ElementAtIndex(slice, index)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
		return
	}
	t.Log("success")
}

func TestElementAtIndex_NonExistingIndex(t *testing.T) {
	// arrange
	slice := []int{1, 2, 3, 4, 5}
	idx := len(slice)

	expected := 5

	// act
	result := internal.ElementAtIndex(slice, idx)

	// assert
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
		return
	}
}
