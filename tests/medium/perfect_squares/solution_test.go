package perfect_squares

import (
	"testing"
)

func Test_numSquares_1(t *testing.T) {
	t.SkipNow()
	got := numSquares(12)
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_numSquares_2(t *testing.T) {
	// t.SkipNow()
	got := numSquares(13)
	expected := 2
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
