package main

import "testing"

// TestVectorEquality tests that the Equals method returns true when the vectors are equal
func TestVectorEquality(t *testing.T) {

	x := 2.33
	y := 3.0
	z := -4.76

	v := Vector{x: x, y: y, z: z}

	q := Vector{x: x, y: y, z: z}

	equal := v.Equals(q)

	if !equal {
		t.Errorf("Vectors are equal. Got: %t, want: %t.", equal, true)
	}
}

// TestSubtractingVectors tests that the Subtract method returns the correct vector
func TestSubtractingVectors(t *testing.T) {

	v := Vector{x: 3, y: 2, z: 1}
	q := Vector{x: 5, y: 6, z: 7}

	result := v.Subtract(q)

	expected := Vector{x: -2, y: -4, z: -6}

	if !result.Equals(expected) {
		t.Errorf("Vector 'Subtract' method failed. Got: %v, want: %v.", result, expected)
	}
}

// TestOpposite tests that the Opposite method returns the correct vector
func TestOpposite(t *testing.T) {

	v := Vector{x: 1, y: -2, z: 3}

	result := v.Opposite()

	expected := Vector{x: -1, y: 2, z: -3}

	if !result.Equals(expected) {
		t.Errorf("Vector 'Opposite' method failed. Got: %v, want: %v.", result, expected)
	}
}

// TestMultiplyingVectorByScalar tests that the Multiply method returns the correct vector.
func TestMultiplyingVectorByScalar(t *testing.T) {

	v := Vector{x: 1, y: -2, z: 3}

	result := v.Multiply(0.5)

	expected := Vector{x: 0.5, y: -1, z: 1.5}

	if !result.Equals(expected) {
		t.Errorf("Vector 'Multiply' method failed. Got: %v, want: %v.", result, expected)
	}
}

// TestMagnitude tests that the Magnitude method returns the correct magnitude.
func TestMagnitude(t *testing.T) {

	v := Vector{x: 1, y: 2, z: -2}

	result := v.Magnitude()

	expected := float64(3)

	if !FloatEqual(result, expected) {
		t.Errorf("Vector 'Magnitude' method failed. Got: %v, want: %v.", result, expected)
	}
}

// TestNormalise tests that the Normalise method returns the correct vector.
func TestNormalise(t *testing.T) {

	x := float64(1)
	y := float64(2)
	z := float64(3)

	v := Vector{x: x, y: y, z: z}

	result := v.Normalise()

	m := v.Magnitude()
	expected := Vector{x: 1 / m, y: 2 / m, z: 3 / m}

	if !result.Equals(expected) {
		t.Errorf("Vector 'Normalise' method failed. Got: %v, want: %v.", result, expected)
	}

	mag := result.Magnitude()
	if !FloatEqual(mag, float64(1)) {
		t.Errorf("Vector 'Normalise' method failed. The magnitude of a normalised Vector should be 1. Got: %v, want: %v.", mag, 1)
	}
}

// TestDotProduct tests that the DotProduct method returns the correct value.
func TestDotProduct(t *testing.T) {

	v := Vector{x: 1, y: 2, z: 3}
	q := Vector{x: 2, y: 3, z: 4}

	result := v.DotProduct(q)

	expected := float64(20)

	if !FloatEqual(result, expected) {
		t.Errorf("Vector 'DotProduct' method failed. Got: %v, want: %v.", result, expected)
	}
}

// TestCrossProduct tests that the CrossProduct method returns the correct vector.
func TestCrossProduct(t *testing.T) {

	v := Vector{x: 1, y: 2, z: 3}
	q := Vector{x: 2, y: 3, z: 4}

	result := v.CrossProduct(q)

	expected := Vector{x: -1, y: 2, z: -1}

	if !result.Equals(expected) {
		t.Errorf("Vector 'CrossProduct' method failed. Got: %v, want: %v.", result, expected)
	}
}
