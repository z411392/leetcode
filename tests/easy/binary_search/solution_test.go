package binary_search

import (
	"fmt"
	"testing"
)

func Test_search_1(t *testing.T) {
	t.SkipNow()
	got := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	expected := 4
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_search_2(t *testing.T) {
	t.SkipNow()
	got := search([]int{-1, 0, 3, 5, 9, 12}, 2)
	expected := -1
	if got != expected {
		fmt.Printf("%v\n", got)
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
