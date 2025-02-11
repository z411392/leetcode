package sort_list

//lint:file-ignore ST1001 _

import (
	"reflect"
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func Test_sortList_1(t *testing.T) {
	// t.SkipNow()
	got := ConvertLinkedListToSlice(sortList(NewLinkedListFromSlice([]int{4, 2, 1, 3})))
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_sortList_2(t *testing.T) {
	// t.SkipNow()
	got := ConvertLinkedListToSlice(sortList(NewLinkedListFromSlice([]int{-1, 5, 3, 4, 0})))
	expected := []int{-1, 0, 3, 4, 5}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_sortList_3(t *testing.T) {
	// t.SkipNow()
	got := ConvertLinkedListToSlice(sortList(NewLinkedListFromSlice([]int{})))
	expected := []int{}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
