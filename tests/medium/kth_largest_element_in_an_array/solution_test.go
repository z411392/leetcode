package kth_largest_element_in_an_array

import (
	"reflect"
	"testing"
)

func Test_findKthLargest_1(t *testing.T) {
	// t.SkipNow()
	got := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	expected := 5
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_findKthLargest_2(t *testing.T) {
	t.SkipNow()
	got := findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
	expected := 4
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
