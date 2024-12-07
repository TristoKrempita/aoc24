package day1_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	day1 "github.com/TristoKrempita/aoc_go_24/day_1"
)

func TestDay1(t *testing.T) {
	t.Run("slices from file", func(t *testing.T) {
		fs := fstest.MapFS{
			"simple": {Data: []byte("12345   54321\n10452   44444")},
		}
		got_left, got_right, _ := day1.SlicesFromFile("simple", fs)
		want_left := []int{12345, 10452}
		want_right := []int{54321, 44444}
		assertSlices(t, want_left, got_left)
		assertSlices(t, want_right, got_right)
	})

	t.Run("sorting", func(t *testing.T) {
		got := []int{4, 7, 1, 7, 3, 8, 2, 5}
		day1.SortSlice(got)
		want := []int{1, 2, 3, 4, 5, 7, 7, 8}
		assertSlices(t, got, want)
	})

	t.Run("ACCEPTANCE: entire output", func(t *testing.T) {
		got := day1.SumAllDifferences()
		want := 2166959
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("similar elements", func(t *testing.T) {
		want := 23741109
		got := day1.SumSimilarity()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func assertSlices(tb testing.TB, want, got []int) {
	if !reflect.DeepEqual(want, got) {
		tb.Errorf("got %+v want %+v", got, want)
	}
}
