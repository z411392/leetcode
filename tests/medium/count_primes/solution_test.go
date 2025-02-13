package count_primes

import (
	"testing"
)

func Test_countPrimes_1(t *testing.T) {
	// t.SkipNow()
	got := countPrimes(10)
	expected := 4
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_countPrimes_2(t *testing.T) {
	// t.SkipNow()
	got := countPrimes(0)
	expected := 0
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_countPrimes_3(t *testing.T) {
	// t.SkipNow()
	got := countPrimes(1)
	expected := 0
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
