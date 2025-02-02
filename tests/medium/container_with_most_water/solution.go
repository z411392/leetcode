package container_with_most_water

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.
*/
func maxArea(height []int) int {
	// https://leetcode.com/problems/container-with-most-water/solutions/6308492/simple-solution-0-ms-beats-100-00-go
	left, right := 0, len(height)-1 // 從最外圍開始找
	maxArea := 0
	// 過程中總是保持最佳策略避免極值被略過
	for {
		if left >= right { // 柱子重疊或者左邊超過右邊就停止
			break
		}
		w := right - left
		h := min(height[left], height[right])
		area := w * h
		if area > maxArea {
			maxArea = area
		}
		// 如果左邊比較高就移動右邊反之亦然
		if height[left] > height[right] {
			right -= 1
		} else {
			left += 1
		}
	}
	return maxArea
}
