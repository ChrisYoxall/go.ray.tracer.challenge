package main

type Point struct {
	x float64
	y float64
	z float64
}

// Add returns a Point describing the destination arrived at by adding the Vector v to the Point p
func (p Point) Add(v Vector) Point {
	return Point{
		x: p.x + v.x,
		y: p.y + v.y,
		z: p.z + v.z,
	}
}

// Subtract returns a Point describing the destination arrived at by subtracting the Vector v from the Point p
func (p Point) Subtract(v Vector) Point {
	return Point{
		x: p.x - v.x,
		y: p.y - v.y,
		z: p.z - v.z,
	}
}

// VectorTo returns a Vector that points from the Point p to the Point q
func (p Point) VectorTo(q Point) Vector {
	return Vector{
		x: q.x - p.x,
		y: q.y - p.y,
		z: q.z - p.z,
	}
}
