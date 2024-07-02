package main

func maxArea(heights []int) int {
	left, right, max_amount := 0, len(heights)-1, -1
	for left < right {
		width := right - left
		if width*min(heights[left], heights[right]) > max_amount {
			max_amount = width * min(heights[left], heights[right])
		}
		if heights[left] < heights[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return max_amount
}
