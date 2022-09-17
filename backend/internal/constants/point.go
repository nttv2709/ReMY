package constants

type Point struct {
	X float32
	Y float32
}

func NewPoint(x, y float32) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}
