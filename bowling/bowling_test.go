package bowling_test

import (
	"bytes"
	"testing"

	"github.com/gaborszakacs/bowling-golang/bowling"
)

func TestScoreTable(t *testing.T) {
	testCases := []struct {
		desc    string
		RollsFn func(g *bowling.Game)
		want    int
	}{
		{
			desc: "Gutter game",
			RollsFn: func(g *bowling.Game) {
				rollMany(g, 10, 0, 0)
			},
			want: 0,
		},
		{
			desc: "Always 1",
			RollsFn: func(g *bowling.Game) {
				rollMany(g, 10, 1, 0)
			},
			want: 10,
		},
		{
			desc: "Spare game",
			RollsFn: func(g *bowling.Game) {
				g.Roll(1)
				g.Roll(9)
				rollMany(g, 9, 1, 0)
			},
			want: 20,
		},
		{
			desc: "Strike game",
			RollsFn: func(g *bowling.Game) {
				g.Roll(10)
				g.Roll(2)
				g.Roll(2)
				rollMany(g, 8, 0, 0)
			},
			want: 18,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			g := bowling.Game{}
			tc.RollsFn(&g)
			got := g.Score()
			want := tc.want
			if got != want {
				t.Errorf("got: %d, want: %d", got, want)
			}
		})
	}
}

func TestPrintScoreCard(t *testing.T) {
	g := bowling.Game{}
	rollMany(&g, 10, 1, 2)

	b := bytes.Buffer{}
	g.PrintScoreCard(&b)
	got := b.String()
	want := `1 | 2
1 | 2
1 | 2
1 | 2
1 | 2
1 | 2
1 | 2
1 | 2
1 | 2
1 | 2
`

	if got != want {
		t.Errorf("got:\n%s\n want:\n%s", got, want)
	}
}

func rollMany(g *bowling.Game, times, score1, score2 int) {
	for i := 0; i < times; i++ {
		g.Roll(score1)
		g.Roll(score2)
	}
}
