package largest_number

import (
	"testing"
)

func Test_largestNumber_1(t *testing.T) {
	// t.SkipNow()
	got := largestNumber([]int{10, 2})
	expected := "210"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_largestNumber_2(t *testing.T) {
	// t.SkipNow()
	got := largestNumber([]int{3, 30, 34, 5, 9})
	expected := "9534330"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
