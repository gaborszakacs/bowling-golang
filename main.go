package main

import (
	"github.com/gaborszakacs/bowling-golang/bowling"
)

func main() {
	g := bowling.Game{}
	for i := 0; i < 9; i++ {
		g.Roll(2)
		g.Roll(3)
	}
	g.Roll(10)
	g.PrintRolls()
}
