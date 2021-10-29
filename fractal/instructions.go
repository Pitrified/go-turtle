package fractal

import (
	"fmt"

	"github.com/Pitrified/go-turtle"
)

// Generate instructions for a general Lindenmayer system.
//
// level: recursion level to reach.
// instructions: channel where the instructions will be sent.
// remaining: set to the axiom.
// rules: production rules.
//
// https://en.wikipedia.org/wiki/L-system
func Instructions(
	level int,
	instructions chan<- turtle.Instruction,
	remaining string,
	rules map[string]string,
	angle float64,
	forward float64,
) string {

	for len(remaining) > 0 {
		curChar := remaining[0]
		remaining = remaining[1:]
		fmt.Printf("%3d %c %+v\n", level, curChar, remaining)

		switch curChar {

		case '|':
			return remaining

		case '+':
			instructions <- turtle.Instruction{Cmd: turtle.CmdLeft, Amount: angle}
		case '-':
			instructions <- turtle.Instruction{Cmd: turtle.CmdRight, Amount: angle}

		case 'F':
			instructions <- turtle.Instruction{Cmd: turtle.CmdForward, Amount: forward}

		case 'A':
			if level >= 0 {
				remaining = rules["A"] + "|" + remaining
				remaining = Instructions(level-1, instructions, remaining, rules, angle, forward)
			}
		case 'B':
			if level >= 0 {
				remaining = rules["B"] + "|" + remaining
				remaining = Instructions(level-1, instructions, remaining, rules, angle, forward)
			}

		case 'X':
			if level == -1 {
				instructions <- turtle.Instruction{Cmd: turtle.CmdForward, Amount: forward}
			}
			if level >= 0 {
				remaining = rules["X"] + "|" + remaining
				remaining = Instructions(level-1, instructions, remaining, rules, angle, forward)
			}
		case 'Y':
			if level == -1 {
				instructions <- turtle.Instruction{Cmd: turtle.CmdForward, Amount: forward}
			}
			if level >= 0 {
				remaining = rules["Y"] + "|" + remaining
				remaining = Instructions(level-1, instructions, remaining, rules, angle, forward)
			}
		}
	}

	close(instructions)
	return ""
}

// Generate instructions to draw a Hilbert curve.
//
// of the requested level on channel instructions.
//
// The instructions received can be:
// * "F": move forward.
// * "L": rotate left 90 degrees.
// * "R": rotate rigth 90 degrees.
//
// The channel will be closed to signal the end of the instructions.
//
// For more information:
// https://en.wikipedia.org/wiki/Hilbert_curve#Representation_as_Lindenmayer_system
func GenerateHilbert(
	level int,
	instructions chan<- turtle.Instruction,
	forward float64,
) {
	// rewrite rules
	// https://en.wikipedia.org/wiki/Hilbert_curve#Representation_as_Lindenmayer_system
	rules := map[string]string{
		"A": "+BF-AFA-FB+",
		"B": "-AF+BFB+FA-",
	}
	// initial remaining commands to do
	remaining := "A"
	// will produce instructions on the channel
	angle := 90.0
	Instructions(level, instructions, remaining, rules, angle, forward)
}

// The dragon curve drawn using an L-system.
// variables : A B
// constants : + −
// start  : A
// rules  : (A → A+B), (B → A-B)
// angle  : 90°
// A and B both mean "draw forward",
// + means "turn left by angle", and − means "turn right by angle".
//
// https://en.wikipedia.org/wiki/L-system#Example_6:_Dragon_curve
func GenerateDragon(
	level int,
	instructions chan<- turtle.Instruction,
	forward float64,
) {
	rules := map[string]string{
		"A": "AF+B",
		"B": "AF-B",
		// "X": "X+Y",
		// "Y": "X-Y",
		// "X": "Y-X-Y",
		// "Y": "X+Y+X",
		// "X": "X-Y+X+Y-X",
		// "Y": "YY",
	}
	// initial remaining commands to do
	// remaining := "X"
	remaining := "A"
	// remaining := "X-Y-Y"
	// will produce instructions on the channel
	angle := 90.0
	Instructions(level, instructions, remaining, rules, angle, forward)
}

// TODO
// rules[byte]string, use a single case
// ABCD, XYWZ
// Document that there are different way of moving
