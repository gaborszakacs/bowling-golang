package bowling

// Game ...
type Game struct {
	frames []frame
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
	return g.sumOfFrames() + g.spareBonus()
}

type frame struct {
	first  *int
	second *int
}

func (f *frame) isFinished() bool {
	return f.second != nil
}

func (f *frame) sum() int {
	return *f.first + *f.second
}

func (f *frame) isSpare() bool {
	return *f.first+*f.second == 10
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
