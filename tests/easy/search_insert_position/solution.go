package search_insert_position

/*
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.
*/
func searchInsert(nums []int, target int) int {
	index := 0
	for _, num := range nums {
		if num >= target {
			break
		}
		index += 1
	}
	return index
}
