package tests

import (
	"testing"
	"goai/core"
)

func TestSimpleComputation(t *testing.T) {
	result := core.SimpleComputation(3, 4)
	expected := 12

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
