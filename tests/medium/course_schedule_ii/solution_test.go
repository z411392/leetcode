package course_schedule_ii

import (
	"reflect"
	"testing"
)

func Test_findOrder_1(t *testing.T) {
	// t.SkipNow()
	got := findOrder(2, [][]int{{1, 0}})
	expected := []int{0, 1}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_findOrder_2(t *testing.T) {
	// t.SkipNow()
	got := findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	expected := []int{0, 1, 2, 3}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
