package merge_two_sorted_lists

//lint:file-ignore ST1001 _

import (
	"reflect"
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func Test_mergeTwoLists_1(t *testing.T) {
	t.SkipNow()
	list1 := NewLinkedListFromSlice([]int{1, 2, 4})
	list2 := NewLinkedListFromSlice([]int{1, 3, 4})
	got := mergeTwoLists(list1, list2)
	if !reflect.DeepEqual(ConvertLinkedListToSlice(got), []int{1, 1, 2, 3, 4, 4}) {
		t.FailNow()
	}
}

func Test_mergeTwoLists_2(t *testing.T) {
	t.SkipNow()
	list1 := NewLinkedListFromSlice([]int{})
	list2 := NewLinkedListFromSlice([]int{})
	got := mergeTwoLists(list1, list2)
	if !reflect.DeepEqual(ConvertLinkedListToSlice(got), []int{}) {
		t.FailNow()
	}
}

func Test_mergeTwoLists_3(t *testing.T) {
	t.SkipNow()
	list1 := NewLinkedListFromSlice([]int{})
	list2 := NewLinkedListFromSlice([]int{0})
	got := mergeTwoLists(list1, list2)
	if !reflect.DeepEqual(ConvertLinkedListToSlice(got), []int{0}) {
		t.FailNow()
	}
}
