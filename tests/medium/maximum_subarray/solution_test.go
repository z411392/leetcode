package maximum_subarray

import (
	"testing"
)

func Test_maxSubArray_1(t *testing.T) {
	t.SkipNow()
	got := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	expected := 6
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_maxSubArray_2(t *testing.T) {
	t.SkipNow()
	got := maxSubArray([]int{1})
	expected := 1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_maxSubArray_3(t *testing.T) {
	t.SkipNow()
	got := maxSubArray([]int{5, 4, -1, 7, 8})
	expected := 23
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
