package two_sum

import (
	"reflect"
	"testing"
)

func Test_twoSum_1(t *testing.T) {
	t.SkipNow()
	got := twoSum([]int{2, 7, 11, 15}, 9)
	expected := []int{0, 1}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_twoSum_2(t *testing.T) {
	t.SkipNow()
	got := twoSum([]int{3, 2, 4}, 6)
	expected := []int{1, 2}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_twoSum_3(t *testing.T) {
	t.SkipNow()
	got := twoSum([]int{3, 3}, 6)
	expected := []int{0, 1}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
