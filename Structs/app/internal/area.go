package internal

import (
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

/*its a copy*
func (c Circle) SetRadius(newRadius float64) {
	c.Radius = newRadius
} */

func (c Circle) SetRadius(newRadius *float64, value float64) {
	*newRadius = value
}
