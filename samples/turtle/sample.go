package main

import (
	"fmt"

	"github.com/Pitrified/go-turtle"
)

func main() {
	t := turtle.New()
	fmt.Println("T:", t)

	t.Forward(5)
	fmt.Println("T:", t)

	t.Left(45)
	fmt.Println("T:", t)

	t.Forward(5)
	fmt.Println("T:", t)

	t.Right(45)
	fmt.Println("T:", t)

	t.Backward(5)
	fmt.Println("T:", t)
	fmt.Println(t.X, t.Y, t.Deg)

	t.SetPos(4, 4)
	fmt.Println("T:", t)

	t.SetHeading(120)
	fmt.Println("T:", t)
}
