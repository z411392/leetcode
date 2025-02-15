package product_of_array_except_self

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf_1(t *testing.T) {
	// t.SkipNow()
	got := productExceptSelf([]int{1, 2, 3, 4})
	expected := []int{24, 12, 8, 6}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_productExceptSelf_2(t *testing.T) {
	// t.SkipNow()
	got := productExceptSelf([]int{-1, 1, 0, -3, 3})
	expected := []int{0, 0, 9, 0, 0}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
