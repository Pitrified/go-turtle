package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/Pitrified/go-turtle"
)

// https://en.wikipedia.org/wiki/Hilbert_curve#Representation_as_Lindenmayer_system
func hilbertInstructions(
	level int,
	remaining string,
	instructions chan<- string,
	rules map[string]string,
) string {

	for len(remaining) > 0 {
		curChar := remaining[0]
		remaining = remaining[1:]

		switch curChar {

		case '|':
			return remaining

		case '+':
			instructions <- "L"
		case '-':
			instructions <- "R"

		case 'F':
			instructions <- "F"

		case 'A':
			if level > 0 {
				remaining = rules["A"] + "|" + remaining
				remaining = hilbertInstructions(
					level-1, remaining, instructions, rules)
			}
		case 'B':
			if level > 0 {
				remaining = rules["B"] + "|" + remaining
				remaining = hilbertInstructions(
					level-1, remaining, instructions, rules)
			}
		}
	}

	close(instructions)
	return ""
}

func hilbertSingle(
	level int,
	rules map[string]string,
) {
	// initial remaining commands to do
	remaining := rules["A"]

	// receive the instructions here
	instructions := make(chan string)

	// will produce instructions on the channel
	go hilbertInstructions(level, remaining, instructions, rules)

	// the size of the image
	imgRes := 1080

	// the step for each Hilbert curve segment
	segLen := getSegmentLen(level, 80, imgRes)

	// create a new world to draw in
	w := turtle.NewWorld(imgRes, imgRes)

	// create and setup a turtle
	td := turtle.NewTurtleDraw(w)
	td.SetPos(40, 40)
	td.PenDown()
	td.SetColor(color.RGBA{150, 75, 0, 255})

	for cmd := range instructions {
		switch cmd {
		case "F":
			td.Forward(segLen)
		case "R":
			td.Right(90)
		case "L":
			td.Left(90)
		}
	}

	outImgName := fmt.Sprintf("hilbert_single_%02d_%d.png", level, imgRes)
	w.SaveImage(outImgName)
}

func hilbertFancy(
	level int,
	sides int,
	rules map[string]string,
) {
	// initial remaining commands to do
	remaining := rules["A"]

	// receive the instructions here
	instructions := make(chan string)

	// will produce instructions on the channel
	go hilbertInstructions(level, remaining, instructions, rules)

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
	segLen := getSegmentLen(level, 0, int(side))

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
			switch cmd {
			case "F":
				tDraw[i].Forward(segLen)
			case "R":
				tDraw[i].Right(90)
			case "L":
				tDraw[i].Left(90)
			}
		}
	}

	// save the image
	outImgName := fmt.Sprintf("hilbert_fancy_%02d_%02d_%d.png", sides, level, imgWidth)
	fmt.Printf("outImgName = %+v\n", outImgName)
	w.SaveImage(outImgName)
}

func getSegmentLen(level, pad, imgRes int) float64 {
	return float64(imgRes-pad) / (math.Exp2(float64(level-1))*4 - 1)
}

func main() {

	// recursion level
	level := 2

	// rewrite rules
	rules := map[string]string{
		"A": "+BF-AFA-FB+",
		"B": "-AF+BFB+FA-",
	}

	// draw a single Hilbert curve
	hilbertSingle(level, rules)

	// draw a zillion of them
	sides := 19
	hilbertFancy(level, sides, rules)

}
