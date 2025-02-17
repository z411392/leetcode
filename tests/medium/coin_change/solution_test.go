package coin_change

import (
	"testing"
)

func Test_coinChange_1(t *testing.T) {
	t.SkipNow()
	got := coinChange([]int{1, 2, 5}, 11)
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_coinChange_2(t *testing.T) {
	t.SkipNow()
	got := coinChange([]int{2}, 3)
	expected := -1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_coinChange_3(t *testing.T) {
	t.SkipNow()
	got := coinChange([]int{1}, 0)
	expected := 0
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_coinChange_4(t *testing.T) {
	t.SkipNow()
	got := coinChange([]int{2, 5, 10, 1}, 27)
	expected := 4
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_coinChange_5(t *testing.T) {
	// t.SkipNow()
	got := coinChange([]int{186, 419, 83, 408}, 20)
	expected := 4
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
