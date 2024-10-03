package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	longestLength := 0
	for i := 0; i < len(s); {
		hash := make(map[byte]bool)
		for i < len(s) && !hash[s[i]] {
			hash[s[i]] = true
			i++
		}
		if len(hash) > longestLength {
			longestLength = len(hash)
		}
	}
	return longestLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dfvdf"))
}
