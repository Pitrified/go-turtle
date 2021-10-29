package fractal

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
	instructions chan<- string,
	remaining string,
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
			if level >= 0 {
				remaining = rules["A"] + "|" + remaining
				remaining = Instructions(level-1, instructions, remaining, rules)
			}
		case 'B':
			if level >= 0 {
				remaining = rules["B"] + "|" + remaining
				remaining = Instructions(level-1, instructions, remaining, rules)
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
	instructions chan<- string,
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
	Instructions(level, instructions, remaining, rules)
}
