package coin_change

import "sort"

func coinChange(coins []int, amount int) int {
	// 1. 預先排序硬幣，大的在前面
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	// 2. 優化初始化
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	// 3. 對每個金額進行運算
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			// 4. 提前終止
			if coin > i {
				continue
			}
			// 5. 剪枝：如果已經無法得到更好的解，就跳過
			if dp[i-coin] == amount+1 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
