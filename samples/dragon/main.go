package main

import (
	"fmt"

	"github.com/Pitrified/go-turtle"
	"github.com/Pitrified/go-turtle/fractal"
)

func dragonSingle(level int) {

	// receive the instructions here
	instructions := make(chan turtle.Instruction)

	// segment size for the dragon
	segLen := 10.0

	// will produce instructions on the channel
	go fractal.GenerateDragon(level, instructions, segLen)

	// the size of the image
	imgRes := 1080

	// create a new world to draw in
	w := turtle.NewWorld(imgRes, imgRes)

	// create and setup a turtle
	td := turtle.NewTurtleDraw(w)
	td.SetPos(500, 500)
	td.PenDown()
	td.SetColor(turtle.DarkOrange)

	for i := range instructions {
		td.DoInstruction(i)
	}

	outImgName := fmt.Sprintf("dragon_single_%02d_%d.png", level, imgRes)
	w.SaveImage(outImgName)
}

func main() {
	fmt.Println("vim-go")

	// recursion level
	level := 10

	// draw a single Hilbert curve
	dragonSingle(level)
}
