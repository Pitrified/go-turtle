package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/Pitrified/go-turtle"
)

func squiggly(td *turtle.TurtleDraw) {
	td.SetPos(100, 300)
	td.SetHeading(turtle.East + 80)
	td.PenDown()

	// must be a float64 distance if used in a variable
	// magic untyped constants allows for td.Forward(50)
	// same for pretty much every value, they are all floats and are converted
	// to int as late as possible (when drawing the line)
	segLen := 150.0
	td.SetSize(1)
	td.Forward(segLen)

	td.SetColor(turtle.Red)
	td.Right(160)
	td.SetSize(2)
	td.Forward(segLen)

	td.SetColor(turtle.Green)
	td.Left(160)
	td.SetSize(3)
	td.Forward(segLen)

	td.SetColor(turtle.Blue)
	td.Right(160)
	td.SetSize(4)
	td.Forward(segLen)

	td.SetColor(turtle.Cyan)
	td.Left(160)
	td.SetSize(5)
	td.Forward(segLen)

	td.SetColor(turtle.Magenta)
	td.Right(160)
	td.SetSize(6)
	td.Forward(segLen)

	td.SetColor(turtle.Yellow)
	td.Left(160)
	td.SetSize(7)
	td.Forward(segLen)

	td.SetColor(color.RGBA{30, 200, 100, 255})
	td.Right(160)
	td.SetSize(8)
	td.Forward(segLen)
}

func circle(td *turtle.TurtleDraw) {
	// move somewhere else
	td.PenUp()
	td.SetPos(450, 300)

	// draw a circle with increasing brightness
	td.PenDown()
	td.SetHeading(turtle.North)
	td.SetSize(5)
	for i := 0; i < 360; i++ {
		val := uint8(float64(i) * 255 / 360)
		td.SetColor(color.RGBA{val, val / 2, 0, 255})
		td.Right(1)
		td.Forward(3)
	}
}

// Forward(0) draws the point on the current position
func dot(td *turtle.TurtleDraw, x, y float64, s int) {
	td.PenUp()
	td.SetPos(x, y)
	td.PenDown()

	td.SetSize(s)
	td.SetColor(turtle.White)
	td.Forward(0)

	td.SetSize(1)
	td.SetColor(turtle.Green)
	td.Forward(0)
}

func dots(td *turtle.TurtleDraw) {
	td.SetHeading(turtle.North)
	for i := 1; i < 10; i++ {
		dot(td, float64(10*i+120), 0, i)
		dot(td, float64(10*i+120), 30, i)
	}
}

func general() {
	// create a new world to draw in
	w := turtle.NewWorld(900, 600)

	// create and setup a turtle
	td := turtle.NewTurtleDraw(w)
	fmt.Println("TD:", td)

	// draw a squiggly line
	squiggly(td)

	// draw a circle with increasing brightness
	circle(td)

	// draw dots to show how the points are drawn
	dots(td)

	// save the current image
	err := w.SaveImage("world.png")
	if err != nil {
		fmt.Println("Could not save the image: ", err)
	}

	// close the world (you might want to defer this)
	w.Close()

	// this is an error: the turtle tries to send the line
	// to the world input channel that has been closed
	// td.Forward(50)
}

func constructor() {

	// pass an image to the world
	img := image.NewRGBA(image.Rect(0, 0, 900, 600))
	draw.Draw(img, img.Bounds(), &image.Uniform{turtle.Cyan}, image.Point{0, 0}, draw.Src)
	wi := turtle.NewWorldWithImage(img)
	defer wi.Close()
	tdi := turtle.NewTurtleDraw(wi)
	circle(tdi)
	err := wi.SaveImage("cyan_world.png")
	if err != nil {
		fmt.Println("Could not save the image: ", err)
	}

	// create a world with color
	wc := turtle.NewWorldWithColor(900, 600, turtle.Yellow)
	defer wc.Close()
	tdc := turtle.NewTurtleDraw(wc)
	circle(tdc)
	err = wc.SaveImage("yellow_world.png")
	if err != nil {
		fmt.Println("Could not save the image: ", err)
	}

}

func reset() {
	// create a World and draw something
	w := turtle.NewWorld(900, 600)
	td := turtle.NewTurtleDraw(w)
	squiggly(td)

	// reset with same size
	w.ResetImage()
	squiggly(td)
	_ = w.SaveImage("resetImage.png")

	// reset with new size
	w.ResetImageWithSize(1200, 1200)
	squiggly(td)
	_ = w.SaveImage("resetSize.png")

	// reset with new size and color
	w.ResetImageWithSizeColor(1200, 1200, turtle.Magenta)
	squiggly(td)
	_ = w.SaveImage("resetSizeColor.png")

	// reset with new custom image
	img := image.NewRGBA(image.Rect(0, 0, 600, 900))
	draw.Draw(img, img.Bounds(), &image.Uniform{turtle.Cyan}, image.Point{0, 0}, draw.Src)
	w.ResetImageWithImage(img)
	squiggly(td)
	_ = w.SaveImage("resetImageImage.png")
}

func main() {
	general()
	constructor()
	reset()
}
