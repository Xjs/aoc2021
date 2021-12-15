package grid

// A Point represents a point on a grid.
type Point struct {
	X, Y uint
}

func P(x, y uint) Point {
	return Point{X: x, Y: y}
}
