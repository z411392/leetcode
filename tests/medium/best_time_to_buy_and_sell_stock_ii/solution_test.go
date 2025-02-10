package best_time_to_buy_and_sell_stock_ii

import (
	"testing"
)

func Test_maxProfit_1(t *testing.T) {
	// t.SkipNow()
	got := maxProfit([]int{7, 1, 5, 3, 6, 4})
	expected := 7
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_maxProfit_2(t *testing.T) {
	// t.SkipNow()
	got := maxProfit([]int{1, 2, 3, 4, 5})
	expected := 4
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_maxProfit_3(t *testing.T) {
	// t.SkipNow()
	got := maxProfit([]int{7, 6, 4, 3, 1})
	expected := 0
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
