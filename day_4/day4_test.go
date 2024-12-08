package day4_test

import (
	"testing"

	day4 "github.com/TristoKrempita/aoc_go_24/day_4"
)

func TestWordSearch(t *testing.T) {
	t.Run("find XMAS in example string", func(t *testing.T) {
		exampleSearch := `....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX`
		got := day4.WordSearch(exampleSearch)
		want := 18

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("find X-MAS in example string", func(t *testing.T) {
		exampleSearch := `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`
		got := day4.CrossSearch(exampleSearch)
		want := 9

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
