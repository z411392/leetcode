package wiggle_sort_ii

import (
	"reflect"
	"testing"
)

func Test_wiggleSort_1(t *testing.T) {
	// t.SkipNow()
	got := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(got)
	expected := []int{1, 4, 1, 5, 1, 6}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_wiggleSort_2(t *testing.T) {
	// t.SkipNow()
	got := []int{1, 3, 2, 2, 3, 1}
	wiggleSort(got)
	expected := []int{1, 2, 1, 3, 2, 3}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
