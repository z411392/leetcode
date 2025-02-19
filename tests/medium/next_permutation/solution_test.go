package next_permutation

import (
	"reflect"
	"testing"
)

func Test_nextPermutation_1(t *testing.T) {
	t.SkipNow()
	got := []int{1, 2, 3}
	nextPermutation(got)
	expected := []int{1, 3, 2}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_nextPermutation_2(t *testing.T) {
	t.SkipNow()
	got := []int{3, 2, 1}
	nextPermutation(got)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_nextPermutation_3(t *testing.T) {
	t.SkipNow()
	got := []int{1, 1, 5}
	nextPermutation(got)
	expected := []int{1, 5, 1}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_nextPermutation_4(t *testing.T) {
	// t.SkipNow()
	got := []int{1, 5, 1}
	nextPermutation(got)
	expected := []int{5, 1, 1}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
