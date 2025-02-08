package binary_tree_level_order_traversal

import (
	"reflect"
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

func Test_levelOrder_1(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{3, 9, 20, null, null, 15, 7}, 0, null)
	got := levelOrder(root)
	expected := [][]int{{3}, {9, 20}, {15, 7}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_levelOrder_2(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{1}, 0, null)
	got := levelOrder(root)
	expected := [][]int{{1}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_levelOrder_3(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{}, 0, null)
	got := levelOrder(root)
	expected := [][]int{}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
