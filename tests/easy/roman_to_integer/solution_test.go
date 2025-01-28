package roman_to_integer

import (
	"testing"
)

func Test_romanToInt_1(t *testing.T) {
	t.SkipNow()
	got := romanToInt("III")
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_romanToInt_2(t *testing.T) {
	t.SkipNow()
	got := romanToInt("LVIII")
	expected := 58
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_romanToInt_3(t *testing.T) {
	t.SkipNow()
	got := romanToInt("MCMXCIV")
	expected := 1994
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
