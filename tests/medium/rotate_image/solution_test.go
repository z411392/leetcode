package rotate_image

import (
	"reflect"
	"testing"
)

func Test_rotate_1(t *testing.T) {
	t.SkipNow()
	got := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(got)
	expected := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_rotate_2(t *testing.T) {
	t.SkipNow()
	got := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(got)
	expected := [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
