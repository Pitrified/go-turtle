package fractal

import (
	"github.com/Pitrified/go-turtle"
)

// Generate instructions for a general Lindenmayer system.
//
// level: recursion level to reach.
// instructions: channel where the instructions will be sent.
// remaining: set to the axiom.
// rules: production rules.
// angle: how much to rotate.
// forward: how much to move forward.
//
// Two mildly different rewrite rules can be used:
// using ABCD, the forward movement must be explicit, using an F.
// using XYWZ, the forward movement is done when the base of the recursion is reached.
//
// https://en.wikipedia.org/wiki/L-system
func Instructions(
	level int,
	instructions chan<- turtle.Instruction,
	remaining string,
	rules map[byte]string,
	angle float64,
	forward float64,
) string {

	for len(remaining) > 0 {
		curChar := remaining[0]
		remaining = remaining[1:]
		// fmt.Printf("%3d %c %+v\n", level, curChar, remaining)

		switch curChar {

		case '|':
			return remaining

		case '+':
			instructions <- turtle.Instruction{Cmd: turtle.CmdLeft, Amount: angle}
		case '-':
			instructions <- turtle.Instruction{Cmd: turtle.CmdRight, Amount: angle}

		case 'F':
			instructions <- turtle.Instruction{Cmd: turtle.CmdForward, Amount: forward}

		// move forward explicitly when an 'F' is encountered
		case 'A', 'B', 'C', 'D':
			if level > 0 {
				remaining = rules[curChar] + "|" + remaining
				remaining = Instructions(level-1, instructions, remaining, rules, angle, forward)
			}

		// move forward when the base of the recursion is reached
		case 'X', 'Y', 'W', 'Z':
			if level == 0 {
				instructions <- turtle.Instruction{Cmd: turtle.CmdForward, Amount: forward}
			} else if level > 0 {
				remaining = rules[curChar] + "|" + remaining
				remaining = Instructions(level-1, instructions, remaining, rules, angle, forward)
			}
		}
	}

	close(instructions)
	return ""
}

// Generate instructions to draw a Hilbert curve,
// with the requested recursion level,
// receiving Instruction on the channel instructions.
//
// The channel will be closed to signal the end of the stream.
//
// For more information:
// https://en.wikipedia.org/wiki/Hilbert_curve#Representation_as_Lindenmayer_system
func GenerateHilbert(level int, instructions chan<- turtle.Instruction, forward float64) {
	rules := map[byte]string{'A': "+BF-AFA-FB+", 'B': "-AF+BFB+FA-"}
	Instructions(level, instructions, "A", rules, 90, forward)
}

// Generate instructions to draw a dragon curve.
//
// https://en.wikipedia.org/wiki/Dragon_curve
// https://en.wikipedia.org/wiki/L-system#Example_6:_Dragon_curve
func GenerateDragon(level int, instructions chan<- turtle.Instruction, forward float64) {
	rules := map[byte]string{
		// 'A': "AF+B", 'B': "AF-B", // identical
		'X': "X+Y",
		'Y': "X-Y",
		// 'X': "Y-X-Y",
		// 'Y': "X+Y+X",
		// 'X': "X-Y+X+Y-X",
		// 'Y': "YY",
	}
	// initial remaining commands to do
	remaining := "X"
	// remaining := "A"
	// remaining := "X-Y-Y"
	// will produce instructions on the channel
	angle := 90.0
	Instructions(level, instructions, remaining, rules, angle, forward)
}

// Generate instructions to draw a Sierpinski arrowhead curve.
//
// https://en.wikipedia.org/wiki/Sierpi%C5%84ski_curve#Arrowhead_curve
func GenerateSierpinskiArrowhead(level int, instructions chan<- turtle.Instruction, forward float64) {
	rules := map[byte]string{'X': "Y-X-Y", 'Y': "X+Y+X"}
	Instructions(level, instructions, "X", rules, 60, forward)
}

// Generate instructions to draw a Sierpinski triangle.
//
// https://en.wikipedia.org/wiki/L-system#Example_5:_Sierpinski_triangle
func GenerateSierpinskiTriangle(level int, instructions chan<- turtle.Instruction, forward float64) {
	rules := map[byte]string{'X': "X-Y+X+Y-X", 'Y': "YY"}
	Instructions(level, instructions, "X-Y-Y", rules, 120, forward)
}
