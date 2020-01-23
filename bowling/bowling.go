package bowling

import (
	"fmt"
	"io"
)

// Game ...
type Game struct {
	frames []frame
	Out    io.Writer
}

// Roll ...
func (g *Game) Roll(n int) {
	if g.hasJustStarted() || g.lastFrame().isFinished() {
		g.startNewFrame(n)
	} else {
		g.lastFrame().second = &n
	}
}

// Score ...
func (g *Game) Score() int {
	return g.sumOfFrames() + g.spareBonus() + g.strikeBonus()
}

type Gipher interface {
	Random(string) (string, error)
}

// Celebrate ...
func (g *Game) Celebrate(giphy Gipher) (string, error) {
	url, err := giphy.Random("bowling")
	if err != nil {
		return "", err
	}
	return url, nil
}

// PrintRolls ...
func (g *Game) PrintRolls() {
	for _, frame := range g.frames {
		if frame.second == nil {
			fmt.Fprintln(g.Out, *frame.first)
		} else {
			fmt.Fprintf(g.Out, "%d | %d\n", *frame.first, *frame.second)
		}
	}
}

type frame struct {
	first  *int
	second *int
}

func (f *frame) isFinished() bool {
	if *f.first == 10 {
		return true
	}
	return f.second != nil
}

func (f *frame) sum() int {
	if f.second == nil {
		return *f.first
	}
	return *f.first + *f.second
}

func (f *frame) isSpare() bool {
	if f.second == nil {
		return false
	}
	return *f.first+*f.second == 10
}

func (f *frame) isStrike() bool {
	return *f.first == 10
}

func (g *Game) lastFrame() *frame {
	return &g.frames[len(g.frames)-1]
}

func (g *Game) hasJustStarted() bool {
	return len(g.frames) == 0
}

func (g *Game) startNewFrame(n int) {
	g.frames = append(g.frames, frame{first: &n})
}

func (g *Game) sumOfFrames() int {
	sum := 0
	for _, frame := range g.frames {
		sum += frame.sum()
	}
	return sum
}

func (g *Game) spareBonus() int {
	spareBonus := 0
	for i, frame := range g.frames {
		if frame.isSpare() {
			spareBonus += *g.frames[i+1].first
		}
	}
	return spareBonus
}

func (g *Game) strikeBonus() int {
	strikeBonus := 0
	for i, frame := range g.frames {
		if frame.isStrike() {
			strikeBonus += *g.frames[i+1].first
			strikeBonus += *g.frames[i+1].second
		}
	}
	return strikeBonus
}
