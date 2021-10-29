package main

import (
	"fmt"
	"image/color"

	"github.com/Pitrified/go-turtle"
	"github.com/Pitrified/go-turtle/fractal"
)

func dragonSingle(level int) {

	// receive the instructions here
	instructions := make(chan string)

	// will produce instructions on the channel
	go fractal.GenerateDragon(level, instructions)

	// the size of the image
	imgRes := 1080

	// segment size for the dragon
	segLen := 10.0

	// create a new world to draw in
	w := turtle.NewWorld(imgRes, imgRes)

	// create and setup a turtle
	td := turtle.NewTurtleDraw(w)
	td.SetPos(500, 500)
	td.PenDown()
	td.SetColor(color.RGBA{150, 75, 0, 255})

	for cmd := range instructions {
		switch cmd {
		case "F":
			td.Forward(segLen)
		case "R":
			td.Right(90)
			// td.Right(60)
		case "L":
			td.Left(90)
			// td.Left(60)
		}
	}

	// outImgName := fmt.Sprintf("dragon_single_%02d_%d.png", level, imgRes)
	outImgName := "serp.png"
	w.SaveImage(outImgName)
}

func main() {
	fmt.Println("vim-go")

	// recursion level
	level := 10

	// draw a single Hilbert curve
	dragonSingle(level)
}
