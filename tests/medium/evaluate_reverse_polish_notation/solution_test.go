package evaluate_reverse_polish_notation

import (
	"testing"
)

func Test_evalRPN_1(t *testing.T) {
	t.SkipNow()
	got := evalRPN([]string{"2", "1", "+", "3", "*"})
	expected := 9
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_evalRPN_2(t *testing.T) {
	t.SkipNow()
	got := evalRPN([]string{"4", "13", "5", "/", "+"})
	expected := 6
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_evalRPN_3(t *testing.T) {
	// t.SkipNow()
	got := evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	expected := 22
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
