package gas_station

import (
	"testing"
)

func Test_canCompleteCircuit_1(t *testing.T) {
	// t.SkipNow()
	got := canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_canCompleteCircuit_2(t *testing.T) {
	// t.SkipNow()
	got := canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
