package day2_test

import (
	"reflect"
	"testing"

	day2 "github.com/TristoKrempita/aoc_go_24/day_2"
)

func TestDay2(t *testing.T) {
	t.Run("ACCEPTANCE: final output", func(t *testing.T) {
		levelList := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
		want := 4
		got := day2.CountSafeLevels(levelList)
		if want != got {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("string to slice of ints", func(t *testing.T) {
		stringLine := "1 2 7 8 9"
		want := []int{1, 2, 7, 8, 9}
		got := day2.LevelToSlice(stringLine)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})
}
