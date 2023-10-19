package main

import "testing"

// TestSubtractingVectors tests that the Subtract method returns the correct vector
func TestSubtractingVectors(t *testing.T) {

	v := Vector{3, 2, 1}
	q := Vector{5, 6, 7}

	result := v.Subtract(q)

	expected := Vector{-2, -4, -6}

	if result != expected {
		t.Errorf("Vector 'Subtract' method failed. Got: %v, want: %v.", result, expected)
	}
}

// TestOpposite tests that the Opposite method returns the correct vector
func TestOpposite(t *testing.T) {

	v := Vector{1, -2, 3}

	result := v.Opposite()

	expected := Vector{-1, 2, -3}

	if result != expected {
		t.Errorf("Vector 'Opposite' method failed. Got: %v, want: %v.", result, expected)
	}
}
