package shapes

import "math"

// We create a function name with capital letters so it can be exported
func AreaOfCircle(radius float64) float64 {
	return math.Pi * radius * radius
}

// This can not be exported
func diameter(radius float64) float64 {
	return 2 * radius
}
