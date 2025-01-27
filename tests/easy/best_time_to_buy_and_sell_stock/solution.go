package best_time_to_buy_and_sell_stock

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/
func maxProfit(prices []int) int {
	min := 10001
	max := 0
	maxProfit := 0
	for _, price := range prices {
		if price > max {
			max = price
			// 價格創新高時，應該計算目前的獲利，當超過以往最大獲利時要儲存起來
			profit := max - min
			if profit > maxProfit {
				maxProfit = profit
			}
		}
		if price < min {
			min = price
			max = price // 最低價一但更新，過去的最高價便不能用，且當前計算得的最大獲利應該歸零
		}
	}
	return maxProfit
}
