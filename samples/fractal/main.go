package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/Pitrified/go-turtle"
	"github.com/Pitrified/go-turtle/fractal"
)

func drawFractals(which, imgShape string, level int) {

	var imgWidth, imgHeight int
	var startX, startY, startD float64

	switch imgShape {
	case "4K":
		imgWidth = 1920 * 2
		imgHeight = 1080 * 2
	case "1200":
		imgWidth = 1200
		imgHeight = 1200
	}

	// receive the instructions here
	instructions := make(chan turtle.Instruction)

	switch which {

	case "hilbert":
		// compute the segLen to fill the image in height
		pad := 80
		segLen := getHilbertSegmentLen(level, imgHeight-pad)

		// center the drawing
		hp := float64(pad) / 2
		startX = float64(imgWidth-imgHeight)/2 + hp
		startY = hp
		startD = 0

		// will produce instructions on the channel
		go fractal.GenerateHilbert(level, instructions, segLen)

	case "dragon":
		segLen := 20.0
		// start at the center
		startX = float64(imgWidth) / 2
		startY = float64(imgHeight) / 2
		go fractal.GenerateDragon(level, instructions, segLen)

	case "sierpTri":

	case "sierpArrow":
		segLen := 20.0
		startX = float64(imgWidth) / 2
		startY = float64(imgHeight) / 2
		if level%2 == 0 {
			startD = 60.0
		}
		go fractal.GenerateSierpinskiArrowhead(level, instructions, segLen)

	}

	// create a new world to draw in
	w := turtle.NewWorld(imgWidth, imgHeight)

	// create and setup a turtle in the right place
	td := turtle.NewTurtleDraw(w)
	td.SetPos(startX, startY)
	td.SetHeading(startD)
	td.PenDown()
	td.SetColor(turtle.DarkOrange)

	// draw the fractal
	for i := range instructions {
		td.DoInstruction(i)
	}

	outImgName := fmt.Sprintf("%s_%02d_%s.png", which, level, imgShape)
	w.SaveImage(outImgName)
}

func getHilbertSegmentLen(level, size int) float64 {
	return float64(size) / (math.Exp2(float64(level-1))*4 - 1)
}

// Nice output:
// go run main.go -f dragon -l 11 -i 4K
// go run main.go -f hilbert -l 6 -i 4K
func main() {
	which := flag.String("f", "hilbert", "Type of fractal to generate.")
	imgShape := flag.String("i", "4K", "Shape of the image to generate.")
	level := flag.Int("l", 4, "Recursion level to reach.")
	flag.Parse()
	drawFractals(*which, *imgShape, *level)
}
