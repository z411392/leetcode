package diameter_of_binary_tree

import (
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

func Test_diameterOfBinaryTree_1(t *testing.T) {
	t.SkipNow()
	got := diameterOfBinaryTree(NewTreeFromSlice([]int{1, 2, 3, 4, 5}, 0, null))
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_diameterOfBinaryTree_2(t *testing.T) {
	t.SkipNow()
	got := diameterOfBinaryTree(NewTreeFromSlice([]int{1, 2}, 0, null))
	expected := 1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
