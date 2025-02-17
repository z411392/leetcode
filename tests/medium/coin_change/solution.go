package coin_change

import (
	"slices"
)

func coinChange(coins []int, amount int) int {
	slices.Sort(coins)
	// fmt.Printf("%v\n", coins)
	if amount == 0 {
		return 0
	}
	fewestNumberOfCoins := 0
	for j := len(coins) - 1; j > -1; j-- {
		if amount < coins[j] {
			continue
		}
		fewestNumberOfCoins += amount / coins[j]
		amount %= coins[j]
	}
	if amount > 0 {
		return -1
	}
	return fewestNumberOfCoins
}
