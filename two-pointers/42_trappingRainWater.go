package main

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	left, right := 0, len(height)-1
	maxLeft, maxRight := -1, -1
	waterAmount := 0

	for left < right {
		if height[left] <= height[right] {
			if maxLeft < height[left] {
				maxLeft = height[left]
			}
			waterAmount += maxLeft - height[left]
			left++
		} else {
			if maxRight < height[right] {
				maxRight = height[right]
			}
			waterAmount += maxRight - height[right]
			right--
		}
	}
	return waterAmount
}
