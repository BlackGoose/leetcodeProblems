package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	letterCount := make(map[rune]int, 0)
	for _, lt := range s {
		letterCount[lt]++
	}
	for _, lt := range t {
		_, ok := letterCount[lt]
		if !ok {
			return false
		}
		letterCount[lt]--
		if letterCount[lt] < 0 {
			return false
		}
	}
	return true
}
