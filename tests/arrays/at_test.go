package arrays_test

import (
	"testing"

	"github.com/w1tchCrafter/arrays/pkg/arrays"
)

func TestAt(t *testing.T) {
	arr := arrays.New(1.2, 102.187, 45.0, 94.3)

	third, err := arr.At(2)
	if third != 45.0 || err != nil {
		t.Errorf("Expected 45.0, got %v (err: %v)", third, err)
	}

	// Test case: Access out of bounds (high index)
	nan, err := arr.At(11)
	if nan != 0 || err == nil {
		t.Errorf("Expected 0 and an error, got %v (err: %v)", nan, err)
	}

	// Test case: Access out of bounds (negative index)
	nan, err = arr.At(-1)
	if nan != 0 || err == nil {
		t.Errorf("Expected 0 and an error for negative index, got %v (err: %v)", nan, err)
	}

	// Test case: Access first element
	first, err := arr.At(0)
	if first != 1.2 || err != nil {
		t.Errorf("Expected 1.2, got %v (err: %v)", first, err)
	}

	// Test case: Access last element
	last, err := arr.At(3)
	if last != 94.3 || err != nil {
		t.Errorf("Expected 94.3, got %v (err: %v)", last, err)
	}
}
