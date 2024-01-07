package internal_test

import (
	"testing"

	internal "github.com/jcanonbenavi/app/internal/test_example_two"
)

func TestElementAtIndex_ExistentIndex(t *testing.T) {
	// arrange
	slice := []int{1, 2, 3, 4, 5}
	index := 9

	expected := 5

	// act
	result, err := internal.ElementAtIndexVersionTwo(slice, index)

	// assert
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
	t.Log("success")
}
