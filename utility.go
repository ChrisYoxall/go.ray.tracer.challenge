package main

import "math"

var float64EqualityThreshold = 1e-15

// FloatEqual returns true if the floats are within the equality threshold of each other
func FloatEqual(a, b float64) bool {
	return math.Abs(a-b) < float64EqualityThreshold
}
