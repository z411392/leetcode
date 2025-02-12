package fraction_to_recurring_decimal

import (
	"testing"
)

func Test_fractionToDecimal_1(t *testing.T) {
	// t.SkipNow()
	got := fractionToDecimal(1, 2)
	expected := "0.5"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_fractionToDecimal_2(t *testing.T) {
	// t.SkipNow()
	got := fractionToDecimal(2, 1)
	expected := "2"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_fractionToDecimal_3(t *testing.T) {
	// t.SkipNow()
	got := fractionToDecimal(4, 333)
	expected := "0.(012)"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
