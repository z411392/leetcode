package add_two_numbers

//lint:file-ignore ST1001 _

import (
	"reflect"
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func Test_addTwoNumbers_1(t *testing.T) {
	// t.SkipNow()
	l1 := NewLinkedListFromSlice([]int{2, 4, 3})
	l2 := NewLinkedListFromSlice([]int{5, 6, 4})
	got := ConvertLinkedListToSlice(addTwoNumbers(l1, l2))
	expected := []int{7, 0, 8}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_addTwoNumbers_2(t *testing.T) {
	// t.SkipNow()
	l1 := NewLinkedListFromSlice([]int{0})
	l2 := NewLinkedListFromSlice([]int{0})
	got := ConvertLinkedListToSlice(addTwoNumbers(l1, l2))
	expected := []int{0}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_addTwoNumbers_3(t *testing.T) {
	// t.SkipNow()
	l1 := NewLinkedListFromSlice([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := NewLinkedListFromSlice([]int{9, 9, 9, 9})
	got := ConvertLinkedListToSlice(addTwoNumbers(l1, l2))
	expected := []int{8, 9, 9, 9, 0, 0, 0, 1}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
