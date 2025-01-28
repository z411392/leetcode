package maximum_number_of_fish_in_a_grid

import (
	"testing"
)

func Test_findMaxFish_1(t *testing.T) {
	t.SkipNow()
	got := findMaxFish([][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}})
	expected := 7
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_findMaxFish_2(t *testing.T) {
	t.SkipNow()
	got := findMaxFish([][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 1}})
	expected := 1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
