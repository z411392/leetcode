package find_the_duplicate_number

/*
Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and using only constant extra space.
*/
func findDuplicate(nums []int) int {
	// https://englishandcoding.pixnet.net/blog/post/30106939
	turtoise := nums[0]
	hare := nums[0]
	for {
		turtoise = nums[turtoise]
		hare = nums[nums[hare]]
		// fmt.Printf("turtoise=%v, hare=%v\n", turtoise, hare)
		if turtoise == hare {
			// 此時 turtoise 走了 2(L + x) 而 hare 走了 L + x + nC。此時是 L + x + nC 而還不是 L + x！
			break
		}
	}
	// 所求發生在烏龜重走一次 L + x 時
	// 此時兔子會走 nC，正因為兔子原本停在交會點，而 nC 是 C 的倍數，兔子與烏龜會重新交會
	turtoise = nums[0]
	for {
		if turtoise == hare {
			break
		}
		turtoise = nums[turtoise]
		hare = nums[hare]
	}
	return turtoise
}
