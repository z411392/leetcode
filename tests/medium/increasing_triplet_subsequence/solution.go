package increasing_triplet_subsequence

import (
	"math"
)

func increasingTriplet(nums []int) bool {
	// https://leetcode.com/problems/increasing-triplet-subsequence/solutions/2689261/go-python-o-n-time-o-1-space/
	smallest := math.MaxInt
	secondSmallest := math.MaxInt
	for _, num := range nums {
		if num <= smallest { // 必須下 <=，否則 1, 2, 1 時，第二個 1 會更新到 secondSmallest，產生 smallest 和 secondSmallest 是同一個數字的情形
			smallest = num
			// fmt.Printf("smallest=%v\n", smallest)
		} else if num <= secondSmallest {
			secondSmallest = num
			// fmt.Printf("secondSmallest=%v\n", secondSmallest)
		} else {
			return true
		}
	}
	return false
}
