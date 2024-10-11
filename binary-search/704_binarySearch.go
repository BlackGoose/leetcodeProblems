package main

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (right-left)/2 + left
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}
