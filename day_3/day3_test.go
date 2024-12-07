package day3_test

import (
	"testing"

	day3 "github.com/TristoKrempita/aoc_go_24/day_3"
)

func TestParser(t *testing.T) {
	t.Run("test parsing one line", func(t *testing.T) {
		inputString := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
		got := day3.ParseString(inputString)
		want := 48

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
