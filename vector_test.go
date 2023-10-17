package main

import "testing"

// TestSubtractingVectors tests that the Difference method of the Vector type returns the correct vector
func TestSubtractingVectors(t *testing.T) {

	v := Vector{3, 2, 1}
	q := Vector{5, 6, 7}

	result := v.Difference(q)

	expected := Vector{-2, -4, -6}

	if result != expected {
		t.Errorf("Vector 'Difference' method failed. Got: %v, want: %v.", result, expected)
	}
}
