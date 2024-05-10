package main

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	frequencyMap := make(map[int]int, 0)
	for _, v := range nums {
		frequencyMap[v]++
	}
	sortNums := make([]int, 0)
	for v := range frequencyMap {
		sortNums = append(sortNums, v)
	}
	sort.Slice(sortNums, func(i, j int) bool {
		return frequencyMap[sortNums[i]] > frequencyMap[sortNums[j]]
	})
	return sortNums[:k]
}
