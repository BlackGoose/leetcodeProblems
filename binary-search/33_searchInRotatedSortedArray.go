package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (left + right) / 2

		if nums[middle] == target {
			return middle
		}

		if nums[left] < nums[middle] {
			if nums[left] > target || nums[middle] < target {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			if nums[right] < target || nums[middle] > target {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return -1
}
