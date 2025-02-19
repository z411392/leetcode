package swap_nodes_in_pairs

//lint:file-ignore ST1001 _

import (
	"reflect"
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func Test_swapPairs_1(t *testing.T) {
	// t.SkipNow()
	got := ConvertLinkedListToSlice(swapPairs(NewLinkedListFromSlice([]int{1, 2, 3, 4})))
	expected := []int{2, 1, 4, 3}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_swapPairs_2(t *testing.T) {
	// t.SkipNow()
	got := ConvertLinkedListToSlice(swapPairs(NewLinkedListFromSlice([]int{})))
	expected := []int{}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_swapPairs_3(t *testing.T) {
	// t.SkipNow()
	got := ConvertLinkedListToSlice(swapPairs(NewLinkedListFromSlice([]int{1})))
	expected := []int{1}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_swapPairs_4(t *testing.T) {
	// t.SkipNow()
	got := ConvertLinkedListToSlice(swapPairs(NewLinkedListFromSlice([]int{1, 2, 3})))
	expected := []int{2, 1, 3}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
