package longest_increasing_subsequence

import (
	"testing"
)

func Test_lengthOfLIS_1(t *testing.T) {
	// t.SkipNow()
	got := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	expected := 4
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_lengthOfLIS_2(t *testing.T) {
	// t.SkipNow()
	got := lengthOfLIS([]int{0, 1, 0, 3, 2, 3})
	expected := 4
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_lengthOfLIS_3(t *testing.T) {
	// t.SkipNow()
	got := lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7})
	expected := 1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
