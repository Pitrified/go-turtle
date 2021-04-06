package turtle

// A simple Line with a Pen to send around channels.
type Line struct {
	X0, Y0 float64
	X1, Y1 float64
	p      *Pen
}
