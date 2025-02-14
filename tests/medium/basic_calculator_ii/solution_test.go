package basic_calculator_ii

import (
	"testing"
)

func Test_calculate_1(t *testing.T) {
	t.SkipNow()
	got := calculate("3+2*2")
	expected := 7
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_calculate_2(t *testing.T) {
	t.SkipNow()
	got := calculate(" 3/2 ")
	expected := 1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_calculate_3(t *testing.T) {
	t.SkipNow()
	got := calculate(" 3+5 / 2 ")
	expected := 5
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_calculate_4(t *testing.T) {
	t.SkipNow()
	got := calculate("1-1+1")
	expected := 1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_calculate_5(t *testing.T) {
	// t.SkipNow()
	got := calculate("1*2-3/4+5*6-7*8+9/10")
	expected := -24
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
