package main

import "math"

type Vector struct {
	x float64
	y float64
	z float64
}

// Equals returns true if the Vector v is equal to the Vector w
func (v Vector) Equals(w Vector) bool {
	return FloatEqual(v.x, w.x) && FloatEqual(v.y, w.y) && FloatEqual(v.z, w.z)
}

// Subtract is the result of subtracting Vector v from Vector w
func (v Vector) Subtract(w Vector) Vector {
	return Vector{
		x: v.x - w.x,
		y: v.y - w.y,
		z: v.z - w.z,
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

// DotProduct returns the dot product of Vector v and Vector w.
// The smaller the dot product, the larger the angle between the two vectors.
func (v Vector) DotProduct(w Vector) float64 {
	return v.x*w.x + v.y*w.y + v.z*w.z
}

// CrossProduct returns the cross product of Vector v and Vector w.
// The cross product is a vector perpendicular to both v and w.
func (v Vector) CrossProduct(w Vector) Vector {
	return Vector{
		x: v.y*w.z - v.z*w.y,
		y: v.z*w.x - v.x*w.z,
		z: v.x*w.y - v.y*w.x,
	}
}
