package main

import (
	"strings"
	"unicode"
)

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i := 0
	j := len(s) - 1
	for i < j {
		if !isAlphanumeric(rune(s[i])) {
			i++
			continue
		}
		if !isAlphanumeric(rune(s[j])) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
