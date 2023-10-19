package main

import "testing"

// TestPointForEquality tests the use of the == operator to confirm Points are equal
func TestPointForEquality(t *testing.T) {

	x := 2.33
	y := 3.0
	z := -4.76

	x1 := Point{x: x, y: y, z: z}

	x2 := Point{x: x, y: y, z: z}

	// Most sites suggest checking for equality for floats by taking the absolute value of the difference
	// of the two floats and checking that they are within a certain tolerance. There is also the Cmp function
	// in the math/big package that can be used to compare floats, refer https://pkg.go.dev/math/big#Float.Cmp.

	equal := x1 == x2

	if !equal {
		t.Errorf("Points are equal. Got: %t, want: %t.", equal, true)
	}
}

// TestPointForInequality tests the use of the == operator to confirm Points are not equal
func TestPointForInequality(t *testing.T) {

	x1 := Point{x: -0.2, y: 987687.3234324, z: -13.70}

	x2 := Point{x: 0, y: 0, z: 0}

	equal := x1 == x2

	if equal {
		t.Errorf("Points are not equal. Got: %t, want: %t.", equal, false)
	}
}

// TestAddingVectorToPoint checks that the Add method returns the correct Point
func TestAddingVectorToPoint(t *testing.T) {

	p := Point{x: 3, y: -2, z: 5}

	v := Vector{x: -2, y: 3, z: 1}

	destination := p.Add(v)

	expected := Point{x: 1, y: 1, z: 6}

	if destination != expected {
		t.Errorf("Point 'Add' method failed. Got: %v, want: %v.", destination, expected)
	}
}

// TestSubtractingVectorFromPoint checks that the Subtract method returns the correct Point
func TestSubtractingVectorFromPoint(t *testing.T) {

	p := Point{x: 3, y: 2, z: 1}

	v := Vector{x: 5, y: 6, z: 7}

	destination := p.Subtract(v)

	expected := Point{x: -2, y: -4, z: -6}

	if destination != expected {
		t.Errorf("Point 'Subtract' method failed. Got: %v, want: %v.", destination, expected)
	}
}

// TestDirectTo checks that the VectorTo method returns the correct Vector to navigate from one Point to another
func TestTestVectorTo(t *testing.T) {

	start := Point{x: 5, y: 6, z: 7}

	end := Point{x: 3, y: 2, z: 1}

	result := start.VectorTo(end)

	expected := Vector{x: -2, y: -4, z: -6}

	if result != expected {
		t.Errorf("Point 'VectorTo' methond failed. Got: %v, want: %v.", result, expected)
	}
}
