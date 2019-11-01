package bowling

// Game ...
type Game struct {
	rolls []int
}

// Roll ...
func (g *Game) Roll(n int) {
	g.rolls = append(g.rolls, n)
}

// Score ...
func (g *Game) Score() int {
	return g.sumOfRolls() + g.spareBonus()
}

func (g *Game) sumOfRolls() int {
	sum := 0
	for _, n := range g.rolls {
		sum += n
	}
	return sum
}

func (g *Game) spareBonus() int {
	spareBonus := 0
	for i := 0; i < 20; i += 2 {
		if g.rolls[i]+g.rolls[i+1] == 10 {
			spareBonus += g.rolls[i+2]
		}
	}
	return spareBonus
}
