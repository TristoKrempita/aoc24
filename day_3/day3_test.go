package day3_test

import (
	"testing"

	day3 "github.com/TristoKrempita/aoc_go_24/day_3"
)

func TestParser(t *testing.T) {
	t.Run("test parsing one line", func(t *testing.T) {
		inputString := "why()}''(!how()$~mul(420,484) ]}}mul(218,461),]"
		got := day3.ParseString(inputString)
		want := 303778

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
