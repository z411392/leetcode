package subsets

import (
	"reflect"
	"testing"
)

func Test_subsets_1(t *testing.T) {
	t.SkipNow()
	got := subsets([]int{1, 2, 3})
	expected := [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_subsets_2(t *testing.T) {
	t.SkipNow()
	got := subsets([]int{0})
	expected := [][]int{{}, {0}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
