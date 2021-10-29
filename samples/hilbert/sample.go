package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/Pitrified/go-turtle"
	"github.com/Pitrified/go-turtle/fractal"
)

func hilbertSingle(level int) {

	// receive the instructions here
	instructions := make(chan turtle.Instruction)

	// the size of the image
	imgRes := 1080

	// the step for each Hilbert curve segment
	pad := 80
	segLen := getSegmentLen(level, imgRes-pad)

	// will produce instructions on the channel
	go fractal.GenerateHilbert(level, instructions, segLen)

	// create a new world to draw in
	w := turtle.NewWorld(imgRes, imgRes)

	// create and setup a turtle
	td := turtle.NewTurtleDraw(w)
	td.SetPos(40, 40)
	td.PenDown()
	td.SetColor(color.RGBA{150, 75, 0, 255})

	for i := range instructions {
		td.DoInstruction(i)
		// switch cmd {
		// case "F":
		// 	td.Forward(segLen)
		// case "R":
		// 	td.Right(90)
		// case "L":
		// 	td.Left(90)
		// }
	}

	outImgName := fmt.Sprintf("hilbert_single_%02d_%d.png", level, imgRes)
	w.SaveImage(outImgName)
}

func hilbertFancy(level int, sides int) {
	// receive the instructions here
	instructions := make(chan turtle.Instruction)

	// the size of the image
	imgHeight := 1080 * 2
	imgWidth := 1920 * 2
	// imgHeight := 2000
	// imgWidth := 2000

	// half the height
	midHeight := float64(imgHeight) / 2
	midWidth := float64(imgWidth) / 2

	// radius of the circumscribed (? is that a word in English?) circle
	radius := float64(imgHeight) / 4

	// angle of each sector
	secAngleDeg := 360 / float64(sides)
	secAngle := turtle.Deg2rad(secAngleDeg)

	// side length of the sidesAgon
	side := radius * 2 * math.Sin(secAngle/2)

	// segment length for the Hilbert curve
	segLen := getSegmentLen(level, int(side))

	// will produce instructions on the channel
	go fractal.GenerateHilbert(level, instructions, segLen)

	// create a new world to draw in
	w := turtle.NewWorld(imgWidth, imgHeight)

	// a helper turtle to find the corners of the sidesAgon
	// hexagon is the bestagon
	tHelp := turtle.New()
	tHelp.SetPos(midWidth, midHeight)
	tHelp.SetHeading(turtle.South + secAngleDeg/2)
	tHelp.Forward(radius)
	tHelp.SetHeading(turtle.West)
	// we are now in the bottom right vertex

	// one drawing turtle per side
	tDraw := make([]*turtle.TurtleDraw, sides)

	for i := 0; i < sides; i++ {
		// setup the turtle
		tDraw[i] = turtle.NewTurtleDraw(w)
		tDraw[i].SetHeading(tHelp.Deg)
		tDraw[i].SetPos(tHelp.X, tHelp.Y)
		tDraw[i].PenDown()
		tDraw[i].SetColor(turtle.DarkOrange)
		fmt.Printf("%2d : %+v\n", i, tDraw[i])

		// go to the next vertex
		tHelp.Forward(side)
		tHelp.Right(secAngleDeg)
	}

	// generate the instructions once and move all the turtles!
	for cmd := range instructions {
		for i := 0; i < sides; i++ {
			tDraw[i].DoInstruction(cmd)
			// switch cmd {
			// case "F":
			// 	tDraw[i].Forward(segLen)
			// case "R":
			// 	tDraw[i].Right(90)
			// case "L":
			// 	tDraw[i].Left(90)
			// }
		}
	}

	// save the image
	outImgName := fmt.Sprintf("hilbert_fancy_%02d_%02d_%d.png", sides, level, imgWidth)
	fmt.Printf("outImgName = %+v\n", outImgName)
	w.SaveImage(outImgName)
}

func getSegmentLen(level, size int) float64 {
	return float64(size) / (math.Exp2(float64(level-1))*4 - 1)
}

func main() {

	// recursion level
	level := 2

	// draw a single Hilbert curve
	hilbertSingle(level)

	// draw a zillion of them
	sides := 19
	hilbertFancy(level, sides)

}
