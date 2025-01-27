package invert_binary_tree

//lint:file-ignore ST1001 _

import (
	"reflect"
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

func Test_invertTree_1(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{4, 2, 7, 1, 3, 6, 9}, 0, null)
	got := ConvertTreeToSlice(invertTree(root))
	expected := []int{4, 7, 2, 9, 6, 3, 1}
	if !reflect.DeepEqual(got, expected) {
		t.FailNow()
	}
}

func Test_invertTree_2(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{2, 1, 3}, 0, null)
	got := ConvertTreeToSlice(invertTree(root))
	expected := []int{2, 3, 1}
	if !reflect.DeepEqual(got, expected) {
		t.FailNow()
	}
}

func Test_invertTree_3(t *testing.T) {
	t.SkipNow()
	root := NewTreeFromSlice([]int{}, 0, null)
	got := ConvertTreeToSlice(invertTree(root))
	expected := []int{}
	if !reflect.DeepEqual(got, expected) {
		t.FailNow()
	}
}
