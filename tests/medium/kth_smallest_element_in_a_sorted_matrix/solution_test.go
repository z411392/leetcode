package kth_smallest_element_in_a_sorted_matrix

import (
	"reflect"
	"testing"
)

func Test_kthSmallest_1(t *testing.T) {
	// t.SkipNow()
	got := kthSmallest([][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}, 8)
	expected := 13
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_kthSmallest_2(t *testing.T) {
	// t.SkipNow()
	got := kthSmallest([][]int{{-5}}, 1)
	expected := 5
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
