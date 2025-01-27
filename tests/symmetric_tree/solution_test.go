package symmetric_tree

import (
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

func Test_isSymmetric_1(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{1, 2, 2, 3, 4, 4, 3}, 0, null)
	got := isSymmetric(root)
	expected := true
	if got != expected {
		t.FailNow()
	}
}

func Test_isSymmetric_2(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{1, 2, 2, null, 3, null, 3}, 0, null)
	got := isSymmetric(root)
	expected := false
	if got != expected {
		t.FailNow()
	}
}

func Test_isSymmetric_3(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{1, 2, 2, 2, null, 2}, 0, null)
	got := isSymmetric(root)
	expected := false
	if got != expected {
		t.FailNow()
	}
}
