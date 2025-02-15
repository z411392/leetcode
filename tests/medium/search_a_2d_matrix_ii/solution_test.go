package search_a_2d_matrix_ii

import (
	"testing"
)

func Test_isearchMatrix_1(t *testing.T) {
	// t.SkipNow()
	got := searchMatrix([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 5)
	expected := true
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_isearchMatrix_2(t *testing.T) {
	// t.SkipNow()
	got := searchMatrix([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 5)
	expected := true
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
