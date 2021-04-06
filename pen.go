package turtle

import (
	"fmt"
	"image/color"
)

// A simple Pen.
type Pen struct {
	Color color.Color // Line color.
	Size  int         // Line width.
	On    bool        // State of the Pen.
}

// Create a new Pen.
func NewPen() *Pen {
	p := new(Pen)
	p.Color = White
	p.Size = 3
	return p
}

// Start writing.
func (p *Pen) PenDown() {
	p.On = true
}

// Stop writing.
func (p *Pen) PenUp() {
	p.On = false
}

// Change the Pen color.
func (p *Pen) SetColor(c color.Color) {
	p.Color = c
}

// Change the Pen size.
func (p *Pen) SetSize(s int) {
	p.Size = s
}

// Write the Pen state.
func (p *Pen) String() string {
	return fmt.Sprintf("%v, %d, %t", p.Color, p.Size, p.On)
}
