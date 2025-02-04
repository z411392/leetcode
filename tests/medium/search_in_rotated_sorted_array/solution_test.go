package search_in_rotated_sorted_array

import (
	"testing"
)

func Test_search_1(t *testing.T) {
	// t.SkipNow()
	got := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	expected := 4
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_search_2(t *testing.T) {
	// t.SkipNow()
	got := search([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	expected := -1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_search_3(t *testing.T) {
	// t.SkipNow()
	got := search([]int{1}, 0)
	expected := -1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
