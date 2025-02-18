package four_sum_ii

import (
	"testing"
)

func Test_fourSumCount_1(t *testing.T) {
	// t.SkipNow()
	got := fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2})
	expected := 2
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_fourSumCount_2(t *testing.T) {
	// t.SkipNow()
	got := fourSumCount([]int{0}, []int{0}, []int{0}, []int{0})
	expected := 1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
