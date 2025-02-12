package find_peak_element

import (
	"testing"
)

func Test_findPeakElement_1(t *testing.T) {
	// t.SkipNow()
	got := findPeakElement([]int{1, 2, 3, 1})
	expected := 2
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_findPeakElement_2(t *testing.T) {
	// t.SkipNow()
	got := findPeakElement([]int{1, 2, 1, 3, 5, 6, 4})
	expected := 5
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
