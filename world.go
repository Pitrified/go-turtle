package turtle

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

// A world to draw on.
type World struct {
	Image         *image.RGBA
	Width, Height int

	DrawLineCh chan Line
	closeCh    chan bool
}

// Create a new empty World.
func NewWorld(width, height int) *World {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{SoftBlack}, image.Point{0, 0}, draw.Src)
	return NewWorldImage(img)
}

// Create a new World attached to the Image img,
// and start listening on w.DrawLineCh for lines to draw.
func NewWorldImage(img *image.RGBA) *World {
	drawCh := make(chan Line)
	closeCh := make(chan bool)
	w := &World{
		Image:      img,
		Width:      img.Bounds().Max.X,
		Height:     img.Bounds().Max.Y,
		DrawLineCh: drawCh,
		closeCh:    closeCh,
	}
	go w.listen()
	return w
}

// Save output
func (w *World) SaveImage(filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	err = png.Encode(f, w.Image)
	return err
}

// Close the world channels, and stop the listen goroutine.
func (w *World) Close() {
	w.closeCh <- true
}

// listen for draw commands on drawLineCh.
func (w *World) listen() {
	for {
		select {

		// draw the received line
		case line := <-w.DrawLineCh:
			w.drawLine(line)
			w.DrawLineCh <- line

		// close the channels and exit the func
		case <-w.closeCh:
			close(w.closeCh)
			close(w.DrawLineCh)
			return
		}
	}
}

// Draw a line on the image.
func (w *World) drawLine(l Line) {
	x0 := int(l.X0)
	y0 := int(l.Y0)
	x1 := int(l.X1)
	y1 := int(l.Y1)

	// line is vertical
	if x0 == x1 {
		if y0 > y1 {
			y1, y0 = y0, y1
		}
		for i := y0; i <= y1; i++ {
			w.setPoint(x0, i, l.p)
		}
		return
	}

	// line is horizontal
	if y0 == y1 {
		if x0 > x1 {
			x1, x0 = x0, x1
		}
		for i := x0; i <= x1; i++ {
			w.setPoint(i, y0, l.p)
		}
		return
	}

	// line is diagonal, draw it with Bresenham algo
	dx := intAbs(x1 - x0)
	dy := -intAbs(y1 - y0)
	var sx, sy int
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}
	err := dx + dy

	var e2 int
	for {
		w.setPoint(x0, y0, l.p)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 = 2 * err
		if e2 >= dy {
			err += dy
			x0 += sx
		}
		if e2 <= dx {
			err += dx
			y0 += sy
		}
	}
}

// Draw a point on the image.
func (w *World) setPoint(x, y int, p *Pen) {
	// the y in the reference frame of the image
	yr := w.Height - y - 1

	// always draw at least one pixel
	if p.Size <= 1 {
		w.Image.Set(x, yr, p.Color)
		return
	}

	half := p.Size / 2
	before := half
	// if the size is even, remove a pixel from the left/bottom
	// in the cartesian coord
	if p.Size%2 == 0 {
		before = half - 1
	}
	// fill the square
	for i := -before; i <= half; i++ {
		for ii := -before; ii <= half; ii++ {
			// yr-ii because before/half are in cartesian coord
			// so we move to image coord by flipping the y axis
			w.Image.Set(x+i, yr-ii, p.Color)
		}
	}
}