package search_insert_position

import (
	"testing"
)

func Test_searchInsert_1(t *testing.T) {
	t.SkipNow()
	got := searchInsert([]int{1, 3, 5, 6}, 5)
	expected := 2
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_searchInsert_2(t *testing.T) {
	t.SkipNow()
	got := searchInsert([]int{1, 3, 5, 6}, 2)
	expected := 1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_searchInsert_3(t *testing.T) {
	t.SkipNow()
	got := searchInsert([]int{1, 3, 5, 6}, 7)
	expected := 4
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
