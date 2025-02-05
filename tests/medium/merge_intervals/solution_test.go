package merge_intervals

import (
	"reflect"
	"testing"
)

func Test_merge_1(t *testing.T) {
	t.SkipNow()
	got := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_merge_2(t *testing.T) {
	t.SkipNow()
	got := merge([][]int{{1, 4}, {4, 5}})
	expected := [][]int{{1, 5}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_merge_3(t *testing.T) {
	t.SkipNow()
	got := merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}})
	expected := [][]int{{1, 10}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_merge_4(t *testing.T) {
	t.SkipNow()
	got := merge([][]int{{1, 4}, {2, 3}})
	expected := [][]int{{1, 4}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
