package main

import "math"

type Vector struct {
	x float64
	y float64
	z float64
}

// Subtract is the result of subtracting Vector q from Vector v
func (v Vector) Subtract(q Vector) Vector {
	return Vector{
		x: v.x - q.x,
		y: v.y - q.y,
		z: v.z - q.z,
	}
}

// Opposite returns the opposite of Vector v
func (v Vector) Opposite() Vector {
	return Vector{
		x: -v.x,
		y: -v.y,
		z: -v.z,
	}
}

// Multiply returns the result of multiplying Vector v by scalar s
func (v Vector) Multiply(s float64) Vector {
	return Vector{
		x: v.x * s,
		y: v.y * s,
		z: v.z * s,
	}
}

// Magnitude returns the magnitude of Vector v
func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

// Normalise returns a Vector with the same direction as Vector v but with a magnitude of 1.
func (v Vector) Normalise() Vector {
	// Need to divide each component by the magnitude of the vector
	return v.Multiply(1 / v.Magnitude())
}

// DotProduct returns the dot product of Vector v and Vector q.
// The smaller the dot product, the larger the angle between the two vectors.
func (v Vector) DotProduct(q Vector) float64 {
	return v.x*q.x + v.y*q.y + v.z*q.z
}

// CrossProduct returns the cross product of Vector v and Vector q.
// The cross product is a vector perpendicular to both v and q.
func (v Vector) CrossProduct(q Vector) Vector {
	return Vector{
		x: v.y*q.z - v.z*q.y,
		y: v.z*q.x - v.x*q.z,
		z: v.x*q.y - v.y*q.x,
	}
}
