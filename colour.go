package main

type Colour struct {
	r float64
	g float64
	b float64
}

// Equals returns true if the Colour c is equal to the Colour d
func (c Colour) Equals(d Colour) bool {
	return FloatEqual(c.r, d.r) && FloatEqual(c.g, d.g) && FloatEqual(c.b, d.b)
}

// Add returns a Colour created by adding the Colour c to the Colour d
func (c Colour) Add(d Colour) Colour {
	return Colour{
		r: c.r + d.r,
		g: c.g + d.g,
		b: c.b + d.b,
	}
}

// Subtract returns a Colour created by subtracting the Colour d from the Colour c
func (c Colour) Subtract(d Colour) Colour {
	return Colour{
		r: c.r - d.r,
		g: c.g - d.g,
		b: c.b - d.b,
	}
}

// Multiply returns a Colour created by multiplying the Colour c by the scalar s
func (c Colour) Multiply(s float64) Colour {
	return Colour{
		r: c.r * s,
		g: c.g * s,
		b: c.b * s,
	}
}

// Blend returns a Colour created by blending the Colour c with the Colour d. This is otherwise knows as
// the Hadamard or Schur product.
func (c Colour) Blend(d Colour) Colour {
	return Colour{
		r: c.r * d.r,
		g: c.g * d.g,
		b: c.b * d.b,
	}
}
