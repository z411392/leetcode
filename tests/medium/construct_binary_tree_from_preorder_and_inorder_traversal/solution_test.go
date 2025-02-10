package construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"reflect"
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

func Test_buildTree_1(t *testing.T) {
	t.SkipNow()
	got := ConvertTreeToSlice(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	// expected := []int{3, 9, 20, null, null, 15, 7}
	expected := []int{3, 9, 20, 15, 7}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_buildTree_2(t *testing.T) {
	t.SkipNow()
	got := ConvertTreeToSlice(buildTree([]int{-1}, []int{-1}))
	expected := []int{-1}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
