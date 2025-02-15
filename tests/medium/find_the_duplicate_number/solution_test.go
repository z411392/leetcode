package find_the_duplicate_number

import (
	"testing"
)

func Test_findDuplicate_1(t *testing.T) {
	// t.SkipNow()
	got := findDuplicate([]int{1, 3, 4, 2, 2})
	expected := 2
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_findDuplicate_2(t *testing.T) {
	// t.SkipNow()
	got := findDuplicate([]int{3, 1, 3, 4, 2})
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_findDuplicate_3(t *testing.T) {
	// t.SkipNow()
	got := findDuplicate([]int{3, 3, 3, 3, 3})
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_findDuplicate_4(t *testing.T) {
	// t.SkipNow()
	got := findDuplicate([]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1})
	expected := 9
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
