package merge_two_sorted_lists

//lint:file-ignore ST1001 Dot imports are used for convenience in this test file

import (
	"reflect"
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func newLinkedListFromSlice[T any](slice []T) *LinkedList[T] {
	if len(slice) == 0 {
		return nil
	}

	head := &LinkedList[T]{Val: slice[0]}
	current := head

	for _, value := range slice[1:] {
		current.Next = &LinkedList[T]{Val: value}
		current = current.Next
	}

	return head
}

func convertLinkedListToSlice[T any](list *LinkedList[T]) []T {
	if list == nil {
		return []T{}
	}

	slice := []T{}
	current := list
	for current != nil {
		slice = append(slice, current.Val)
		current = current.Next
	}

	return slice
}

func Test_mergeTwoLists_1(t *testing.T) {
	t.SkipNow()
	list1 := newLinkedListFromSlice([]int{1, 2, 4})
	list2 := newLinkedListFromSlice([]int{1, 3, 4})
	got := mergeTwoLists(list1, list2)
	if !reflect.DeepEqual(convertLinkedListToSlice(got), []int{1, 1, 2, 3, 4, 4}) {
		t.FailNow()
	}
}

func Test_mergeTwoLists_2(t *testing.T) {
	t.SkipNow()
	list1 := newLinkedListFromSlice([]int{})
	list2 := newLinkedListFromSlice([]int{})
	got := mergeTwoLists(list1, list2)
	if !reflect.DeepEqual(convertLinkedListToSlice(got), []int{}) {
		t.FailNow()
	}
}

func Test_mergeTwoLists_3(t *testing.T) {
	t.SkipNow()
	list1 := newLinkedListFromSlice([]int{})
	list2 := newLinkedListFromSlice([]int{0})
	got := mergeTwoLists(list1, list2)
	if !reflect.DeepEqual(convertLinkedListToSlice(got), []int{0}) {
		t.FailNow()
	}
}
