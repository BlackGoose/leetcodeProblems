package main

func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]bool, 0)
	maxString := 0
	l, r := 0, 0
	for r < len(s) {
		if set[s[r]] {
			for s[r] != s[l] {
				set[s[l]] = false
				l++
			}
			l++
		} else {
			set[s[r]] = true
		}
		if maxString < r-l+1 {
			maxString = r - l + 1
		}
		r++
	}
	return maxString
}
