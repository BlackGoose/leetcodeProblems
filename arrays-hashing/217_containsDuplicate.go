package main

func containsDuplicate(nums []int) bool {
	setNums := make(map[int]bool, 0)
	for _, num := range nums {
		contain := setNums[num]
		if contain {
			return true
		} else {
			setNums[num] = true
		}
	}
	return false
}
