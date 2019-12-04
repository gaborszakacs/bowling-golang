package bowling

import (
	"fmt"
	"io"
)

// Game ...
type Game struct {
	frames []frame
}

type frame struct {
	firstRoll  int
	secondRoll int
	rollsCount int
}

// Roll ...
func (g *Game) Roll(n int) {
	if g.shouldStartNewFrame() {
		g.startFrameWith(n)
	} else {
		g.addToLastFrame(n)
	}
}

// Score ...
func (g *Game) Score() int {
	// fmt.Printf("%+v\n", g.frames)
	// fmt.Printf("base: %d\n", g.baseScore())
	// fmt.Printf("spare: %d\n", g.spareBonus())
	// fmt.Printf("strike: %d\n", g.strikeBonus())
	// panic("hi mom")
	return g.baseScore() + g.spareBonus() + g.strikeBonus()
}

// PrintScoreCard ...
func (g *Game) PrintScoreCard(w io.Writer) {
	for _, f := range g.frames {
		// fmt.Printf("%d | %d\n", f.firstRoll, f.secondRoll)
		fmt.Fprintf(w, "%d | %d\n", f.firstRoll, f.secondRoll)
	}
}
func (g *Game) shouldStartNewFrame() bool {
	if len(g.frames) == 0 {
		return true
	}
	if g.lastFrame().firstRoll == 10 {
		return true
	}

	return g.lastFrame().rollsCount == 2
}

func (g *Game) startFrameWith(n int) {
	g.frames = append(g.frames, frame{firstRoll: n, rollsCount: 1})
}

func (g *Game) addToLastFrame(n int) {
	g.lastFrame().secondRoll = n
	g.lastFrame().rollsCount++
}

func (g *Game) baseScore() int {
	score := 0
	for _, frame := range g.frames {
		score += frame.score()
	}

	return score
}

func (g *Game) spareBonus() int {
	spareBonus := 0
	for i, frame := range g.frames {
		if frame.isSpare() {
			spareBonus += g.frames[i+1].firstRoll
		}
	}

	return spareBonus
}

func (g *Game) strikeBonus() int {
	strikeBonus := 0
	for i, frame := range g.frames {
		if frame.isStrike() {
			strikeBonus += g.frames[i+1].firstRoll + g.frames[i+1].secondRoll
		}
	}

	return strikeBonus
}

func (g *Game) lastFrame() *frame {
	return &g.frames[len(g.frames)-1]
}

func (f frame) isSpare() bool {
	if f.isStrike() {
		return false
	}

	return f.firstRoll+f.secondRoll == 10
}

func (f frame) isStrike() bool {
	return f.firstRoll == 10
}

func (f frame) score() int {
	return f.firstRoll + f.secondRoll
}
