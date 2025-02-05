package count_and_say

import (
	"testing"
)

func Test_myPow_1(t *testing.T) {
	t.SkipNow()
	got := myPow(2.00000, 10)
	expected := 1024.00000
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_myPow_2(t *testing.T) {
	t.SkipNow()
	got := myPow(2.10000, 3)
	expected := 9.26100
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_myPow_3(t *testing.T) {
	t.SkipNow()
	got := myPow(2.00000, -2)
	expected := 0.25000
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_myPow_4(t *testing.T) {
	t.SkipNow()
	got := myPow(2.00, -200000000)
	expected := 0.0
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
