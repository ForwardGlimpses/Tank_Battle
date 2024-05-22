package vector2

import "math"

type Vector struct {
	X, Y float64
}

func NewVector(x, y float64) Vector {
	return Vector{X: x, Y: y}
}

func (vec Vector) Add(other Vector) Vector {
	vec.X += other.X
	vec.Y += other.Y
	return vec
}

func (vec Vector) Sub(other Vector) Vector {
	vec.X -= other.X
	vec.Y -= other.Y
	return vec
}

func (vec Vector) Invert() Vector {
	vec.X = -vec.X
	vec.Y = -vec.Y
	return vec
}

func (vec Vector) Magnitude() float64 {
	return math.Sqrt(vec.X*vec.X + vec.Y*vec.Y)
}

func (vec Vector) MagnitudeSquared() float64 {
	return vec.X*vec.X + vec.Y*vec.Y
}

func (vec Vector) DistanceSquared(other Vector) float64 {
	return vec.Sub(other).MagnitudeSquared()
}

func (vec Vector) Unit() Vector {
	l := vec.Magnitude()
	if l < 1e-8 || l == 1 {
		// If it's 0, then don't modify the vector
		return vec
	}
	vec.X, vec.Y = vec.X/l, vec.Y/l
	return vec
}

func (vec Vector) Rotate(angle float64) Vector {
	x := vec.X
	y := vec.Y
	vec.X = x*math.Cos(angle) - y*math.Sin(angle)
	vec.Y = x*math.Sin(angle) + y*math.Cos(angle)
	return vec
}

func (vec Vector) MulScale(scalar float64) Vector {
	vec.X *= scalar
	vec.Y *= scalar
	return vec
}

func (vec Vector) Dot(other Vector) float64 {
	return vec.X*other.X + vec.Y*other.Y
}
