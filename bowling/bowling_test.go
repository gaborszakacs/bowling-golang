package bowling_test

import (
	"testing"

	"github.com/gaborszakacs/bowling-golang/bowling"
)

func TestScore(t *testing.T) {
	testCases := []struct {
		desc    string
		rollAll func(g *bowling.Game)
		want    int
	}{
		{
			desc: "Gutter game",
			rollAll: func(g *bowling.Game) {
				for i := 0; i < 10; i++ {
					g.Roll(0)
					g.Roll(0)
				}
			},
			want: 0,
		},
		{
			desc: "Simple game",
			rollAll: func(g *bowling.Game) {
				for i := 0; i < 10; i++ {
					g.Roll(1)
					g.Roll(1)
				}
			},
			want: 20,
		},
		{
			desc: "Spare",
			rollAll: func(g *bowling.Game) {
				g.Roll(1)
				g.Roll(1)

				g.Roll(3)
				g.Roll(7)

				for i := 0; i < 8; i++ {
					g.Roll(1)
					g.Roll(1)
				}
			},
			want: 29,
		},
		{
			desc: "Strike",
			rollAll: func(g *bowling.Game) {
				g.Roll(1)
				g.Roll(1)

				g.Roll(10)

				g.Roll(2)
				g.Roll(1)
				for i := 0; i < 7; i++ {
					g.Roll(1)
					g.Roll(1)
				}
			},
			want: 32,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			g := bowling.Game{}
			tc.rollAll(&g)
			got := g.Score()
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
