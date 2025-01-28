package best_time_to_buy_and_sell_stock

import (
	"testing"
)

func Test_maxProfit_1(t *testing.T) {
	t.SkipNow()
	got := maxProfit([]int{7, 1, 5, 3, 6, 4})
	expected := 5
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_maxProfit_2(t *testing.T) {
	t.SkipNow()
	got := maxProfit([]int{7, 6, 4, 3, 1})
	expected := 0
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
