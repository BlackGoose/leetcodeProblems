package main

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int, 0)
	result := make([]int, 2)
	for i, num := range nums {
		v, ok := dict[target-num]
		if ok {
			result = []int{i, v}
		}
		dict[num] = i
	}
	return result
}
