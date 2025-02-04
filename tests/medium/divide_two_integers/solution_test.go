package divide_two_integers

import (
	"testing"
)

func Test_divide_1(t *testing.T) {
	// t.SkipNow()
	got := divide(10, 3)
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_divide_2(t *testing.T) {
	// t.SkipNow()
	got := divide(7, -3)
	expected := -2
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
