package sum_of_two_integers

import (
	"testing"
)

func Test_getSum_1(t *testing.T) {
	// t.SkipNow()
	got := getSum(1, 2)
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_getSum_2(t *testing.T) {
	// t.SkipNow()
	got := getSum(2, 3)
	expected := 5
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
