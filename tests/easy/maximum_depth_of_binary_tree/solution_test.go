package maximum_depth_of_binary_tree

import (
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

func Test_maxDepth_1(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{2, 7, 11, 15}, 0, null)
	got := maxDepth(root)
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_maxDepth_2(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{1, null, 2}, 0, null)
	got := maxDepth(root)
	expected := 2
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
