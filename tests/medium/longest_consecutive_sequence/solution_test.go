package longest_consecutive_sequence

import (
	"testing"
)

func Test_longestConsecutive_1(t *testing.T) {
	// t.SkipNow()
	got := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	expected := 4
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_longestConsecutive_2(t *testing.T) {
	// t.SkipNow()
	got := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	expected := 9
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
