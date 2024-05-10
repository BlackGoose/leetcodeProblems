package main

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	prefix := 1
	for i, v := range nums {
		result[i] = prefix
		prefix *= v
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}
