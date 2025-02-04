package permutations

import (
	"reflect"
	"testing"
)

func Test_permute_1(t *testing.T) {
	t.SkipNow()
	got := permute([]int{1, 2, 3})
	expected := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_permute_2(t *testing.T) {
	t.SkipNow()
	got := permute([]int{0, 1})
	expected := [][]int{{0, 1}, {1, 0}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_permute_3(t *testing.T) {
	t.SkipNow()
	got := permute([]int{1})
	expected := [][]int{{1}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
