package majority_element

import (
	"testing"
)

func Test_majorityElement_1(t *testing.T) {
	t.SkipNow()
	got := majorityElement([]int{3, 2, 3})
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_majorityElement_2(t *testing.T) {
	t.SkipNow()
	got := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	expected := 2
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
