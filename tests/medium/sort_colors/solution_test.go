package sort_colors

import (
	"reflect"
	"testing"
)

func Test_sortColors_1(t *testing.T) {
	// t.SkipNow()
	got := []int{2, 0, 2, 1, 1, 0}
	sortColors(got)
	expected := []int{0, 0, 1, 1, 2, 2}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_sortColors_2(t *testing.T) {
	// t.SkipNow()
	got := []int{2, 0, 1}
	sortColors(got)
	expected := []int{0, 1, 2}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
