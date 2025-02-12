package maximum_product_subarray

import (
	"testing"
)

func Test_maxProduct_1(t *testing.T) {
	// t.SkipNow()
	got := maxProduct([]int{2, 3, -2, 4})
	expected := 6
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_maxProduct_2(t *testing.T) {
	// t.SkipNow()
	got := maxProduct([]int{-2, 0, -1})
	expected := 0
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
