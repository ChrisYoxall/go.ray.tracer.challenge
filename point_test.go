package main

import "testing"

// TestPointEquality tests that the Equals method returns true when the points are equal
func TestPointForEquality(t *testing.T) {

	x := 2.33
	y := 3.0
	z := -4.76

	x1 := Point{x: x, y: y, z: z}

	x2 := Point{x: x, y: y, z: z}

	equal := x1.Equals(x2)

	if !equal {
		t.Errorf("Points are equal. Got: %t, want: %t.", equal, true)
	}
}

// TestAddingVectorToPoint checks that the Add method returns the correct Point
func TestAddingVectorToPoint(t *testing.T) {

	p := Point{x: 3, y: -2, z: 5}

	v := Vector{x: -2, y: 3, z: 1}

	destination := p.Add(v)

	expected := Point{x: 1, y: 1, z: 6}

	if !destination.Equals(expected) {
		t.Errorf("Point 'Add' method failed. Got: %v, want: %v.", destination, expected)
	}
}

// TestSubtractingVectorFromPoint checks that the Subtract method returns the correct Point
func TestSubtractingVectorFromPoint(t *testing.T) {

	p := Point{x: 3, y: 2, z: 1}

	v := Vector{x: 5, y: 6, z: 7}

	destination := p.Subtract(v)

	expected := Point{x: -2, y: -4, z: -6}

	if !destination.Equals(expected) {
		t.Errorf("Point 'Subtract' method failed. Got: %v, want: %v.", destination, expected)
	}
}

// TestDirectTo checks that the VectorTo method returns the correct Vector to navigate from one Point to another
func TestTestVectorTo(t *testing.T) {

	start := Point{x: 5, y: 6, z: 7}

	end := Point{x: 3, y: 2, z: 1}

	result := start.VectorTo(end)

	expected := Vector{x: -2, y: -4, z: -6}

	if !result.Equals(expected) {
		t.Errorf("Point 'VectorTo' methond failed. Got: %v, want: %v.", result, expected)
	}
}
