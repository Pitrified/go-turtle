package turtle

// Possible commands to send inside an Instruction.
type CmdType byte

const (
	CmdForward CmdType = iota
	CmdBackward
	CmdLeft
	CmdRight
)

// An action for the turtle.
type Instruction struct {
	Cmd    CmdType
	Amount float64
}
