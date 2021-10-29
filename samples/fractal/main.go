package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/Pitrified/go-turtle"
	"github.com/Pitrified/go-turtle/fractal"
)

func drawFractals(which, imgShape string, level int) {

	var imgWidth, imgHeight float64
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
		pad := 80.0
		segLen := getHilbertSegmentLen(level-1, imgHeight-pad)

		// center the drawing
		hp := pad / 2
		startX = float64(imgWidth-imgHeight)/2 + hp
		startY = hp
		startD = 0

		// will produce instructions on the channel
		go fractal.GenerateHilbert(level, instructions, segLen)

	case "dragon":
		// type 1: start at the center and spiral
		segLen := 20.0
		startX = imgWidth / 2
		startY = imgHeight / 2
		// type 2 could rotate by 45 deg per level,
		// to keep the main design fixed
		// and reduce the length by a factor of sqrt(2)
		go fractal.GenerateDragon(level, instructions, segLen)

	case "sierpTri":
		pad := 80.0
		hp := pad / 2
		startX = float64(imgWidth-imgHeight)/2 + hp + imgHeight - pad
		startY = (imgHeight - (imgHeight-pad)*math.Sin(math.Pi/3)) / 2
		startD = 180.0
		segLen := (imgHeight - pad) / math.Exp2(float64(level))
		go fractal.GenerateSierpinskiTriangle(level, instructions, segLen)

	case "sierpArrow":
		pad := 80.0
		hp := pad / 2
		startX = float64(imgWidth-imgHeight)/2 + hp
		startY = (imgHeight - (imgHeight-pad)*math.Sin(math.Pi/3)) / 2
		if level%2 != 0 {
			startD = 60.0
		}
		segLen := (imgHeight - pad) / math.Exp2(float64(level))
		go fractal.GenerateSierpinskiArrowhead(level, instructions, segLen)

	}

	// create a new world to draw in
	w := turtle.NewWorld(int(imgWidth), int(imgHeight))

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

func getHilbertSegmentLen(level int, size float64) float64 {
	return size / (math.Exp2(float64(level-1))*4 - 1)
}

// Nice images:
// go run main.go -f dragon -l 12 -i 4K
// go run main.go -f hilbert -l 7 -i 4K
// go run main.go -f sierpArrow -l 7 -i 4K
// go run main.go -f sierpTri -l 7 -i 4K
func main() {
	which := flag.String("f", "hilbert", "Type of fractal to generate.")
	imgShape := flag.String("i", "4K", "Shape of the image to generate.")
	level := flag.Int("l", 4, "Recursion level to reach.")
	flag.Parse()
	drawFractals(*which, *imgShape, *level)
}
