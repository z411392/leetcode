package set_matrix_zeroes

import (
	"reflect"
	"testing"
)

func Test_setZeroes_1(t *testing.T) {
	// t.SkipNow()
	got := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(got)
	expected := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_setZeroes_2(t *testing.T) {
	// t.SkipNow()
	got := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(got)
	expected := [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
