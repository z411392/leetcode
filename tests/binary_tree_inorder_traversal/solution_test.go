package binary_tree_inorder_traversal

import (
	"reflect"
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

func Test_inorderTraversal_1(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{1, null, 2, null, null, 3, null}, 0, null)
	got := inorderTraversal(root)
	expected := []int{1, 3, 2}
	if !reflect.DeepEqual(got, expected) {
		t.FailNow()
	}
}

func Test_inorderTraversal_2(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{1, 2, 3, 4, 5, null, 8, null, null, 6, 7, null, null, 9, null}, 0, null)
	got := inorderTraversal(root)
	expected := []int{4, 2, 6, 5, 7, 1, 3, 9, 8}
	if !reflect.DeepEqual(got, expected) {
		t.FailNow()
	}
}

func Test_inorderTraversal_3(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{}, 0, null)
	got := inorderTraversal(root)
	expected := []int{}
	if !reflect.DeepEqual(got, expected) {
		t.FailNow()
	}
}

func Test_inorderTraversal_4(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{1}, 0, null)
	got := inorderTraversal(root)
	expected := []int{1}
	if !reflect.DeepEqual(got, expected) {
		t.FailNow()
	}
}
