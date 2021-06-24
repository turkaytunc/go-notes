package inter

import (
	"math"
)

type Double float64

type AreaCalculator interface {
	CalculateArea() Double
}

type Circle struct {
	R float64
}

func (c *Circle) CalculateArea() Double {
	return Double(math.Pow(c.R, 2) * math.Pi)
}

type Square struct {
	Height float64
}

func (s *Square) CalculateArea() Double {
	return Double(math.Pow(s.Height, 2))
}
