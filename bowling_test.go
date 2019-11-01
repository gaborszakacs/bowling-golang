package bowling_test

import (
	"testing"

	bowling "github.com/gaborszakacs/bowling-golang"
)

func TestScore(t *testing.T) {
	t.Run("Gutter game", func(t *testing.T) {
		b := bowling.Game{}
		for i := 0; i < 10; i++ {
			b.Roll(0)
			b.Roll(0)
		}
		got := b.Score()
		want := 0

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Simple game", func(t *testing.T) {

		b := bowling.Game{}
		for i := 0; i < 10; i++ {
			b.Roll(1)
			b.Roll(0)
		}
		got := b.Score()
		want := 10

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Spare", func(t *testing.T) {
		b := bowling.Game{}
		b.Roll(1)
		b.Roll(0)

		b.Roll(3)
		b.Roll(7)

		for i := 0; i < 8; i++ {
			b.Roll(1)
			b.Roll(0)
		}
		got := b.Score()
		want := 20

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
