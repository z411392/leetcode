package jump_game

/*
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.
*/
func canJump(nums []int) bool {
	// https://leetcode.com/problems/jump-game/submissions/1531955492/
	scope := 0
	for i := 0; i < len(nums); i++ {
		if i > scope {
			return false
		}
		reached := i + nums[i]
		if reached > scope {
			scope = reached
		}
	}
	return true
}
