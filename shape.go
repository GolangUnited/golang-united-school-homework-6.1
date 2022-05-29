package golang_united_school_homework

import (
	"math"
)

type Shape interface {
	// CalcPerimeter returns calculation result of perimeter
	CalcPerimeter() float64
	// CalcArea returns calculation result of area
	CalcArea() float64
}

func (c Circle) CalcPerimeter() float64 {
	return c.Radius * 2 * math.Pi
}

func (r Rectangle) CalcPerimeter() float64 {
	return 2 * (r.Height + r.Weight)
}

func (t Triangle) CalcPerimeter() float64 {
	return 3 * t.Side
}

func (c Circle) CalcArea() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (r Rectangle) CalcArea() float64 {
	return r.Height * r.Weight
}

func (t Triangle) CalcArea() float64 {
	return math.Pow(t.Side, 2) * math.Sqrt(3) / 4
}
