package container_with_most_water

import (
	"testing"
)

func Test_maxArea_1(t *testing.T) {
	t.SkipNow()
	got := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	expected := 49
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_maxArea_2(t *testing.T) {
	t.SkipNow()
	got := maxArea([]int{1, 1})
	expected := 1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
