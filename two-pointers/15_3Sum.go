package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	result := [][]int{}
	for i, v := range nums {
		if i > 0 && nums[i-1] == v {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if v+nums[left]+nums[right] == 0 {
				result = append(result, []int{v, nums[left], nums[right]})
				left++
				right--
				for nums[left-1] == nums[left] && left < right {
					left++
				}
			} else if v+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
}
