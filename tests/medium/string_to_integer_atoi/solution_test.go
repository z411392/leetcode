package string_to_integer_atoi

import (
	"testing"
)

func Test_myAtoi_1(t *testing.T) {
	t.SkipNow()
	got := myAtoi("42")
	expected := 42
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_myAtoi_2(t *testing.T) {
	t.SkipNow()
	got := myAtoi("-042")
	expected := -42
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_myAtoi_3(t *testing.T) {
	t.SkipNow()
	got := myAtoi("1337c0d3")
	expected := 1337
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_myAtoi_4(t *testing.T) {
	t.SkipNow()
	got := myAtoi("0-1")
	expected := 0
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_myAtoi_5(t *testing.T) {
	t.SkipNow()
	got := myAtoi("words and 987")
	expected := 0
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_myAtoi_6(t *testing.T) {
	t.SkipNow()
	got := myAtoi("+-12")
	expected := 0
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_myAtoi_7(t *testing.T) {
	t.SkipNow()
	got := myAtoi("   -042")
	expected := -42
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
