package odd_even_linked_list

//lint:file-ignore ST1001 _

import (
	"reflect"
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func Test_oddEvenList_1(t *testing.T) {
	t.SkipNow()
	got := ConvertLinkedListToSlice(oddEvenList(NewLinkedListFromSlice([]int{1, 2, 3, 4, 5})))
	expected := []int{1, 3, 5, 2, 4}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_oddEvenList_2(t *testing.T) {
	t.SkipNow()
	got := ConvertLinkedListToSlice(oddEvenList(NewLinkedListFromSlice([]int{2, 1, 3, 5, 6, 4, 7})))
	expected := []int{2, 3, 6, 7, 1, 5, 4}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_oddEvenList_3(t *testing.T) {
	// t.SkipNow()
	got := ConvertLinkedListToSlice(oddEvenList(NewLinkedListFromSlice([]int{1, 2, 3})))
	expected := []int{1, 3, 2}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
