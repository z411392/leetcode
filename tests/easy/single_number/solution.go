package single_number

/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.
*/
// func singleNumber(nums []int) int {
// 	var counters = map[int]int{}
// 	for _, num := range nums {
// 		_, exists := counters[num]
// 		if exists {
// 			counters[num] += 1
// 		} else {
// 			counters[num] = 1
// 		}
// 	}
// 	for num, counter := range counters {
// 		// fmt.Printf("%v %v\n", num, counter)
// 		if counter == 1 {
// 			return num
// 		}
// 	}
// 	return 0
// }

// https://leetcode.com/problems/single-number/solutions/1771771/think-it-through-time-o-n-space-o-1-python-go-explained/
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result = result ^ num // 只有在其他數字都是偶數個的情況才適用（譬如 3, 5 ...）
		// fmt.Printf("%v %v\n", num, result)
	}
	return result
}
