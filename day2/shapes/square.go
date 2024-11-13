package shapes

// This can be exported
func AreaOfSquare(side float64) float64 {
	return side * side
}

//This can not be exported

func permitere(side float64) float64 {
	return 4 * side
}
