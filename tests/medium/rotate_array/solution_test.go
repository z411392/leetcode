package rotate_array

import (
	"reflect"
	"testing"
)

func Test_rotate_1(t *testing.T) {
	// t.SkipNow()
	got := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(got, 3)
	expected := []int{5, 6, 7, 1, 2, 3, 4}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_rotate_2(t *testing.T) {
	// t.SkipNow()
	got := []int{-1, -100, 3, 99}
	rotate(got, 2)
	expected := []int{3, 99, -1, -100}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
