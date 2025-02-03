package remove_nth_node_from_end_of_list

//lint:file-ignore ST1001 _

import (
	"reflect"
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func Test_removeNthFromEnd_1(t *testing.T) {
	t.SkipNow()
	got := ConvertLinkedListToSlice(removeNthFromEnd(NewLinkedListFromSlice([]int{1, 2, 3, 4, 5}), 2))
	expected := []int{1, 2, 3, 5}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_removeNthFromEnd_2(t *testing.T) {
	t.SkipNow()
	got := ConvertLinkedListToSlice(removeNthFromEnd(NewLinkedListFromSlice([]int{1}), 1))
	expected := []int{}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_removeNthFromEnd_3(t *testing.T) {
	t.SkipNow()
	got := ConvertLinkedListToSlice(removeNthFromEnd(NewLinkedListFromSlice([]int{1, 2}), 1))
	expected := []int{1}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
