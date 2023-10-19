package main

import "testing"

// TestColourEquality tests that the Equals method returns true when the colours are equal
func TestColourEquality(t *testing.T) {

	c := Colour{r: 0.9, g: 0.6, b: 0.75}
	d := Colour{r: 0.9, g: 0.6, b: 0.75}

	equal := c.Equals(d)

	if !equal {
		t.Errorf("Colours are equal. Got: %t, want: %t.", equal, true)
	}
}

// TestAddingColours tests that the Add method returns the correct colour
func TestAddingColours(t *testing.T) {

	c := Colour{r: 0.9, g: 0.6, b: 0.75}
	d := Colour{r: 0.7, g: 0.1, b: 0.25}

	result := c.Add(d)

	expected := Colour{r: 1.6, g: 0.7, b: 1.0}

	if !result.Equals(expected) {
		t.Errorf("Colour 'Add' method failed. Got: %v, want: %v.", result, expected)
	}
}

// TestSubtractingColours tests that the Subtract method returns the correct colour
func TestSubtractingColours(t *testing.T) {

	c := Colour{r: 0.9, g: 0.6, b: 0.75}
	d := Colour{r: 0.7, g: 0.1, b: 0.25}

	result := c.Subtract(d)

	expected := Colour{r: 0.2, g: 0.5, b: 0.5}

	if !result.Equals(expected) {
		t.Errorf("Colour 'Subtract' method failed. Got: %v, want: %v.", result, expected)
	}
}

// TestMultiplyingColours tests that the Multiply method returns the correct colour
func TestMultiplyingColours(t *testing.T) {

	c := Colour{r: 0.2, g: 0.3, b: 0.4}

	result := c.Multiply(2)

	expected := Colour{r: 0.4, g: 0.6, b: 0.8}

	if !result.Equals(expected) {
		t.Errorf("Colour 'Multiply' method failed. Got: %v, want: %v.", result, expected)
	}
}

// TestBlendingColours tests that the Blend method returns the correct colour
func TestBlendingColours(t *testing.T) {

	c := Colour{r: 1, g: 0.2, b: 0.4}
	d := Colour{r: 0.9, g: 1, b: 0.1}

	result := c.Blend(d)

	expected := Colour{r: 0.9, g: 0.2, b: 0.04}

	if !result.Equals(expected) {
		t.Errorf("Colour 'Blend' method failed. Got: %v, want: %v.", result, expected)
	}
}
