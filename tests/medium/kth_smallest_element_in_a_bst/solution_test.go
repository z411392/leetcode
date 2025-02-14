package kth_smallest_element_in_a_bst

import (
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

func Test_kthSmallest_1(t *testing.T) {
	// t.SkipNow()
	root := NewTreeFromSlice([]int{3, 1, 4, null, 2}, 0, null)
	got := kthSmallest(root, 1)
	expected := 1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_kthSmallest_2(t *testing.T) {
	// t.SkipNow()
	root := NewTreeFromSlice([]int{5, 3, 6, 2, 4, null, null, 1}, 0, null)
	got := kthSmallest(root, 3)
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
