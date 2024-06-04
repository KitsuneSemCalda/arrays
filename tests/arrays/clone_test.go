package arrays_test

import (
	"testing"

	"github.com/w1tchCrafter/arrays/pkg/arrays"
)

func TestClone(t *testing.T) {
	arr := arrays.New(1, 2, 3, 4, 5)

	// Create a clone of the array
	slc := arr.Clone()

	// Append additional elements to the original array
	arr.Push(6, 7, 8)

	// Check if the length of the clone remains unchanged
	if slc.Len() == arr.Len() {
		t.Errorf("Expected length of clone to be different than %d but the length is the same", arr.Len())
	}

	// Check if the content of the clone remains unchanged
	for i := 0; i < slc.Len(); i++ {
		slcValue, _ := slc.At(i)
		arrValue, _ := slc.At(i)

		if slcValue != arrValue {
			t.Errorf("Content of the clone differs from the original array at index %d", i)
		}
	}
}
