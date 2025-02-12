package factorial_trailing_zeroes

import (
	"testing"
)

func Test_trailingZeroes_1(t *testing.T) {
	t.SkipNow()
	got := trailingZeroes(3)
	expected := 0
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_trailingZeroes_2(t *testing.T) {
	t.SkipNow()
	got := trailingZeroes(5)
	expected := 1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_trailingZeroes_3(t *testing.T) {
	t.SkipNow()
	got := trailingZeroes(0)
	expected := 0
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_trailingZeroes_4(t *testing.T) {
	// t.SkipNow()
	got := trailingZeroes(30)
	expected := 7
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
