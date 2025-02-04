package find_first_and_last_position_of_element_in_sorted_array

import (
	"reflect"
	"testing"
)

func Test_searchRange_1(t *testing.T) {
	t.SkipNow()
	got := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	expected := []int{3, 4}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_searchRange_2(t *testing.T) {
	t.SkipNow()
	got := searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	expected := []int{-1, -1}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_searchRange_3(t *testing.T) {
	// t.SkipNow()
	got := searchRange([]int{}, 0)
	expected := []int{-1, -1}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
