package best_time_to_buy_and_sell_stock_ii

/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.
*/
func maxProfit(prices []int) int {
	profit := 0
	prev := prices[0]
	// 如果存在有價差的組合不會因為早賣或晚賣而總額有所不同；也就是說，一但有利潤就直接賣出的結果會跟各種組合串遍之後得到的最佳結果是一樣的
	for i := 1; i < len(prices); i++ {
		if prices[i]-prev > 0 {
			profit += prices[i] - prev
		}
		prev = prices[i]
	}

	return profit
}
