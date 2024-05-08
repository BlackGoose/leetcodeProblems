package main

func longestConsecutive(nums []int) int {
	uniqNums := map[int]bool{}
	longestLength := 0

	for _, v := range nums {
		uniqNums[v] = true
	}

	for k := range uniqNums {
		if !uniqNums[k-1] {
			length := 1
			for uniqNums[k+length] {
				length++
			}
			if length > longestLength {
				longestLength = length
			}
		}
	}
	return longestLength
}
