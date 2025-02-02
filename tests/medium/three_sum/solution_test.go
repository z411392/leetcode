package three_sum

import (
	"reflect"
	"testing"
)

func Test_threeSum_1(t *testing.T) {
	// t.SkipNow()
	got := threeSum([]int{-1, 0, 1, 2, -1, -4})
	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_threeSum_2(t *testing.T) {
	// t.SkipNow()
	got := threeSum([]int{0, 1, 1})
	expected := [][]int{}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_threeSum_3(t *testing.T) {
	// t.SkipNow()
	got := threeSum([]int{0, 0, 0})
	expected := [][]int{{0, 0, 0}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
