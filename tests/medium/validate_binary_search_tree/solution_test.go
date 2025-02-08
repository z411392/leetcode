package validate_binary_search_tree

import (
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

func Test_isValidBST_1(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{2, 1, 3}, 0, null)
	got := isValidBST(root)
	expected := true
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_isValidBST_2(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{5, 1, 4, null, null, 3, 6}, 0, null)
	got := isValidBST(root)
	expected := false
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
