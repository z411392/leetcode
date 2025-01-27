package reverse_linked_list

import (
	"reflect"
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func Test_reverseList_1(t *testing.T) {
	t.SkipNow()
	got := ConvertLinkedListToSlice(reverseList(NewLinkedListFromSlice([]int{1, 2, 3, 4, 5})))
	expected := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(got, expected) {
		t.FailNow()
	}
}

func Test_reverseList_2(t *testing.T) {
	t.SkipNow()
	got := ConvertLinkedListToSlice(reverseList(NewLinkedListFromSlice([]int{1, 2})))
	expected := []int{2, 1}
	if !reflect.DeepEqual(got, expected) {
		t.FailNow()
	}
}

func Test_reverseList_3(t *testing.T) {
	t.SkipNow()
	got := ConvertLinkedListToSlice(reverseList(NewLinkedListFromSlice([]int{})))
	expected := []int{}
	if !reflect.DeepEqual(got, expected) {
		t.FailNow()
	}
}
