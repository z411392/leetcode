package spiral_matrix

import (
	"reflect"
	"testing"
)

func Test_spiralOrder_1(t *testing.T) {
	t.SkipNow()
	got := spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_spiralOrder_2(t *testing.T) {
	t.SkipNow()
	got := spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})
	expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_spiralOrder_3(t *testing.T) {
	t.SkipNow()
	got := spiralOrder([][]int{{3}, {2}})
	expected := []int{3, 2}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
