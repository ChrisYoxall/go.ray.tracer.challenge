package main

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
