package turtle

import (
	"fmt"
	"math"
)

// A minimal Turtle agent, moving on a cartesian plane.
//
// https://en.wikipedia.org/wiki/Turtle_graphics
type Turtle struct {
	X, Y float64 // Position.
	Deg  float64 // Orientation in degrees.
}

// Create a new Turtle.
func New() *Turtle {
	return new(Turtle)
}

// Move the Turtle forward by dist.
func (t *Turtle) Forward(dist float64) {
	rad := Deg2rad(t.Deg)
	t.X += dist * math.Cos(rad)
	t.Y += dist * math.Sin(rad)
}

// Move the Turtle backward by dist.
func (t *Turtle) Backward(dist float64) {
	t.Forward(-dist)
}

// Rotate the Turtle counter clockwise by deg degrees.
func (t *Turtle) Left(deg float64) {
	t.Deg += deg
}

// Rotate the Turtle clockwise by deg degrees.
func (t *Turtle) Right(deg float64) {
	t.Deg -= deg
}

// Teleport the Turtle to (x, y).
func (t *Turtle) SetPos(x, y float64) {
	t.X = x
	t.Y = y
}

// Orient the Turtle towards deg.
func (t *Turtle) SetHeading(deg float64) {
	t.Deg = deg
}

// Execute the received instruction.
func (t *Turtle) DoInstruction(i Instruction) {
	switch i.Cmd {
	case CmdForward:
		t.Forward(i.Amount)
	case CmdBackward:
		t.Backward(i.Amount)
	case CmdLeft:
		t.Left(i.Amount)
	case CmdRight:
		t.Right(i.Amount)
	}
}

// Write the Turtle state.
func (t *Turtle) String() string {
	return fmt.Sprintf("(%9.4f, %9.4f) ^ %9.4f", t.X, t.Y, t.Deg)
}
