package fraction_to_recurring_decimal

import (
	"testing"
)

func Test_fractionToDecimal_1(t *testing.T) {
	t.SkipNow()
	got := fractionToDecimal(1, 2)
	expected := "0.5"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_fractionToDecimal_2(t *testing.T) {
	t.SkipNow()
	got := fractionToDecimal(2, 1)
	expected := "2"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_fractionToDecimal_3(t *testing.T) {
	t.SkipNow()
	got := fractionToDecimal(4, 333)
	expected := "0.(012)"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_fractionToDecimal_4(t *testing.T) {
	t.SkipNow()
	got := fractionToDecimal(22, 7)
	expected := "3.(142857)"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_fractionToDecimal_5(t *testing.T) {
	t.SkipNow()
	got := fractionToDecimal(-50, 8)
	expected := "-6.25"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_fractionToDecimal_6(t *testing.T) {
	// t.SkipNow()
	got := fractionToDecimal(420, 226)
	expected := "1.(8584070796460176991150442477876106194690265486725663716814159292035398230088495575221238938053097345132743362831"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
