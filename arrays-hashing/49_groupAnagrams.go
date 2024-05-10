package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	srtsMap := make(map[string][]string, 0)

	for _, str := range strs {
		sortStr := []byte(str)
		sort.Slice(sortStr, func(i, j int) bool {
			return sortStr[i] < sortStr[j]
		})
		srtsMap[string(sortStr)] = append(srtsMap[string(sortStr)], str)
	}
	result := [][]string{}
	for _, v := range srtsMap {
		result = append(result, v)
	}
	return result
}
