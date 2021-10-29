package turtle

import "fmt"

// A drawing Turtle.
type TurtleDraw struct {
	Turtle // Turtle agent to move around.
	Pen    // Pen used when drawing.

	W *World // World to draw on.
}

// Create a new TurtleDraw, attached to the World w.
func NewTurtleDraw(w *World) *TurtleDraw {
	t := *New()
	p := *NewPen()
	td := &TurtleDraw{t, p, w}
	return td
}

// Move the turtle forward and draw the line if the Pen is On.
func (td *TurtleDraw) Forward(dist float64) {
	x0, y0 := td.X, td.Y
	td.Turtle.Forward(dist)
	x1, y1 := td.X, td.Y
	line := Line{x0, y0, x1, y1, &td.Pen}
	if td.On {
		td.drawLine(line)
	}
}

// Move the turtle backward and draw the line if the Pen is On.
func (td *TurtleDraw) Backward(dist float64) {
	td.Forward(-dist)
}

// Teleport the Turtle to (x, y) and draw the line if the Pen is On.
func (td *TurtleDraw) SetPos(x, y float64) {
	x0, y0 := td.X, td.Y
	td.Turtle.SetPos(x, y)
	x1, y1 := td.X, td.Y
	line := Line{x0, y0, x1, y1, &td.Pen}
	if td.On {
		td.drawLine(line)
	}
}

// Execute the received instruction.
func (td *TurtleDraw) DoInstruction(i Instruction) {
	switch i.Cmd {
	case CmdForward:
		td.Forward(i.Amount)
	case CmdBackward:
		td.Backward(i.Amount)
	case CmdLeft:
		td.Left(i.Amount)
	case CmdRight:
		td.Right(i.Amount)
	}
}

// Write the TurtleDraw state.
func (td *TurtleDraw) String() string {
	sT := td.Turtle.String()
	sP := td.Pen.String()
	return fmt.Sprintf("Turtle: %s Pen: %s", sT, sP)
}

// Send the line to the world and wait for it to be drawn
func (td *TurtleDraw) drawLine(l Line) {
	td.W.DrawLineCh <- l
	<-td.W.doneLineCh
}
